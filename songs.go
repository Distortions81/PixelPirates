package main

type playlistData []songData

type songData struct {
	name                    string
	ins                     []insData
	bpm                     int
	reverb, delay, feedback float32
}

type insData struct {
	name, data     string
	volume, square float32
	/*
		Attack: The time it takes for a sound to go from silence to its full volume when a key is first pressed.
		Decay: The time it takes for the sound to drop from its peak volume to the sustain level.
		Sustain: The constant volume level maintained while a key is held down.
		Release: The time it takes for the sound to fade from the sustain level to silence when the key is released.
	*/
	attack, decay, sustain, release float32
}

var (
	gameModePlaylists = [GAME_MAX]playlistData{
		GAME_TITLE: titleSongList,
		GAME_PLAY:  gameSongList,
	}

	titleSongList = []songData{
		voyageOfTheAbyss,
	}

	gameSongList = []songData{
		voyageOfTheAbyss,
	}
)

// An epic two-minute nautical theme called "Voyage of the Abyss"
// BPM: 80  Total beats: 160 (i.e. 40 measures of 4/4)
// Sections:
//  • Section 1 (Measures 1–8): Gentle introduction (all Dm chords)
//  • Section 2 (Measures 9–24): Main theme – a repeating progression (Dm – Bb – C – A)
//  • Section 3 (Measures 25–32): A contrasting bridge (using Gm, A, F, then Dm)
//  • Section 4 (Measures 33–40): Recapitulation of the main theme with a final cadence

var voyageOfTheAbyss = songData{
	name: "Voyage of the Abyss",
	bpm:  80,
	ins: []insData{
		{
			name:   "Harmony",
			volume: 0.6,
			// For each measure the chord lasts 4 beats.
			// Section 1: 8 measures of Dm
			// Section 2: 16 measures – repeating [Dm, Bb, C, A] four times
			// Section 3: 8 measures – [Gm, A, Gm, A, F, Gm, A, Dm]
			// Section 4: 8 measures – [Dm, Bb, C, A, Dm, Bb, C, Dm]
			data: "D4/F4/A4 4, D4/F4/A4 4, D4/F4/A4 4, D4/F4/A4 4, D4/F4/A4 4, D4/F4/A4 4, D4/F4/A4 4, D4/F4/A4 4," +
				// Section 2:
				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, A3/C#4/E4 4, " +
				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, A3/C#4/E4 4, " +
				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, A3/C#4/E4 4, " +
				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, A3/C#4/E4 4," +
				// Section 3:
				"G3/Bb3/D4 4, A3/C#4/E4 4, G3/Bb3/D4 4, A3/C#4/E4 4, F3/A3/C4 4, G3/Bb3/D4 4, A3/C#4/E4 4, D4/F4/A4 4," +
				// Section 4:
				"D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, A3/C#4/E4 4, D4/F4/A4 4, Bb3/D4/F4 4, C4/E4/G4 4, D4/F4/A4 4",
			attack:  0.05,
			decay:   0.10,
			sustain: 0.80,
			release: 0.20,
		},
		{
			name:   "Bass",
			volume: 0.5,
			// The bass underpins the chords:
			// Section 1: 8 measures of two-note patterns (D2 then A2 per measure)
			// Section 2: 16 measures – one bass note per measure matching the chord roots:
			//  Dm: D2, Bb: Bb1, C: C2, A: A1 (repeated 4×)
			// Section 3: 8 measures – [G2, A1, G2, A1, F2, G2, A1, D2]
			// Section 4: 8 measures – [Dm: D2, Bb: Bb1, C: C2, A: A1, D2, Bb1, C2, D2]
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
			name:   "Melody",
			volume: 0.8,
			// The melody unfolds its tale over 40 measures.
			// Section 1 (Measures 1–8): A rising and falling arpeggio and gentle motif.
			//  M1: D4 1, F4 1, A4 1, D5 1
			//  M2: A4 1, F4 1, D4 1, A3 1
			//  M3: D4 0.5, E4 0.5, F4 1, G4 1, A4 1
			//  M4: A4 0.5, G4 0.5, F4 0.5, E4 0.5, D4 1, D4 1
			//  M5: D4 1, E4 1, F4 1, G4 1
			//  M6: G4 0.5, A4 0.5, Bb4 1, A4 1, G4 1
			//  M7: F4 1, E4 1, D4 1, C4 1
			//  M8: D4 2, D4 2
			//
			// Section 2 (Measures 9–24): The main theme – a buoyant, memorable motif
			//  (M9)  Dm:   F4 0.5, A4 0.5, D5 1, C5 1, A4 1
			//  (M10) Bb:   Bb4 0.5, D5 0.5, F5 1, D5 1, Bb4 1
			//  (M11) C:    C5 1, E5 1, G5 1, E5 1
			//  (M12) A:    C#5 0.5, E5 0.5, A5 1, G5 1, E5 1
			//  (M13) Dm Variation: D5 0.5, F5 0.5, A5 1, F5 1, D5 1
			//  (M14) Bb:   Bb4 0.5, D5 0.5, F5 1, D5 1, Bb4 1
			//  (M15) C:    C5 1, E5 1, G5 1, E5 1
			//  (M16) A:    A4 0.5, C#5 0.5, E5 1, C#5 1, A4 1
			//  (M17–24) – Repeat a similar eight–measure pattern (with slight variations)
			// Section 3 (Measures 25–32): A contrasting bridge
			//  M25: Gm:  G4 1, Bb4 1, D5 1, Bb4 1
			//  M26: A:   A4 1, C#5 1, E5 1, C#5 1
			//  M27: Gm:  F4 0.5, G4 0.5, Bb4 1, D5 1, Bb4 1
			//  M28: A:   A4 1, C#5 1, E5 1, A4 1
			//  M29: F:   F4 1, A4 1, C5 1, A4 1
			//  M30: Gm:  G4 1, Bb4 1, D5 1, Bb4 1
			//  M31: A:   A4 0.5, C#5 0.5, E5 1, C#5 1, A4 1
			//  M32: Dm:  D4 1, F4 1, A4 1, F4 1
			// Section 4 (Measures 33–40): Return of the theme with a final cadence
			//  M33: Dm:  F4 0.5, A4 0.5, D5 1, C5 1, A4 1
			//  M34: Bb:  Bb4 0.5, D5 0.5, F5 1, D5 1, Bb4 1
			//  M35: C:   C5 1, E5 1, G5 1, E5 1
			//  M36: A:   C#5 0.5, E5 0.5, A5 1, G5 1, E5 1
			//  M37: Dm:  D5 0.5, F5 0.5, A5 1, F5 1, D5 1
			//  M38: Bb:  Bb4 0.5, D5 0.5, F5 1, D5 1, Bb4 1
			//  M39: C:   C5 1, E5 1, G5 1, E5 1
			//  M40: Final cadence: D4 2, F4 2
			data: "D4 1, F4 1, A4 1, D5 1, " +
				"A4 1, F4 1, D4 1, A3 1, " +
				"D4 0.5, E4 0.5, F4 1, G4 1, A4 1, " +
				"A4 0.5, G4 0.5, F4 0.5, E4 0.5, D4 1, D4 1, " +
				"D4 1, E4 1, F4 1, G4 1, " +
				"G4 0.5, A4 0.5, Bb4 1, A4 1, G4 1, " +
				"F4 1, E4 1, D4 1, C4 1, " +
				"D4 2, D4 2," +
				// Section 2 (measures 9–16)
				"F4 0.5, A4 0.5, D5 1, C5 1, A4 1, " +
				"Bb4 0.5, D5 0.5, F5 1, D5 1, Bb4 1, " +
				"C5 1, E5 1, G5 1, E5 1, " +
				"C#5 0.5, E5 0.5, A5 1, G5 1, E5 1, " +
				"D5 0.5, F5 0.5, A5 1, F5 1, D5 1, " +
				"Bb4 0.5, D5 0.5, F5 1, D5 1, Bb4 1, " +
				"C5 1, E5 1, G5 1, E5 1, " +
				"A4 0.5, C#5 0.5, E5 1, C#5 1, A4 1, " +
				// Section 2 (measures 17–24) – repeat similar motifs
				"F4 0.5, A4 0.5, D5 1, C5 1, A4 1, " +
				"Bb4 0.5, D5 0.5, F5 1, D5 1, Bb4 1, " +
				"C5 1, E5 1, G5 1, E5 1, " +
				"C#5 0.5, E5 0.5, A5 1, G5 1, E5 1, " +
				"D5 0.5, F5 0.5, A5 1, F5 1, D5 1, " +
				"Bb4 0.5, D5 0.5, F5 1, D5 1, Bb4 1, " +
				"C5 1, E5 1, G5 1, E5 1, " +
				"A4 0.5, C#5 0.5, E5 1, C#5 1, A4 1," +
				// Section 3 (measures 25–32)
				"G4 1, Bb4 1, D5 1, Bb4 1, " +
				"A4 1, C#5 1, E5 1, C#5 1, " +
				"F4 0.5, G4 0.5, Bb4 1, D5 1, Bb4 1, " +
				"A4 1, C#5 1, E5 1, A4 1, " +
				"F4 1, A4 1, C5 1, A4 1, " +
				"G4 1, Bb4 1, D5 1, Bb4 1, " +
				"A4 0.5, C#5 0.5, E5 1, C#5 1, A4 1, " +
				"D4 1, F4 1, A4 1, F4 1," +
				// Section 4 (measures 33–40)
				"F4 0.5, A4 0.5, D5 1, C5 1, A4 1, " +
				"Bb4 0.5, D5 0.5, F5 1, D5 1, Bb4 1, " +
				"C5 1, E5 1, G5 1, E5 1, " +
				"C#5 0.5, E5 0.5, A5 1, G5 1, E5 1, " +
				"D5 0.5, F5 0.5, A5 1, F5 1, D5 1, " +
				"Bb4 0.5, D5 0.5, F5 1, D5 1, Bb4 1, " +
				"C5 1, E5 1, G5 1, E5 1, " +
				"D4 2, F4 2",
			attack:  0.01,
			decay:   0.05,
			sustain: 0.90,
			release: 0.30,
			square:  0.2,
		},
	},
}
