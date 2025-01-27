package main

// ApplyReverb takes an input slice of samples, applies a simple delay + feedback effect,
// and returns the processed slice.
//
//   - input:       slice of float64 audio samples
//   - sampleRate:  samples per second (e.g., 44100)
//   - delaySec:    delay time in seconds (e.g., 0.3 for 300ms)
//   - feedback:    how much of the delayed signal is fed back into the effect (0.0 - 1.0)
func ApplyReverb(input audioData, volume, delaySec, feedback float32) audioData {
	// Calculate number of samples corresponding to the delay time
	delaySamples := int(delaySec * float32(sampleRate))
	if delaySamples <= 0 {
		return input
	}

	// Allocate output and a buffer to store delayed samples
	output := make(audioData, len(input))
	delayBuffer := make(audioData, delaySamples)

	// Indices into the delay buffer
	readIndex := 0
	writeIndex := 0

	for i := 0; i < len(input); i++ {
		// Current delayed sample
		delayedSample := delayBuffer[readIndex]

		// Output = dry signal + feedback * delayedSample
		out := input[i] + ((feedback * delayedSample) * volume)
		output[i] = out

		// Write the output into the delay line (for future feedback)
		delayBuffer[writeIndex] = out

		// Move indices forward (wrap around with modulus)
		readIndex = (readIndex + 1) % delaySamples
		writeIndex = (writeIndex + 1) % delaySamples
	}

	return output
}
