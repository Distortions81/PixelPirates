package main

import (
	"encoding/binary"
	"math"
	"math/rand/v2"
	"time"
)

const (
	noiseSmoothing = 7
	maxVolume      = 0.5

	preRenderNoiseSamples = sampleRate * 6
)

var noiseSamples audioData

func init() {
	noiseSamples = make(audioData, preRenderNoiseSamples)
	for i := 0; i < preRenderNoiseSamples; i++ {
		sample := float32(rand.Float32()*2.0 - 1.0)
		noiseSamples[i] = sample
		// Repeat the sample for smoothing.
		for x := 0; x < noiseSmoothing; x++ {
			i++
			if i < preRenderNoiseSamples {
				noiseSamples[i] = sample
			}
		}
	}
}

func generateNoise(duration time.Duration) audioData {
	// Calculate the number of samples needed for the duration.
	numSamples := int(float64(sampleRate) * duration.Seconds())

	// Create a slice to hold the output noise.
	output := make(audioData, numSamples)

	// Pick a random starting offset in the pre-rendered noise slice.
	offset := rand.IntN(len(noiseSamples))

	// Fill the output slice using modulo arithmetic to loop through noiseSamples.
	for i := 0; i < numSamples; i++ {
		output[i] = noiseSamples[(offset+i)%len(noiseSamples)]
	}

	return output
}

// generateWave accepts a frequency, duration, and waveform type and returns the corresponding audio data.
// It pre-renders one period of the waveform and then re-uses that data.
func generateWave(freq float64, duration time.Duration, waveform int) audioData {
	// Total number of samples for the given duration.
	totalSamples := int(float64(sampleRate) * duration.Seconds())

	// Calculate the period in seconds and the number of samples in one period.
	period := 1.0 / freq
	periodSamples := int(math.Round(period * float64(sampleRate)))
	if periodSamples <= 0 {
		periodSamples = 1
	}

	// RMS volume correction offsets for different waveforms.
	var dbOffset float64
	switch waveform {
	case WAVE_SINE:
		dbOffset = 0
	case WAVE_SQUARE:
		dbOffset = -3.0
	case WAVE_TRIANGLE:
		dbOffset = -1.76
	case WAVE_SAW:
		dbOffset = -1.75
	default:
		dbOffset = 0
	}
	compensation := math.Pow(10, dbOffset/20)

	// Pre-render one period of the waveform.
	onePeriod := make([]float64, periodSamples)
	for j := 0; j < periodSamples; j++ {
		// Compute the normalized phase [0, 1) for the sample.
		phase := float64(j) / float64(periodSamples)

		// Pre-calculate the sine value (one cycle) for use in different waveforms.
		sineVal := math.Sin(2 * math.Pi * phase)

		var sample float64
		switch waveform {
		case WAVE_SINE:
			sample = sineVal
		case WAVE_SQUARE:
			// A square wave: 1 for positive half cycle, -1 for negative.
			if sineVal < 0 {
				sample = -1.0
			} else {
				sample = 1.0
			}
		case WAVE_TRIANGLE:
			// A triangle wave: linear ramp up then down.
			// Map phase from [0,1) to [-1, 1] in a triangle shape.
			sample = 4*math.Abs(phase-0.5) - 1.0
		case WAVE_SAW:
			// A sawtooth wave: a linear ramp from -1 to 1.
			sample = 2*phase - 1.0
		default:
			// Default to sine wave.
			sample = sineVal
		}

		// Apply the compensation factor.
		onePeriod[j] = sample * compensation
	}

	// Allocate the output buffer and fill it by looping over the pre-rendered period.
	output := make(audioData, totalSamples)
	for i := 0; i < totalSamples; i++ {
		// Use modulo to repeat the one period cycle.
		output[i] = float32(onePeriod[i%periodSamples])
	}

	return output
}

func playWave(g *Game, music bool, wave audioData) {
	soundData := make([]byte, len(wave)*4)

	for i, s := range wave {
		binary.LittleEndian.PutUint32(soundData[i*4:], math.Float32bits(s))
	}
	player := g.audioContext.NewPlayerF32FromBytes(soundData)
	player.SetBufferSize(time.Second * 10)
	player.Play()
}

// / applyADSR applies an ADSR envelope to the input waveform 'wave'
// and guarantees that the returned output is exactly the same length as 'wave'.
// If the sum of attack, decay, and release (in samples) is less than the total length,
// the remaining samples are assigned to sustain.
// If the sum is greater than the total length, the attack, decay, and release are scaled down proportionally.
func applyADSR(wave audioData, ins *insData, volume float32) audioData {
	totalSamples := len(wave)

	// Calculate the nominal phase lengths.
	nomAttackSamples := int(float64(sampleRate) * ins.attack)
	nomDecaySamples := int(float64(sampleRate) * ins.decay)
	nomReleaseSamples := int(float64(sampleRate) * ins.release)

	// The sum of the non-sustain phases.
	phaseSum := nomAttackSamples + nomDecaySamples + nomReleaseSamples

	var attackSamples, decaySamples, releaseSamples, sustainSamples int

	if phaseSum > totalSamples {
		// Not enough samples for full phases; scale them down proportionally.
		scale := float32(totalSamples) / float32(phaseSum)
		attackSamples = int(float32(nomAttackSamples) * scale)
		decaySamples = int(float32(nomDecaySamples) * scale)
		releaseSamples = int(float32(nomReleaseSamples) * scale)
		sustainSamples = 0 // No room for sustain.
	} else {
		attackSamples = nomAttackSamples
		decaySamples = nomDecaySamples
		releaseSamples = nomReleaseSamples
		// All remaining samples become sustain.
		sustainSamples = totalSamples - phaseSum
	}

	// Create the envelope array.
	envelope := make([]float32, totalSamples)
	index := 0

	// Attack phase: ramp linearly from 0 to 1.
	for i := 0; i < attackSamples && index < totalSamples; i++ {
		envelope[index] = float32(i) / float32(attackSamples)
		index++
	}

	// Decay phase: ramp from 1 down to ins.sustain.
	for i := 0; i < decaySamples && index < totalSamples; i++ {
		t := float32(i) / float32(decaySamples)
		envelope[index] = 1 - float32(1-ins.sustain)*t
		index++
	}

	// Sustain phase: constant at ins.sustain.
	for i := 0; i < sustainSamples && index < totalSamples; i++ {
		envelope[index] = float32(ins.sustain)
		index++
	}

	// Release phase: ramp from ins.sustain down to 0.
	for i := 0; i < releaseSamples && index < totalSamples; i++ {
		t := float64(i) / float64(releaseSamples)
		if t > 1 {
			t = 1
		}
		envelope[index] = float32(ins.sustain * (1 - t))
		index++
	}

	// If for any reason we have not filled all samples (due to rounding), fill the remainder with 0.
	for index < totalSamples {
		envelope[index] = 0
		index++
	}

	// Apply the envelope to each sample.
	output := make(audioData, totalSamples)
	for i := 0; i < totalSamples; i++ {
		output[i] = wave[i] * envelope[i]
		output[i] *= volume
	}

	return output
}

// mixWaves using the maximum length.
func mixWaves(numNotes, maxLen int, waves ...audioData) audioData {

	mixed := make(audioData, maxLen)
	for _, w := range waves {
		for i := 0; i < maxLen; i++ {
			if i < len(w) {
				mixed[i] += w[i]
			}
		}
	}
	if numNotes > 1 {
		for i := 0; i < maxLen; i++ {
			mixed[i] /= float32(numNotes)
		}
	}
	return mixed
}
