package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

// Main function to set up Ebiten and audio
func playMusic() {
	sampleRate := 48000
	audioContext := audio.NewContext(sampleRate)

	time.Sleep(time.Second)
	for {
		for _, song := range songList {
			startTime := time.Now()
			fmt.Printf("Rendering: '%v'", song.name)
			output := playSong(song, sampleRate)

			if song.reverb > 0 {
				fmt.Printf(" (Took: %v)\nAdding reverb: ", time.Since(startTime).Round(time.Millisecond))
				output = ApplyReverb(output, sampleRate, song.delay, song.feedback, song.reverb)
			}
			fmt.Printf(" (Took: %v)\nNow Playing: %v.\n\n", time.Since(startTime).Round(time.Millisecond), song.name)

			PlayWave(output, audioContext, sampleRate)
			//SaveMono16BitWav("songs/"+song.name+".wav", sampleRate, output)
			time.Sleep(time.Second)

		}
		fmt.Println("\nRestarting playlist...")
	}
}

func playSong(song songData, sampleRate int) []float64 {
	var waves [][]float64

	for _, instrument := range song.ins {
		if instrument.volume <= 0 {
			continue
		}
		// We'll assume we stored volume & waveBlend in instrument.volume, instrument.waveBlend
		insWave := GenerateWaveFromTextWithParams(
			sampleRate,
			&song,
			&instrument,
		)
		waves = append(waves, insWave)
	}

	// Mix all instrument waves
	mixed := MixWaves(waves...) // your no-clipping version

	return mixed
}

// GenerateWaveFromTextWithParams creates a single wave for one instrument.
// We pass waveBlend & volume from insData to shape the tone and loudness.
func GenerateWaveFromTextWithParams(sampleRate int, song *songData, ins *insData) []float64 {
	beatDuration := time.Minute / time.Duration(song.bpm)
	var finalWave []float64

	for _, noteStr := range strings.Split(ins.data, ",") {
		note, duration := ParseNote(noteStr)
		if note == "" {
			continue
		}
		noteDuration := time.Duration(beatDuration.Seconds() * duration * float64(time.Second))

		// Check for chord
		chordNotes := strings.Split(note, "/")
		if len(chordNotes) > 1 {
			chordWave := PlayChordOfflineWithParams(chordNotes, noteDuration, sampleRate, ins)
			finalWave = append(finalWave, chordWave...)
		} else {
			freq := CalculateFrequency(note)
			noteWave := PlayNoteOfflineWithParams(freq, noteDuration, sampleRate, ins)
			finalWave = append(finalWave, noteWave...)
		}
	}

	return finalWave
}

// PlayNoteOfflineWithParams: generates wave data for a single note, applying volume & wave blend.
func PlayNoteOfflineWithParams(freq float64, duration time.Duration, sampleRate int, ins *insData) []float64 {
	// Handle rest
	if freq == 0 {
		return make([]float64, int(float64(sampleRate)*duration.Seconds()))
	}

	wave := GenerateCustomWave(freq, duration, sampleRate, ins.square)
	wave = ApplyADSR(wave, sampleRate, ins)

	// Apply per-instrument volume
	for i := range wave {
		wave[i] *= float64(ins.volume)
	}

	return wave
}

// GenerateCustomWave blends square & sine waves based on waveBlend (0.0 to 1.0).
func GenerateCustomWave(freq float64, duration time.Duration, sampleRate int, waveBlend float64) []float64 {
	samples := int(float64(sampleRate) * duration.Seconds())
	wave := make([]float64, samples)
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
			wave[i] = float64(mix)
		}
	}
	return wave
}

// PlayChordOfflineWithParams: generates wave data for a chord, applying volume & wave blend.
func PlayChordOfflineWithParams(chord []string, duration time.Duration, sampleRate int, ins *insData) []float64 {
	// Generate wave for each note in the chord
	var waves [][]float64
	for _, note := range chord {
		freq := CalculateFrequency(note)
		noteWave := GenerateCustomWave(freq, duration, sampleRate, ins.square)
		noteWave = ApplyADSR(noteWave, sampleRate, ins)
		// Apply volume
		for i := range noteWave {
			noteWave[i] *= float64(ins.volume)
		}
		waves = append(waves, noteWave)
	}

	// Sum waves
	chordWave := make([]float64, len(waves[0]))
	for i := range chordWave {
		var sum float64
		for _, w := range waves {
			sum += w[i]
		}
		// average to prevent single chord from ballooning amplitude
		chordWave[i] = sum / float64(len(waves))
	}

	return chordWave
}

// PlayNoteOffline generates wave data (without playing) for a single note.
func PlayNoteOffline(freq float64, duration time.Duration, sampleRate int, ins *insData) []float64 {
	if freq == 0 {
		// rest
		return make([]float64, int(float64(sampleRate)*duration.Seconds()))
	}
	wave := GenerateCustomWave(freq, duration, sampleRate, ins.square)
	wave = ApplyADSR(wave, sampleRate, ins)
	return wave
}

// PlayChordOffline generates wave data (without playing) for a chord.
func PlayChordOffline(chord []string, duration time.Duration, sampleRate int, ins *insData) []float64 {
	// Generate wave for each note in the chord
	var waves [][]float64
	for _, note := range chord {
		freq := CalculateFrequency(note)
		noteWave := GenerateCustomWave(freq, duration, sampleRate, ins.square)
		noteWave = ApplyADSR(noteWave, sampleRate, ins)
		waves = append(waves, noteWave)
	}

	// Sum waves
	chordWave := make([]float64, len(waves[0]))
	for i := range chordWave {
		var sum float64
		for _, w := range waves {
			sum += w[i]
		}
		sum /= float64(len(waves)) // average
		chordWave[i] = sum
	}
	return chordWave
}

// MixWaves sums multiple mono wave slices (all same sample rate)
// 1) Averages by number of wave inputs
// 2) Scales further only if needed to prevent clipping
func MixWaves(waves ...[]float64) []float64 {
	// 1. Determine the maximum length among all input waves
	var maxLen int
	for _, w := range waves {
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	// 2. Sum the waves
	mixed := make([]float64, maxLen)
	for _, w := range waves {
		for i := 0; i < len(w); i++ {
			mixed[i] += w[i]
		}
	}

	// 3. Average by number of waves
	numWaves := float64(len(waves))
	if numWaves > 1.0 {
		for i := 0; i < maxLen; i++ {
			mixed[i] /= numWaves
		}
	}

	// 4. Find the peak amplitude (absolute value)
	var maxAmp float64
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

func PlayWave(wave []float64, audioContext *audio.Context, sampleRate int) {

	// 2) Convert float64 samples to raw bytes (16-bit PCM)
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

// MakeSilence inserts a specified pause (in seconds) of silence
// at the start of the wave.
func MakeSilence(sampleRate int, pauseSeconds float64) []float64 {
	// Calculate how many samples correspond to the pause
	numSilenceSamples := int(float64(sampleRate) * pauseSeconds)

	// Create a slice of zeros (silence)
	silence := make([]float64, numSilenceSamples)

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

func ApplyADSR(wave []float64, sampleRate int, ins *insData) []float64 {
	length := len(wave)
	adsrWave := make([]float64, length)

	if ins.attack < 0.01 {
		ins.attack = 0.01
	}
	if ins.decay < 0.01 {
		ins.decay = 0.01
	}
	if ins.sustain < 0.01 {
		ins.sustain = 0.01
	}

	// Calculate the number of samples for each phase
	attackSamples := int(float64(sampleRate) * ins.attack)
	decaySamples := int(float64(sampleRate) * ins.decay)
	releaseSamples := int(float64(sampleRate) * ins.release)
	sustainSamples := length - attackSamples - decaySamples - releaseSamples

	if sustainSamples < 0 {
		sustainSamples = 0
	}

	// Attack
	for i := 0; i < attackSamples && i < length; i++ {
		adsrWave[i] = wave[i] * float64(i) / float64(attackSamples)
	}

	// Decay
	for i := attackSamples; i < attackSamples+decaySamples && i < length; i++ {
		t := float64(i-attackSamples) / float64(decaySamples)
		adsrWave[i] = wave[i] * (1.0 - (1.0-ins.sustain)*t)
	}

	// Sustain
	for i := attackSamples + decaySamples; i < attackSamples+decaySamples+sustainSamples && i < length; i++ {
		adsrWave[i] = wave[i] * ins.sustain
	}

	// Release
	releaseStart := attackSamples + decaySamples + sustainSamples
	for i := releaseStart; i < length; i++ {
		t := float64(i-releaseStart) / float64(releaseSamples)
		adsrWave[i] = wave[i] * ins.sustain * (1.0 - t)
	}

	// ---- Quick fade-out fix to ensure zero at the end ----
	fadeOutDurationSec := 0.01 // 10 ms
	fadeOutSamples := int(float64(sampleRate) * fadeOutDurationSec)
	if fadeOutSamples > length {
		fadeOutSamples = length
	}

	startFade := length - fadeOutSamples
	for i := startFade; i < length; i++ {
		factor := 1.0 - float64(i-startFade)/float64(fadeOutSamples)
		adsrWave[i] *= factor
	}

	return adsrWave
}

func PlayNote(freq float64, duration time.Duration, sampleRate int, audioContext *audio.Context, ins *insData) {
	if freq == 0 {
		time.Sleep(duration) // Handle rests by sleeping
		return
	}

	// Generate a smoother wave
	wave := GenerateCustomWave(freq, duration, sampleRate, ins.square)

	// Apply an ADSR envelope (adjust parameters for smoother transitions)
	wave = ApplyADSR(wave, sampleRate, ins)

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

func PlayChord(chord []string, duration time.Duration, sampleRate int, audioContext *audio.Context, ins *insData) {
	// Generate frequencies for the chord notes
	var waves [][]float64
	for _, note := range chord {
		freq := CalculateFrequency(note)
		wave := GenerateCustomWave(freq, duration, sampleRate, ins.square)
		waves = append(waves, wave)
	}

	// Combine the waves by adding their samples together
	chordWave := make([]float64, len(waves[0]))
	for i := 0; i < len(chordWave); i++ {
		var sampleSum float64
		for _, wave := range waves {
			sampleSum += wave[i]
		}
		chordWave[i] = sampleSum / float64(len(waves)) // Average for multiple notes
	}

	// Apply ADSR envelope
	chordWave = ApplyADSR(chordWave, sampleRate, ins)

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
func PlayTextAsNotes(text string, bpm int, sampleRate int, audioContext *audio.Context, ins *insData) {
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
			PlayChord(chordNotes, time.Duration(beatDuration.Seconds()*duration*float64(time.Second)), sampleRate, audioContext, ins)
		} else {
			// It's a single note, play it
			freq := CalculateFrequency(note)
			PlayNote(freq, time.Duration(beatDuration.Seconds()*duration*float64(time.Second)), sampleRate, audioContext, ins)
		}
	}
}
