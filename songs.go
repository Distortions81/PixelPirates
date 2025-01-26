package main

var songList []*songData = []*songData{
	&GuildedStorm,
	&EbonyGale,
	&SailsOfDusk,
	&DarklingWaters,
	&NightfallApproch,
	&SerpentsWake,
	&AshenCrosswinds,
}

type songData struct {
	name                string
	lead, harmony, bass string
	bpm                 int
}

var EbonyGale = songData{
	name: "Aboard the Ebony Gale",
	bpm:  90,
	lead: `Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Gb4 2,
Fb4 1, Gb4 1, Ab4 2,
Bb4 2, Ab4 2,

Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Db5 2,
Cb5 1, Bb4 1, Ab4 2,
Gb4 4`,

	harmony: `Eb4/Gb4/Bb4 4,
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
Eb4/Gb4/Bb4 4`,

	bass: `Eb2 4,
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
}

var SailsOfDusk = songData{
	name: "Sails of Dusk",
	bpm:  120,
	lead: `
Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Gb4 2,
Fb4 1, Gb4 1, Ab4 2,
Bb4 2, Ab4 2,

Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Db5 2,
Cb5 1, Bb4 1, Ab4 2,
Gb4 4
`,
	harmony: `
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
`,
	bass: `
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
`,
}

var DarklingWaters = songData{
	name: "Darkling Waters",
	bpm:  130,
	lead: `
Eb4 1, Fb4 1, Gb4 2,
Ab4 1, Gb4 1, Fb4 2,
Eb4 1, Gb4 1, Bb4 2,
Db5 2, Bb4 2,

Eb4 1, Fb4 1, Gb4 2,
Ab4 1, Gb4 1, Eb4 2,
Fb4 1, Gb4 1, Ab4 2,
Bb4 4
`,
	harmony: `
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
`,
	bass: `
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
`,
}

var NightfallApproch = songData{
	name: "Nightfall Approach",
	bpm:  110,
	lead: `
Eb4 1, Gb4 1, Bb4 2,
Ab4 1, Gb4 1, Eb4 2,
Fb4 2, Gb4 2,
Ab4 1, Gb4 1, Fb4 2,
Eb4 2, Gb4 2,

Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Gb4 2,
Eb4 1, Gb4 1, Ab4 2,
Bb4 4
`,
	harmony: `
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
`,
	bass: `
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
`,
}

var SerpentsWake = songData{
	name: "Serpent's Wake",
	bpm:  115,
	lead: `
Eb4 1, Fb4 1, Gb4 2,
Ab4 1, Gb4 1, Fb4 2,
Eb4 1, Gb4 1, Ab4 2,
Bb4 2, Ab4 2,

Eb4 1, Fb4 1, Gb4 2,
Ab4 1, Gb4 1, Db5 2,
Cb5 1, Bb4 1, Ab4 2,
Gb4 4
`,
	harmony: `
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
`,
	bass: `
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
`,
}

var AshenCrosswinds = songData{
	name: "Ashen Crosswinds",
	bpm:  125,
	lead: `
Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Gb4 2,
Fb4 1, Gb4 1, Bb4 2,
Db5 2, Bb4 2,

Eb4 1, Gb4 1, Ab4 2,
Bb4 1, Ab4 1, Eb4 2,
Fb4 1, Gb4 1, Ab4 2,
Bb4 4
`,
	harmony: `
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
`,
	bass: `
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
`,
}

var GuildedStorm = songData{
	name: "The Gilded Storm",
	bpm:  110,

	lead: `
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
`,

	harmony: `
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
`,

	bass: `
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
`,
}
