package main

var songList []*songData = []*songData{
	&EbonyGale,
	&MidnightDepthsPercussion,

	&TwilightReef,
	&IronWake,
	&SpectralKeel,
	&CursedMeridian,
	&SilverMaelstrom,

	&SunkenArmada,
	&MaroonedTwilight,
	&BloodyCorsair,
	&LostTidepool,
	&BlackReef,
}

type songData struct {
	name string
	ins  []insData
	bpm  int
}

type insData struct {
	name, data    string
	volume, blend float64
}

var EbonyGale = songData{
	name: "Aboard the Ebony Gale",
	bpm:  70,
	ins: []insData{
		{name: "lead",
			volume: 0.5, blend: 1,
			data: `Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Gb4 2,
Fb4 1, Gb4 1, Ab4 2,
Bb4 2, Ab4 2,

Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Db5 2,
Cb5 1, Bb4 1, Ab4 2,
Gb4 4`},

		{name: "harmony",
			volume: 0.5, blend: 1,
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
Eb4/Gb4/Bb4 4`},

		{name: "bass",
			volume: 1, blend: 0.5,
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
Eb2 4`,
		}, {
			// 6) Fake Percussion:
			// We'll use Eb1 as a 'kick' on beats 1 & 3,
			// and Bb1 as a 'snare' on beats 2 & 4 (very short notes).
			name:   "percussion",
			volume: 0.5, blend: 0.7,
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
		}}}

var MidnightDepthsPercussion = songData{
	name: "Midnight Depths",
	bpm:  70,
	ins: []insData{
		{
			name:   "lead",
			volume: 0.5, blend: 1,
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
			name:   "harmony",
			volume: 0.5, blend: 1,
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
			name:   "bass",
			volume: 1, blend: 0.5,
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
			name:   "descant",
			volume: 1, blend: 0,
			data: `
NN 8, // (rests for first 8 measures)

Gb5 0.5, Ab5 0.5, Bb5 1, NN 2.5,
Db6 1, Bb5 1, NN 2,
Ab5 0.5, Gb5 0.5, Eb5 1, NN 2.5
`,
		},
		{
			name:   "pad",
			volume: 1, blend: 0.5,
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
			// 6) Fake Percussion:
			// We'll use Eb1 as a 'kick' on beats 1 & 3,
			// and Bb1 as a 'snare' on beats 2 & 4 (very short notes).
			name:   "percussion",
			volume: 2, blend: 0.7,
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

var SilverMaelstrom = songData{
	name: "Silver Maelstrom",
	bpm:  100,
	ins: []insData{
		{
			name:   "lead",
			volume: 1.0,
			blend:  0.5, // balanced sine/square
			data: `
Eb4 1, Gb4 1, Ab4 2, 
Bb4 2, Ab4 2, 
Db5 1, Bb4 1, Gb4 2,
Eb4 4,

Eb4 1, Fb4 1, Gb4 2,
Ab4 1, Bb4 1, Db5 2,
Cb5 2, Ab4 2,
Gb4 2, Eb4 2
`,
		},
		{
			name:   "harmony",
			volume: 0.8,
			blend:  0.2, // mostly sine
			data: `
Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Gb4/Bb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bb4/Db5 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Gb4/Bb4 4,
Eb4/Gb4/Bb4 4
`,
		},
		{
			name:   "bass",
			volume: 1.1,
			blend:  0.3, // mellow but slightly square
			data: `
Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Bb2 4,
Ab2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Bb2 4,
Ab2 4,
Eb2 4
`,
		},
		{
			name:   "descant",
			volume: 0.7,
			blend:  0.6, // a bit edgier than pure sine
			data: `
NN 8, // silent first 8 measures for drama

Gb5 0.5, Ab5 0.5, Bb5 1, NN 2.5,
Db6 1, Bb5 1, NN 2,
Ab5 0.5, Gb5 0.5, Eb5 1, NN 2.5
`,
		},
		{
			name:   "pad",
			volume: 0.5,
			blend:  0.1, // almost pure sine, very soft
			data: `
Eb4/Gb4/Bb4 0.5, NN 3.5,
Cb4/Eb4/Gb4 0.5, NN 3.5,
Db4/Fb4/Ab4 0.5, NN 3.5,
Eb4/Gb4/Bb4 0.5, NN 3.5,

Eb4/Gb4/Bb4 0.5, NN 3.5,
Cb4/Eb4/Gb4 0.5, NN 3.5,
Bb3/Db4/Fb4 0.5, NN 3.5,
Eb4/Gb4/Bb4 0.5, NN 3.5,

Eb4/Gb4/Bb4 0.5, NN 3.5,
Db4/Fb4/Ab4 0.5, NN 3.5,
Gb4/Bb4/Db5 0.5, NN 3.5,
Eb4/Gb4/Bb4 0.5, NN 3.5,

Eb4/Gb4/Bb4 0.5, NN 3.5,
Cb4/Eb4/Gb4 0.5, NN 3.5,
Db4/Fb4/Ab4 0.5, NN 3.5,
Eb4/Gb4/Bb4 0.5, NN 3.5
`,
		},
		{
			// Fake percussion: Kick on Eb1, Snare on Bb1
			name:   "percussion",
			volume: 1.3, // louder
			blend:  1.0, // pure square = more punch
			data: `
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75
`,
		},
	},
}

var CursedMeridian = songData{
	name: "Cursed Meridian",
	bpm:  80,
	ins: []insData{
		{
			name:   "lead",
			volume: 1.0,
			blend:  0.4,
			data: `
Eb4 2, Gb4 2,
Ab4 1, Bb4 1, Db5 2,
Cb5 2, Ab4 2,
Gb4 2, Eb4 2,

Eb4 2, Fb4 2,
Gb4 2, Ab4 2,
Bb4 1, Ab4 1, Gb4 2,
Eb4 4
`,
		},
		{
			name:   "harmony",
			volume: 0.9,
			blend:  0.3,
			data: `
Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Gb4/Bb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bb4/Db5 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bb4/Db5 4,
Eb4/Gb4/Bb4 4
`,
		},
		{
			name:   "bass",
			volume: 1.1,
			blend:  0.2,
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
			name:   "descant",
			volume: 0.8,
			blend:  0.5,
			data: `
NN 8,

Gb5 1, Ab5 1, Bb5 2,
Db6 1, Bb5 1, NN 2,
Ab5 0.5, Gb5 0.5, Eb5 1, NN 2.5
`,
		},
		{
			name:   "pad",
			volume: 0.6,
			blend:  0.0, // pure sine wave for a haunting vibe
			data: `
Eb4/Gb4/Bb4 2, NN 2,
Cb4/Eb4/Gb4 2, NN 2,
Db4/Fb4/Ab4 2, NN 2,
Eb4/Gb4/Bb4 2, NN 2,

Eb4/Gb4/Bb4 2, NN 2,
Db4/Fb4/Ab4 2, NN 2,
Gb4/Bb4/Db5 2, NN 2,
Eb4/Gb4/Bb4 2, NN 2
`,
		},
		{
			name:   "percussion",
			volume: 1.4,
			blend:  1.0,
			data: `
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75
`,
		},
	},
}

var SpectralKeel = songData{
	name: "Spectral Keel",
	bpm:  95,
	ins: []insData{
		{
			name:   "lead",
			volume: 1.0,
			blend:  0.5,
			data: `
Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Gb4 2,
Eb4 1, Fb4 1, Gb4 2,
Ab4 2, Gb4 2,

Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Db5 2,
Cb5 1, Bb4 1, Ab4 2,
Gb4 4
`,
		},
		{
			name:   "harmony",
			volume: 0.75,
			blend:  0.2,
			data: `
Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Fb4/Ab4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bb4/Db5 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Fb4/Ab4 4,
Eb4/Gb4/Bb4 4
`,
		},
		{
			name:   "bass",
			volume: 1.2,
			blend:  0.4,
			data: `
Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Bb2 4,
Ab2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Bb2 4,
Ab2 4,
Eb2 4
`,
		},
		{
			name:   "descant",
			volume: 0.6,
			blend:  0.6,
			data: `
NN 8,
Gb5 0.5, Ab5 0.5, Bb5 1, NN 2.5,
Db6 1, Bb5 1, NN 2,
Ab5 0.5, Gb5 0.5, Eb5 1, NN 2.5
`,
		},
		{
			name:   "pad",
			volume: 0.5,
			blend:  0.0, // pure sine
			data: `
Eb4/Gb4/Bb4 0.5, NN 3.5,
Db4/Fb4/Ab4 0.5, NN 3.5,
Gb4/Bb4/Db5 0.5, NN 3.5,
Eb4/Gb4/Bb4 0.5, NN 3.5,

Eb4/Gb4/Bb4 0.5, NN 3.5,
Cb4/Eb4/Gb4 0.5, NN 3.5,
Bb3/Db4/Fb4 0.5, NN 3.5,
Eb4/Gb4/Bb4 0.5, NN 3.5
`,
		},
		{
			name:   "percussion",
			volume: 1.4,
			blend:  1.0,
			data: `
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75
`,
		},
	},
}

var IronWake = songData{
	name: "Iron Wake",
	bpm:  105,
	ins: []insData{
		{
			name:   "lead",
			volume: 1.0,
			blend:  0.5,
			data: `
Eb4 1, Fb4 1, Gb4 2,
Ab4 1, Gb4 1, Fb4 2,
Eb4 1, Gb4 1, Ab4 2,
Bb4 2, Ab4 2,

Eb4 1, Fb4 1, Gb4 2,
Ab4 1, Gb4 1, Db5 2,
Cb5 1, Bb4 1, Ab4 2,
Gb4 4
`,
		},
		{
			name:   "harmony",
			volume: 0.8,
			blend:  0.3,
			data: `
Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Fb4/Ab4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bb4/Db5 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Fb4/Ab4 4,
Eb4/Gb4/Bb4 4
`,
		},
		{
			name:   "bass",
			volume: 1.2,
			blend:  0.3,
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
			name:   "descant",
			volume: 0.7,
			blend:  0.5,
			data: `
NN 8,

Gb5 1, Ab5 1, Bb5 2,
Db6 1, Bb5 1, NN 2,
Ab5 0.5, Gb5 0.5, Eb5 1, NN 2.5
`,
		},
		{
			name:   "pad",
			volume: 0.5,
			blend:  0.1,
			data: `
Eb4/Gb4/Bb4 1, NN 1,
Cb4/Eb4/Gb4 1, NN 1,
Bb3/Db4/Fb4 1, NN 1,
Eb4/Gb4/Bb4 1, NN 1,

Eb4/Gb4/Bb4 1, NN 1,
Db4/Fb4/Ab4 1, NN 1,
Gb4/Bb4/Db5 1, NN 1,
Eb4/Gb4/Bb4 1, NN 1,

Eb4/Gb4/Bb4 1, NN 1,
Cb4/Eb4/Gb4 1, NN 1,
Bb3/Db4/Fb4 1, NN 1,
Eb4/Gb4/Bb4 1, NN 1,

Eb4/Gb4/Bb4 1, NN 1,
Db4/Fb4/Ab4 1, NN 1,
Gb4/Bb4/Db5 1, NN 1,
Eb4/Gb4/Bb4 1, NN 1
`,
		},
		{
			name:   "percussion",
			volume: 1.3,
			blend:  1.0,
			data: `
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75
`,
		},
	},
}

var TwilightReef = songData{
	name: "Twilight Reef",
	bpm:  90,
	ins: []insData{
		{
			name:   "lead",
			volume: 1.0,
			blend:  0.5,
			data: `
Eb4 2, NN 2,
Gb4 2, Ab4 2,
Bb4 1, Ab4 1, Gb4 2,
Eb4 4,

Eb4 2, Fb4 2,
Gb4 2, Ab4 2,
Db5 2, Bb4 2,
Ab4 2, Gb4 2
`,
		},
		{
			name:   "harmony",
			volume: 0.8,
			blend:  0.2,
			data: `
Eb4/Gb4/Bb4 4,
Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,

Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bb4/Db5 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Fb4/Ab4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bb4/Db5 4,
Eb4/Gb4/Bb4 4
`,
		},
		{
			name:   "bass",
			volume: 1.1,
			blend:  0.25,
			data: `
Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Bb2 4,
Ab2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Bb2 4,
Ab2 4,
Eb2 4
`,
		},
		{
			name:   "descant",
			volume: 0.7,
			blend:  0.6,
			data: `
NN 8,

Gb5 1, Ab5 1, Bb5 2,
Db6 1, Bb5 1, NN 2,
Ab5 0.5, Gb5 0.5, Eb5 1, NN 2.5
`,
		},
		{
			name:   "pad",
			volume: 0.4,
			blend:  0.0,
			data: `
Eb4/Gb4/Bb4 0.5, NN 3.5,
Cb4/Eb4/Gb4 0.5, NN 3.5,
Db4/Fb4/Ab4 0.5, NN 3.5,
Eb4/Gb4/Bb4 0.5, NN 3.5,

Eb4/Gb4/Bb4 0.5, NN 3.5,
Db4/Fb4/Ab4 0.5, NN 3.5,
Gb4/Bb4/Db5 0.5, NN 3.5,
Eb4/Gb4/Bb4 0.5, NN 3.5
`,
		},
		{
			name:   "percussion",
			volume: 1.5,
			blend:  1.0,
			data: `
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75
`,
		},
	},
}

var SunkenArmada = songData{
	name: "Sunken Armada",
	bpm:  85,
	ins: []insData{
		{
			name:   "lead",
			volume: 1.0,
			blend:  0.4,
			data: `
Eb4 2, Gb4 2,
Ab4 1, Bb4 1, Db5 2,
Bb4 2, Ab4 2,
Gb4 4,

Eb4 2, Fb4 2,
Gb4 1, Ab4 1, Eb5 2,
Db5 2, Bb4 2,
Eb4 2, NN 2
`,
		},
		{
			name:   "harmony",
			volume: 0.8,
			blend:  0.2,
			data: `
Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Gb4/Bb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bb4/Db5 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bb4/Db5 4,
Eb4/Gb4/Bb4 4
`,
		},
		{
			name:   "bass",
			volume: 1.1,
			blend:  0.3,
			data: `
Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Bb2 4,
Ab2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Bb2 4,
Ab2 4,
Eb2 4
`,
		},
		{
			name:   "descant",
			volume: 0.7,
			blend:  0.5,
			data: `
NN 8,
Gb5 1, Ab5 1, Bb5 2,
Db6 1, Bb5 1, NN 2,
Ab5 0.5, Gb5 0.5, Eb5 1, NN 2.5
`,
		},
		{
			name:   "pad",
			volume: 0.5,
			blend:  0.0, // mostly sine, giving dark hum
			data: `
Eb4/Gb4/Bb4 1, NN 3,
Cb4/Eb4/Gb4 1, NN 3,
Db4/Fb4/Ab4 1, NN 3,
Eb4/Gb4/Bb4 1, NN 3,

Eb4/Gb4/Bb4 1, NN 3,
Db4/Fb4/Ab4 1, NN 3,
Gb4/Bb4/Db5 1, NN 3,
Eb4/Gb4/Bb4 1, NN 3
`,
		},
		{
			name:   "percussion",
			volume: 1.4,
			blend:  1.0, // pure square for punch
			data: `
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75
`,
		},
	},
}

var MaroonedTwilight = songData{
	name: "Marooned Twilight",
	bpm:  75,
	ins: []insData{
		{
			name:   "lead",
			volume: 1.0,
			blend:  0.3,
			data: `
Eb4 2, NN 2,
Gb4 2, Ab4 2,
Bb4 1, Ab4 1, Gb4 2,
Eb4 4,

Eb4 2, Fb4 2,
Gb4 1, Ab4 1, Db5 2,
Cb5 2, Bb4 2,
Ab4 2, Gb4 2
`,
		},
		{
			name:   "harmony",
			volume: 0.75,
			blend:  0.2,
			data: `
Eb4/Gb4/Bb4 4,
Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bb4/Db5 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Fb4/Ab4 4,
Eb4/Gb4/Bb4 4
`,
		},
		{
			name:   "bass",
			volume: 1.1,
			blend:  0.25,
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
			name:   "descant",
			volume: 0.7,
			blend:  0.6,
			data: `
NN 8,
Gb5 1, Ab5 1, Bb5 2,
Db6 1, Bb5 1, NN 2,
Ab5 0.5, Gb5 0.5, Eb5 1, NN 2.5
`,
		},
		{
			name:   "pad",
			volume: 0.5,
			blend:  0.0,
			data: `
Eb4/Gb4/Bb4 0.5, NN 3.5,
Db4/Fb4/Ab4 0.5, NN 3.5,
Gb4/Bb4/Db5 0.5, NN 3.5,
Eb4/Gb4/Bb4 0.5, NN 3.5,

Eb4/Gb4/Bb4 0.5, NN 3.5,
Cb4/Eb4/Gb4 0.5, NN 3.5,
Bb3/Db4/Fb4 0.5, NN 3.5,
Eb4/Gb4/Bb4 0.5, NN 3.5
`,
		},
		{
			name:   "percussion",
			volume: 1.3,
			blend:  1.0,
			data: `
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75
`,
		},
	},
}

var BloodyCorsair = songData{
	name: "Bloody Corsair",
	bpm:  110,
	ins: []insData{
		{
			name:   "lead",
			volume: 1.0,
			blend:  0.5,
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
			name:   "harmony",
			volume: 0.85,
			blend:  0.3,
			data: `
Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bb4/Db5 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Fb4/Ab4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Bb3/Db4/Fb4 4,
Ab3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4
`,
		},
		{
			name:   "bass",
			volume: 1.2,
			blend:  0.4,
			data: `
Eb2 4,
Cb2 4,
Bb2 4,
Eb2 4,

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
			name:   "descant",
			volume: 0.7,
			blend:  0.6,
			data: `
NN 8,
Gb5 0.5, Ab5 0.5, Bb5 1, NN 2.5,
Db6 1, Bb5 1, NN 2,
Ab5 0.5, Gb5 0.5, Eb5 1, NN 2.5
`,
		},
		{
			name:   "pad",
			volume: 0.4,
			blend:  0.1,
			data: `
Eb4/Gb4/Bb4 1, NN 1,
Cb4/Eb4/Gb4 1, NN 1,
Bb3/Db4/Fb4 1, NN 1,
Eb4/Gb4/Bb4 1, NN 1,

Eb4/Gb4/Bb4 1, NN 1,
Db4/Fb4/Ab4 1, NN 1,
Gb4/Bb4/Db5 1, NN 1,
Eb4/Gb4/Bb4 1, NN 1,

Eb4/Gb4/Bb4 1, NN 1,
Cb4/Eb4/Gb4 1, NN 1,
Db4/Fb4/Ab4 1, NN 1,
Eb4/Gb4/Bb4 1, NN 1,

Eb4/Gb4/Bb4 1, NN 1,
Bb3/Db4/Fb4 1, NN 1,
Ab3/Db4/Fb4 1, NN 1,
Eb4/Gb4/Bb4 1, NN 1
`,
		},
		{
			name:   "percussion",
			volume: 1.4,
			blend:  1.0,
			data: `
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75
`,
		},
	},
}

var LostTidepool = songData{
	name: "Lost Tidepool",
	bpm:  78,
	ins: []insData{
		{
			name:   "lead",
			volume: 1.0,
			blend:  0.4,
			data: `
Eb4 2, Gb4 2,
Ab4 2, Fb4 2,
Eb4 1, Gb4 1, Ab4 2,
Bb4 4,

Eb4 2, Fb4 2,
Gb4 1, Ab4 1, Bb4 2,
Db5 2, Ab4 2,
Gb4 1, Eb4 1, NN 2
`,
		},
		{
			name:   "harmony",
			volume: 0.75,
			blend:  0.2,
			data: `
Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bb4/Db5 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bb4/Db5 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4
`,
		},
		{
			name:   "bass",
			volume: 1.2,
			blend:  0.3,
			data: `
Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Bb2 4,
Ab2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Bb2 4,
Ab2 4,
Eb2 4
`,
		},
		{
			name:   "descant",
			volume: 0.7,
			blend:  0.6,
			data: `
NN 8,
Gb5 1, Ab5 1, Bb5 2,
Db6 1, Bb5 1, NN 2,
Ab5 0.5, Gb5 0.5, Eb5 1, NN 2.5
`,
		},
		{
			name:   "pad",
			volume: 0.5,
			blend:  0.1,
			data: `
Eb4/Gb4/Bb4 0.75, NN 3.25,
Db4/Fb4/Ab4 0.75, NN 3.25,
Gb4/Bb4/Db5 0.75, NN 3.25,
Eb4/Gb4/Bb4 0.75, NN 3.25,

Eb4/Gb4/Bb4 0.75, NN 3.25,
Cb4/Eb4/Gb4 0.75, NN 3.25,
Bb3/Db4/Fb4 0.75, NN 3.25,
Eb4/Gb4/Bb4 0.75, NN 3.25
`,
		},
		{
			name:   "percussion",
			volume: 1.4,
			blend:  1.0,
			data: `
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75
`,
		},
	},
}

var BlackReef = songData{
	name: "Black Reef",
	bpm:  92,
	ins: []insData{
		{
			name:   "lead",
			volume: 1.0,
			blend:  0.5,
			data: `
Eb4 1, Gb4 1, Ab4 1, Bb4 1,
Db5 2, Bb4 2,
Ab4 1, Gb4 1, Eb4 2,
Fb4 2, Gb4 2,

Eb4 1, Fb4 1, Gb4 2,
Ab4 1, Gb4 1, Db5 2,
Cb5 2, Ab4 2,
Bb4 1, Ab4 1, Gb4 2
`,
		},
		{
			name:   "harmony",
			volume: 0.8,
			blend:  0.2,
			data: `
Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bb4/Db5 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bb4/Db5 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4
`,
		},
		{
			name:   "bass",
			volume: 1.2,
			blend:  0.3,
			data: `
Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Bb2 4,
Ab2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Bb2 4,
Ab2 4,
Eb2 4
`,
		},
		{
			name:   "descant",
			volume: 0.7,
			blend:  0.6,
			data: `
NN 8,
Gb5 0.5, Ab5 0.5, Bb5 1, NN 2.5,
Db6 1, Bb5 1, NN 2,
Ab5 0.5, Gb5 0.5, Eb5 1, NN 2.5
`,
		},
		{
			name:   "pad",
			volume: 0.5,
			blend:  0.0,
			data: `
Eb4/Gb4/Bb4 0.5, NN 3.5,
Db4/Fb4/Ab4 0.5, NN 3.5,
Gb4/Bb4/Db5 0.5, NN 3.5,
Eb4/Gb4/Bb4 0.5, NN 3.5,

Eb4/Gb4/Bb4 0.5, NN 3.5,
Cb4/Eb4/Gb4 0.5, NN 3.5,
Bb3/Db4/Fb4 0.5, NN 3.5,
Eb4/Gb4/Bb4 0.5, NN 3.5
`,
		},
		{
			name:   "percussion",
			volume: 1.4,
			blend:  1.0,
			data: `
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75,
Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75, Eb1 0.25, NN 0.75, Bb1 0.25, NN 0.75
`,
		},
	},
}
