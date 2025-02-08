package main

type playlistData []songData

type songData struct {
	name                    string
	ins                     []insData
	bpm                     int
	reverb, delay, feedback float32
}

type insData struct {
	name, waveform, data string
	volume, square       float32
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
		twilightReflections,
		voyageOfTheAbyss,
	}

	gameSongList = []songData{}
)

// A new example song for the same custom Go synthesizer.
// This one is structured into four sections of 8 measures each (total 32 measures).
// The BPM is slower (90), giving a more tranquil vibe.

var twilightReflections = songData{
	name: "Twilight Reflections - Nocturnal Journey",
	bpm:  90, // 90 BPM → 32 measures of 4/4 is 128 beats total
	ins: []insData{
		//---------------------------------------------------------------------
		// 1. Harmony: Smooth chordal pad using a sawtooth for a soft shimmer.
		//---------------------------------------------------------------------
		{
			name:     "Harmony",
			volume:   0.55,
			waveform: "sawtooth",
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
			waveform: "sine",
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
			waveform: "triangle",
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
			waveform: "sine",
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
	name:   "Voyage of the Abyss - Nautical Odyssey",
	bpm:    80, // 80 BPM → 160 beats (40 measures of 4/4)
	reverb: 0.0,
	ins: []insData{
		// 1. Harmony: Chordal support with a sawtooth waveform for a shimmering texture.
		{
			name:     "Harmony",
			volume:   0.6,
			waveform: "sawtooth",
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
			waveform: "sine",
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
			waveform: "triangle",
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
			// The "square" parameter is ignored here because we’re not using "mix".
			square: 0.05,
		},
		// 4. Kick: A deep, punchy hit using a sine wave.
		{
			name:     "Kick",
			volume:   0.5,
			waveform: "sine",
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
