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

// ApplyADSR applies an ADSR envelope to the wave to smooth the note transitions.
func ApplyADSR(wave []float32, sampleRate int, attack, decay, sustain, release float64) []float32 {
	length := len(wave)
	adsrWave := make([]float32, length)

	// Calculate the number of samples for each phase
	attackSamples := int(float64(sampleRate) * attack)
	decaySamples := int(float64(sampleRate) * decay)
	releaseSamples := int(float64(sampleRate) * release)
	sustainSamples := length - attackSamples - decaySamples - releaseSamples

	// Ensure sustainSamples is non-negative; if not, adjust to 0 to avoid out-of-bounds
	if sustainSamples < 0 {
		sustainSamples = 0
	}

	// Apply the Attack phase
	for i := 0; i < attackSamples && i < length; i++ {
		adsrWave[i] = wave[i] * float32(float64(i)/float64(attackSamples))
	}

	// Apply the Decay phase
	for i := attackSamples; i < attackSamples+decaySamples && i < length; i++ {
		adsrWave[i] = wave[i] * float32(1-(1-sustain)*(float64(i-attackSamples)/float64(decaySamples)))
	}

	// Apply the Sustain phase
	for i := attackSamples + decaySamples; i < attackSamples+decaySamples+sustainSamples && i < length; i++ {
		adsrWave[i] = wave[i] * float32(sustain)
	}

	// Apply the Release phase
	for i := attackSamples + decaySamples + sustainSamples; i < length; i++ {
		// Prevent any out-of-bounds access by capping index properly
		adsrWave[i] = wave[i] * float32(sustain*(1-float64(i-(attackSamples+decaySamples+sustainSamples))/float64(releaseSamples)))
	}

	return adsrWave
}

func GenerateSmoothWave(freq float64, duration time.Duration, sampleRate int) []float32 {
	samples := int(float64(sampleRate) * duration.Seconds())
	wave := make([]float32, samples)
	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)

		if freq == 0 {
			wave[i] = 0
		} else {
			// Blend square and sine waves for a smoother sound
			squareWave := 1.0
			if math.Sin(2*math.Pi*freq*t) < 0 {
				squareWave = -1.0
			}
			sineWave := math.Sin(2 * math.Pi * freq * t)
			wave[i] = float32(0.5*squareWave + 0.5*sineWave) // 50% square, 50% sine
		}
	}
	return wave
}

func PlayNote(freq float64, duration time.Duration, sampleRate int, audioContext *audio.Context) {
	if freq == 0 {
		time.Sleep(duration) // Handle rests by sleeping
		return
	}

	// Generate a smoother wave
	wave := GenerateSmoothWave(freq, duration, sampleRate)

	// Apply an ADSR envelope (adjust parameters for smoother transitions)
	wave = ApplyADSR(wave, sampleRate, 0.05, 0.1, 0.5, 0.5)

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

	for {
		time.Sleep(time.Millisecond * 500)

		// Sea shanty melody with variable note lengths (e.g., "A4 1" for whole notes)
		text := `Ab4 1, Bb4 1, Db5 0.5, Eb5 0.5, Gb5 1, Fb5 1, Ab5 1, Bb5 0.5, Fb5 0.5, Db5 1,
Cb5 1, Gb4 0.5, Ab4 0.5, Bb4 1, Db5 1, Fb5 0.5, Gb5 0.5, A#5 1, Bb5 1,

Ab5 1, Gb5 1, Fb5 0.75, Db5 0.25, Bb4 0.5, Cb5 0.5, Gb5 1, Db5 1, A#5 0.75, Bb5 0.25,
Cb5 1, D#5 1, Fb5 1, Gb5 0.5, Ab5 0.5, Bb5 1, Cb5 1,

Db5 1, Eb5 0.5, Fb5 0.5, Gb5 1, Fb5 1, Cb5 1, Ab4 1, A#4 0.5, Bb4 0.5, Db5 1,
Ab4 1, Gb4 0.75, A#4 0.25, Cb5 1, D#5 0.5, Fb5 0.5, Gb5 1, Bb5 1,

Ab5 1, Gb5 1, Fb5 0.75, Db5 0.25, Bb4 0.5, Cb5 0.5, Gb5 1, Db5 1, A#5 0.75, Bb5 0.25,
Db5 1, Fb5 1, Ab5 1, Gb5 0.5, Fb5 0.5, Db5 1, Cb5 1,

// Key Change - Higher Energy
A#5 1, Bb5 1, Db6 0.5, Fb6 0.5, Gb6 1, Fb6 1, Ab6 1, Bb6 0.5, Db6 0.5, Fb6 1,
Bb5 1, Ab5 0.5, Gb5 0.5, Db6 1, Fb6 0.5, Gb6 0.5, A#6 1, Bb6 1,

Ab6 1, Gb6 1, Fb6 0.75, Db6 0.25, Bb5 0.5, Cb6 0.5, Gb6 1, Db6 1, A#6 0.75, Bb6 0.25,
Cb6 1, D#6 1, Fb6 1, Gb6 0.5, Ab6 0.5, Bb6 1, Cb6 1,

Db6 1, Eb6 0.5, Fb6 0.5, Gb6 1, Fb6 1, Cb6 1, Ab5 1, A#5 0.5, Bb5 0.5, Db6 1,
Ab5 1, Gb5 0.75, A#5 0.25, Cb6 1, D#6 0.5, Fb6 0.5, Gb6 1, Bb6 1,

// Final Chorus (Dramatic Finale)
Ab6 1, Gb6 1, Fb6 0.75, Db6 0.25, Bb5 0.5, Cb6 0.5, Gb6 1, Db6 1, A#6 0.75, Bb6 0.25,
Db6 1, Fb6 1, Ab6 1, Gb6 0.5, Fb6 0.5, Db6 1, Cb6 1`

		bpm := 120          // 120 beats per minute
		sampleRate := 44100 // Standard audio sample rate

		// Play the sea shanty notes with variable lengths
		fmt.Println("Playing Sea Shanty with Variable Length Notes...")
		PlayTextAsNotes(text, bpm, sampleRate, audioContext)

		time.Sleep(time.Second * 5)
	}
}
