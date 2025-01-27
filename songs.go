package main

type songData struct {
	name                    string
	ins                     []insData
	bpm                     int
	reverb, delay, feedback float64
}

type insData struct {
	name, data     string
	volume, square float64
	/*
		Attack: The time it takes for a sound to go from silence to its full volume when a key is first pressed.
		Decay: The time it takes for the sound to drop from its peak volume to the sustain level.
		Sustain: The constant volume level maintained while a key is held down.
		Release: The time it takes for the sound to fade from the sustain level to silence when the key is released.
	*/
	attack, decay, sustain, release float64
}

var songList = []songData{
	EbonyGale,
	MidnightDepthsPercussion,

	introspectiveTheme,
	MerrySailingTheme,
	PirateDramaticTheme,
	PirateEpicTheme,
}

var introspectiveTheme = songData{
	name:     "Introspective Reflection",
	bpm:      70,
	reverb:   0.5, // moderate, for spacious introspection
	delay:    0.3,
	feedback: 0.3,
	ins: []insData{
		{
			// 1) Soft Piano (Chords)
			name:   "Soft Piano (Chords)",
			volume: 0.6,
			square: 0.1, // mostly sine-like but a bit of "bite"
			// Longer attack/decay for a gentle swell, moderate sustain/release
			attack:  0.5,
			decay:   0.5,
			sustain: 0.7,
			release: 1.0,
			// One 16-measure chord progression:
			//   i (Eb minor) -> VI (Cb major) -> iv (Ab minor) -> i (Eb minor)
			// Each chord is 4 measures, total 16. Then repeat 2 more times => 48 measures (~2:45 at 70 BPM).
			data: `
////////////////////////////////////////////////////////////////////////
// 16-MEASURE CYCLE (4 chords × 4 measures each), REPEATED 3x
////////////////////////////////////////////////////////////////////////

// MEASURES 1-4: Eb minor (Eb, Gb, Bb)
Eb3 1, Gb3 1, Bb3 1, Eb4 1,
Gb3 1, Bb3 1, Eb4 1, Gb4 1,
Bb3 1, Eb4 1, Gb4 1, Bb4 1,
Eb4 1, Gb4 1, Bb4 1, Eb5 1,

// MEASURES 5-8: Cb major (Cb, Eb, Gb)
Cb3 1, Eb3 1, Gb3 1, Cb4 1,
Eb3 1, Gb3 1, Cb4 1, Eb4 1,
Gb3 1, Cb4 1, Eb4 1, Gb4 1,
Cb4 1, Eb4 1, Gb4 1, Cb5 1,

// MEASURES 9-12: Ab minor (Ab, Cb, Eb)
Ab3 1, Cb4 1, Eb4 1, Ab4 1,
Cb4 1, Eb4 1, Ab4 1, Cb5 1,
Eb4 1, Ab4 1, Cb5 1, Eb5 1,
Ab4 1, Cb5 1, Eb5 1, Ab5 1,

// MEASURES 13-16: Eb minor again
Eb3 1, Gb3 1, Bb3 1, Eb4 1,
Gb3 1, Bb3 1, Eb4 1, Gb4 1,
Bb3 1, Eb4 1, Gb4 1, Bb4 1,
Eb4 1, Gb4 1, Bb4 1, Eb5 1,

////////////////////////////////////////////////////////////////////////
// REPEAT THE ABOVE 16 MEASURES 2 MORE TIMES => 48 TOTAL
////////////////////////////////////////////////////////////////////////

// Second pass
Eb3 1, Gb3 1, Bb3 1, Eb4 1,
Gb3 1, Bb3 1, Eb4 1, Gb4 1,
Bb3 1, Eb4 1, Gb4 1, Bb4 1,
Eb4 1, Gb4 1, Bb4 1, Eb5 1,

Cb3 1, Eb3 1, Gb3 1, Cb4 1,
Eb3 1, Gb3 1, Cb4 1, Eb4 1,
Gb3 1, Cb4 1, Eb4 1, Gb4 1,
Cb4 1, Eb4 1, Gb4 1, Cb5 1,

Ab3 1, Cb4 1, Eb4 1, Ab4 1,
Cb4 1, Eb4 1, Ab4 1, Cb5 1,
Eb4 1, Ab4 1, Cb5 1, Eb5 1,
Ab4 1, Cb5 1, Eb5 1, Ab5 1,

Eb3 1, Gb3 1, Bb3 1, Eb4 1,
Gb3 1, Bb3 1, Eb4 1, Gb4 1,
Bb3 1, Eb4 1, Gb4 1, Bb4 1,
Eb4 1, Gb4 1, Bb4 1, Eb5 1,

// Third pass
Eb3 1, Gb3 1, Bb3 1, Eb4 1,
Gb3 1, Bb3 1, Eb4 1, Gb4 1,
Bb3 1, Eb4 1, Gb4 1, Bb4 1,
Eb4 1, Gb4 1, Bb4 1, Eb5 1,

Cb3 1, Eb3 1, Gb3 1, Cb4 1,
Eb3 1, Gb3 1, Cb4 1, Eb4 1,
Gb3 1, Cb4 1, Eb4 1, Gb4 1,
Cb4 1, Eb4 1, Gb4 1, Cb5 1,

Ab3 1, Cb4 1, Eb4 1, Ab4 1,
Cb4 1, Eb4 1, Ab4 1, Cb5 1,
Eb4 1, Ab4 1, Cb5 1, Eb5 1,
Ab4 1, Cb5 1, Eb5 1, Ab5 1,

Eb3 1, Gb3 1, Bb3 1, Eb4 1,
Gb3 1, Bb3 1, Eb4 1, Gb4 1,
Bb3 1, Eb4 1, Gb4 1, Bb4 1,
Eb4 1, Gb4 1, Bb4 1, Eb5 1,
`,
		},
		{
			// 2) Distant Strings (Melody or sustained lines)
			name:   "Distant Strings",
			volume: 0.5,
			square: 0.0, // pure sine-like, smooth/soft
			// Slow, gentle envelope
			attack:  1.0,
			decay:   0.5,
			sustain: 0.7,
			release: 1.5,
			// We'll use half notes (2 beats) to create a slow, reflective melody.
			// Each measure has 2 half notes. The 16-measure phrase is repeated 3x.
			data: `
////////////////////////////////////////////////////////////////////////
// 16 MEASURES (2 half notes each), REPEATED 3x
////////////////////////////////////////////////////////////////////////

// MEASURES 1-4 (over Eb minor)
Eb4 2, Gb4 2,
Bb4 2, Gb4 2,
Eb4 2, Db4 2,
Bb3 2, Eb4 2,

// MEASURES 5-8 (over Cb major)
Cb4 2, Eb4 2,
Gb4 2, Eb4 2,
Cb4 2, Bb3 2,
Ab3 2, Eb4 2,

// MEASURES 9-12 (over Ab minor)
Ab3 2, Cb4 2,
Eb4 2, Cb4 2,
Ab4 2, Gb4 2,
Eb4 2, Cb4 2,

// MEASURES 13-16 (over Eb minor)
Bb3 2, Eb4 2,
Gb4 2, Eb4 2,
Bb4 2, Gb4 2,
Eb4 2, Eb3 2,

////////////////////////////////////////////////////////////////////////
// REPEAT THOSE 16 MEASURES 2 MORE TIMES => 48 TOTAL
////////////////////////////////////////////////////////////////////////

// second pass
Eb4 2, Gb4 2,
Bb4 2, Gb4 2,
Eb4 2, Db4 2,
Bb3 2, Eb4 2,

Cb4 2, Eb4 2,
Gb4 2, Eb4 2,
Cb4 2, Bb3 2,
Ab3 2, Eb4 2,

Ab3 2, Cb4 2,
Eb4 2, Cb4 2,
Ab4 2, Gb4 2,
Eb4 2, Cb4 2,

Bb3 2, Eb4 2,
Gb4 2, Eb4 2,
Bb4 2, Gb4 2,
Eb4 2, Eb3 2,

// third pass
Eb4 2, Gb4 2,
Bb4 2, Gb4 2,
Eb4 2, Db4 2,
Bb3 2, Eb4 2,

Cb4 2, Eb4 2,
Gb4 2, Eb4 2,
Cb4 2, Bb3 2,
Ab3 2, Eb4 2,

Ab3 2, Cb4 2,
Eb4 2, Cb4 2,
Ab4 2, Gb4 2,
Eb4 2, Cb4 2,

Bb3 2, Eb4 2,
Gb4 2, Eb4 2,
Bb4 2, Gb4 2,
Eb4 2, Eb3 2,
`,
		},
	},
}

var EbonyGale = songData{
	name:     "Aboard the Ebony Gale",
	bpm:      70,
	reverb:   0.4,  // Slightly cavernous
	delay:    0.25, // 250ms delay
	feedback: 0.3,  // Mild echo repeats
	ins: []insData{
		{
			// Bright, prominent lead
			name:    "lead",
			volume:  0.8,
			square:  0.2, // Mostly sine with a bit of buzz
			attack:  0.1,
			decay:   0.3,
			sustain: 0.7,
			release: 0.4,
			data: `
Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Gb4 2,
Fb4 1, Gb4 1, Ab4 2,
Bb4 2, Ab4 2,

Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Db5 2,
Cb5 1, Bb4 1, Ab4 2,
Gb4 4
`,
		},
		{
			// Gentle accompaniment chords
			name:    "harmony",
			volume:  0.6,
			square:  0.1, // Softer timbre
			attack:  0.2,
			decay:   0.3,
			sustain: 0.6,
			release: 0.5,
			data: `
Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Ab4/Db5 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Gb4/Bb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Ab4/Db5 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Gb4/Bb4 4,
Eb4/Gb4/Bb4 4
`,
		},
		{
			// Steady low-end anchor
			name:    "bass",
			volume:  0.7,
			square:  0.3, // Slightly more edge in the low range
			attack:  0.05,
			decay:   0.2,
			sustain: 0.7,
			release: 0.3,
			data: `
Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Db2 4,
Eb2 4
`,
		},
		{
			// Simple percussive hits (kick/snare imitation)
			name:    "percussion",
			volume:  0.9,
			square:  1.0, // Pure square for a sharper click
			attack:  0.0,
			decay:   0.1,
			sustain: 0.0,
			release: 0.05,
			data: `
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,

Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75
`,
		},
	},
}

var MidnightDepthsPercussion = songData{
	name:     "Midnight Depths",
	bpm:      70,
	reverb:   0.6, // Larger reverb for a darker, echoic atmosphere
	delay:    0.3, // 300ms delay
	feedback: 0.4, // Enough feedback for an eerie repeat
	ins: []insData{
		{
			name:    "lead",
			volume:  0.8,
			square:  0.2, // Mostly sine, slight edge
			attack:  0.2, // Mild swell
			decay:   0.3,
			sustain: 0.7,
			release: 0.5,
			data: `
Eb4 2, NN 2,
Gb4 1, Ab4 1, Bb4 2,
NN 2, Db5 2,
Bb4 1, Ab4 1, Gb4 2,

Eb4 2, NN 2,
Fb4 1, Gb4 1, Ab4 2,
NN 2, Db5 2,
Bb4 2, Ab4 2,

Gb4 1, Ab4 1, Bb4 2,
Db5 2, Bb4 2,
Eb5 2, NN 2,
Db5 1, Bb4 1, Ab4 2
`,
		},
		{
			name:    "harmony",
			volume:  0.6,
			square:  0.1, // Softer timbre
			attack:  0.3,
			decay:   0.4,
			sustain: 0.6,
			release: 0.5,
			data: `
Eb4/Gb4/Bb4 4,
Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,

Eb4/Gb4/Bb4 4,
Db4/Gb4/Bb4 4,
Ab3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Db4/Gb4/Bb4 4,
Ab3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4
`,
		},
		{
			name:    "bass",
			volume:  0.6,
			square:  0.3, // More bite in low range
			attack:  0.1,
			decay:   0.2,
			sustain: 0.7,
			release: 0.3,
			data: `
Eb2 4,
Eb2 4,
Cb2 4,
Bb2 4,

Eb2 4,
Db2 4,
Ab2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Bb2 4,
Eb2 4,

Eb2 4,
Db2 4,
Ab2 4,
Eb2 4
`,
		},
		{
			name:    "descant",
			volume:  0.7,
			square:  0.0, // Pure sine, ethereal
			attack:  0.4,
			decay:   0.2,
			sustain: 0.8,
			release: 0.6,
			data: `
NN 8, // (rests for first 8 measures)

Gb5 0.5, Ab5 0.5, Bb5 1, NN 2.5,
Db6 1, Bb5 1, NN 2,
Ab5 0.5, Gb5 0.5, Eb5 1, NN 2.5
`,
		},
		{
			name:    "pad",
			volume:  0.5,
			square:  0.0, // Warm, pad-like
			attack:  1.0, // Long swell
			decay:   0.5,
			sustain: 0.7,
			release: 1.5,
			data: `
Eb4/Gb4/Bb4 0.5, NN 0.5, Eb4/Gb4/Bb4 0.5, NN 2.5,
Cb4/Eb4/Gb4 1, NN 3,
Bb3/Db4/Fb4 1, NN 3,
Eb4/Gb4/Bb4 0.5, NN 3.5,

Eb4/Gb4/Bb4 0.5, NN 0.5, Eb4/Gb4/Bb4 0.5, NN 2.5,
Db4/Gb4/Bb4 1, NN 3,
Ab3/Db4/Fb4 1, NN 3,
Eb4/Gb4/Bb4 0.5, NN 3.5,

Eb4/Gb4/Bb4 0.5, NN 3.5,
Cb4/Eb4/Gb4 0.5, NN 3.5,
Bb3/Db4/Fb4 0.5, NN 3.5,
Eb4/Gb4/Bb4 0.5, NN 3.5
`,
		},
		{
			name:    "percussion",
			volume:  0.9,
			square:  1.0, // Sharp, clicky
			attack:  0.0,
			decay:   0.05,
			sustain: 0.0,
			release: 0.1,
			data: `
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,

Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75
`,
		},
	},
}

var PirateEpicTheme = songData{
	name:     "Epic Pirate Saga",
	bpm:      100, // ~64 measures => ~2.5 minutes
	reverb:   0.6,
	delay:    0.3,
	feedback: 0.4,
	ins: []insData{
		{
			// --------------------
			// 1) CHORDS INSTRUMENT
			// --------------------
			name: "Chords",
			// We'll stitch together 4 "sections" of 16 measures each = 64 total.
			// Each measure has 4 beats. BPM=100 => 1 beat = 0.6s, so 1 measure = ~2.4s
			// 64 measures => ~153.6s => ~2m34s total.
			//
			// A rough chord progression in E-flat minor territory:
			//   Section A (16 measures):  Ebm → Gb → Db → Abm
			//   Section B (16 measures):  Ebm → Db → Gb → Bbm
			//   Section C (16 measures):  Ebm → Abm → Bbm → Db
			//   Section D (16 measures):  Gb → Db → Abm → Ebm
			//
			// Each chord is held for 2 measures (8 measures total per “progression”),
			// repeated for 16 in each section. Then we string all sections together.
			//
			// For clarity, we’ll do 2 measures per chord, with a small arpeggio pattern:
			//   - measure 1: (Eb4, Gb4, Bb4, Eb5)
			//   - measure 2: (Gb4, Bb4, Eb5, Gb5)
			// etc.
			//
			// This big string *looks* repetitive, but is valid for a roughly 2.5 min track.
			data: `
////////////////////////////////////////////////////////////////////////
// SECTION A (16 measures): Ebm -> Gb -> Db -> Abm
////////////////////////////////////////////////////////////////////////
// Eb minor (measures 1-2)
Eb4 1, Gb4 1, Bb4 1, Eb5 1,
Gb4 1, Bb4 1, Eb5 1, Gb5 1,
// Eb minor repeated (measures 3-4)
Eb4 1, Gb4 1, Bb4 1, Eb5 1,
Gb4 1, Bb4 1, Eb5 1, Gb5 1,

// Gb major (measures 5-6)
Gb4 1, Bb4 1, Db5 1, Gb5 1,
Bb4 1, Db5 1, Gb5 1, Bb5 1,
// Gb major repeated (measures 7-8)
Gb4 1, Bb4 1, Db5 1, Gb5 1,
Bb4 1, Db5 1, Gb5 1, Bb5 1,

// Db major (enharmonic with E# as F) (measures 9-10)
Db4 1, E#4 1, Ab4 1, Db5 1,
E#4 1, Ab4 1, Db5 1, E#5 1,
// Db repeated (measures 11-12)
Db4 1, E#4 1, Ab4 1, Db5 1,
E#4 1, Ab4 1, Db5 1, E#5 1,

// Ab minor (Ab, Cb, Eb) (measures 13-14)
Ab3 1, Cb4 1, Eb4 1, Ab4 1,
Cb4 1, Eb4 1, Ab4 1, Cb5 1,
// Ab minor repeated (measures 15-16)
Ab3 1, Cb4 1, Eb4 1, Ab4 1,
Cb4 1, Eb4 1, Ab4 1, Cb5 1,

////////////////////////////////////////////////////////////////////////
// SECTION B (16 measures): Ebm -> Db -> Gb -> Bbm
////////////////////////////////////////////////////////////////////////
// Eb minor (measures 17-18)
Eb4 1, Gb4 1, Bb4 1, Eb5 1,
Gb4 1, Bb4 1, Eb5 1, Gb5 1,
// Eb minor repeated (measures 19-20)
Eb4 1, Gb4 1, Bb4 1, Eb5 1,
Gb4 1, Bb4 1, Eb5 1, Gb5 1,

// Db major (measures 21-22)
Db4 1, E#4 1, Ab4 1, Db5 1,
E#4 1, Ab4 1, Db5 1, E#5 1,
// Db repeated (measures 23-24)
Db4 1, E#4 1, Ab4 1, Db5 1,
E#4 1, Ab4 1, Db5 1, E#5 1,

// Gb major (measures 25-26)
Gb4 1, Bb4 1, Db5 1, Gb5 1,
Bb4 1, Db5 1, Gb5 1, Bb5 1,
// Gb repeated (measures 27-28)
Gb4 1, Bb4 1, Db5 1, Gb5 1,
Bb4 1, Db5 1, Gb5 1, Bb5 1,

// Bb minor (Bb, Db, E#) (measures 29-30)
Bb3 1, Db4 1, E#4 1, Bb4 1,
Db4 1, E#4 1, Bb4 1, Db5 1,
// Bbm repeated (measures 31-32)
Bb3 1, Db4 1, E#4 1, Bb4 1,
Db4 1, E#4 1, Bb4 1, Db5 1,

////////////////////////////////////////////////////////////////////////
// SECTION C (16 measures): Ebm -> Abm -> Bbm -> Db
////////////////////////////////////////////////////////////////////////
// Eb minor (33-34)
Eb4 1, Gb4 1, Bb4 1, Eb5 1,
Gb4 1, Bb4 1, Eb5 1, Gb5 1,
// Eb minor repeated (35-36)
Eb4 1, Gb4 1, Bb4 1, Eb5 1,
Gb4 1, Bb4 1, Eb5 1, Gb5 1,

// Ab minor (37-38)
Ab3 1, Cb4 1, Eb4 1, Ab4 1,
Cb4 1, Eb4 1, Ab4 1, Cb5 1,
// Ab minor repeated (39-40)
Ab3 1, Cb4 1, Eb4 1, Ab4 1,
Cb4 1, Eb4 1, Ab4 1, Cb5 1,

// Bb minor (41-42)
Bb3 1, Db4 1, E#4 1, Bb4 1,
Db4 1, E#4 1, Bb4 1, Db5 1,
// Bbm repeated (43-44)
Bb3 1, Db4 1, E#4 1, Bb4 1,
Db4 1, E#4 1, Bb4 1, Db5 1,

// Db major (45-46)
Db4 1, E#4 1, Ab4 1, Db5 1,
E#4 1, Ab4 1, Db5 1, E#5 1,
// Db repeated (47-48)
Db4 1, E#4 1, Ab4 1, Db5 1,
E#4 1, Ab4 1, Db5 1, E#5 1,

////////////////////////////////////////////////////////////////////////
// SECTION D (16 measures): Gb -> Db -> Abm -> Ebm
////////////////////////////////////////////////////////////////////////
// Gb major (49-50)
Gb4 1, Bb4 1, Db5 1, Gb5 1,
Bb4 1, Db5 1, Gb5 1, Bb5 1,
// Gb repeated (51-52)
Gb4 1, Bb4 1, Db5 1, Gb5 1,
Bb4 1, Db5 1, Gb5 1, Bb5 1,

// Db major (53-54)
Db4 1, E#4 1, Ab4 1, Db5 1,
E#4 1, Ab4 1, Db5 1, E#5 1,
// Db repeated (55-56)
Db4 1, E#4 1, Ab4 1, Db5 1,
E#4 1, Ab4 1, Db5 1, E#5 1,

// Ab minor (57-58)
Ab3 1, Cb4 1, Eb4 1, Ab4 1,
Cb4 1, Eb4 1, Ab4 1, Cb5 1,
// Abm repeated (59-60)
Ab3 1, Cb4 1, Eb4 1, Ab4 1,
Cb4 1, Eb4 1, Ab4 1, Cb5 1,

// Eb minor (61-62)
Eb4 1, Gb4 1, Bb4 1, Eb5 1,
Gb4 1, Bb4 1, Eb5 1, Gb5 1,
// Ebm repeated (63-64)
Eb4 1, Gb4 1, Bb4 1, Eb5 1,
Gb4 1, Bb4 1, Eb5 1, Gb5 1,
`,
			volume:  0.7, // Full-sounding chord instrument
			square:  0.3, // Part sine, part square
			attack:  0.1,
			decay:   0.3,
			sustain: 0.7,
			release: 0.5,
		},
		{
			// --------------------
			// 2) MELODY INSTRUMENT
			// --------------------
			name: "Lead Violin",
			// We'll do 64 measures of melodic lines that loosely match the chord progression.
			// For simplicity, each measure has 4 notes, each note = 1 beat,
			// so they line up measure-by-measure with the chord track.
			// We vary between Eb, Gb, Ab, Bb, Db, E# (F), Cb (B).
			data: `
////////////////////////////////////////////////////////////////////////
// 64-Measure Melody, aligns with chord changes above
////////////////////////////////////////////////////////////////////////
// Section A: measures 1-16
Eb5 1, Gb5 1, Eb5 1, Bb4 1,
Gb5 1, Ab5 1, Gb5 1, Eb5 1,
Bb4 1, Db5 1, Gb5 1, Bb4 1,
Eb5 1, Gb5 1, Ab5 1, Gb5 1,

Eb5 1, Eb5 1, Gb5 1, Bb4 1,
Ab5 1, Gb5 1, Eb5 1, Db5 1,
Gb5 1, Ab5 1, Bb5 1, Ab5 1,
Gb5 1, Eb5 1, Db5 1, Bb4 1,

// Section B: measures 17-32
Eb5 1, Gb5 1, Eb5 1, Bb4 1,
Gb5 1, Ab5 1, Gb5 1, Eb5 1,
Db5 1, Eb5 1, Gb5 1, Db5 1,
Bb4 1, Db5 1, Eb5 1, Gb5 1,

Bb4 1, Db5 1, E#5 1, Bb4 1,
Db5 1, E#5 1, Ab5 1, Gb5 1,
Eb5 1, Bb4 1, Gb5 1, Eb5 1,
Db5 1, Ab5 1, Gb5 1, Eb5 1,

// Section C: measures 33-48
Eb5 1, Eb5 1, Gb5 1, Ab5 1,
Gb5 1, Eb5 1, Bb5 1, Ab5 1,
Db5 1, Eb5 1, Gb5 1, Db5 1,
Eb5 1, Gb5 1, Ab5 1, Bb5 1,

Bb5 1, Ab5 1, Gb5 1, Eb5 1,
Db5 1, Bb4 1, Eb5 1, Gb5 1,
Ab5 1, Gb5 1, Eb5 1, Db5 1,
Bb4 1, Ab4 1, Gb4 1, Eb4 1,

// Section D: measures 49-64
Gb5 1, Bb5 1, Db6 1, Gb5 1,
Db5 1, Bb4 1, Gb5 1, Bb5 1,
Db5 1, E#5 1, Ab5 1, Db5 1,
Ab5 1, Gb5 1, Eb5 1, Db5 1,

Eb5 1, Gb5 1, Bb5 1, Eb5 1,
Gb5 1, Bb5 1, Eb6 1, Gb5 1,
Ab5 1, Gb5 1, Eb5 1, Db5 1,
Bb4 1, Gb5 1, Eb5 1, Eb5 1,
`,
			volume:  0.6, // A bit softer than the chords
			square:  0.0, // Pure sine wave for a gentler violin-ish tone
			attack:  0.05,
			decay:   0.2,
			sustain: 0.8,
			release: 0.4,
		},
	},
}

// Dramatic, sweeping ~3:12 pirate theme at 80 BPM, 64 measures total.
var PirateDramaticTheme = songData{
	name:     "Dramatic Sweeping Pirate Theme",
	bpm:      80,  // 80 BPM => 1 measure (4 beats) ~ 3 seconds => 64 measures ~ 192s = 3m12s
	reverb:   0.8, // Large reverb
	delay:    0.4, // 0.4s delay
	feedback: 0.5, // Strong feedback for a big echo
	ins: []insData{
		{
			//-----------------------------------------------------------------------
			// 1) Chords: A "string ensemble" style with slow attack & long release
			//-----------------------------------------------------------------------
			name:   "String Ensemble (Chords)",
			volume: 0.8,
			square: 0.1, // Mostly sine-like, slightly “grainy”
			// Big ADSR for a swelling effect
			attack:  0.8,
			decay:   1.0,
			sustain: 0.7,
			release: 2.0,
			// 16 measures repeated 4 times => 64 measures total
			data: `
////////////////////////////////////////////////////////////////////////
// ONE “CYCLE” OF 16 MEASURES (4 chords × 4 measures each), REPEATED 4X
////////////////////////////////////////////////////////////////////////

////////////////////////////////////////
// CHORD 1 (A♭ minor) - measures 1-4
////////////////////////////////////////
// measure 1
Ab3 1, Cb4 1, Eb4 1, Ab4 1,
// measure 2
Cb4 1, Eb4 1, Ab4 1, Cb5 1,
// measure 3
Eb4 1, Ab4 1, Cb5 1, Eb5 1,
// measure 4
Ab4 1, Cb5 1, Eb5 1, Ab5 1,

////////////////////////////////////////
// CHORD 2 (F♭ major) - measures 5-8
////////////////////////////////////////
// (F♭ = E, A♭ = G#, C♭ = B)
Fb3 1, Ab3 1, Cb4 1, Fb4 1,
Ab3 1, Cb4 1, Fb4 1, Ab4 1,
Cb4 1, Fb4 1, Ab4 1, Cb5 1,
Fb4 1, Ab4 1, Cb5 1, Fb5 1,

////////////////////////////////////////
// CHORD 3 (D♭ major) - measures 9-12
////////////////////////////////////////
// (D♭, E♯ = F, A♭)
Db3 1, E#3 1, Ab3 1, Db4 1,
E#3 1, Ab3 1, Db4 1, E#4 1,
Ab3 1, Db4 1, E#4 1, Ab4 1,
Db4 1, E#4 1, Ab4 1, Db5 1,

////////////////////////////////////////
// CHORD 4 (E♭ minor) - measures 13-16
////////////////////////////////////////
// (E♭, G♭, B♭)
Eb3 1, Gb3 1, Bb3 1, Eb4 1,
Gb3 1, Bb3 1, Eb4 1, Gb4 1,
Bb3 1, Eb4 1, Gb4 1, Bb4 1,
Eb4 1, Gb4 1, Bb4 1, Eb5 1,

////////////////////////////////////////////////////////////////////////
// REPEAT THE ABOVE 16 MEASURES 4 TIMES FOR A TOTAL OF 64 MEASURES
////////////////////////////////////////////////////////////////////////
Ab3 1, Cb4 1, Eb4 1, Ab4 1,
Cb4 1, Eb4 1, Ab4 1, Cb5 1,
Eb4 1, Ab4 1, Cb5 1, Eb5 1,
Ab4 1, Cb5 1, Eb5 1, Ab5 1,

Fb3 1, Ab3 1, Cb4 1, Fb4 1,
Ab3 1, Cb4 1, Fb4 1, Ab4 1,
Cb4 1, Fb4 1, Ab4 1, Cb5 1,
Fb4 1, Ab4 1, Cb5 1, Fb5 1,

Db3 1, E#3 1, Ab3 1, Db4 1,
E#3 1, Ab3 1, Db4 1, E#4 1,
Ab3 1, Db4 1, E#4 1, Ab4 1,
Db4 1, E#4 1, Ab4 1, Db5 1,

Eb3 1, Gb3 1, Bb3 1, Eb4 1,
Gb3 1, Bb3 1, Eb4 1, Gb4 1,
Bb3 1, Eb4 1, Gb4 1, Bb4 1,
Eb4 1, Gb4 1, Bb4 1, Eb5 1,

Ab3 1, Cb4 1, Eb4 1, Ab4 1,
Cb4 1, Eb4 1, Ab4 1, Cb5 1,
Eb4 1, Ab4 1, Cb5 1, Eb5 1,
Ab4 1, Cb5 1, Eb5 1, Ab5 1,

Fb3 1, Ab3 1, Cb4 1, Fb4 1,
Ab3 1, Cb4 1, Fb4 1, Ab4 1,
Cb4 1, Fb4 1, Ab4 1, Cb5 1,
Fb4 1, Ab4 1, Cb5 1, Fb5 1,

Db3 1, E#3 1, Ab3 1, Db4 1,
E#3 1, Ab3 1, Db4 1, E#4 1,
Ab3 1, Db4 1, E#4 1, Ab4 1,
Db4 1, E#4 1, Ab4 1, Db5 1,

Eb3 1, Gb3 1, Bb3 1, Eb4 1,
Gb3 1, Bb3 1, Eb4 1, Gb4 1,
Bb3 1, Eb4 1, Gb4 1, Bb4 1,
Eb4 1, Gb4 1, Bb4 1, Eb5 1,
`,
		},
		{
			//-----------------------------------------------------------------------
			// 2) Choir: Big sustained notes under the chords, to add depth/drama
			//-----------------------------------------------------------------------
			name:   "Choir",
			volume: 0.5,
			square: 0.0, // Pure sine-ish (smooth)
			// Very slow attack, long release
			attack:  1.5,
			decay:   1.0,
			sustain: 0.8,
			release: 2.0,
			// We'll do half-notes that follow the chord roots for 16 measures,
			// repeated 4 times. Each chord root changes every 4 measures, so
			// we’ll hold each root for 2 beats, do 2 half-notes per measure, etc.
			// For variety, we’ll step from root to 5th or minor 3rd.
			data: `
////////////////////////////////////////////////////////////////////////
// 16 MEASURES, REPEATED 4X
////////////////////////////////////////////////////////////////////////

////////////////////
// CHORD 1: Ab minor
////////////////////
// measures 1-4, each measure has 2 half-notes = 4 beats
Ab2 2, Eb2 2,
Ab2 2, Cb3 2,
Ab2 2, Eb2 2,
Ab2 2, Cb3 2,

////////////////////
// CHORD 2: Fb major
////////////////////
Fb2 2, Cb3 2,
Fb2 2, Ab2 2,
Fb2 2, Cb3 2,
Fb2 2, Ab2 2,

////////////////////
// CHORD 3: Db major
////////////////////
Db2 2, Ab2 2,
Db2 2, E#2 2,
Db2 2, Ab2 2,
Db2 2, E#2 2,

////////////////////
// CHORD 4: Eb minor
////////////////////
Eb2 2, Bb2 2,
Eb2 2, Gb2 2,
Eb2 2, Bb2 2,
Eb2 2, Gb2 2,

////////////////////////////////////////////////////////////////////////
// REPEAT THE ABOVE 16 MEASURES 4 TIMES
////////////////////////////////////////////////////////////////////////

Ab2 2, Eb2 2,
Ab2 2, Cb3 2,
Ab2 2, Eb2 2,
Ab2 2, Cb3 2,

Fb2 2, Cb3 2,
Fb2 2, Ab2 2,
Fb2 2, Cb3 2,
Fb2 2, Ab2 2,

Db2 2, Ab2 2,
Db2 2, E#2 2,
Db2 2, Ab2 2,
Db2 2, E#2 2,

Eb2 2, Bb2 2,
Eb2 2, Gb2 2,
Eb2 2, Bb2 2,
Eb2 2, Gb2 2,
`,
		},
		{
			//-----------------------------------------------------------------------
			// 3) Lead Cello: A melodic line that soars over the chords
			//-----------------------------------------------------------------------
			name:   "Lead Cello",
			volume: 0.7,
			square: 0.2, // Slightly more bite than the choir
			// Somewhat shorter attack than chords, but still a bit of swell
			attack:  0.5,
			decay:   0.5,
			sustain: 0.8,
			release: 1.5,
			// We'll do quarter-note melodies in each measure (4 notes/measure),
			// 16 measures, repeated 4x
			data: `
////////////////////////////////////////////////////////////////////////
// 16-MEASURE MELODY, REPEATED 4X
////////////////////////////////////////////////////////////////////////

// measures 1-4: climbing around Ab minor (Ab, Cb, Eb)
Ab3 1, Cb4 1, Eb4 1, Ab4 1,
Cb4 1, Eb4 1, Ab4 1, Gb4 1,
Eb4 1, Ab4 1, Cb4 1, Eb4 1,
Ab4 1, Gb4 1, Eb4 1, Cb4 1,

// measures 5-8: around Fb major (Fb, Ab, Cb)
Fb4 1, Ab4 1, Cb4 1, Fb4 1,
Ab4 1, Fb4 1, Cb4 1, Ab4 1,
Fb4 1, Cb4 1, Ab4 1, Fb4 1,
Cb4 1, Ab4 1, Fb4 1, Cb4 1,

// measures 9-12: around Db major (Db, E#, Ab)
Db4 1, E#4 1, Ab4 1, Db5 1,
E#4 1, Db5 1, Ab4 1, E#4 1,
Ab4 1, Db5 1, E#4 1, Ab4 1,
Db5 1, E#4 1, Ab4 1, Db4 1,

// measures 13-16: around Eb minor (Eb, Gb, Bb)
Eb4 1, Gb4 1, Bb4 1, Eb5 1,
Gb4 1, Eb5 1, Bb4 1, Gb4 1,
Bb4 1, Eb5 1, Gb4 1, Bb4 1,
Eb5 1, Gb4 1, Bb4 1, Eb4 1,

////////////////////////////////////////////////////////////////////////
// REPEAT THAT BLOCK 4 TIMES FOR 64 MEASURES
////////////////////////////////////////////////////////////////////////
Ab3 1, Cb4 1, Eb4 1, Ab4 1,
Cb4 1, Eb4 1, Ab4 1, Gb4 1,
Eb4 1, Ab4 1, Cb4 1, Eb4 1,
Ab4 1, Gb4 1, Eb4 1, Cb4 1,

Fb4 1, Ab4 1, Cb4 1, Fb4 1,
Ab4 1, Fb4 1, Cb4 1, Ab4 1,
Fb4 1, Cb4 1, Ab4 1, Fb4 1,
Cb4 1, Ab4 1, Fb4 1, Cb4 1,

Db4 1, E#4 1, Ab4 1, Db5 1,
E#4 1, Db5 1, Ab4 1, E#4 1,
Ab4 1, Db5 1, E#4 1, Ab4 1,
Db5 1, E#4 1, Ab4 1, Db4 1,

Eb4 1, Gb4 1, Bb4 1, Eb5 1,
Gb4 1, Eb5 1, Bb4 1, Gb4 1,
Bb4 1, Eb5 1, Gb4 1, Bb4 1,
Eb5 1, Gb4 1, Bb4 1, Eb4 1,
`,
		},
	},
}

// Merry, playful sea-sailing tune (~2 min) in "Bb major" spelled enharmonically
var MerrySailingTheme = songData{
	name: "Merry Sea Shanty",
	bpm:  120,
	// Subtle, pleasant space
	reverb:   0.3,
	delay:    0.2,
	feedback: 0.25,
	ins: []insData{
		{
			//-------------------------------------------------------------------
			// 1) Strumming Guitar (Chords)
			//-------------------------------------------------------------------
			name:   "Strumming Guitar",
			volume: 0.8,
			square: 0.4, // somewhat bright/buzzy
			// Short attack for a strum, moderate decay, lower sustain, short release
			attack:  0.05,
			decay:   0.2,
			sustain: 0.5,
			release: 0.3,
			// We'll define a 4-measure loop (I–IV–V–I) repeated 16 times => 64 measures total.
			//
			// Each measure has 4 quarter notes. The chord shape is repeated each beat
			// to simulate a quick "strum" on each beat.
			//
			// B♭ major => (Bb, C##, E#)
			// E♭ major => (Eb, F##, Bb)
			// F major  => (E#, G##, B#)
			data: `
////////////////////////////////////////////////////////////////////////
// 4-measure chord pattern, repeated 16x => 64 total measures
////////////////////////////////////////////////////////////////////////

// MEASURE 1: Bb major
Bb3 1, C##4 1, E#4 1, Bb4 1,

// MEASURE 2: Eb major
Eb3 1, F##3 1, Bb3 1, Eb4 1,

// MEASURE 3: F major
E#3 1, G##3 1, B#3 1, E#4 1,

// MEASURE 4: back to Bb major
Bb3 1, C##4 1, E#4 1, Bb4 1,

////////////////////////////////////////////////////////////////////////
// Repeat above 4-measure block 15 more times to reach 64 measures
////////////////////////////////////////////////////////////////////////
Bb3 1, C##4 1, E#4 1, Bb4 1,
Eb3 1, F##3 1, Bb3 1, Eb4 1,
E#3 1, G##3 1, B#3 1, E#4 1,
Bb3 1, C##4 1, E#4 1, Bb4 1,

Bb3 1, C##4 1, E#4 1, Bb4 1,
Eb3 1, F##3 1, Bb3 1, Eb4 1,
E#3 1, G##3 1, B#3 1, E#4 1,
Bb3 1, C##4 1, E#4 1, Bb4 1,

Bb3 1, C##4 1, E#4 1, Bb4 1,
Eb3 1, F##3 1, Bb3 1, Eb4 1,
E#3 1, G##3 1, B#3 1, E#4 1,
Bb3 1, C##4 1, E#4 1, Bb4 1,

Bb3 1, C##4 1, E#4 1, Bb4 1,
Eb3 1, F##3 1, Bb3 1, Eb4 1,
E#3 1, G##3 1, B#3 1, E#4 1,
Bb3 1, C##4 1, E#4 1, Bb4 1,

Bb3 1, C##4 1, E#4 1, Bb4 1,
Eb3 1, F##3 1, Bb3 1, Eb4 1,
E#3 1, G##3 1, B#3 1, E#4 1,
Bb3 1, C##4 1, E#4 1, Bb4 1,

Bb3 1, C##4 1, E#4 1, Bb4 1,
Eb3 1, F##3 1, Bb3 1, Eb4 1,
E#3 1, G##3 1, B#3 1, E#4 1,
Bb3 1, C##4 1, E#4 1, Bb4 1,
`,
		},
		{
			//-------------------------------------------------------------------
			// 2) Flute (Melody)
			//-------------------------------------------------------------------
			name:   "Flute",
			volume: 0.6,
			square: 0.0, // pure sine (soft flute-like)
			// A slightly gentler ADSR than the guitar
			attack:  0.1,
			decay:   0.2,
			sustain: 0.8,
			release: 0.4,
			// We'll do a 4-measure melodic phrase repeated 16 times => 64 measures.
			// Each measure has 4 quarter notes. They loosely trace the chords.
			data: `
////////////////////////////////////////////////////////////////////////
// 4-measure melodic phrase, repeated 16 times
////////////////////////////////////////////////////////////////////////

// MEASURE 1 (over Bb major): notes revolve around Bb, C##, E# 
Bb4 1, C##5 1, Bb4 1, E#5 1,

// MEASURE 2 (over Eb major): revolve around Eb, F##, Bb
Eb5 1, Bb4 1, F##5 1, Eb5 1,

// MEASURE 3 (over F major): revolve around E#, G##, B#
E#5 1, G##4 1, B#4 1, G##5 1,

// MEASURE 4 (back to Bb major)
Bb4 1, E#5 1, C##5 1, Bb4 1,

////////////////////////////////////////////////////////////////////////
// Repeat that 4-measure block 15 more times => total 64 measures
////////////////////////////////////////////////////////////////////////
Bb4 1, C##5 1, Bb4 1, E#5 1,
Eb5 1, Bb4 1, F##5 1, Eb5 1,
E#5 1, G##4 1, B#4 1, G##5 1,
Bb4 1, E#5 1, C##5 1, Bb4 1,

Bb4 1, C##5 1, Bb4 1, E#5 1,
Eb5 1, Bb4 1, F##5 1, Eb5 1,
E#5 1, G##4 1, B#4 1, G##5 1,
Bb4 1, E#5 1, C##5 1, Bb4 1,

Bb4 1, C##5 1, Bb4 1, E#5 1,
Eb5 1, Bb4 1, F##5 1, Eb5 1,
E#5 1, G##4 1, B#4 1, G##5 1,
Bb4 1, E#5 1, C##5 1, Bb4 1,

Bb4 1, C##5 1, Bb4 1, E#5 1,
Eb5 1, Bb4 1, F##5 1, Eb5 1,
E#5 1, G##4 1, B#4 1, G##5 1,
Bb4 1, E#5 1, C##5 1, Bb4 1,

Bb4 1, C##5 1, Bb4 1, E#5 1,
Eb5 1, Bb4 1, F##5 1, Eb5 1,
E#5 1, G##4 1, B#4 1, G##5 1,
Bb4 1, E#5 1, C##5 1, Bb4 1,

Bb4 1, C##5 1, Bb4 1, E#5 1,
Eb5 1, Bb4 1, F##5 1, Eb5 1,
E#5 1, G##4 1, B#4 1, G##5 1,
Bb4 1, E#5 1, C##5 1, Bb4 1,
`,
		},
	},
}
