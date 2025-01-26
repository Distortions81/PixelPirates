package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

// GenerateWaveFromTextWithParams creates a single wave for one instrument.
// We pass waveBlend & volume from insData to shape the tone and loudness.
func GenerateWaveFromTextWithParams(text string, bpm, sampleRate int, waveBlend, volume float64) []float32 {
	beatDuration := time.Minute / time.Duration(bpm)
	var finalWave []float32

	if volume == 0 {
		volume = 1
	}

	for _, noteStr := range strings.Split(text, ",") {
		note, duration := ParseNote(noteStr)
		if note == "" {
			continue
		}
		noteDuration := time.Duration(beatDuration.Seconds() * duration * float64(time.Second))

		// Check for chord
		chordNotes := strings.Split(note, "/")
		if len(chordNotes) > 1 {
			chordWave := PlayChordOfflineWithParams(chordNotes, noteDuration, sampleRate, waveBlend, volume)
			finalWave = append(finalWave, chordWave...)
		} else {
			freq := CalculateFrequency(note)
			noteWave := PlayNoteOfflineWithParams(freq, noteDuration, sampleRate, waveBlend, volume)
			finalWave = append(finalWave, noteWave...)
		}
	}

	return finalWave
}

// PlayNoteOfflineWithParams: generates wave data for a single note, applying volume & wave blend.
func PlayNoteOfflineWithParams(freq float64, duration time.Duration, sampleRate int, waveBlend, volume float64) []float32 {
	// Handle rest
	if freq == 0 {
		return make([]float32, int(float64(sampleRate)*duration.Seconds()))
	}

	wave := GenerateCustomWave(freq, duration, sampleRate, waveBlend)
	wave = ApplyADSR(wave, sampleRate, 0.05, 0.1, 0.5, 0.5)

	// Apply per-instrument volume
	for i := range wave {
		wave[i] *= float32(volume)
	}

	return wave
}

// GenerateCustomWave blends square & sine waves based on waveBlend (0.0 to 1.0).
func GenerateCustomWave(freq float64, duration time.Duration, sampleRate int, waveBlend float64) []float32 {
	samples := int(float64(sampleRate) * duration.Seconds())
	wave := make([]float32, samples)
	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)
		if freq == 0 {
			wave[i] = 0
		} else {
			sinVal := math.Sin(2 * math.Pi * freq * t)
			sqrVal := 1.0
			if sinVal < 0 {
				sqrVal = -1.0
			}
			// waveBlend: 0.0 = pure sine, 1.0 = pure square
			mix := waveBlend*sqrVal + (1.0-waveBlend)*sinVal
			wave[i] = float32(mix)
		}
	}
	return wave
}

// PlayChordOfflineWithParams: generates wave data for a chord, applying volume & wave blend.
func PlayChordOfflineWithParams(chord []string, duration time.Duration, sampleRate int, waveBlend, volume float64) []float32 {
	// Generate wave for each note in the chord
	var waves [][]float32
	for _, note := range chord {
		freq := CalculateFrequency(note)
		noteWave := GenerateCustomWave(freq, duration, sampleRate, waveBlend)
		noteWave = ApplyADSR(noteWave, sampleRate, 0.05, 0.1, 0.5, 0.5)
		// Apply volume
		for i := range noteWave {
			noteWave[i] *= float32(volume)
		}
		waves = append(waves, noteWave)
	}

	// Sum waves
	chordWave := make([]float32, len(waves[0]))
	for i := range chordWave {
		var sum float32
		for _, w := range waves {
			sum += w[i]
		}
		// average to prevent single chord from ballooning amplitude
		chordWave[i] = sum / float32(len(waves))
	}

	return chordWave
}

// PlayNoteOffline generates wave data (without playing) for a single note.
func PlayNoteOffline(freq float64, duration time.Duration, sampleRate int) []float32 {
	if freq == 0 {
		// rest
		return make([]float32, int(float64(sampleRate)*duration.Seconds()))
	}
	wave := GenerateSmoothWave(freq, duration, sampleRate)
	wave = ApplyADSR(wave, sampleRate, 0.05, 0.1, 0.5, 0.5)
	return wave
}

// PlayChordOffline generates wave data (without playing) for a chord.
func PlayChordOffline(chord []string, duration time.Duration, sampleRate int) []float32 {
	// Generate wave for each note in the chord
	var waves [][]float32
	for _, note := range chord {
		freq := CalculateFrequency(note)
		noteWave := GenerateSmoothWave(freq, duration, sampleRate)
		noteWave = ApplyADSR(noteWave, sampleRate, 0.05, 0.1, 0.5, 0.5)
		waves = append(waves, noteWave)
	}

	// Sum waves
	chordWave := make([]float32, len(waves[0]))
	for i := range chordWave {
		var sum float32
		for _, w := range waves {
			sum += w[i]
		}
		sum /= float32(len(waves)) // average
		chordWave[i] = sum
	}
	return chordWave
}

// MixWaves sums multiple mono wave slices (all same sample rate)
// 1) Averages by number of wave inputs
// 2) Scales further only if needed to prevent clipping
func MixWaves(waves ...[]float32) []float32 {
	// 1. Determine the maximum length among all input waves
	var maxLen int
	for _, w := range waves {
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	// 2. Sum the waves
	mixed := make([]float32, maxLen)
	for _, w := range waves {
		for i := 0; i < len(w); i++ {
			mixed[i] += w[i]
		}
	}

	// 3. Average by number of waves
	numWaves := float32(len(waves))
	if numWaves > 1.0 {
		for i := 0; i < maxLen; i++ {
			mixed[i] /= numWaves + 1
		}
	}

	// 4. Find the peak amplitude (absolute value)
	var maxAmp float32
	for _, sample := range mixed {
		absVal := sample
		if absVal < 0 {
			absVal = -absVal
		}
		if absVal > maxAmp {
			maxAmp = absVal
		}
	}

	// 5. If the peak amplitude exceeds 1.0, scale the whole wave down
	if maxAmp > 1.0 {
		scale := 1.0 / maxAmp
		for i := range mixed {
			mixed[i] *= scale
		}
	}

	return mixed
}

func PlayWave(wave []float32, audioContext *audio.Context, sampleRate int) {
	// 1) Trim trailing silence
	wave = trimTrailingSilence(wave, 0.0001) // remove samples < 0.0001 near the end

	// 2) Convert float32 samples to raw bytes (16-bit PCM)
	soundData := make([]byte, len(wave)*2)
	for i, sample := range wave {
		val := int16(sample * 32767)
		soundData[i*2] = byte(val)
		soundData[i*2+1] = byte(val >> 8)
	}

	// 3) Create a player and play
	player := audioContext.NewPlayerFromBytes(soundData)
	player.Play()

	// 4) Wait for playback to finish
	duration := time.Duration(float64(len(wave)/2) / float64(sampleRate) * float64(time.Second))
	time.Sleep(duration)
}

// trimTrailingSilence cuts off samples at the end of `wave` that are
// below the specified `threshold`. This helps remove long silent tails.
func trimTrailingSilence(wave []float32, threshold float32) []float32 {
	end := len(wave) - 1
	for end >= 0 {
		if wave[end] > threshold || wave[end] < -threshold {
			break
		}
		end--
	}
	// If 'end' is -1, it means the entire wave is silence
	if end < 0 {
		return []float32{} // or wave[:0] to keep same slice reference
	}
	return wave[:end+1]
}

// AddLeMakeSilenceader inserts a specified pause (in seconds) of silence
// at the start of the wave.
func MakeSilence(sampleRate int, pauseSeconds float64) []float32 {
	// Calculate how many samples correspond to the pause
	numSilenceSamples := int(float64(sampleRate) * pauseSeconds)

	// Create a slice of zeros (silence)
	silence := make([]float32, numSilenceSamples)

	return silence
}

// CalculateFrequency calculates the frequency of a note based on its name and octave.
func CalculateFrequency(note string) float64 {
	// Base note A4
	baseFrequency := 440.0
	// Note names (A-G), standard equal temperament tuning
	noteNames := map[string]int{
		"NN": -1, "Ab": 0, "A#": 1, "Bb": 2, "Cb": 3, "C#": 4, "Db": 5,
		"D#": 6, "Eb": 7, "Fb": 8, "F#": 9, "Gb": 10, "G#": 11,
	}

	// Note names are of the form "A1", "C#4", etc.
	// First, extract the note (A, B, C, etc.) and the octave number
	var noteName string
	var octave int
	fmt.Sscanf(note, "%2s%d", &noteName, &octave)

	// Find the index of the note (A, A#, B, etc.)
	halfSteps := noteNames[noteName]
	if halfSteps == -1 {
		return 0
	}

	// Calculate the number of half-steps from A4 (which is the 49th note)
	halfStepsFromA4 := (octave-6)*12 + halfSteps
	// Frequency of the note
	frequency := baseFrequency * math.Pow(2, float64(halfStepsFromA4)/12)
	return frequency
}

// ApplyADSR applies an ADSR envelope to the wave to smooth the note transitions.
func ApplyADSR(wave []float32, sampleRate int, attack, decay, sustain, release float64) []float32 {
	length := len(wave)
	adsrWave := make([]float32, length)

	// Calculate the number of samples for each phase
	attackSamples := int(float64(sampleRate) * attack)
	decaySamples := int(float64(sampleRate) * decay)
	releaseSamples := int(float64(sampleRate) * release)
	sustainSamples := length - attackSamples - decaySamples - releaseSamples

	// Ensure sustainSamples is non-negative; if not, adjust to 0 to avoid out-of-bounds
	if sustainSamples < 0 {
		sustainSamples = 0
	}

	// Apply the Attack phase
	for i := 0; i < attackSamples && i < length; i++ {
		adsrWave[i] = wave[i] * float32(float64(i)/float64(attackSamples))
	}

	// Apply the Decay phase
	for i := attackSamples; i < attackSamples+decaySamples && i < length; i++ {
		adsrWave[i] = wave[i] * float32(1-(1-sustain)*(float64(i-attackSamples)/float64(decaySamples)))
	}

	// Apply the Sustain phase
	for i := attackSamples + decaySamples; i < attackSamples+decaySamples+sustainSamples && i < length; i++ {
		adsrWave[i] = wave[i] * float32(sustain)
	}

	// Apply the Release phase
	for i := attackSamples + decaySamples + sustainSamples; i < length; i++ {
		// Prevent any out-of-bounds access by capping index properly
		adsrWave[i] = wave[i] * float32(sustain*(1-float64(i-(attackSamples+decaySamples+sustainSamples))/float64(releaseSamples)))
	}

	return adsrWave
}

func GenerateSmoothWave(freq float64, duration time.Duration, sampleRate int) []float32 {
	samples := int(float64(sampleRate) * duration.Seconds())
	wave := make([]float32, samples)
	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)

		if freq == 0 {
			wave[i] = 0
		} else {
			// Blend square and sine waves for a smoother sound
			squareWave := 1.0
			if math.Sin(2*math.Pi*freq*t) < 0 {
				squareWave = -1.0
			}
			sineWave := math.Sin(2 * math.Pi * freq * t)
			wave[i] = float32(0.3*squareWave + 0.5*sineWave) // 50% square, 50% sine
		}
	}
	return wave
}

func PlayNote(freq float64, duration time.Duration, sampleRate int, audioContext *audio.Context) {
	if freq == 0 {
		time.Sleep(duration) // Handle rests by sleeping
		return
	}

	// Generate a smoother wave
	wave := GenerateSmoothWave(freq, duration, sampleRate)

	// Apply an ADSR envelope (adjust parameters for smoother transitions)
	wave = ApplyADSR(wave, sampleRate, 0.05, 0.1, 0.5, 0.5)

	// Convert wave to []byte suitable for ebiten
	soundData := make([]byte, len(wave)*2)
	for i, sample := range wave {
		// Convert to int16 (2 bytes per sample)
		val := int16(sample * 32767)
		soundData[i*2] = byte(val)
		soundData[i*2+1] = byte(val >> 8)
	}

	// Load the sound data into an AudioPlayer
	player := audioContext.NewPlayerFromBytes(soundData)

	// Play the sound
	player.Play()

	// Overlap the notes slightly by reducing sleep duration
	time.Sleep(duration - (duration / 8)) // Reduce sleep time to avoid sharp note cuts
}

func PlayChord(chord []string, duration time.Duration, sampleRate int, audioContext *audio.Context) {
	// Generate frequencies for the chord notes
	var waves [][]float32
	for _, note := range chord {
		freq := CalculateFrequency(note)
		wave := GenerateSmoothWave(freq, duration, sampleRate)
		waves = append(waves, wave)
	}

	// Combine the waves by adding their samples together
	chordWave := make([]float32, len(waves[0]))
	for i := 0; i < len(chordWave); i++ {
		var sampleSum float32
		for _, wave := range waves {
			sampleSum += wave[i]
		}
		chordWave[i] = sampleSum / float32(len(waves)) // Average for multiple notes
	}

	// Apply ADSR envelope
	chordWave = ApplyADSR(chordWave, sampleRate, 0.02, 0.1, 0.5, 0.5)

	// Convert wave to []byte suitable for ebiten
	soundData := make([]byte, len(chordWave)*2)
	for i, sample := range chordWave {
		val := int16(sample * 32767)
		soundData[i*2] = byte(val)
		soundData[i*2+1] = byte(val >> 8)
	}

	// Load the sound data into an AudioPlayer
	player := audioContext.NewPlayerFromBytes(soundData)

	// Play the chord
	player.Play()

	// Overlap the notes slightly by reducing sleep duration
	time.Sleep(duration - (duration / 8)) // Reduce sleep time to avoid sharp note cuts
}

// ParseNote parses the note and its duration from a string like "A4 1".
func ParseNote(noteStr string) (string, float64) {
	parts := strings.Fields(noteStr)
	if len(parts) != 2 {
		return "", 0 // Invalid input
	}
	note := parts[0]
	var duration float64
	_, err := fmt.Sscanf(parts[1], "%f", &duration)
	if err != nil {
		return "", 0 // Invalid duration
	}
	return note, duration
}

// PlayTextAsNotes converts text to notes and plays them in 4/4 time with variable note lengths.
func PlayTextAsNotes(text string, bpm int, sampleRate int, audioContext *audio.Context) {
	// 4/4 time: one quarter note per beat
	beatDuration := time.Minute / time.Duration(bpm) // duration of one beat (quarter note)

	for _, noteStr := range strings.Split(text, ",") {
		// Parse the note and its duration (e.g., "A4 1" -> note "A4", duration 1)
		note, duration := ParseNote(noteStr)
		if note == "" {
			continue
		}

		// Check if the note is a chord (comma-separated notes, e.g., "C4,E4,G4")
		chordNotes := strings.Split(note, "/")
		if len(chordNotes) > 1 {
			// It's a chord, play all notes simultaneously
			PlayChord(chordNotes, time.Duration(beatDuration.Seconds()*duration*float64(time.Second)), sampleRate, audioContext)
		} else {
			// It's a single note, play it
			freq := CalculateFrequency(note)
			PlayNote(freq, time.Duration(beatDuration.Seconds()*duration*float64(time.Second)), sampleRate, audioContext)
		}
	}
}

// Main function to set up Ebiten and audio
func playMusic() {
	sampleRate := 44100
	audioContext := audio.NewContext(sampleRate)

	time.Sleep(time.Second)
	for {
		for _, song := range songList {
			startTime := time.Now()
			output := playSong(*song, sampleRate)
			fmt.Printf("Render: %v, Playing %v...\n", time.Since(startTime).Round(time.Millisecond), song.name)
			PlayWave(output, audioContext, sampleRate)
			//SaveMono16BitWav("songs/"+song.name+".wav", sampleRate, output)
			time.Sleep(time.Second)

		}
		fmt.Println("\nRestarting playlist...")
	}
}

func playSong(song songData, sampleRate int) []float32 {
	var waves [][]float32

	for _, instrument := range song.ins {
		// We'll assume we stored volume & waveBlend in instrument.volume, instrument.waveBlend
		insWave := GenerateWaveFromTextWithParams(
			instrument.data,
			song.bpm,
			sampleRate,
			instrument.blend,
			instrument.volume,
		)
		waves = append(waves, insWave)
	}

	// Mix all instrument waves
	mixed := MixWaves(waves...) // your no-clipping version

	return mixed
}
