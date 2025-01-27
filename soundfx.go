package main

import (
	"math"
	"math/rand/v2"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

func PlayFx() {
	sampleRate := 48000
	audioContext := audio.NewContext(sampleRate)

	const repeatTime = 1
	const nextTime = 3
	const fxTime = 1

	for {
		for x := 0; x < 3; x++ {
			PlayWave(CannonSound(sampleRate, 3), audioContext, sampleRate, fxTime)
			time.Sleep(repeatTime * time.Second)
		}
		time.Sleep(time.Second * nextTime)
	}

	for {
		for x := 0; x < 3; x++ {
			PlayWave(SeagullSound(sampleRate, 5), audioContext, sampleRate, fxTime)
			time.Sleep(repeatTime * time.Second)
		}
		time.Sleep(time.Second * nextTime)

		for x := 0; x < 3; x++ {
			PlayWave(CannonSound(sampleRate, 1), audioContext, sampleRate, fxTime)
			time.Sleep(repeatTime * time.Second)
		}
		time.Sleep(time.Second * nextTime)

		for x := 0; x < 3; x++ {
			PlayWave(GoldCoinsSound(sampleRate, 1), audioContext, sampleRate, fxTime)
			time.Sleep(repeatTime * time.Second)
		}
		time.Sleep(time.Second * nextTime)
	}
}

// ------------------------------------------------------------------------------
// 1) Seagull Sound
// ------------------------------------------------------------------------------
func SeagullSound(sampleRate int, durationSec float64) []float32 {
	numSamples := int(durationSec * float64(sampleRate))
	out := make([]float32, numSamples)

	// We'll produce 3 short calls, each ~0.3s, with small gaps in between
	call := 0.2
	gap := 0.1
	numCalls := 3

	position := 0

	var direction bool = rand.Float64() > 0.5

	shift := rand.Float64() * 50
	rise := rand.Float64() * 20
	for c := 0; c < numCalls; c++ {
		callDuration := call + (rand.Float64())
		gapDuration := gap + (rand.Float64() / 2)

		shift += rise
		startSample := position
		endSample := startSample + int(callDuration*float64(sampleRate))

		for i := startSample; i < endSample && i < numSamples; i++ {
			t := float64(i-startSample) / float64(sampleRate)

			//Sweep
			start := 600.0 + shift
			var increase float64
			if direction {
				increase = start - 20.0
			} else {
				increase = start + 20.0
			}
			fr := start + (increase-start)*(t/callDuration)

			//timbre
			timbre := 0.5 * math.Sin(2.0*math.Pi*20.0*t)

			// Final frequency
			f := fr

			// Sine wave
			sample := math.Sin(2.0*math.Pi*f*t) - timbre

			// Amplitude envelope (quick attack, quick decay)
			ampEnv := seagullEnv(t, callDuration)

			out[i] = float32(sample * ampEnv)
		}

		position += int((callDuration + gapDuration) * float64(sampleRate))
		if position >= numSamples {
			break
		}
	}

	return out
}

// A simple amplitude envelope for the "gull call"
func seagullEnv(t, total float64) float64 {
	attack := 0.05
	decay := total - attack

	if t < attack {
		// ramp up
		return t / attack
	} else if t < total {
		// ramp down
		return 1.0 - (t-attack)/decay
	}
	return 0
}

// ------------------------------------------------------------------------------
// 3) Cannon Sound
// ------------------------------------------------------------------------------
func CannonSound(sampleRate int, durationSec float64) []float32 {
	numSamples := int(durationSec * float64(sampleRate))
	out := make([]float32, numSamples)

	noiseBurstLen := 0.2
	sineLen := 6.0
	freq := 12.0

	// 1) Noise burst
	noiseEnd := int(noiseBurstLen * float64(sampleRate))
	for i := 0; i < noiseEnd && i < numSamples; i++ {
		val := 4.0*rand.Float64() - 1.0
		env := 1.0 - float64(i)/float64(noiseEnd)
		out[i] = float32(val * env)
	}

	// 2) Low frequency boom
	sineStart := 0
	sineEnd := int(sineLen * float64(sampleRate)) // up to 1s
	for i := sineStart; i < sineEnd && i < numSamples; i++ {

		t := float64(i) / float64(sampleRate)
		cannonFreq := freq + (rand.Float64() / 3)
		sample := math.Sin(2.0 * math.Pi * cannonFreq * t)
		var sqrVal float64 = 1.0
		if sample < 0 {
			sqrVal = -1.0
		}
		decay := math.Exp(-3.0 * t)
		waveBlend := 0.5
		mix := waveBlend*sqrVal + (1.0-waveBlend)*sample
		out[i] += float32(mix * decay * 2.0)
	}

	return out
}

// ------------------------------------------------------------------------------
// 4) Gold Coins Sound
// ------------------------------------------------------------------------------
func GoldCoinsSound(sampleRate int, durationSec float64) []float32 {
	numSamples := int(durationSec * float64(sampleRate))
	out := make([]float32, numSamples)

	// We'll simulate a few "clinks" at random times.
	// Each clink is a combination of a few short ringing frequencies
	// with exponential decay. This is extremely simplified,
	// but can give a vaguely metallic feel.

	numClinks := 4

	for c := 0; c < numClinks; c++ {
		// Random start time for each clink
		startTime := rand.Float64() * (durationSec - 0.2)
		startSample := int(startTime * float64(sampleRate))

		// Each clink ~200 ms long
		clinkLen := 0.2
		endSample := startSample + int(clinkLen*float64(sampleRate))
		if endSample > numSamples {
			endSample = numSamples
		}

		// We’ll pick 3-4 partial frequencies in a metallic range
		partials := []float64{
			1000 + 3000*rand.Float64(),
			1000 + 3000*rand.Float64(),
			1000 + 3000*rand.Float64(),
		}
		// Optionally add a 4th partial
		if rand.Float64() < 0.5 {
			partials = append(partials, 1000+3000*rand.Float64())
		}

		for i := startSample; i < endSample; i++ {
			t := float64(i-startSample) / float64(sampleRate)

			var sampleVal float64
			for _, f := range partials {
				sampleVal += math.Sin(2.0 * math.Pi * f * t)
			}

			// Exponential decay
			decay := math.Exp(-12.0 * t) // short ring
			sampleVal *= decay

			// Average across partials
			sampleVal /= float64(len(partials))

			out[i] += float32(sampleVal * 0.8) // scale amplitude
		}
	}

	return out
}
