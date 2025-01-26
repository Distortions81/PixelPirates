package main

import (
	"math"
)

// Generate a sine wave sample
func sineWave(freq, t float64) float64 {
	return math.Sin(2 * math.Pi * freq * t)
}

// FM synthesis function
func fmSynthesize(carrierFreq, modulatorFreq, modIndex, duration float64, sampleRate int) []float64 {
	samples := make([]float64, int(duration*float64(sampleRate)))

	for i := 0; i < len(samples); i++ {
		t := float64(i) / float64(sampleRate)
		modulator := sineWave(modulatorFreq, t)
		carrier := sineWave(carrierFreq+modIndex*modulator, t)
		samples[i] = carrier
	}

	return samples
}

func beep() {
	// Example usage
	carrierFreq := 440.0
	modulatorFreq := 10.0
	modIndex := 50.0
	duration := 2.0
	sampleRate := 44100

	samples := fmSynthesize(carrierFreq, modulatorFreq, modIndex, duration, sampleRate)

	if samples != nil {
		//
	}
	// TODO: Output the samples to an audio device or file
}
