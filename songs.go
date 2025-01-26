package main

var songList []*songData = []*songData{
	//&MidnightDepthsPercussion,
	&MidnightDepths,
	&CrimsonTempest,
	&GuildedStorm,
	&EbonyGale,
	&SailsOfDusk,
	&DarklingWaters,
	&NightfallApproch,
	&SerpentsWake,
	&AshenCrosswinds,
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
	bpm:  90,
	ins: []insData{
		insData{name: "lead",
			volume: 0.5, blend: 1,
			data: `Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Gb4 2,
Fb4 1, Gb4 1, Ab4 2,
Bb4 2, Ab4 2,

Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Db5 2,
Cb5 1, Bb4 1, Ab4 2,
Gb4 4`},

		insData{name: "harmony",
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

		insData{name: "bass",
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
		}}}

var SailsOfDusk = songData{
	name: "Sails of Dusk",
	bpm:  120,
	ins: []insData{
		insData{name: "lead",
			volume: 0.5, blend: 1,
			data: `
Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Gb4 2,
Fb4 1, Gb4 1, Ab4 2,
Bb4 2, Ab4 2,

Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Db5 2,
Cb5 1, Bb4 1, Ab4 2,
Gb4 4
`},
		insData{name: "harmony",
			volume: 0.5, blend: 1,
			data: `
Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Gb4/Bb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Gb4/Bb4 4,
Eb4/Gb4/Bb4 4
`},
		insData{name: "bass",
			volume: 1, blend: 0.5,
			data: `
Eb2 4,
Cb2 4,
Bb2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Bb2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Db2 4,
Eb2 4
`},
	}}

var DarklingWaters = songData{
	name: "Darkling Waters",
	bpm:  130,
	ins: []insData{
		insData{name: "lead",
			volume: 0.5, blend: 1,
			data: `
Eb4 1, Fb4 1, Gb4 2,
Ab4 1, Gb4 1, Fb4 2,
Eb4 1, Gb4 1, Bb4 2,
Db5 2, Bb4 2,

Eb4 1, Fb4 1, Gb4 2,
Ab4 1, Gb4 1, Eb4 2,
Fb4 1, Gb4 1, Ab4 2,
Bb4 4
`},
		insData{name: "harmony",
			volume: 0.5, blend: 1,
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
Cb4/Eb4/Gb4 4,
Db4/Fb4/Ab4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4
`},
		insData{name: "bass",
			volume: 1, blend: 0.5,
			data: `
Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Bb2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Bb2 4,
Eb2 4
`},
	}}

var NightfallApproch = songData{
	name: "Nightfall Approach",
	bpm:  110,
	ins: []insData{
		insData{name: "lead",
			volume: 0.5, blend: 1,
			data: `
Eb4 1, Gb4 1, Bb4 2,
Ab4 1, Gb4 1, Eb4 2,
Fb4 2, Gb4 2,
Ab4 1, Gb4 1, Fb4 2,
Eb4 2, Gb4 2,

Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Gb4 2,
Eb4 1, Gb4 1, Ab4 2,
Bb4 4
`},
		insData{name: "harmony",
			volume: 0.5, blend: 1,
			data: `
Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Gb4/Bb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Db4/Gb4/Bb4 4,
Eb4/Gb4/Bb4 4
`},
		insData{name: "bass",
			volume: 1, blend: 0.5,
			data: `
Eb2 4,
Cb2 4,
Bb2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Bb2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Db2 4,
Eb2 4
`},
	}}

var SerpentsWake = songData{
	name: "Serpent's Wake",
	bpm:  115,
	ins: []insData{
		insData{name: "lead",
			volume: 0.5, blend: 1,
			data: `
Eb4 1, Fb4 1, Gb4 2,
Ab4 1, Gb4 1, Fb4 2,
Eb4 1, Gb4 1, Ab4 2,
Bb4 2, Ab4 2,

Eb4 1, Fb4 1, Gb4 2,
Ab4 1, Gb4 1, Db5 2,
Cb5 1, Bb4 1, Ab4 2,
Gb4 4
`},
		insData{name: "harmony",
			volume: 0.5, blend: 1,
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
Cb4/Eb4/Gb4 4,
Db4/Fb4/Ab4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4
`},
		insData{name: "bass",
			volume: 1, blend: 0.5,
			data: `
Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Bb2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Bb2 4,
Eb2 4
`},
	}}

var AshenCrosswinds = songData{
	name: "Ashen Crosswinds",
	bpm:  125,
	ins: []insData{
		insData{name: "lead",
			volume: 0.5, blend: 1,
			data: `
Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Gb4 2,
Fb4 1, Gb4 1, Bb4 2,
Db5 2, Bb4 2,

Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Eb4 2,
Fb4 1, Gb4 1, Ab4 2,
Bb4 4
`},
		insData{name: "harmony",
			volume: 0.5, blend: 1,
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
Cb4/Eb4/Gb4 4,
Db4/Fb4/Ab4 4,
Eb4/Gb4/Bb4 4,

Eb4/Gb4/Bb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4
`},
		insData{name: "bass",
			volume: 1, blend: 0.5,
			data: `
Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Bb2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Db2 4,
Eb2 4,

Eb2 4,
Cb2 4,
Bb2 4,
Eb2 4
`},
	}}

var GuildedStorm = songData{
	name: "The Gilded Storm",
	bpm:  110,
	ins: []insData{
		insData{name: "lead",
			volume: 0.5, blend: 1,
			data: `
Eb4 1, Gb4 1, Ab4 1, Bb4 1,
Db5 1, Bb4 1, Ab4 1, Gb4 1,
Eb4 2, Fb4 2,
Gb4 1, Ab4 1, Bb4 1, Db5 1,
Bb4 1.5, Ab4 0.5, Gb4 1, Eb4 1,
Fb4 2, Gb4 2,
Ab4 1, Bb4 1, Db5 2,
Bb4 4,

Eb4 1, Gb4 1, Ab4 1, Bb4 1,
Db5 1, Bb4 1, Ab4 1, Gb4 1,
Eb4 2, Fb4 2,
Gb4 1, Ab4 1, Bb4 1, Db5 1,
Bb4 1.5, Ab4 0.5, Gb4 1, Eb4 1,
Fb4 2, Gb4 2,
Ab4 1, Bb4 1, Db5 2,
Eb5 4
`},

		insData{name: "harmony",
			volume: 0.5, blend: 1,
			data: `
Eb4/Gb4/Bb4 4,
Eb4/Gb4/Bb4 4,
Gb4/Bb4/Db5 4,
Gb4/Bb4/Db5 4,
Cb4/Eb4/Gb4 4,
Cb4/Eb4/Gb4 4,
Bb3/Db4/Fb4 4,
Bb3/Db4/Fb4 4,
Eb4/Gb4/Bb4 4,
Eb4/Gb4/Bb4 4,
Gb4/Bb4/Db5 4,
Gb4/Bb4/Db5 4,
Cb4/Eb4/Gb4 4,
Cb4/Eb4/Gb4 4,
Eb4/Gb4/Bb4 4,
Eb4/Gb4/Bb4 4
`},

		insData{name: "bass",
			volume: 1, blend: 0.5,
			data: `
Eb2 4,
Eb2 4,
Gb2 4,
Gb2 4,
Cb2 4,
Cb2 4,
Bb2 4,
Bb2 4,
Eb2 4,
Eb2 4,
Gb2 4,
Gb2 4,
Cb2 4,
Cb2 4,
Eb2 4,
Eb2 4
`},
	}}

var CrimsonTempest = songData{
	name: "Crimson Tempest",
	bpm:  100,
	ins: []insData{
		{
			name:   "lead",
			volume: 0.5, blend: 1,
			data: `
	Eb4 1, Gb4 1, Ab4 1, Bb4 1,
	Db5 2, Bb4 2,
	Ab4 1, Gb4 1, Eb4 2,
	Fb4 2, Gb4 2,
	
	Eb4 1, Gb4 1, Ab4 1, Bb4 1,
	Db5 1, Fb5 1, Eb5 2,
	Db5 1, Bb4 1, Ab4 2,
	Gb4 4
	`,
		},
		{
			name:   "harmony",
			volume: 0.5, blend: 1,
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
	Gb4/Bb4/Db5 4,
	Ab4/Cb5/Eb5 4,
	Eb4/Gb4/Bb4 4,
	
	Eb4/Gb4/Bb4 4,
	Cb4/Eb4/Gb4 4,
	Db4/Gb4/Bb4 4,
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
	Cb2 4,
	
	Db2 4,
	Db2 4,
	Bb2 4,
	Bb2 4,
	
	Eb2 4,
	Eb2 4,
	Gb2 4,
	Gb2 4,
	
	Ab2 4,
	Ab2 4,
	Eb2 4,
	Eb2 4
	`,
		},
		{
			name:   "descant",
			volume: 1, blend: 0,
			data: `
	Eb5 0.5, Gb5 0.5, Ab5 0.5, Bb5 0.5,
	Db6 0.25, Bb5 0.25, Ab5 0.25, Gb5 0.25,
	NN 0.5, Eb5 0.5, Fb5 0.5, Gb5 0.5,
	Ab5 1, NN 1,
	
	Eb5 0.25, Gb5 0.25, Ab5 0.25, Bb5 0.25,
	Db6 0.25, Fb5 0.25, Eb5 0.25, Db5 0.25,
	NN 1,
	Bb5 0.5, Ab5 0.5, Gb5 0.5, Eb5 0.5,
	NN 1,
	Fb5 2
	`,
		},
		{
			name:   "pad",
			volume: 1, blend: 0.5,
			data: `
	Eb4/Gb4/Bb4 1, NN 1,
	Eb4/Gb4/Bb4 1, NN 1,
	Gb4/Bb4/Db5 0.5, NN 0.5, Gb4/Bb4/Db5 0.5, NN 0.5,
	Eb4/Gb4/Bb4 2,
	
	Eb4/Gb4/Bb4 1, NN 1,
	Cb4/Eb4/Gb4 1, NN 1,
	Db4/Gb4/Bb4 1, NN 1,
	Eb4/Gb4/Bb4 2,
	
	Eb4/Gb4/Bb4 0.5, NN 0.5, Eb4/Gb4/Bb4 0.5, NN 0.5,
	Gb4/Bb4/Db5 1, NN 1,
	Ab4/Cb5/Eb5 1, NN 1,
	Eb4/Gb4/Bb4 2,
	
	Eb4/Gb4/Bb4 1, NN 1,
	Cb4/Eb4/Gb4 1, NN 1,
	Db4/Gb4/Bb4 1, NN 1,
	Eb4/Gb4/Bb4 2
	`,
		},
	},
}

var MidnightDepths = songData{
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
			name: "bass",
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
	},
}

var MidnightDepthsPercussion = songData{
	name: "Midnight Depths (Percussive)",
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
