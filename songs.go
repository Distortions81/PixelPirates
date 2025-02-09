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
	fmt.Printf("Parsed %v songs.\n", count)
}

var (
	gameModePlaylists = [GAME_MAX]playlistData{
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

// --------------------------------------------------------------------------------
// Chord Progression (72 measures total)
// We’ll break it into 4 sections of 18 chords each, in E minor with occasional
// borrowed chords (F, B major, etc.) for a moody, haunting vibe.
// --------------------------------------------------------------------------------

var chordProgression2 = []string{
	// ------------------------------------------------------------------
	// SECTION 1 (Measures 1-18)
	// ------------------------------------------------------------------
	"E2/G2/B2",   // 1:  Em
	"E2/G2/B2",   // 2:  Em
	"D2/F#2/A2",  // 3:  D
	"D2/F#2/A2",  // 4:  D
	"F2/A2/C3",   // 5:  F
	"E2/G2/B2",   // 6:  Em
	"G2/B2/D3",   // 7:  G
	"E2/G2/B2",   // 8:  Em
	"E2/G2/B2",   // 9:  Em
	"B2/D#3/F#3", //10: B major
	"C2/E2/G2",   // 11: C
	"B2/D#3/F#3", //12: B
	"E2/G2/B2",   // 13: Em
	"E2/G2/B2",   // 14: Em
	"A2/C3/E3",   // 15: Am
	"B2/D#3/F#3", //16: B
	"E2/G2/B2",   // 17: Em
	"E2/G2/B2",   // 18: Em

	// ------------------------------------------------------------------
	// SECTION 2 (Measures 19-36)
	// ------------------------------------------------------------------
	"E2/G2/B2",   //19: Em
	"G2/B2/D3",   //20: G
	"C2/E2/G2",   //21: C
	"D2/F#2/A2",  //22: D
	"F2/A2/C3",   //23: F
	"E2/G2/B2",   //24: Em
	"B2/D#3/F#3", //25: B
	"B2/D#3/F#3", //26: B
	"F2/A2/C3",   //27: F
	"G2/B2/D3",   //28: G
	"A2/C3/E3",   //29: Am
	"B2/D#3/F#3", //30: B
	"E2/G2/B2",   //31: Em
	"D2/F#2/A2",  //32: D
	"C2/E2/G2",   //33: C
	"F2/A2/C3",   //34: F
	"E2/G2/B2",   //35: Em
	"E2/G2/B2",   //36: Em

	// ------------------------------------------------------------------
	// SECTION 3 (Measures 37-54)
	// ------------------------------------------------------------------
	"E2/G2/B2",   //37: Em
	"D2/F#2/A2",  //38: D
	"D2/F#2/A2",  //39: D
	"F2/A2/C3",   //40: F
	"E2/G2/B2",   //41: Em
	"E2/G2/B2",   //42: Em
	"B2/D#3/F#3", //43: B
	"B2/D#3/F#3", //44: B
	"C2/E2/G2",   //45: C
	"A2/C3/E3",   //46: Am
	"D2/F#2/A2",  //47: D
	"G2/B2/D3",   //48: G
	"F2/A2/C3",   //49: F
	"E2/G2/B2",   //50: Em
	"D2/F#2/A2",  //51: D
	"C2/E2/G2",   //52: C
	"B2/D#3/F#3", //53: B
	"E2/G2/B2",   //54: Em

	// ------------------------------------------------------------------
	// SECTION 4 (Measures 55-72)
	// ------------------------------------------------------------------
	"F2/A2/C3",   //55: F
	"F2/A2/C3",   //56: F
	"A2/C3/E3",   //57: Am
	"B2/D#3/F#3", //58: B
	"E2/G2/B2",   //59: Em
	"E2/G2/B2",   //60: Em
	"G2/B2/D3",   //61: G
	"F2/A2/C3",   //62: F
	"F2/A2/C3",   //63: F
	"G2/B2/D3",   //64: G
	"B2/D#3/F#3", //65: B
	"A2/C3/E3",   //66: Am
	"E2/G2/B2",   //67: Em
	"D2/F#2/A2",  //68: D
	"C2/E2/G2",   //69: C
	"B2/D#3/F#3", //70: B
	"E2/G2/B2",   //71: Em
	"E2/G2/B2",   //72: Em
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
	{0, 1, 2, 3}, // up
	{3, 2, 1, 0}, // down
	{0, 2, 1, 3}, // skip
	{1, 0, 3, 2}, // skip alt
	{0, 1, 3, 2}, // partial up
	{2, 3, 1, 0}, // partial down
	{1, 2, 0, 3}, // swirl
	{3, 0, 2, 1}, // swirl alt
}

// Our new “haunting” track definition:
var spectersOfAshenTwilight = songData{
	name:     "Specters of Ashen Twilight - A Haunting Saga",
	bpm:      80, // 72 measures => 288 beats => 3.6 mins at 80 BPM
	reverb:   0.50,
	delay:    0.2,
	feedback: 0.5,

	ins: []insData{
		//----------------------------------------------------------------------
		// 1. GhostPad (sawtooth) - Held chords each measure, slower attack.
		//----------------------------------------------------------------------
		{
			name:     "GhostPad",
			volume:   0.50,
			waveform: WAVE_SAW,
			data: func() string {
				var s strings.Builder
				for _, chord := range chordProgression2 {
					s.WriteString(chordToMeasure2(chord, 4)) // hold each chord for 4 beats
				}
				return s.String()
			}(),
			attack:  0.15,
			decay:   0.2,
			sustain: 0.8,
			release: 0.4,
		},
		//----------------------------------------------------------------------
		// 2. Choir (square) - Bold secondary layer, half-measure hits.
		//----------------------------------------------------------------------
		{
			name:     "Choir",
			volume:   0.45,
			waveform: WAVE_SQUARE,
			data: func() string {
				var s strings.Builder
				// Each chord for 2 beats, then 2 beats of rest
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
		//----------------------------------------------------------------------
		// 3. Bass (sine) - Root note for each measure
		//----------------------------------------------------------------------
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
		//----------------------------------------------------------------------
		// 4. Arp (triangle) - Evolving patterns each measure for movement.
		//----------------------------------------------------------------------
		{
			name:     "Arp",
			volume:   0.33,
			waveform: WAVE_TRIANGLE,
			data: func() string {
				var s strings.Builder
				for measureIndex, chord := range chordProgression2 {
					notes := strings.Split(chord, "/")
					// fallback if chord is malformed
					if len(notes) < 3 {
						notes = []string{"E2", "G2", "B2"}
					}
					// pick pattern
					patIndex := measureIndex % len(arpPatterns2)
					pattern := arpPatterns2[patIndex]

					// Build chordNotes with an octave-lifted root as the 4th index
					chordNotes := []string{
						notes[0],
						notes[1],
						notes[2],
						incrementOctave2(notes[0]),
					}
					// 4 beats => 8 sub-beats @ 0.5 each.
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
		//----------------------------------------------------------------------
		// 5. DarkBell (sine) - Occasional eerie bell hits each measure, short & soft.
		//----------------------------------------------------------------------
		{
			name:     "DarkBell",
			volume:   0.25,
			waveform: WAVE_SINE,
			data: func() string {
				var s strings.Builder
				// Let’s place 1 or 2 quick bell tones each measure at random offsets.
				// For simplicity, we’ll do a small repeating pattern:
				// measure => bell at beat 1.5 and maybe beat 3.5, then next measure silent or less frequent hits.
				// We'll alternate every measure so it doesn't get too repetitive.
				for i := 0; i < 72; i++ {
					if i%2 == 0 {
						// Even measure: 2 bell hits
						s.WriteString("NN 1.5, A4 0.25, NN 0.25, NN 1, D5 0.25, NN 0.75, ")
					} else {
						// Odd measure: 1 bell hit
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
		//----------------------------------------------------------------------
		// 6. Lead (sawtooth) - Minimal, haunting lines, repeated in short phrases.
		//----------------------------------------------------------------------
		{
			name:     "Lead",
			volume:   0.35,
			waveform: WAVE_SAW,
			data: func() string {
				// We'll define short 4-measure phrases repeated or varied.
				// Each measure is 4 beats => define a minimal line with half/quarter notes.
				phrase1 := []string{
					"E4 2, G4 2, ",             // m1
					"B3 1, C4 1, A3 1, G3 1, ", // m2
					"F3 2, E4 2, ",             // m3
					"E3 2, NN 2, ",             // m4 (rests)
				}
				phrase2 := make([]string, len(phrase1))
				copy(phrase2, phrase1)
				// small variation in measure 2
				phrase2[1] = "D4 1, C4 1, A3 1, G3 1, "

				phrase3 := make([]string, len(phrase1))
				copy(phrase3, phrase1)
				phrase3[2] = "F4 1, E4 1, D4 1, C4 1, " // measure 3 changed
				phrase3[3] = "NN 2, G4 2, "             // measure 4 changed

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
		//----------------------------------------------------------------------
		// 7. Kick (sine) - Sparse, half-time hits for a slow, heavy feel.
		//----------------------------------------------------------------------
		{
			name:     "Kick",
			volume:   0.45,
			waveform: WAVE_SINE,
			data: func() string {
				var s strings.Builder
				// For each measure (4 beats), let's place a kick on beat 1 and beat 3 for a half-time vibe:
				// measure => Kick(1), silence(1), Kick(1), silence(1)
				for i := 0; i < 72; i++ {
					// Beat 1
					s.WriteString("C1 0.5, NN 0.5, ")
					// Beat 2
					s.WriteString("NN 1, ")
					// Beat 3
					s.WriteString("C1 0.5, NN 0.5, ")
					// Beat 4
					s.WriteString("NN 1, ")
				}
				return s.String()
			}(),
			attack:  0.01,
			decay:   0.07,
			sustain: 0.2,
			release: 0.1,
		},
		//----------------------------------------------------------------------
		// 8. Snare (white noise) - Crisp backbeat on beats 2 and 4.
		//----------------------------------------------------------------------
		{
			name:   "Snare",
			volume: 0.36,
			data: func() string {
				var s strings.Builder
				// Each measure:
				//   Beat 1: silent
				//   Beat 2: snare
				//   Beat 3: silent
				//   Beat 4: snare
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
		//----------------------------------------------------------------------
		// 9. Wind (noise) - A swirling noise bed, fades in/out every few measures.
		//----------------------------------------------------------------------
		{
			name:   "Wind",
			volume: 0.28,
			data: func() string {
				var s strings.Builder
				// 72 measures total => let’s do 9 cycles of 8 measures each.
				// Each cycle: 4 measures (16 beats) of noise, 4 measures (16 beats) silence.
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
	//-------------------------------------------------------------------
	// SECTION 1 (Measures 1-16)
	//-------------------------------------------------------------------
	"A2/C3/E3",  // 1: Am
	"A2/C3/E3",  // 2: Am
	"G2/B2/D3",  // 3: G
	"G2/B2/D3",  // 4: G
	"F2/A2/C3",  // 5: F
	"E2/G#2/B2", // 6: E major
	"D2/F2/A2",  // 7: Dm
	"G2/B2/D3",  // 8: G
	"A2/C3/E3",  // 9: Am
	"A2/C3/E3",  // 10: Am
	"F2/A2/C3",  // 11: F
	"E2/G#2/B2", // 12: E
	"D2/F2/A2",  // 13: Dm
	"G2/B2/D3",  // 14: G
	"F2/A2/C3",  // 15: F
	"E2/G#2/B2", // 16: E

	//-------------------------------------------------------------------
	// SECTION 2 (Measures 17-32)
	//-------------------------------------------------------------------
	"F2/A2/C3",  // 17
	"G2/B2/D3",  // 18
	"A2/C3/E3",  // 19: Am
	"A2/C3/E3",  // 20: Am
	"F2/A2/C3",  // 21
	"G2/B2/D3",  // 22
	"C2/E2/G2",  // 23: C major
	"E2/G#2/B2", // 24: E major
	"D2/F2/A2",  // 25: Dm
	"F2/A2/C3",  // 26
	"G2/B2/D3",  // 27
	"E2/G#2/B2", // 28
	"A2/C3/E3",  // 29: Am
	"G2/B2/D3",  // 30
	"F2/A2/C3",  // 31
	"E2/G#2/B2", // 32

	//-------------------------------------------------------------------
	// SECTION 3 (Measures 33-48)
	//-------------------------------------------------------------------
	"A2/C3/E3",  // 33
	"A2/C3/E3",  // 34
	"G2/B2/D3",  // 35
	"G2/B2/D3",  // 36
	"F2/A2/C3",  // 37
	"F2/A2/C3",  // 38
	"E2/G#2/B2", // 39
	"E2/G#2/B2", // 40
	"D2/F2/A2",  // 41
	"D2/F2/A2",  // 42
	"F2/A2/C3",  // 43
	"F2/A2/C3",  // 44
	"G2/B2/D3",  // 45
	"G2/B2/D3",  // 46
	"A2/C3/E3",  // 47
	"A2/C3/E3",  // 48

	//-------------------------------------------------------------------
	// SECTION 4 (Measures 49-64)
	//-------------------------------------------------------------------
	"A2/C3/E3",  // 49
	"F2/A2/C3",  // 50
	"G2/B2/D3",  // 51
	"E2/G#2/B2", // 52
	"D2/F2/A2",  // 53
	"G2/B2/D3",  // 54
	"A2/C3/E3",  // 55
	"A2/C3/E3",  // 56
	"F2/A2/C3",  // 57
	"G2/B2/D3",  // 58
	"A2/C3/E3",  // 59
	"C2/E2/G2",  // 60
	"D2/F2/A2",  // 61
	"E2/G#2/B2", // 62
	"F2/A2/C3",  // 63
	"G2/B2/D3",  // 64

	//-------------------------------------------------------------------
	// SECTION 5 (Measures 65-80)
	//-------------------------------------------------------------------
	"A2/C3/E3",  // 65
	"A2/C3/E3",  // 66
	"A2/C3/E3",  // 67
	"G2/B2/D3",  // 68
	"F2/A2/C3",  // 69
	"E2/G#2/B2", // 70
	"D2/F2/A2",  // 71
	"G2/B2/D3",  // 72
	"F2/A2/C3",  // 73
	"G2/B2/D3",  // 74
	"A2/C3/E3",  // 75
	"E2/G#2/B2", // 76
	"D2/F2/A2",  // 77
	"G2/B2/D3",  // 78
	"A2/C3/E3",  // 79
	"E2/G#2/B2", // 80

	//-------------------------------------------------------------------
	// SECTION 6 (Measures 81-96)
	//-------------------------------------------------------------------
	"A2/C3/E3",  // 81
	"A2/C3/E3",  // 82
	"G2/B2/D3",  // 83
	"G2/B2/D3",  // 84
	"F2/A2/C3",  // 85
	"F2/A2/C3",  // 86
	"E2/G#2/B2", // 87
	"E2/G#2/B2", // 88
	"A2/C3/E3",  // 89
	"A2/C3/E3",  // 90
	"G2/B2/D3",  // 91
	"F2/A2/C3",  // 92
	"E2/G#2/B2", // 93
	"E2/G#2/B2", // 94
	"A2/C3/E3",  // 95
	"A2/C3/E3",  // 96
}

// Generate chord-based strings for each measure (each measure = 4 beats).
// Our chordProgression slice has 96 items, one per measure.

// For convenience in each instrument, we define a helper to get full measure chord data:

func chordToMeasure(chord string, beats float64) string {
	// Example usage: chordToMeasure("A2/C3/E3", 4) => "A2/C3/E3 4, "
	// This means hold that chord for 'beats' beats.
	return chord + fmt.Sprintf(" %.2f, ", beats)
}

var infinitoRealms = songData{
	name:     "Infinito Realms - A Grand 3-Minute Odyssey",
	bpm:      120, // 96 measures * 2s each = ~3:12 total
	reverb:   0.50,
	delay:    0.2,
	feedback: 0.5,
	ins: []insData{
		//----------------------------------------------------------------------
		// 1. Strings: Warm sawtooth pad, holding each measure's chord for 4 beats.
		//----------------------------------------------------------------------
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
		//----------------------------------------------------------------------
		// 2. Brass: Square wave, half-measure hits (2 beats), then 2 beats rest.
		//----------------------------------------------------------------------
		{
			name:     "Brass",
			volume:   0.5,
			waveform: WAVE_SQUARE,
			data: func() string {
				var s strings.Builder
				for _, chord := range chordProgression {
					// 1 measure = 4 beats => (chord 2, rest 2)
					s.WriteString(chord + " 2, NN 2, ")
				}
				return s.String()
			}(),
			attack:  0.04,
			decay:   0.12,
			sustain: 0.75,
			release: 0.3,
		},
		//----------------------------------------------------------------------
		// 3. Bass: Sine wave, each chord root for the entire measure.
		//----------------------------------------------------------------------
		{
			name:     "Bass",
			volume:   0.5,
			waveform: WAVE_SINE,
			data: func() string {
				var s strings.Builder
				// We'll just take the first note from each chord’s string (root).
				// E.g. "A2/C3/E3" => "A2" for the measure.
				for _, chord := range chordProgression {
					parts := strings.Split(chord, "/")
					// parts[0] is the root, like "A2"
					s.WriteString(parts[0] + " 4, ")
				}
				return s.String()
			}(),
			attack:  0.02,
			decay:   0.10,
			sustain: 0.70,
			release: 0.25,
		},
		//----------------------------------------------------------------------
		// 4. Arpeggio: A repeating 4-note pattern, subdividing each measure.
		//----------------------------------------------------------------------
		// 4. Arpeggio: A more varied repeating pattern, to avoid incessant repetition.
		{
			name:     "Arp",
			volume:   0.35,
			waveform: WAVE_TRIANGLE,
			data: func() string {
				var s strings.Builder

				// We'll define multiple patterns that each specify an order for the chord notes.
				// For example, if the chord is A2/C3/E3, we'll interpret them as indexes [0,1,2].
				// Some patterns also include an octave-up of the root as a 4th note for variety.
				// We'll apply 8 sub-beats (each 0.5) per measure, but the note *sequence* changes.

				// Patterns are arrays of note indices. Index 3 means "root note, one octave higher."
				// We'll handle the octave-lift in code further down.
				// For example, pattern [0, 1, 2, 3] means:
				//   chord[0], chord[1], chord[2], chord[0]+octaveUp
				// pattern [3, 2, 1, 0] means:
				//   chord[0]+octaveUp, chord[2], chord[1], chord[0]
				// We'll define 6 patterns (feel free to add more).
				patterns := [][]int{
					{0, 1, 2, 3}, // up
					{3, 2, 1, 0}, // down
					{0, 2, 1, 3}, // skip pattern
					{1, 0, 3, 2}, // another skip
					{0, 1, 3, 2}, // partial up then skip
					{2, 3, 1, 0}, // partial down
				}

				// We have 96 measures total in the piece (as in the original example).
				// For each measure, pick a pattern based on (measureIndex mod len(patterns)).
				// That’ll cycle through patterns 0→1→2→3→4→5→0→1, etc.
				for measureIndex, chord := range chordProgression {
					// Parse the chord slash-string into something like ["A2", "C3", "E3"]
					notes := strings.Split(chord, "/")
					if len(notes) < 3 {
						// fallback if chord is malformed for some reason
						notes = []string{"A2", "C3", "E3"}
					}

					// patternIndex cycles through 0..(len(patterns)-1)
					patternIndex := measureIndex % len(patterns)
					pattern := patterns[patternIndex]

					// We'll collect the 4 core notes from chord: root, third, fifth,
					// plus an octave-lift of the root for index=3 if we want it:
					chordNotes := make([]string, 4)
					chordNotes[0] = notes[0]                  // root
					chordNotes[1] = notes[1]                  // 2nd chord tone
					chordNotes[2] = notes[2]                  // 3rd chord tone
					chordNotes[3] = incrementOctave(notes[0]) // root + 1 octave

					// We want 8 sub-beats of 0.5 each, total 4 beats per measure.
					// Let's fill them by applying the pattern's 4 steps, repeated twice.
					// e.g. if pattern = [0,1,2,3], then the final 8 steps are:
					// chordNotes[0], chordNotes[1], chordNotes[2], chordNotes[3], chordNotes[0], ...
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
		//----------------------------------------------------------------------
		// 5. Lead: A melodic theme that changes slightly every 16 measures.
		//----------------------------------------------------------------------
		{
			name:     "Lead",
			volume:   0.40,
			waveform: WAVE_SAW,
			data: func() string {
				// We'll define 6 small melodic phrases (each 16 measures).
				// Each measure has 4 beats, so let's define each phrase as a string array of length 16.
				// Then we’ll concatenate them.
				phrase1 := []string{
					"A3 1, C4 1, E4 1, A4 1, ", // measure 1
					"G4 2, E4 2, ",             // measure 2
					"C4 1, D4 1, E4 1, G4 1, ", // measure 3
					"E4 2, D4 2, ",             // measure 4
					"F4 1, E4 1, D4 1, C4 1, ", // measure 5
					"B3 2, C4 2, ",             // measure 6
					"A3 1, C4 1, D4 1, E4 1, ", // measure 7
					"E4 2, G4 2, ",             // measure 8
					"A3 1, C4 1, E4 1, A4 1, ", // measure 9
					"G4 2, E4 2, ",             // measure 10
					"F4 1, E4 1, D4 1, C4 1, ", // measure 11
					"B3 2, A3 2, ",             // measure 12
					"A3 1, C4 1, D4 1, E4 1, ", // measure 13
					"E4 2, D4 2, ",             // measure 14
					"F4 1, E4 1, D4 1, C4 1, ", // measure 15
					"B3 2, B3 2, ",             // measure 16
				}

				// We'll vary phrase2, phrase3, etc. slightly.
				// For brevity, let's reuse phrase1 with some small edits at the end or middle.
				phrase2 := make([]string, 16)
				copy(phrase2, phrase1)
				phrase2[7] = "E4 2, A4 2, "              // measure 8 changed
				phrase2[15] = "B3 1, C4 1, D4 1, E4 1, " // measure 16 changed
				phrase3 := make([]string, 16)
				copy(phrase3, phrase1)
				phrase3[0] = "A3 2, C4 2, "  // measure 1 changed
				phrase3[15] = "B3 2, A3 2, " // measure 16 changed
				phrase4 := make([]string, 16)
				copy(phrase4, phrase1)
				phrase4[14] = "F4 2, E4 2, " // measure 15 changed
				phrase5 := make([]string, 16)
				copy(phrase5, phrase1)
				phrase5[7] = "G4 1, A4 1, G4 1, E4 1, " // measure 8 changed
				phrase6 := make([]string, 16)
				copy(phrase6, phrase1)
				phrase6[15] = "A4 1, G4 1, E4 1, C4 1, " // measure 16 changed

				// Concatenate them: 6 phrases * 16 measures each = 96 measures
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
		//----------------------------------------------------------------------
		// 6. Kick: 4-on-the-floor for the entire 96 measures.
		//----------------------------------------------------------------------
		{
			name:     "Kick",
			volume:   0.5,
			waveform: WAVE_SINE,
			data: func() string {
				var s strings.Builder
				// Each measure = 4 beats. We’ll do a kick on each beat:
				// C1 0.5, NN 0.5 repeated 4 times per measure => 4 beats
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
		//----------------------------------------------------------------------
		// 7. Snare: Standard backbeat on beats 2 and 4. Let's do it all 96 measures.
		//----------------------------------------------------------------------
		{
			name:   "Snare",
			volume: 0.4,
			data: func() string {
				var s strings.Builder
				// Each measure:
				// Beat 1 (1 beat) => silence
				// Beat 2 => snare
				// Beat 3 => silence
				// Beat 4 => snare
				// So measure = "NN 1, WN 0.25, NN 0.75, NN 1, WN 0.25, NN 0.75"
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
		//----------------------------------------------------------------------
		// 8. Hats: Short white-noise hits on the off-beats ("&" of each beat).
		//----------------------------------------------------------------------
		{
			name:   "Hats",
			volume: 0.3,
			data: func() string {
				var s strings.Builder
				// Each measure = 4 beats. Hats on the “&” => 0.5 offset
				// Then each hat is 0.25, followed by 0.25 silence.
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
		//----------------------------------------------------------------------
		// 9. FX: Periodic noise sweeps every 8 measures for dramatic effect.
		//----------------------------------------------------------------------
		{
			name:   "FX",
			volume: 0.25,
			data: func() string {
				var s strings.Builder
				// We have 96 measures total. Let's do a 1-beat sweep near the end of
				// every 8th measure. That means 96 / 8 = 12 sweeps.
				// Each measure = 4 beats => the sweep can occupy the 4th beat, for instance.
				// Implementation: For each block of 8 measures, we have 32 beats. We'll fill
				// the first 31 beats with silence, then 1 beat of WN.
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
	// Find the last character or two if it’s a digit or digits.
	// E.g., "A2" => base "A", octave "2".
	// We won't handle multi-digit octaves beyond 9 for simplicity.
	if len(note) < 2 {
		return note
	}
	lastChar := note[len(note)-1]
	if lastChar < '0' || lastChar > '9' {
		// no digit at end => just return the note
		return note
	}
	// We have a digit at the end
	octave := int(lastChar - '0')
	octaveUp := octave + 1
	// Rebuild the note but increment the octave
	return note[:len(note)-1] + fmt.Sprintf("%d", octaveUp)
}

var epicWarOfTheAncients = songData{
	name:     "Epic War of the Ancients - Battle Hymn",
	bpm:      100, // 32 measures of 4/4 at 100 BPM => 128 total beats
	reverb:   0.50,
	delay:    0.2,
	feedback: 0.5,
	ins: []insData{
		//---------------------------------------------------------------------
		// 1. Strings: Sawtooth-based chordal pad in D minor / related chords.
		//    Four sections, each 8 measures (for a total of 32).
		//---------------------------------------------------------------------
		{
			name:     "Strings",
			volume:   0.6,
			waveform: WAVE_SAW,
			// Chords laid out in sections of 8 measures each:
			// Section 1 (Measures 1-8):
			//   Dm - Dm - C - C - Bb - F - Gm - A
			// Section 2 (Measures 9-16):
			//   Dm - Dm - F - F - Gm - Gm - A - A
			// Section 3 (Measures 17-24):
			//   Dm - Dm - C - C - Bb - Gm - A - A
			// Section 4 (Measures 25-32):
			//   Dm - C - Bb - F - Gm - A - Dm - Dm
			data: "" +
				// Section 1
				"D3/F3/A3 4, D3/F3/A3 4, C3/E3/G3 4, C3/E3/G3 4, Bb2/D3/F3 4, F2/A2/C3 4, G2/Bb2/D3 4, A2/C#3/E3 4," +
				// Section 2
				"D3/F3/A3 4, D3/F3/A3 4, F2/A2/C3 4, F2/A2/C3 4, G2/Bb2/D3 4, G2/Bb2/D3 4, A2/C#3/E3 4, A2/C#3/E3 4," +
				// Section 3
				"D3/F3/A3 4, D3/F3/A3 4, C3/E3/G3 4, C3/E3/G3 4, Bb2/D3/F3 4, G2/Bb2/D3 4, A2/C#3/E3 4, A2/C#3/E3 4," +
				// Section 4
				"D3/F3/A3 4, C3/E3/G3 4, Bb2/D3/F3 4, F2/A2/C3 4, G2/Bb2/D3 4, A2/C#3/E3 4, D3/F3/A3 4, D3/F3/A3 4",
			attack:  0.08,
			decay:   0.10,
			sustain: 0.80,
			release: 0.25,
		},
		//---------------------------------------------------------------------
		// 2. Brass: Square wave for a bold, heroic backing line (same chords).
		//---------------------------------------------------------------------
		{
			name:     "Brass",
			volume:   0.55,
			waveform: WAVE_SQUARE,
			// Matches the chord progression from Strings, but we’ll sustain each chord for half a measure (2 beats) twice per measure.
			data: "" +
				// Section 1
				"D3/F3/A3 2, NN 2, D3/F3/A3 2, NN 2, C3/E3/G3 2, NN 2, C3/E3/G3 2, NN 2, Bb2/D3/F3 2, NN 2, F2/A2/C3 2, NN 2, G2/Bb2/D3 2, NN 2, A2/C#3/E3 2, NN 2," +
				// Section 2
				"D3/F3/A3 2, NN 2, D3/F3/A3 2, NN 2, F2/A2/C3 2, NN 2, F2/A2/C3 2, NN 2, G2/Bb2/D3 2, NN 2, G2/Bb2/D3 2, NN 2, A2/C#3/E3 2, NN 2, A2/C#3/E3 2, NN 2," +
				// Section 3
				"D3/F3/A3 2, NN 2, D3/F3/A3 2, NN 2, C3/E3/G3 2, NN 2, C3/E3/G3 2, NN 2, Bb2/D3/F3 2, NN 2, G2/Bb2/D3 2, NN 2, A2/C#3/E3 2, NN 2, A2/C#3/E3 2, NN 2," +
				// Section 4
				"D3/F3/A3 2, NN 2, C3/E3/G3 2, NN 2, Bb2/D3/F3 2, NN 2, F2/A2/C3 2, NN 2, G2/Bb2/D3 2, NN 2, A2/C#3/E3 2, NN 2, D3/F3/A3 2, NN 2, D3/F3/A3 2, NN 2",
			attack:  0.06,
			decay:   0.15,
			sustain: 0.75,
			release: 0.3,
		},
		//---------------------------------------------------------------------
		// 3. Bass: Simple sine wave root notes, one measure per note.
		//---------------------------------------------------------------------
		{
			name:     "Bass",
			volume:   0.5,
			waveform: WAVE_SINE,
			// Root notes, each for 1 measure (4 beats), matching the chord roots from above.
			data: "" +
				// Section 1
				"D2 4, D2 4, C2 4, C2 4, Bb1 4, F2 4, G1 4, A1 4," +
				// Section 2
				"D2 4, D2 4, F2 4, F2 4, G1 4, G1 4, A1 4, A1 4," +
				// Section 3
				"D2 4, D2 4, C2 4, C2 4, Bb1 4, G1 4, A1 4, A1 4," +
				// Section 4
				"D2 4, C2 4, Bb1 4, F2 4, G1 4, A1 4, D2 4, D2 4",
			attack:  0.02,
			decay:   0.10,
			sustain: 0.70,
			release: 0.20,
		},
		//---------------------------------------------------------------------
		// 4. Lead: A triumphant triangle-wave melody that varies by section.
		//---------------------------------------------------------------------
		{
			name:     "Lead",
			volume:   0.45,
			waveform: WAVE_TRIANGLE,
			// 32 measures of melody, broken into four 8-measure sections.
			// Each measure is 4 beats. Use a mixture of quarter/half notes.
			data: "" +
				//-----------------------------------------------------------------
				// Section 1 (Measures 1-8)
				//-----------------------------------------------------------------
				// Measure 1
				"D4 1, F4 1, A4 1, D5 1, " +
				// Measure 2
				"D5 2, C5 2, " +
				// Measure 3
				"C5 1, A4 1, F4 1, D4 1, " +
				// Measure 4
				"D4 2, F4 2, " +
				// Measure 5
				"E4 1, F4 1, G4 1, A4 1, " +
				// Measure 6
				"A4 2, G4 2, " +
				// Measure 7
				"Bb4 1, A4 1, G4 1, F4 1, " +
				// Measure 8
				"E4 2, D4 2, " +

				//-----------------------------------------------------------------
				// Section 2 (Measures 9-16)
				//-----------------------------------------------------------------
				// Measure 9
				"D4 1, F4 1, A4 1, C5 1, " +
				// Measure 10
				"C5 2, Bb4 2, " +
				// Measure 11
				"C4 1, D4 1, F4 1, A4 1, " +
				// Measure 12
				"A4 2, G4 2, " +
				// Measure 13
				"G4 1, Bb4 1, C5 1, D5 1, " +
				// Measure 14
				"D5 2, C5 2, " +
				// Measure 15
				"E4 1, G4 1, A4 1, C5 1, " +
				// Measure 16
				"A4 2, D5 2, " +

				//-----------------------------------------------------------------
				// Section 3 (Measures 17-24)
				//-----------------------------------------------------------------
				// Measure 17
				"D4 1, F4 1, A4 1, D5 1, " +
				// Measure 18
				"D5 2, C5 2, " +
				// Measure 19
				"C5 1, A4 1, F4 1, D4 1, " +
				// Measure 20
				"D4 2, F4 2, " +
				// Measure 21
				"E4 1, G4 1, A4 1, Bb4 1, " +
				// Measure 22
				"Bb4 2, A4 2, " +
				// Measure 23
				"C5 1, Bb4 1, G4 1, E4 1, " +
				// Measure 24
				"E4 2, D4 2, " +

				//-----------------------------------------------------------------
				// Section 4 (Measures 25-32)
				//-----------------------------------------------------------------
				// Measure 25
				"D4 1, F4 1, A4 1, D5 1, " +
				// Measure 26
				"C5 2, D5 2, " +
				// Measure 27
				"Bb4 1, A4 1, F4 1, D4 1, " +
				// Measure 28
				"D4 2, F4 2, " +
				// Measure 29
				"G4 1, Bb4 1, C5 1, D5 1, " +
				// Measure 30
				"F5 2, E5 2, " +
				// Measure 31
				"D5 1, C5 1, A4 1, F4 1, " +
				// Measure 32
				"D4 2, D4 2",
			attack:  0.10,
			decay:   0.15,
			sustain: 0.85,
			release: 0.40,
		},
		//---------------------------------------------------------------------
		// 5. Kick: A driving pattern—let's do a basic 4-on-the-floor + extra hits.
		//    (Kick on beats 1, the “&” of 2, and beat 3. This is just an example.)
		//---------------------------------------------------------------------
		{
			name:     "Kick",
			volume:   0.5,
			waveform: WAVE_SINE,
			data: func() string {
				var s string
				// 1 measure = 4 beats. We'll place:
				// - Beat 1: Kick 0.5
				// - Beat 2: no kick
				// - “&” of 2: Kick 0.5 (i.e., 2.5 count)
				// - Beat 3: Kick
				// - Beat 4: no kick
				//
				// Each measure:
				//   1:  C1 0.5
				//   1.5: NN 0.5
				//   2:  NN 1  (beat 2)
				//   2.5: C1 0.5
				//   3:  NN 0.5
				//   3.5: C1 0.5
				//   4:  NN 1
				// This sums up to 4 beats total. Let's do that x 32 measures.
				// The simplest approach is to write it out carefully in code.
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
		//---------------------------------------------------------------------
		// 6. Snare: Crisp hits on beats 2 and 4 (classic backbeat).
		//---------------------------------------------------------------------
		{
			name:   "Snare",
			volume: 0.45,
			data: func() string {
				var s string
				// Each measure:
				// - 1 beat silence
				// - beat 2: WN 0.25 + rest 0.75
				// - 1 beat silence
				// - beat 4: WN 0.25 + rest 0.75
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
		//---------------------------------------------------------------------
		// 7. Cymbals: Quick white-noise accents at the start of each measure.
		//---------------------------------------------------------------------
		{
			name:   "Cymbals",
			volume: 0.3,
			data: func() string {
				var s string
				// For each of the 32 measures, we’ll put a short noise burst on beat 1
				// (0.25 beats) and the rest of the measure is silent (3.75).
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
		//---------------------------------------------------------------------
		// 8. Storm: An atmospheric noise bed, swelling every 4 measures.
		//---------------------------------------------------------------------
		{
			name:   "Storm",
			volume: 0.25,
			data: func() string {
				var s string
				// We have 32 measures total. Let’s do 8 cycles of 4 measures each.
				// Each cycle: 8 beats of noise, 8 beats of silence (2 measures each).
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

// A new example song for the same custom Go synthesizer.
// This one is structured into four sections of 8 measures each (total 32 measures).
// The BPM is slower (90), giving a more tranquil vibe.

var twilightReflections = songData{
	name:     "Twilight Reflections - Nocturnal Journey",
	bpm:      90, // 90 BPM → 32 measures of 4/4 is 128 beats total
	reverb:   0.50,
	delay:    0.2,
	feedback: 0.5,
	ins: []insData{
		//---------------------------------------------------------------------
		// 1. Harmony: Smooth chordal pad using a sawtooth for a soft shimmer.
		//---------------------------------------------------------------------
		{
			name:     "Harmony",
			volume:   0.55,
			waveform: WAVE_SAW,
			// Four sections, each 8 measures (each measure = 4 beats).
			// Repeats are shown as pairs to keep it clear, but it’s just a sequence of 32 chords total.
			data: "" +
				// Section 1 (Measures 1-8):
				"E3/G3/B3 4, E3/G3/B3 4, C3/E3/G3 4, C3/E3/G3 4, " +
				"G3/B3/D4 4, G3/B3/D4 4, D3/F#3/A3 4, D3/F#3/A3 4, " +
				// Section 2 (Measures 9-16):
				"E3/G3/B3 4, E3/G3/B3 4, C3/E3/G3 4, C3/E3/G3 4, " +
				"A3/C3/E3 4, A3/C3/E3 4, B2/D#3/F#3 4, B2/D#3/F#3 4, " +
				// Section 3 (Measures 17-24):
				"G3/B3/D4 4, G3/B3/D4 4, B2/D3/F#3 4, B2/D3/F#3 4, " +
				"E3/G3/B3 4, E3/G3/B3 4, C3/E3/G3 4, C3/E3/G3 4, " +
				// Section 4 (Measures 25-32):
				"E3/G3/B3 4, C3/E3/G3 4, G3/B3/D4 4, D3/F#3/A3 4, " +
				"E3/G3/B3 4, C3/E3/G3 4, G3/B3/D4 4, D3/F#3/A3 4",
			attack:  0.05,
			decay:   0.10,
			sustain: 0.80,
			release: 0.25,
		},
		//-------------------------------------------------
		// 2. Bass: A low sine underpinning each chord root.
		//-------------------------------------------------
		{
			name:     "Bass",
			volume:   0.45,
			waveform: WAVE_SINE,
			// One measure = 4 beats. We simply give the root note of each chord for a full measure.
			// 32 measures total, matching the harmony’s chord progression.
			data: "" +
				// Measures 1-8 (Section 1)
				"E2 4, E2 4, C2 4, C2 4, G2 4, G2 4, D2 4, D2 4," +
				// Measures 9-16 (Section 2)
				"E2 4, E2 4, C2 4, C2 4, A2 4, A2 4, B2 4, B2 4," +
				// Measures 17-24 (Section 3)
				"G2 4, G2 4, B2 4, B2 4, E2 4, E2 4, C2 4, C2 4," +
				// Measures 25-32 (Section 4)
				"E2 4, C2 4, G2 4, D2 4, E2 4, C2 4, G2 4, D2 4",
			attack:  0.01,
			decay:   0.10,
			sustain: 0.70,
			release: 0.20,
		},
		//-----------------------------------------------------
		// 3. Melody: A gentle triangle-wave lead, mid-range.
		//-----------------------------------------------------
		{
			name:     "Melody",
			volume:   0.40,
			waveform: WAVE_TRIANGLE,
			// A simple 2-bar motif repeated or varied across sections.
			// Each measure is 4 beats; we’ll use a mix of quarter (1) and half (2) notes.
			data: "" +
				// Section 1 (Measures 1-8)
				// Measures 1-2 (motif A)
				"E4 1, G4 1, B4 1, E5 1,  E5 2, D5 2, " +
				// Measures 3-4 (motif B)
				"C5 1, B4 1, G4 1, E4 1,  E4 2, G4 2, " +
				// Measures 5-6 (motif A repeated)
				"E4 1, G4 1, B4 1, E5 1,  E5 2, D5 2, " +
				// Measures 7-8 (slight variation)
				"C5 1, D5 1, F#4 1, A4 1, A4 2, G4 2, " +

				// Section 2 (Measures 9-16) - altered to match the new chords
				// Measures 9-10
				"E4 1, G4 1, B4 1, C5 1,  C5 2, B4 2, " +
				// Measures 11-12
				"C4 1, E4 1, A4 1, A4 1,  A4 2, G4 2, " +
				// Measures 13-14
				"A4 1, C5 1, E5 1, D#5 1, D#5 2, B4 2, " +
				// Measures 15-16
				"B4 1, D#5 1, F#4 1, B4 1, B4 2, A4 2, " +

				// Section 3 (Measures 17-24)
				// Measures 17-18
				"G4 1, B4 1, D5 1, B4 1,  G4 2, F#4 2, " +
				// Measures 19-20
				"B4 1, D5 1, F#4 1, D5 1, D5 2, B4 2, " +
				// Measures 21-22
				"E4 1, G4 1, B4 1, E5 1,  E5 2, D5 2, " +
				// Measures 23-24
				"C5 1, B4 1, G4 1, E4 1,  E4 2, G4 2, " +

				// Section 4 (Measures 25-32)
				// Measures 25-26
				"E4 1, G4 1, B4 1, E5 1,  E5 2, D5 2, " +
				// Measures 27-28
				"C5 1, D5 1, F#4 1, A4 1, A4 2, G4 2, " +
				// Measures 29-30
				"E4 1, G4 1, B4 1, E5 1,  E5 2, D5 2, " +
				// Measures 31-32
				"C5 1, B4 1, G4 1, D4 1,  D4 2, E4 2",
			attack:  0.12,
			decay:   0.10,
			sustain: 0.85,
			release: 0.40,
		},
		//-----------------------------------------------------------------------
		// 4. Kick: A simple pattern on beats 1 and 3, repeated for 32 measures.
		//-----------------------------------------------------------------------
		{
			name:     "Kick",
			volume:   0.5,
			waveform: WAVE_SINE,
			data: func() string {
				var s string
				// 32 measures, each measure = 4 beats:
				// Kick on beat 1 (C1 0.5) and beat 3 (C1 0.5).
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
		//--------------------------------------------------------------------
		// 5. Snare: White-noise hits on beats 2 and 4, repeated for 32 measures.
		//--------------------------------------------------------------------
		{
			name:   "Snare",
			volume: 0.4,
			data: func() string {
				var s string
				// Each measure: silence for 1 beat, then snare at beat 2, etc.
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
		//---------------------------------------------------------------------
		// 6. Crickets: A gentle, continuous noise that fades in/out in 4-bar loops.
		//---------------------------------------------------------------------
		{
			name:   "Crickets",
			volume: 0.25,
			data: func() string {
				var s string
				// 8 cycles, each 4 bars = 16 beats:
				// 2 bars (8 beats) of noise, then 2 bars (8 beats) of silence, repeated 8 times = 32 bars total.
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
		//---------------------------------------------------------------------
		// 7. Distant Thunder: Occasional short bursts of noise to hint at thunder.
		//---------------------------------------------------------------------
		{
			name:   "Thunder",
			volume: 0.3,
			data: func() string {
				var s string
				// Every 8 measures (32 beats), add a short noise burst of 1 beat, then silence.
				// We'll place 4 thunder hits across the 32 measures.
				// For simplicity, let's do them near the end of each 8-measure section.
				for i := 0; i < 4; i++ {
					// 7 measures (28 beats) of silence + 1 beat of noise in the 8th measure,
					// then 3 beats of silence to fill the measure.
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
	bpm:      80, // 80 BPM → 160 beats (40 measures of 4/4)
	reverb:   0.50,
	delay:    0.2,
	feedback: 0.5,
	ins: []insData{
		// 1. Harmony: Chordal support with a sawtooth waveform for a shimmering texture.
		{
			name:     "Harmony",
			volume:   0.6,
			waveform: WAVE_SAW,
			data: "D4/F4/A4 4, D4/F4/A4 4, D4/F4/A4 4, D4/F4/A4 4, " +
				"D4/F4/A4 4, D4/F4/A4 4, D4/F4/A4 4, D4/F4/A4 4," +
				// Section 2:
				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, A3/C#4/E4 4, " +
				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, A3/C#4/E4 4, " +
				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, A3/C#4/E4 4, " +
				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, A3/C#4/E4 4," +
				// Section 3:
				"G3/Bb3/D4 4, A3/C#4/E4 4, G3/Bb3/D4 4, A3/C#4/E4 4, " +
				"F3/A3/C4 4, G3/Bb3/D4 4, A3/C#4/E4 4, D4/F4/A4 4," +
				// Section 4:
				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, A3/C#4/E4 4, " +
				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, D4/F4/A4 4",
			attack:  0.05,
			decay:   0.10,
			sustain: 0.80,
			release: 0.20,
		},
		// 2. Bass: Deep, rocking pulse rendered with a pure sine wave.
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
		// 3. Melody (Lead): Lowered for a more serious tone, using a triangle wave.
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
			// The WAVE_SQUARE parameter is ignored here because we’re not using "mix".
			square: 0.05,
		},
		// 4. Kick: A deep, punchy hit using a sine wave.
		{
			name:     "Kick",
			volume:   0.5,
			waveform: WAVE_SINE,
			data: func() string {
				var s string
				// Each measure (4 beats): kick on beats 1 and 3.
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
		// 5. Snare: Crisp, quick white-noise hits on beats 2 and 4.
		{
			name:   "Snare",
			volume: 0.4,
			data: func() string {
				var s string
				// Each measure: silence for 1 beat, then a snare hit ("WN 0.25") then rest (0.75), repeated.
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
		// 6. Waves: A re-triggering noise pattern that swells and recedes (noise).
		{
			name:   "Waves",
			volume: 0.2,
			data: func() string {
				var s string
				// 20 cycles of 8 beats each: 4 beats of noise followed by 4 beats of silence.
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
		// 7. Water Splashes: Occasional bursts evoking crashing waves (noise).
		{
			name:   "Water Splashes",
			volume: 0.3,
			data: func() string {
				var s string
				// 10 cycles of 16 beats each: 2 beats of noise then 14 beats of silence.
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
