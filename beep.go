package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

// CalculateFrequency calculates the frequency of a note based on its name and octave.
func CalculateFrequency(note string) float64 {
	// Base note A4
	baseFrequency := 440.0
	// Note names (A-G), standard equal temperament tuning
	noteNames := map[string]int{
		"A": 0, "A#": 1, "B": 2, "C": 3, "C#": 4, "D": 5,
		"D#": 6, "E": 7, "F": 8, "F#": 9, "G": 10, "G#": 11,
	}

	// Note names are of the form "A1", "C#4", etc.
	// First, extract the note (A, B, C, etc.) and the octave number
	var noteName string
	var octave int
	fmt.Sscanf(note, "%1s%d", &noteName, &octave)

	// Find the index of the note (A, A#, B, etc.)
	halfSteps := noteNames[noteName]
	// Calculate the number of half-steps from A4 (which is the 49th note)
	halfStepsFromA4 := (octave-4)*12 + halfSteps
	// Frequency of the note
	frequency := baseFrequency * math.Pow(2, float64(halfStepsFromA4)/12)
	return frequency
}

// GenerateSquareWave generates a square wave for a given frequency, duration, and sample rate.
func GenerateSquareWave(freq float64, duration time.Duration, sampleRate int) []float32 {
	samples := int(float64(sampleRate) * duration.Seconds())
	wave := make([]float32, samples)
	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)
		// Generate a square wave by alternating between 1 and -1
		if math.Sin(2*math.Pi*freq*t) >= 0 {
			wave[i] = 1
		} else {
			wave[i] = -1
		}
	}
	return wave
}

// PlayNote generates and plays a note as a square wave.
func PlayNote(freq float64, duration time.Duration, sampleRate int, audioContext *audio.Context) {
	if freq == 0 {
		time.Sleep(duration) // Handle rests by sleeping
		return
	}

	fmt.Printf("%v : %v\n", freq, duration)
	wave := GenerateSquareWave(freq, duration, sampleRate)

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

	// Wait for the note to finish playing
	time.Sleep(duration)
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

		// Calculate the frequency for the note directly
		freq := CalculateFrequency(note)

		// Adjust the duration based on the beat duration
		durationInSeconds := beatDuration.Seconds() * duration

		// Play the note (or rest) for the calculated duration
		PlayNote(freq, time.Duration(durationInSeconds*float64(time.Second)), sampleRate, audioContext)
	}
}

// Main function to set up Ebiten and audio
func playMusic() {
	// Initialize Ebiten's audio system
	audioContext := audio.NewContext(44100)

	// Sea shanty melody with variable note lengths (e.g., "A4 1" for whole notes)
	text := "A4 1,B4 0.5,C4 0.25,D4 0.5,E4 1,F4 0.5,G4 0.25,A4 1 " +
		"A4 1,B4 0.5,C4 0.25,D4 0.5,E4 1,F4 0.5,G4 0.25,A4 1"

	bpm := 120          // 120 beats per minute
	sampleRate := 44100 // Standard audio sample rate

	// Play the sea shanty notes with variable lengths
	fmt.Println("Playing Sea Shanty with Variable Length Notes...")
	PlayTextAsNotes(text, bpm, sampleRate, audioContext)

}
