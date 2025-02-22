package main

import (
	"fmt"
	"strings"
)

func init() {
	count := 0
	for _, playlist := range gameModePlaylists {
		for s, _ := range playlist {
			parseSong(&playlist[s])
			count++
		}
	}
}

var (
	gameModePlaylists = [GAME_MAX]playlistData{
		GAME_BOOT:   nil,
		GAME_START:  nil,
		GAME_TITLE:  titleSongList,
		GAME_PLAY:   gameSongList,
		GAME_ISLAND: islandSongList,
	}

	titleSongList = []songData{
		twilightReflections,
	}

	gameSongList = []songData{
		infinitoRealms,
		epicWarOfTheAncients,
		spectersOfAshenTwilight,
	}

	islandSongList = []songData{
		voyageOfTheAbyss,
	}
)

var chordProgression2 = []string{

	"E2/G2/B2",
	"E2/G2/B2",
	"D2/F#2/A2",
	"D2/F#2/A2",
	"F2/A2/C3",
	"E2/G2/B2",
	"G2/B2/D3",
	"E2/G2/B2",
	"E2/G2/B2",
	"B2/D#3/F#3",
	"C2/E2/G2",
	"B2/D#3/F#3",
	"E2/G2/B2",
	"E2/G2/B2",
	"A2/C3/E3",
	"B2/D#3/F#3",
	"E2/G2/B2",
	"E2/G2/B2",

	"E2/G2/B2",
	"G2/B2/D3",
	"C2/E2/G2",
	"D2/F#2/A2",
	"F2/A2/C3",
	"E2/G2/B2",
	"B2/D#3/F#3",
	"B2/D#3/F#3",
	"F2/A2/C3",
	"G2/B2/D3",
	"A2/C3/E3",
	"B2/D#3/F#3",
	"E2/G2/B2",
	"D2/F#2/A2",
	"C2/E2/G2",
	"F2/A2/C3",
	"E2/G2/B2",
	"E2/G2/B2",

	"E2/G2/B2",
	"D2/F#2/A2",
	"D2/F#2/A2",
	"F2/A2/C3",
	"E2/G2/B2",
	"E2/G2/B2",
	"B2/D#3/F#3",
	"B2/D#3/F#3",
	"C2/E2/G2",
	"A2/C3/E3",
	"D2/F#2/A2",
	"G2/B2/D3",
	"F2/A2/C3",
	"E2/G2/B2",
	"D2/F#2/A2",
	"C2/E2/G2",
	"B2/D#3/F#3",
	"E2/G2/B2",

	"F2/A2/C3",
	"F2/A2/C3",
	"A2/C3/E3",
	"B2/D#3/F#3",
	"E2/G2/B2",
	"E2/G2/B2",
	"G2/B2/D3",
	"F2/A2/C3",
	"F2/A2/C3",
	"G2/B2/D3",
	"B2/D#3/F#3",
	"A2/C3/E3",
	"E2/G2/B2",
	"D2/F#2/A2",
	"C2/E2/G2",
	"B2/D#3/F#3",
	"E2/G2/B2",
	"E2/G2/B2",
}

// Helper: writes <chord> <beats>, e.g. "E2/G2/B2 4,"
func chordToMeasure2(chord string, beats float64) string {
	return chord + fmt.Sprintf(" %.2f, ", beats)
}

// Increments note's octave, e.g. "E2" -> "E3".
func incrementOctave2(note string) string {
	if len(note) < 2 {
		return note
	}
	lastChar := note[len(note)-1]
	if lastChar < '0' || lastChar > '9' {
		return note
	}
	octave := int(lastChar - '0')
	octaveUp := octave + 1
	return note[:len(note)-1] + fmt.Sprintf("%d", octaveUp)
}

// A handful of arpeggio patterns, referencing chord notes by index [0,1,2], plus [3]=octave-lift root.
var arpPatterns2 = [][]int{
	{0, 1, 2, 3},
	{3, 2, 1, 0},
	{0, 2, 1, 3},
	{1, 0, 3, 2},
	{0, 1, 3, 2},
	{2, 3, 1, 0},
	{1, 2, 0, 3},
	{3, 0, 2, 1},
}

// Our new “haunting” track definition:
var spectersOfAshenTwilight = songData{
	name:     "Specters of Ashen Twilight - A Haunting Saga",
	bpm:      80,
	reverb:   0.50,
	delay:    0.2,
	feedback: 0.5,

	ins: []insData{

		{
			name:     "GhostPad",
			volume:   0.50,
			waveform: WAVE_SAW,
			data: func() string {
				var s strings.Builder
				for _, chord := range chordProgression2 {
					s.WriteString(chordToMeasure2(chord, 4))
				}
				return s.String()
			}(),
			attack:  0.15,
			decay:   0.2,
			sustain: 0.8,
			release: 0.4,
		},

		{
			name:     "Choir",
			volume:   0.45,
			waveform: WAVE_SQUARE,
			data: func() string {
				var s strings.Builder

				for _, chord := range chordProgression2 {
					s.WriteString(chord + " 2, NN 2, ")
				}
				return s.String()
			}(),
			attack:  0.1,
			decay:   0.15,
			sustain: 0.7,
			release: 0.4,
		},

		{
			name:     "Bass",
			volume:   0.48,
			waveform: WAVE_SINE,
			data: func() string {
				var s strings.Builder
				for _, chord := range chordProgression2 {
					parts := strings.Split(chord, "/")
					root := parts[0]
					s.WriteString(root + " 4, ")
				}
				return s.String()
			}(),
			attack:  0.02,
			decay:   0.1,
			sustain: 0.65,
			release: 0.3,
		},

		{
			name:     "Arp",
			volume:   0.33,
			waveform: WAVE_TRIANGLE,
			data: func() string {
				var s strings.Builder
				for measureIndex, chord := range chordProgression2 {
					notes := strings.Split(chord, "/")

					if len(notes) < 3 {
						notes = []string{"E2", "G2", "B2"}
					}

					patIndex := measureIndex % len(arpPatterns2)
					pattern := arpPatterns2[patIndex]

					chordNotes := []string{
						notes[0],
						notes[1],
						notes[2],
						incrementOctave2(notes[0]),
					}

					for r := 0; r < 2; r++ {
						for _, idx := range pattern {
							s.WriteString(chordNotes[idx] + " 0.5, ")
						}
					}
				}
				return s.String()
			}(),
			attack:  0.03,
			decay:   0.07,
			sustain: 0.8,
			release: 0.15,
		},

		{
			name:     "DarkBell",
			volume:   0.25,
			waveform: WAVE_SINE,
			data: func() string {
				var s strings.Builder

				for i := 0; i < 72; i++ {
					if i%2 == 0 {

						s.WriteString("NN 1.5, A4 0.25, NN 0.25, NN 1, D5 0.25, NN 0.75, ")
					} else {

						s.WriteString("NN 2, E5 0.25, NN 1.75, ")
					}
				}
				return s.String()
			}(),
			attack:  0.01,
			decay:   0.3,
			sustain: 0.0,
			release: 0.5,
		},

		{
			name:     "Lead",
			volume:   0.35,
			waveform: WAVE_SAW,
			data: func() string {

				phrase1 := []string{
					"E4 2, G4 2, ",
					"B3 1, C4 1, A3 1, G3 1, ",
					"F3 2, E4 2, ",
					"E3 2, NN 2, ",
				}
				phrase2 := make([]string, len(phrase1))
				copy(phrase2, phrase1)

				phrase2[1] = "D4 1, C4 1, A3 1, G3 1, "

				phrase3 := make([]string, len(phrase1))
				copy(phrase3, phrase1)
				phrase3[2] = "F4 1, E4 1, D4 1, C4 1, "
				phrase3[3] = "NN 2, G4 2, "

				// We'll just cycle phrase1, phrase2, phrase3 across 72 measures => 18 cycles x 4 measures = 72
				var s strings.Builder
				allPhrases := [][]string{phrase1, phrase2, phrase3}
				phraseIndex := 0
				for i := 0; i < 18; i++ {
					ph := allPhrases[phraseIndex]
					for _, measureData := range ph {
						s.WriteString(measureData)
					}
					phraseIndex = (phraseIndex + 1) % len(allPhrases)
				}
				return s.String()
			}(),
			attack:  0.08,
			decay:   0.1,
			sustain: 0.7,
			release: 0.4,
		},

		{
			name:     "Kick",
			volume:   0.45,
			waveform: WAVE_SINE,
			data: func() string {
				var s strings.Builder

				for i := 0; i < 72; i++ {

					s.WriteString("C1 0.5, NN 0.5, ")

					s.WriteString("NN 1, ")

					s.WriteString("C1 0.5, NN 0.5, ")

					s.WriteString("NN 1, ")
				}
				return s.String()
			}(),
			attack:  0.01,
			decay:   0.07,
			sustain: 0.2,
			release: 0.1,
		},

		{
			name:   "Snare",
			volume: 0.36,
			data: func() string {
				var s strings.Builder

				for i := 0; i < 72; i++ {
					s.WriteString("NN 1, WN 0.25, NN 0.75, NN 1, WN 0.25, NN 0.75, ")
				}
				return s.String()
			}(),
			attack:  0.005,
			decay:   0.03,
			sustain: 0.1,
			release: 0.05,
		},

		{
			name:   "Wind",
			volume: 0.28,
			data: func() string {
				var s strings.Builder

				for i := 0; i < 9; i++ {
					s.WriteString("WN 16, NN 16, ")
				}
				return s.String()
			}(),
			attack:  1.2,
			decay:   0.5,
			sustain: 0.6,
			release: 1.0,
		},
	},
}

// We’ll create a single slice of chord names—one chord per measure for 96 measures.
// Then each instrument will build its data string based on that slice.
//
// For convenience, let's define a chord progression in A minor
// but with variations and passing chords across 6 sections, each 16 measures long.
var chordProgression = []string{

	"A2/C3/E3",
	"A2/C3/E3",
	"G2/B2/D3",
	"G2/B2/D3",
	"F2/A2/C3",
	"E2/G#2/B2",
	"D2/F2/A2",
	"G2/B2/D3",
	"A2/C3/E3",
	"A2/C3/E3",
	"F2/A2/C3",
	"E2/G#2/B2",
	"D2/F2/A2",
	"G2/B2/D3",
	"F2/A2/C3",
	"E2/G#2/B2",

	"F2/A2/C3",
	"G2/B2/D3",
	"A2/C3/E3",
	"A2/C3/E3",
	"F2/A2/C3",
	"G2/B2/D3",
	"C2/E2/G2",
	"E2/G#2/B2",
	"D2/F2/A2",
	"F2/A2/C3",
	"G2/B2/D3",
	"E2/G#2/B2",
	"A2/C3/E3",
	"G2/B2/D3",
	"F2/A2/C3",
	"E2/G#2/B2",

	"A2/C3/E3",
	"A2/C3/E3",
	"G2/B2/D3",
	"G2/B2/D3",
	"F2/A2/C3",
	"F2/A2/C3",
	"E2/G#2/B2",
	"E2/G#2/B2",
	"D2/F2/A2",
	"D2/F2/A2",
	"F2/A2/C3",
	"F2/A2/C3",
	"G2/B2/D3",
	"G2/B2/D3",
	"A2/C3/E3",
	"A2/C3/E3",

	"A2/C3/E3",
	"F2/A2/C3",
	"G2/B2/D3",
	"E2/G#2/B2",
	"D2/F2/A2",
	"G2/B2/D3",
	"A2/C3/E3",
	"A2/C3/E3",
	"F2/A2/C3",
	"G2/B2/D3",
	"A2/C3/E3",
	"C2/E2/G2",
	"D2/F2/A2",
	"E2/G#2/B2",
	"F2/A2/C3",
	"G2/B2/D3",

	"A2/C3/E3",
	"A2/C3/E3",
	"A2/C3/E3",
	"G2/B2/D3",
	"F2/A2/C3",
	"E2/G#2/B2",
	"D2/F2/A2",
	"G2/B2/D3",
	"F2/A2/C3",
	"G2/B2/D3",
	"A2/C3/E3",
	"E2/G#2/B2",
	"D2/F2/A2",
	"G2/B2/D3",
	"A2/C3/E3",
	"E2/G#2/B2",

	"A2/C3/E3",
	"A2/C3/E3",
	"G2/B2/D3",
	"G2/B2/D3",
	"F2/A2/C3",
	"F2/A2/C3",
	"E2/G#2/B2",
	"E2/G#2/B2",
	"A2/C3/E3",
	"A2/C3/E3",
	"G2/B2/D3",
	"F2/A2/C3",
	"E2/G#2/B2",
	"E2/G#2/B2",
	"A2/C3/E3",
	"A2/C3/E3",
}

func chordToMeasure(chord string, beats float64) string {

	return chord + fmt.Sprintf(" %.2f, ", beats)
}

var infinitoRealms = songData{
	name:     "Infinito Realms - A Grand 3-Minute Odyssey",
	bpm:      120,
	reverb:   0.50,
	delay:    0.2,
	feedback: 0.5,
	ins: []insData{

		{
			name:     "Strings",
			volume:   0.55,
			waveform: WAVE_SAW,
			data: func() string {
				var s strings.Builder
				for _, chord := range chordProgression {
					s.WriteString(chordToMeasure(chord, 4))
				}
				return s.String()
			}(),
			attack:  0.08,
			decay:   0.10,
			sustain: 0.80,
			release: 0.3,
		},

		{
			name:     "Brass",
			volume:   0.5,
			waveform: WAVE_SQUARE,
			data: func() string {
				var s strings.Builder
				for _, chord := range chordProgression {

					s.WriteString(chord + " 2, NN 2, ")
				}
				return s.String()
			}(),
			attack:  0.04,
			decay:   0.12,
			sustain: 0.75,
			release: 0.3,
		},

		{
			name:     "Bass",
			volume:   0.5,
			waveform: WAVE_SINE,
			data: func() string {
				var s strings.Builder

				for _, chord := range chordProgression {
					parts := strings.Split(chord, "/")

					s.WriteString(parts[0] + " 4, ")
				}
				return s.String()
			}(),
			attack:  0.02,
			decay:   0.10,
			sustain: 0.70,
			release: 0.25,
		},

		{
			name:     "Arp",
			volume:   0.35,
			waveform: WAVE_TRIANGLE,
			data: func() string {
				var s strings.Builder

				patterns := [][]int{
					{0, 1, 2, 3},
					{3, 2, 1, 0},
					{0, 2, 1, 3},
					{1, 0, 3, 2},
					{0, 1, 3, 2},
					{2, 3, 1, 0},
				}

				for measureIndex, chord := range chordProgression {

					notes := strings.Split(chord, "/")
					if len(notes) < 3 {

						notes = []string{"A2", "C3", "E3"}
					}

					patternIndex := measureIndex % len(patterns)
					pattern := patterns[patternIndex]

					chordNotes := make([]string, 4)
					chordNotes[0] = notes[0]
					chordNotes[1] = notes[1]
					chordNotes[2] = notes[2]
					chordNotes[3] = incrementOctave(notes[0])

					for repeat := 0; repeat < 2; repeat++ {
						for _, noteIndex := range pattern {
							theNote := chordNotes[noteIndex]
							s.WriteString(theNote + " 0.5, ")
						}
					}
				}

				return s.String()
			}(),
			attack:  0.01,
			decay:   0.05,
			sustain: 0.8,
			release: 0.1,
		},

		{
			name:     "Lead",
			volume:   0.40,
			waveform: WAVE_SAW,
			data: func() string {

				phrase1 := []string{
					"A3 1, C4 1, E4 1, A4 1, ",
					"G4 2, E4 2, ",
					"C4 1, D4 1, E4 1, G4 1, ",
					"E4 2, D4 2, ",
					"F4 1, E4 1, D4 1, C4 1, ",
					"B3 2, C4 2, ",
					"A3 1, C4 1, D4 1, E4 1, ",
					"E4 2, G4 2, ",
					"A3 1, C4 1, E4 1, A4 1, ",
					"G4 2, E4 2, ",
					"F4 1, E4 1, D4 1, C4 1, ",
					"B3 2, A3 2, ",
					"A3 1, C4 1, D4 1, E4 1, ",
					"E4 2, D4 2, ",
					"F4 1, E4 1, D4 1, C4 1, ",
					"B3 2, B3 2, ",
				}

				phrase2 := make([]string, 16)
				copy(phrase2, phrase1)
				phrase2[7] = "E4 2, A4 2, "
				phrase2[15] = "B3 1, C4 1, D4 1, E4 1, "
				phrase3 := make([]string, 16)
				copy(phrase3, phrase1)
				phrase3[0] = "A3 2, C4 2, "
				phrase3[15] = "B3 2, A3 2, "
				phrase4 := make([]string, 16)
				copy(phrase4, phrase1)
				phrase4[14] = "F4 2, E4 2, "
				phrase5 := make([]string, 16)
				copy(phrase5, phrase1)
				phrase5[7] = "G4 1, A4 1, G4 1, E4 1, "
				phrase6 := make([]string, 16)
				copy(phrase6, phrase1)
				phrase6[15] = "A4 1, G4 1, E4 1, C4 1, "

				allPhrases := [][]string{phrase1, phrase2, phrase3, phrase4, phrase5, phrase6}
				var s strings.Builder
				for _, ph := range allPhrases {
					for _, measureData := range ph {
						s.WriteString(measureData)
					}
				}
				return s.String()
			}(),
			attack:  0.1,
			decay:   0.1,
			sustain: 0.85,
			release: 0.4,
		},

		{
			name:     "Kick",
			volume:   0.5,
			waveform: WAVE_SINE,
			data: func() string {
				var s strings.Builder

				for i := 0; i < 96; i++ {
					for b := 0; b < 4; b++ {
						s.WriteString("C1 0.5, NN 0.5, ")
					}
				}
				return s.String()
			}(),
			attack:  0.01,
			decay:   0.06,
			sustain: 0.2,
			release: 0.1,
		},

		{
			name:   "Snare",
			volume: 0.4,
			data: func() string {
				var s strings.Builder

				for i := 0; i < 96; i++ {
					s.WriteString("NN 1, WN 0.25, NN 0.75, NN 1, WN 0.25, NN 0.75, ")
				}
				return s.String()
			}(),
			attack:  0.005,
			decay:   0.02,
			sustain: 0.1,
			release: 0.05,
		},

		{
			name:   "Hats",
			volume: 0.3,
			data: func() string {
				var s strings.Builder

				for i := 0; i < 96; i++ {
					for j := 0; j < 4; j++ {
						s.WriteString("NN 0.5, WN 0.25, NN 0.25, ")
					}
				}
				return s.String()
			}(),
			attack:  0.002,
			decay:   0.03,
			sustain: 0.0,
			release: 0.02,
		},

		{
			name:   "FX",
			volume: 0.25,
			data: func() string {
				var s strings.Builder

				for i := 0; i < 12; i++ {
					s.WriteString("NN 31, WN 1, ")
				}
				return s.String()
			}(),
			attack:  0.1,
			decay:   0.4,
			sustain: 0.5,
			release: 1.0,
		},
	},
}

// incrementOctave is a helper to nudge the note's octave up by 1.
// (This is very simplistic string parsing—works for typical note naming like "A2" or "C#3".)
func incrementOctave(note string) string {

	if len(note) < 2 {
		return note
	}
	lastChar := note[len(note)-1]
	if lastChar < '0' || lastChar > '9' {

		return note
	}

	octave := int(lastChar - '0')
	octaveUp := octave + 1

	return note[:len(note)-1] + fmt.Sprintf("%d", octaveUp)
}

var epicWarOfTheAncients = songData{
	name:     "Epic War of the Ancients - Battle Hymn",
	bpm:      100,
	reverb:   0.50,
	delay:    0.2,
	feedback: 0.5,
	ins: []insData{

		{
			name:     "Strings",
			volume:   0.6,
			waveform: WAVE_SAW,

			data: "" +

				"D3/F3/A3 4, D3/F3/A3 4, C3/E3/G3 4, C3/E3/G3 4, Bb2/D3/F3 4, F2/A2/C3 4, G2/Bb2/D3 4, A2/C#3/E3 4," +

				"D3/F3/A3 4, D3/F3/A3 4, F2/A2/C3 4, F2/A2/C3 4, G2/Bb2/D3 4, G2/Bb2/D3 4, A2/C#3/E3 4, A2/C#3/E3 4," +

				"D3/F3/A3 4, D3/F3/A3 4, C3/E3/G3 4, C3/E3/G3 4, Bb2/D3/F3 4, G2/Bb2/D3 4, A2/C#3/E3 4, A2/C#3/E3 4," +

				"D3/F3/A3 4, C3/E3/G3 4, Bb2/D3/F3 4, F2/A2/C3 4, G2/Bb2/D3 4, A2/C#3/E3 4, D3/F3/A3 4, D3/F3/A3 4",
			attack:  0.08,
			decay:   0.10,
			sustain: 0.80,
			release: 0.25,
		},

		{
			name:     "Brass",
			volume:   0.55,
			waveform: WAVE_SQUARE,

			data: "" +

				"D3/F3/A3 2, NN 2, D3/F3/A3 2, NN 2, C3/E3/G3 2, NN 2, C3/E3/G3 2, NN 2, Bb2/D3/F3 2, NN 2, F2/A2/C3 2, NN 2, G2/Bb2/D3 2, NN 2, A2/C#3/E3 2, NN 2," +

				"D3/F3/A3 2, NN 2, D3/F3/A3 2, NN 2, F2/A2/C3 2, NN 2, F2/A2/C3 2, NN 2, G2/Bb2/D3 2, NN 2, G2/Bb2/D3 2, NN 2, A2/C#3/E3 2, NN 2, A2/C#3/E3 2, NN 2," +

				"D3/F3/A3 2, NN 2, D3/F3/A3 2, NN 2, C3/E3/G3 2, NN 2, C3/E3/G3 2, NN 2, Bb2/D3/F3 2, NN 2, G2/Bb2/D3 2, NN 2, A2/C#3/E3 2, NN 2, A2/C#3/E3 2, NN 2," +

				"D3/F3/A3 2, NN 2, C3/E3/G3 2, NN 2, Bb2/D3/F3 2, NN 2, F2/A2/C3 2, NN 2, G2/Bb2/D3 2, NN 2, A2/C#3/E3 2, NN 2, D3/F3/A3 2, NN 2, D3/F3/A3 2, NN 2",
			attack:  0.06,
			decay:   0.15,
			sustain: 0.75,
			release: 0.3,
		},

		{
			name:     "Bass",
			volume:   0.5,
			waveform: WAVE_SINE,

			data: "" +

				"D2 4, D2 4, C2 4, C2 4, Bb1 4, F2 4, G1 4, A1 4," +

				"D2 4, D2 4, F2 4, F2 4, G1 4, G1 4, A1 4, A1 4," +

				"D2 4, D2 4, C2 4, C2 4, Bb1 4, G1 4, A1 4, A1 4," +

				"D2 4, C2 4, Bb1 4, F2 4, G1 4, A1 4, D2 4, D2 4",
			attack:  0.02,
			decay:   0.10,
			sustain: 0.70,
			release: 0.20,
		},

		{
			name:     "Lead",
			volume:   0.45,
			waveform: WAVE_TRIANGLE,

			data: "" +

				"D4 1, F4 1, A4 1, D5 1, " +

				"D5 2, C5 2, " +

				"C5 1, A4 1, F4 1, D4 1, " +

				"D4 2, F4 2, " +

				"E4 1, F4 1, G4 1, A4 1, " +

				"A4 2, G4 2, " +

				"Bb4 1, A4 1, G4 1, F4 1, " +

				"E4 2, D4 2, " +

				"D4 1, F4 1, A4 1, C5 1, " +

				"C5 2, Bb4 2, " +

				"C4 1, D4 1, F4 1, A4 1, " +

				"A4 2, G4 2, " +

				"G4 1, Bb4 1, C5 1, D5 1, " +

				"D5 2, C5 2, " +

				"E4 1, G4 1, A4 1, C5 1, " +

				"A4 2, D5 2, " +

				"D4 1, F4 1, A4 1, D5 1, " +

				"D5 2, C5 2, " +

				"C5 1, A4 1, F4 1, D4 1, " +

				"D4 2, F4 2, " +

				"E4 1, G4 1, A4 1, Bb4 1, " +

				"Bb4 2, A4 2, " +

				"C5 1, Bb4 1, G4 1, E4 1, " +

				"E4 2, D4 2, " +

				"D4 1, F4 1, A4 1, D5 1, " +

				"C5 2, D5 2, " +

				"Bb4 1, A4 1, F4 1, D4 1, " +

				"D4 2, F4 2, " +

				"G4 1, Bb4 1, C5 1, D5 1, " +

				"F5 2, E5 2, " +

				"D5 1, C5 1, A4 1, F4 1, " +

				"D4 2, D4 2",
			attack:  0.10,
			decay:   0.15,
			sustain: 0.85,
			release: 0.40,
		},

		{
			name:     "Kick",
			volume:   0.5,
			waveform: WAVE_SINE,
			data: func() string {
				var s string

				for i := 0; i < 32; i++ {
					s += "C1 0.5, NN 0.5, NN 1, C1 0.5, NN 0.5, C1 0.5, NN 1, "
				}
				return s
			}(),
			attack:  0.01,
			decay:   0.06,
			sustain: 0.2,
			release: 0.1,
		},

		{
			name:   "Snare",
			volume: 0.45,
			data: func() string {
				var s string

				for i := 0; i < 32; i++ {
					s += "NN 1, WN 0.25, NN 0.75, NN 1, WN 0.25, NN 0.75, "
				}
				return s
			}(),
			attack:  0.005,
			decay:   0.02,
			sustain: 0.1,
			release: 0.05,
		},

		{
			name:   "Cymbals",
			volume: 0.3,
			data: func() string {
				var s string

				for i := 0; i < 32; i++ {
					s += "WN 0.25, NN 3.75, "
				}
				return s
			}(),
			attack:  0.001,
			decay:   0.05,
			sustain: 0.0,
			release: 0.02,
		},

		{
			name:   "Storm",
			volume: 0.25,
			data: func() string {
				var s string

				for i := 0; i < 8; i++ {
					s += "WN 8, NN 8, "
				}
				return s
			}(),
			attack:  1.5,
			decay:   0.5,
			sustain: 0.6,
			release: 1.0,
		},
	},
}

var twilightReflections = songData{
	name:     "Twilight Reflections - Nocturnal Journey",
	bpm:      90,
	reverb:   0.50,
	delay:    0.2,
	feedback: 0.5,
	ins: []insData{

		{
			name:     "Harmony",
			volume:   0.55,
			waveform: WAVE_SAW,

			data: "" +

				"E3/G3/B3 4, E3/G3/B3 4, C3/E3/G3 4, C3/E3/G3 4, " +
				"G3/B3/D4 4, G3/B3/D4 4, D3/F#3/A3 4, D3/F#3/A3 4, " +

				"E3/G3/B3 4, E3/G3/B3 4, C3/E3/G3 4, C3/E3/G3 4, " +
				"A3/C3/E3 4, A3/C3/E3 4, B2/D#3/F#3 4, B2/D#3/F#3 4, " +

				"G3/B3/D4 4, G3/B3/D4 4, B2/D3/F#3 4, B2/D3/F#3 4, " +
				"E3/G3/B3 4, E3/G3/B3 4, C3/E3/G3 4, C3/E3/G3 4, " +

				"E3/G3/B3 4, C3/E3/G3 4, G3/B3/D4 4, D3/F#3/A3 4, " +
				"E3/G3/B3 4, C3/E3/G3 4, G3/B3/D4 4, D3/F#3/A3 4",
			attack:  0.05,
			decay:   0.10,
			sustain: 0.80,
			release: 0.25,
		},

		{
			name:     "Bass",
			volume:   0.45,
			waveform: WAVE_SINE,

			data: "" +

				"E2 4, E2 4, C2 4, C2 4, G2 4, G2 4, D2 4, D2 4," +

				"E2 4, E2 4, C2 4, C2 4, A2 4, A2 4, B2 4, B2 4," +

				"G2 4, G2 4, B2 4, B2 4, E2 4, E2 4, C2 4, C2 4," +

				"E2 4, C2 4, G2 4, D2 4, E2 4, C2 4, G2 4, D2 4",
			attack:  0.01,
			decay:   0.10,
			sustain: 0.70,
			release: 0.20,
		},

		{
			name:     "Melody",
			volume:   0.40,
			waveform: WAVE_TRIANGLE,

			data: "" +

				"E4 1, G4 1, B4 1, E5 1,  E5 2, D5 2, " +

				"C5 1, B4 1, G4 1, E4 1,  E4 2, G4 2, " +

				"E4 1, G4 1, B4 1, E5 1,  E5 2, D5 2, " +

				"C5 1, D5 1, F#4 1, A4 1, A4 2, G4 2, " +

				"E4 1, G4 1, B4 1, C5 1,  C5 2, B4 2, " +

				"C4 1, E4 1, A4 1, A4 1,  A4 2, G4 2, " +

				"A4 1, C5 1, E5 1, D#5 1, D#5 2, B4 2, " +

				"B4 1, D#5 1, F#4 1, B4 1, B4 2, A4 2, " +

				"G4 1, B4 1, D5 1, B4 1,  G4 2, F#4 2, " +

				"B4 1, D5 1, F#4 1, D5 1, D5 2, B4 2, " +

				"E4 1, G4 1, B4 1, E5 1,  E5 2, D5 2, " +

				"C5 1, B4 1, G4 1, E4 1,  E4 2, G4 2, " +

				"E4 1, G4 1, B4 1, E5 1,  E5 2, D5 2, " +

				"C5 1, D5 1, F#4 1, A4 1, A4 2, G4 2, " +

				"E4 1, G4 1, B4 1, E5 1,  E5 2, D5 2, " +

				"C5 1, B4 1, G4 1, D4 1,  D4 2, E4 2",
			attack:  0.12,
			decay:   0.10,
			sustain: 0.85,
			release: 0.40,
		},

		{
			name:     "Kick",
			volume:   0.5,
			waveform: WAVE_SINE,
			data: func() string {
				var s string

				for i := 0; i < 32; i++ {
					s += "C1 0.5, NN 0.5, NN 1, C1 0.5, NN 0.5, NN 1, "
				}
				return s
			}(),
			attack:  0.01,
			decay:   0.05,
			sustain: 0.2,
			release: 0.1,
		},

		{
			name:   "Snare",
			volume: 0.4,
			data: func() string {
				var s string

				for i := 0; i < 32; i++ {
					s += "NN 1, WN 0.25, NN 0.75, NN 1, WN 0.25, NN 0.75, "
				}
				return s
			}(),
			attack:  0.005,
			decay:   0.02,
			sustain: 0.1,
			release: 0.05,
		},

		{
			name:   "Crickets",
			volume: 0.25,
			data: func() string {
				var s string

				for i := 0; i < 8; i++ {
					s += "WN 8, NN 8, "
				}
				return s
			}(),
			attack:  1.0,
			decay:   0.4,
			sustain: 0.6,
			release: 1.0,
		},

		{
			name:   "Thunder",
			volume: 0.3,
			data: func() string {
				var s string

				for i := 0; i < 4; i++ {

					s += "NN 28, WN 1, NN 3, "
				}
				return s
			}(),
			attack:  0.05,
			decay:   0.3,
			sustain: 0.0,
			release: 0.5,
		},
	},
}

var voyageOfTheAbyss = songData{
	name:     "Voyage of the Abyss - Nautical Odyssey",
	bpm:      80,
	reverb:   0.50,
	delay:    0.2,
	feedback: 0.5,
	ins: []insData{

		{
			name:     "Harmony",
			volume:   0.6,
			waveform: WAVE_SAW,
			data: "D4/F4/A4 4, D4/F4/A4 4, D4/F4/A4 4, D4/F4/A4 4, " +
				"D4/F4/A4 4, D4/F4/A4 4, D4/F4/A4 4, D4/F4/A4 4," +

				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, A3/C#4/E4 4, " +
				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, A3/C#4/E4 4, " +
				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, A3/C#4/E4 4, " +
				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, A3/C#4/E4 4," +

				"G3/Bb3/D4 4, A3/C#4/E4 4, G3/Bb3/D4 4, A3/C#4/E4 4, " +
				"F3/A3/C4 4, G3/Bb3/D4 4, A3/C#4/E4 4, D4/F4/A4 4," +

				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, A3/C#4/E4 4, " +
				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, D4/F4/A4 4",
			attack:  0.05,
			decay:   0.10,
			sustain: 0.80,
			release: 0.20,
		},

		{
			name:     "Bass",
			volume:   0.5,
			waveform: WAVE_SINE,
			data: "D2 2, A2 2, D2 2, A2 2, D2 2, A2 2, D2 2, A2 2," +
				"D2 4, Bb1 4, C2 4, A1 4, " +
				"D2 4, Bb1 4, C2 4, A1 4, " +
				"D2 4, Bb1 4, C2 4, A1 4, " +
				"D2 4, Bb1 4, C2 4, A1 4," +
				"G2 4, A1 4, G2 4, A1 4, F2 4, G2 4, A1 4, D2 4," +
				"D2 4, Bb1 4, C2 4, A1 4, D2 4, Bb1 4, C2 4, D2 4",
			attack:  0.01,
			decay:   0.10,
			sustain: 0.70,
			release: 0.20,
		},

		{
			name:     "Melody",
			volume:   0.38,
			waveform: WAVE_TRIANGLE,
			data: "D3 1, F3 1, A3 1, D4 1, " +
				"A3 1, F3 1, D3 1, A2 1, " +
				"D3 0.5, E3 0.5, F3 1, G3 1, A3 1, " +
				"A3 0.5, G3 0.5, F3 0.5, E3 0.5, D3 1, D3 1, " +
				"D3 1, E3 1, F3 1, G3 1, " +
				"G3 0.5, A3 0.5, Bb3 1, A3 1, G3 1, " +
				"F3 1, E3 1, D3 1, C3 1, " +
				"D3 2, D3 2, " +
				"F3 0.5, A3 0.5, D4 1, C4 1, A3 1, " +
				"Bb3 0.5, D4 0.5, F4 1, D4 1, Bb3 1, " +
				"C4 1, E4 1, G4 1, E4 1, " +
				"C#4 0.5, E4 0.5, A4 1, G4 1, E4 1, " +
				"D4 0.5, F4 0.5, A4 1, F4 1, D4 1, " +
				"Bb3 0.5, D4 0.5, F4 1, D4 1, Bb3 1, " +
				"C4 1, E4 1, G4 1, E4 1, " +
				"A3 0.5, C#4 0.5, E4 1, C#4 1, A3 1, " +
				"G3 1, Bb3 1, D4 1, Bb3 1, " +
				"A3 1, C#4 1, E4 1, C#4 1, " +
				"F3 0.5, G3 0.5, Bb3 1, D4 1, Bb3 1, " +
				"A3 1, C#4 1, E4 1, A3 1, " +
				"F3 1, A3 1, C4 1, A3 1, " +
				"G3 1, Bb3 1, D4 1, Bb3 1, " +
				"A3 0.5, C#4 0.5, E4 1, C#4 1, A3 1, " +
				"D3 1, F3 1, A3 1, F3 1, " +
				"F3 0.5, A3 0.5, D4 1, C4 1, A3 1, " +
				"Bb3 0.5, D4 0.5, F4 1, D4 1, Bb3 1, " +
				"C4 1, E4 1, G4 1, E4 1, " +
				"C#4 0.5, E4 0.5, A4 1, G4 1, E4 1, " +
				"D4 0.5, F3 0.5, A3 1, F3 1, D3 1, " +
				"Bb3 0.5, D3 0.5, F3 1, D3 1, Bb3 1, " +
				"C3 1, E3 1, G3 1, E3 1, " +
				"D3 2, F3 2",
			attack:  0.15,
			decay:   0.15,
			sustain: 0.85,
			release: 0.35,

			square: 0.05,
		},

		{
			name:     "Kick",
			volume:   0.5,
			waveform: WAVE_SINE,
			data: func() string {
				var s string

				for i := 0; i < 40; i++ {
					s += "C1 0.5, NN 0.5, NN 1, C1 0.5, NN 0.5, NN 1, "
				}
				return s
			}(),
			attack:  0.01,
			decay:   0.05,
			sustain: 0.2,
			release: 0.1,
			square:  0.0,
		},

		{
			name:   "Snare",
			volume: 0.4,
			data: func() string {
				var s string

				for i := 0; i < 40; i++ {
					s += "NN 1, WN 0.25, NN 0.75, NN 1, WN 0.25, NN 0.75, "
				}
				return s
			}(),
			attack:  0.005,
			decay:   0.02,
			sustain: 0.1,
			release: 0.05,
			square:  0.0,
		},

		{
			name:   "Waves",
			volume: 0.2,
			data: func() string {
				var s string

				for i := 0; i < 20; i++ {
					s += "WN 4, NN 4, "
				}
				return s
			}(),
			attack:  1.5,
			decay:   0.5,
			sustain: 0.7,
			release: 1.0,
			square:  0.0,
		},

		{
			name:   "Water Splashes",
			volume: 0.3,
			data: func() string {
				var s string

				for i := 0; i < 10; i++ {
					s += "WN 2, NN 14, "
				}
				return s
			}(),
			attack:  0.05,
			decay:   0.1,
			sustain: 0.0,
			release: 0.2,
			square:  0.0,
		},
	},
}
