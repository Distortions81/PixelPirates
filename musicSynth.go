package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/chewxy/math32"
)

const maxVolume = 0.5

func playMusicPlaylist(g *Game, gameMode int, songList []songData) {
	if *nomusic {
		return
	}
	if len(songList) == 0 {
		return
	}
	for {
		for _, song := range songList {
			if g.gameMode != gameMode {
				return
			}
			startTime := time.Now()
			if *debugMode {
				doLog(true, true, "Rendering: '%v'", song.name)
			}
			output := playSong(song)

			if song.reverb > 0 {
				output = applyReverb(output, song.delay, song.feedback, song.reverb)
			}
			runtime.GC()
			doLog(true, true, "Render took %v\nNow Playing: %v.\n", time.Since(startTime).Round(time.Millisecond), song.name)

			playWave(g, true, output)
		}
		doLog(true, true, "Restarting playlist...")
	}
}

func dumpMusic() {

	os.Mkdir("dump", 0755)

	for _, song := range titleSongList {
		if *debugMode {
			doLog(true, true, "Rendering: '%v'", song.name)
		}
		output := playSong(song)

		if song.reverb > 0 {
			output = applyReverb(output, song.delay, song.feedback, song.reverb)
		}
		saveMono16BitWav("dump/"+song.name+".wav", output)
	}
}

func playSong(song songData) audioData {
	if *nomusic {
		return nil
	}
	var (
		waves    []audioData
		waveLock sync.Mutex
		wg       sync.WaitGroup
	)

	for _, instrument := range song.ins {
		if instrument.volume <= 0 {
			continue
		}

		wg.Add(1)
		go func(ins insData) {
			insWave := generateFromText(
				&song,
				&ins,
			)
			waveLock.Lock()
			waves = append(waves, insWave)
			waveLock.Unlock()
			wg.Done()
		}(instrument)
	}
	wg.Wait()

	return mixWaves(waves...)
}

func generateFromText(song *songData, ins *insData) audioData {
	beatDuration := time.Minute / time.Duration(song.bpm)
	var finalWave audioData

	if *debugMode {
		doLog(true, true, "Rendering: %v", ins.name)
	}
	for _, noteStr := range strings.Split(ins.data, ",") {
		note, duration := parseNote(noteStr)
		if note == "" {
			continue
		}
		noteDuration := time.Duration(beatDuration.Seconds() * duration * float64(time.Second))

		// Check for chord
		chordNotes := strings.Split(note, "/")
		if len(chordNotes) > 1 {
			chordWave := playChord(chordNotes, noteDuration, ins)
			finalWave = append(finalWave, chordWave...)
		} else {
			freq := calculateFrequency(note)
			noteWave := playNote(freq, noteDuration, ins)
			finalWave = append(finalWave, noteWave...)
		}
	}

	return finalWave
}

func playNote(freq float32, duration time.Duration, ins *insData) audioData {
	// Handle rest
	if freq == 0 {
		return make(audioData, int(float64(sampleRate)*duration.Seconds()))
	}

	var wave audioData
	if freq == -1 {
		wave = generateNoise(duration)
	} else {
		wave = generateWave(freq, duration, ins.square)
	}
	wave = applyADSR(wave, ins)

	// Apply per-instrument volume
	for i := range wave {
		wave[i] *= float32(ins.volume)
	}

	return wave
}

func generateNoise(duration time.Duration) audioData {
	numSamples := int(float64(sampleRate) * duration.Seconds())
	wave := make(audioData, numSamples)

	// Generate white noise samples in the range [-1.0, 1.0]
	// and write them as float32 (little-endian)
	for i := 0; i < numSamples; i++ {
		sample := float32((rand.Float64()*2.0 - 1.0)) // in [-1.0, 1.0]
		// Write the float32 sample
		wave[i] = sample
		for x := 0; x < 8; x++ {
			i++
			if i < numSamples {
				wave[i] = sample
			}
		}
	}

	return wave
}

func generateWave(freq float32, duration time.Duration, waveBlend float32) audioData {
	samples := int(float64(sampleRate) * duration.Seconds())
	wave := make(audioData, samples)
	for i := 0; i < samples; i++ {
		t := float32(i) / float32(sampleRate)
		if freq == 0 {
			wave[i] = 0
		} else {
			sinVal := math32.Sin(2 * math32.Pi * freq * t)
			var sqrVal float32 = 1.0
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

func playChord(chord []string, duration time.Duration, ins *insData) audioData {
	// Generate wave for each note in the chord
	var waves []audioData
	for _, note := range chord {
		freq := calculateFrequency(note)
		noteWave := generateWave(freq, duration, ins.square)
		noteWave = applyADSR(noteWave, ins)
		// Apply volume
		for i := range noteWave {
			noteWave[i] *= float32(ins.volume)
		}
		waves = append(waves, noteWave)
	}

	// Sum waves
	chordWave := make(audioData, len(waves[0]))
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

// mixWaves sums multiple mono wave slices (all same sample rate)
// 1) Averages by number of wave inputs
// 2) Scales further only if needed to prevent clipping
func mixWaves(waves ...audioData) audioData {

	// 1. Determine the maximum length among all input waves
	var maxLen int
	for _, w := range waves {
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	// 2. Sum the waves
	mixed := make(audioData, maxLen)
	for _, w := range waves {
		for i := 0; i < len(w); i++ {
			mixed[i] += w[i]
		}
	}

	// 3. Average by number of waves
	/*
		numWaves := float32(len(waves))
		if numWaves > 1.0 {
			for i := 0; i < maxLen; i++ {
				mixed[i] /= numWaves
			}
		}
	*/

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
	if maxAmp > maxVolume {
		scale := maxVolume / maxAmp
		for i := range mixed {
			mixed[i] *= scale
		}
	}

	return mixed
}

func playWave(g *Game, music bool, wave audioData) {

	// 2) Convert float64 samples to raw bytes (16-bit PCM), with noise shaping
	soundData := make([]byte, len(wave)*2)

	// We'll store the quantization error from the previous sample
	var prevError float32

	for i, sample := range wave {
		// Add shaped error from the previous sample.
		// A small feedback factor (like 0.5) is a simple first-order noise shaper.
		shapedSample := sample + 0.5*prevError

		// Hard-clip to -1.0..+1.0 (avoid integer overflow if shapedSample is out of range)
		if shapedSample > 1.0 {
			shapedSample = 1.0
		} else if shapedSample < -1.0 {
			shapedSample = -1.0
		}

		// Convert to 16-bit integer
		intVal := int16(math32.Round(shapedSample * 32767))

		// Store this in the output buffer (little-endian)
		soundData[i*2] = byte(intVal)
		soundData[i*2+1] = byte(intVal >> 8)

		// Calculate the new quantization error:
		// This is the difference between our shapedSample and the quantized integer value.
		actual := float32(intVal) / 32767.0
		prevError = shapedSample - actual
	}

	// 3) Create a player and play
	player := g.audioContext.NewPlayerFromBytes(soundData)
	player.Play()

	for player.IsPlaying() {
		//Handle 'stop music'
		if music && g.stopMusic {
			g.stopMusic = false

			fadeout := 50
			for x := 1; x < fadeout; x++ {
				volume := 1 - float64(x)/float64(fadeout)
				player.SetVolume(volume)
				time.Sleep(time.Millisecond * 10)
			}
			player.Close()
			return
		}
		time.Sleep(time.Millisecond * 10)
	}
}

func calculateFrequency(note string) float32 {
	// Base note A4
	var baseFrequency float32 = 440.0
	// Note names (A-G), standard equal temperament tuning
	noteNames := map[string]int{
		"WN": -2, "NN": -1, "Ab": 0, "A#": 1, "Bb": 2, "Cb": 3, "C#": 4, "Db": 5,
		"D#": 6, "Eb": 7, "Fb": 8, "F#": 9, "Gb": 10, "G#": 11,
	}

	// Note names are of the form "A1", "C#4", etc.
	// First, extract the note (A, B, C, etc.) and the octave number
	var (
		noteName string
		octave   int
	)
	fmt.Sscanf(note, "%2s%d", &noteName, &octave)

	// Find the index of the note (A, A#, B, etc.)
	halfSteps := noteNames[noteName]
	if halfSteps == -1 {
		return 0
	} else if halfSteps == -2 {
		return -1
	}

	// Calculate the number of half-steps from A4 (which is the 49th note)
	halfStepsFromA4 := (octave-6)*12 + halfSteps
	// Frequency of the note
	frequency := baseFrequency * math32.Pow(2, float32(halfStepsFromA4)/12)
	return frequency
}

func applyADSR(wave audioData, ins *insData) audioData {
	length := len(wave)
	adsrWave := make(audioData, length)

	//Prevent clicking
	if ins.attack < 0.01 {
		ins.attack = 0.01
	}
	if ins.decay < 0.01 {
		ins.decay = 0.01
	}
	if ins.sustain < 0.01 {
		ins.sustain = 0.01
	}
	if ins.release < 0.01 {
		ins.release = 0.01
	}

	// Calculate the number of samples for each phase
	attackSamples := int(float32(sampleRate) * ins.attack)
	decaySamples := int(float32(sampleRate) * ins.decay)
	releaseSamples := int(float32(sampleRate) * ins.release)
	sustainSamples := length - attackSamples - decaySamples - releaseSamples

	if sustainSamples < 0 {
		sustainSamples = 0
	}

	// Attack
	for i := 0; i < attackSamples && i < length; i++ {
		adsrWave[i] = wave[i] * float32(i) / float32(attackSamples)
	}

	// Decay
	for i := attackSamples; i < attackSamples+decaySamples && i < length; i++ {
		t := float32(i-attackSamples) / float32(decaySamples)
		adsrWave[i] = wave[i] * (1.0 - (1.0-ins.sustain)*t)
	}

	// Sustain
	for i := attackSamples + decaySamples; i < attackSamples+decaySamples+sustainSamples && i < length; i++ {
		adsrWave[i] = wave[i] * ins.sustain
	}

	// Release
	releaseStart := attackSamples + decaySamples + sustainSamples
	for i := releaseStart; i < length; i++ {
		t := float32(i-releaseStart) / float32(releaseSamples)
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
		factor := 1.0 - float32(i-startFade)/float32(fadeOutSamples)
		adsrWave[i] *= factor
	}

	return adsrWave
}

// parseNote parses the note and its duration from a string like "A4 1".
func parseNote(noteStr string) (string, float64) {
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
