package main

import (
	"encoding/binary"
	"math"
	"math/rand/v2"
	"time"

	"github.com/chewxy/math32"
)

const (
	noiseSmoothing = 7
	maxVolume      = 0.5
)

func generateNoise(duration time.Duration) audioData {
	numSamples := int(float64(sampleRate) * duration.Seconds())
	wave := make(audioData, numSamples)

	for i := 0; i < numSamples; i++ {
		sample := float32(rand.Float32()*2.0 - 1.0)
		wave[i] = sample
		// Repeat the sample for smoothing.
		for x := 0; x < noiseSmoothing; x++ {
			i++
			if i < numSamples {
				wave[i] = sample
			}
		}
	}
	return wave
}

// generateWave accepts a waveform type and blend factor.
func generateWave(freq float64, duration time.Duration, waveform int) audioData {
	samples := int(float64(sampleRate) * duration.Seconds())
	wave := make(audioData, samples)
	period := 1.0 / float64(freq)

	var dbOffset float64
	//We do some RMS volume correction
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

	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)
		sineVal := math.Sin(2 * math.Pi * freq * t)

		var squareVal float64
		if sineVal < 0 {
			squareVal = -1.0
		} else {
			squareVal = 1.0
		}

		tmod := float64(math.Mod(float64(t), period))
		triangleVal := 4*math.Abs(tmod/float64(period)-0.5) - 1.0
		sawtoothVal := 2*float64(math.Mod(float64(t), period))/float64(period) - 1.0

		var sample float64
		switch waveform {
		case WAVE_SINE:
			sample = sineVal
		case WAVE_SQUARE:
			sample = squareVal
		case WAVE_TRIANGLE:
			sample = triangleVal
		case WAVE_SAW:
			sample = sawtoothVal
		default:
			sample = sineVal
		}
		wave[i] = float32(sample * compensation)
	}
	return wave
}

func playWave(g *Game, music bool, wave audioData, fast bool) {
	soundData := make([]byte, len(wave)*2)
	var prevError float32

	if fast {
		for i, s := range wave {
			if s > 1.0 {
				s = 1.0
			} else if s < -1.0 {
				s = -1.0
			}
			intSample := int16(math.Round(float64(s) * 32767.0))
			binary.LittleEndian.PutUint16(soundData[i*2:], uint16(intSample))
		}
	} else {
		for i, sample := range wave {
			shapedSample := sample + 0.5*prevError
			if shapedSample > 1.0 {
				shapedSample = 1.0
			} else if shapedSample < -1.0 {
				shapedSample = -1.0
			}
			intVal := int16(math32.Round(shapedSample * 32767))
			soundData[i*2] = byte(intVal)
			soundData[i*2+1] = byte(intVal >> 8)
			actual := float32(intVal) / 32767.0
			prevError = shapedSample - actual
		}
	}
	player := g.audioContext.NewPlayerFromBytes(soundData)
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
