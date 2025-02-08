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
		voyageOfTheAbyss,
	}

	gameSongList = []songData{}
)

var voyageOfTheAbyss = songData{
	name: "Voyage of the Abyss - Nautical Odyssey",
	bpm:  80, // 80 BPM → 160 beats (40 measures of 4/4)
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
