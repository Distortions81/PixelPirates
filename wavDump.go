package main

import (
	"fmt"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

// saveMono16BitWav uses go-audio/wav to write a single-channel 16-bit WAV.
func saveMono16BitWav(filename string, samples []float32) error {

	// 1. Create the output file
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	// 2. Create a new WAV encoder
	enc := wav.NewEncoder(
		f,
		sampleRate, // sample rate
		16,         // bits per sample
		2,          // number of channels (mono)
		1,          // WAV format (1 = PCM)
	)

	// 3. Convert your float32 samples to 16-bit integer samples
	//    The go-audio library uses an `audio.IntBuffer` internally.
	buf := &audio.IntBuffer{
		Format: &audio.Format{
			NumChannels: 1,
			SampleRate:  sampleRate,
		},
		Data:           make([]int, len(samples)),
		SourceBitDepth: 16,
	}

	for i, s := range samples {
		// Clamp to [-1.0, 1.0]
		if s > 1 {
			s = 1
		} else if s < -1 {
			s = -1
		}
		// Convert float32 to int16, then to 'int'
		buf.Data[i] = int(int16(s * 32767))
	}

	// 4. Write the buffer to the WAV encoder
	if err := enc.Write(buf); err != nil {
		return fmt.Errorf("failed to write samples: %w", err)
	}

	// 5. Close the encoder to finalize the WAV file
	if err := enc.Close(); err != nil {
		return fmt.Errorf("failed to close encoder: %w", err)
	}

	doLog(true, false, "Wrote %v\n", filename)

	return nil
}
