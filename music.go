package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/chewxy/math32"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

// Render takes longer, but higher quality output.
// Recommended: 1 (fast, none), 2 (low), 4 (medium), 8 (high), 16 (very high), 32 (extreme)
// https://theproaudiofiles.com/oversampling/

func PlayMusic() {
	const oversampling = 4
	sampleRate := 48000 * oversampling
	audioContext := audio.NewContext(sampleRate / oversampling)

	for {
		for _, song := range songList {
			startTime := time.Now()
			fmt.Printf("Rendering: '%v'\n", song.name)
			output := PlaySong(song, sampleRate)

			if song.reverb > 0 {
				output = ApplyReverb(output, sampleRate, song.delay, song.feedback, song.reverb)
			}
			runtime.GC()
			fmt.Printf("Render took %v\nNow Playing: %v.\n\n", time.Since(startTime).Round(time.Millisecond), song.name)

			PlayWave(output, audioContext, sampleRate, oversampling)
			//SaveMono16BitWav("songs/"+song.name+".wav", sampleRate/oversampling, output)
		}
		fmt.Println("\nRestarting playlist...")
		//return
	}
}

func DumpMusic() {
	const oversampling = 32
	sampleRate := 48000 * oversampling

	os.Mkdir("dump", 0755)

	for _, song := range songList {
		fmt.Printf("Rendering: '%v'\n", song.name)
		output := PlaySong(song, sampleRate)

		if song.reverb > 0 {
			output = ApplyReverb(output, sampleRate, song.delay, song.feedback, song.reverb)
		}
		SaveMono16BitWav("dump/"+song.name+".wav", sampleRate/oversampling, oversampling, output)
	}
}

func PlaySong(song songData, sampleRate int) audioData {
	var waves []audioData
	var waveLock sync.Mutex

	var wg sync.WaitGroup
	for _, instrument := range song.ins {
		if instrument.volume <= 0 {
			continue
		}

		wg.Add(1)
		go func(ins insData) {
			insWave := GenerateFromText(
				sampleRate,
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

	return MixWaves(waves...)
}

func GenerateFromText(sampleRate int, song *songData, ins *insData) audioData {
	beatDuration := time.Minute / time.Duration(song.bpm)
	var finalWave audioData

	fmt.Printf("Rendering: %v\n", ins.name)
	for _, noteStr := range strings.Split(ins.data, ",") {
		note, duration := ParseNote(noteStr)
		if note == "" {
			continue
		}
		noteDuration := time.Duration(beatDuration.Seconds() * duration * float64(time.Second))

		// Check for chord
		chordNotes := strings.Split(note, "/")
		if len(chordNotes) > 1 {
			chordWave := PlayChord(chordNotes, noteDuration, sampleRate, ins)
			finalWave = append(finalWave, chordWave...)
		} else {
			freq := CalculateFrequency(note)
			noteWave := PlayNote(freq, noteDuration, sampleRate, ins)
			finalWave = append(finalWave, noteWave...)
		}
	}

	return finalWave
}

func PlayNote(freq float32, duration time.Duration, sampleRate int, ins *insData) audioData {
	// Handle rest
	if freq == 0 {
		return make(audioData, int(float64(sampleRate)*duration.Seconds()))
	}

	wave := GenerateWave(freq, duration, sampleRate, ins.square)
	wave = ApplyADSR(wave, sampleRate, ins)

	// Apply per-instrument volume
	for i := range wave {
		wave[i] *= float32(ins.volume)
	}

	return wave
}

func GenerateWave(freq float32, duration time.Duration, sampleRate int, waveBlend float32) audioData {
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

func PlayChord(chord []string, duration time.Duration, sampleRate int, ins *insData) audioData {
	// Generate wave for each note in the chord
	var waves []audioData
	for _, note := range chord {
		freq := CalculateFrequency(note)
		noteWave := GenerateWave(freq, duration, sampleRate, ins.square)
		noteWave = ApplyADSR(noteWave, sampleRate, ins)
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

// MixWaves sums multiple mono wave slices (all same sample rate)
// 1) Averages by number of wave inputs
// 2) Scales further only if needed to prevent clipping
func MixWaves(waves ...audioData) audioData {

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
	if maxAmp > 1.0 {
		scale := 1.0 / maxAmp
		for i := range mixed {
			mixed[i] *= scale
		}
	}

	return mixed
}

// DownsampleLinear takes a slice of samples (wave)
// and returns a new slice at rate/oversample the original sample rate
// using simple linear interpolation.
func DownsampleLinear(wave audioData, oversampling int) audioData {
	oldLen := len(wave)
	// If there's not enough data, or nothing to do, just return the original wave.
	if oldLen < 2 {
		return wave
	}

	// New length will be / oversample of oldLen (integer division).
	newLen := oldLen / oversampling
	if newLen < 2 {
		newLen = 2 // ensure at least 2 samples to avoid edge cases
	}

	out := make(audioData, newLen)

	// We want to cover the entire range of the original wave [0..oldLen-1]
	// and map it onto [0..newLen-1] with linear interpolation.
	//
	// Let's map each index i in [0..newLen-1] to a floating-point index in the old wave:
	//   oldIndexF = i * (float64(oldLen - 1) / float64(newLen - 1))
	// This ensures the first new sample aligns with wave[0]
	// and the last new sample aligns exactly with wave[oldLen - 1].
	scale := float32(oldLen-1) / float32(newLen-1)

	for i := 0; i < newLen; i++ {
		oldIndexF := float32(i) * scale
		idx := int(oldIndexF)
		frac := oldIndexF - float32(idx)

		// Edge case: if idx is at the end, just copy the last sample.
		if idx >= oldLen-1 {
			out[i] = wave[oldLen-1]
		} else {
			// Linear interpolation between wave[idx] and wave[idx+1].
			out[i] = wave[idx]*(1.0-frac) + wave[idx+1]*frac
		}
	}

	return out
}

func PlayWave(wave audioData, audioContext *audio.Context, sampleRate, oversampling int) {

	resampled := DownsampleLinear(wave, oversampling)

	// 2) Convert float64 samples to raw bytes (16-bit PCM), with noise shaping
	soundData := make([]byte, len(resampled)*2)

	// We'll store the quantization error from the previous sample
	var prevError float32

	for i, sample := range resampled {
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
	player := audioContext.NewPlayerFromBytes(soundData)
	player.Play()

	// 4) Wait for playback to finish
	duration := time.Duration(
		float64(len(resampled)/2) / float64(sampleRate/oversampling) * float64(time.Second),
	)
	time.Sleep(duration)
}

func CalculateFrequency(note string) float32 {
	// Base note A4
	var baseFrequency float32 = 440.0
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
	frequency := baseFrequency * math32.Pow(2, float32(halfStepsFromA4)/12)
	return frequency
}

func ApplyADSR(wave audioData, sampleRate int, ins *insData) audioData {
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
