package main

var songList []*songData = []*songData{
	&songOne,
	&songTwo,
	&songThree,
	&songFour,
	&songFive,
}

type songData struct {
	name                string
	lead, harmony, bass string
	bpm                 int
}

var songOne = songData{
	name: "Aboard the Ebony Gale",
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

var songTwo = songData{
	name: "Crimson Tides",
	lead: `Bb4 1, Cb5 1, Db5 2,
Eb5 1, Db5 1, Bb4 2,
Fb5 2, Eb5 2,
Db5 1, Bb4 1, Ab4 2,
Gb4 2, Ab4 2,
Bb4 1, Db5 1, Eb5 2,
Fb5 2, Eb5 2,
Bb4 4`,

	harmony: `Bb4/Db5/Fb5 4,
Gb4/Bb4/Db5 4,
Ab4/Cb5/Eb5 4,
Bb4/Db5/Fb5 4,

Bb4/Db5/Fb5 4,
Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Bb4/Db5/Fb5 4,

Bb4/Db5/Fb5 4,
Gb4/Bb4/Db5 4,
Ab4/Cb5/Eb5 4,
Bb4/Db5/Fb5 4,

Bb4/Db5/Fb5 4,
Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Bb4/Db5/Fb5 4`,

	bass: `Bb4/Db5/Fb5 4,
Gb4/Bb4/Db5 4,
Ab4/Cb5/Eb5 4,
Bb4/Db5/Fb5 4,

Bb4/Db5/Fb5 4,
Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Bb4/Db5/Fb5 4,

Bb4/Db5/Fb5 4,
Gb4/Bb4/Db5 4,
Ab4/Cb5/Eb5 4,
Bb4/Db5/Fb5 4,

Bb4/Db5/Fb5 4,
Eb4/Gb4/Bb4 4,
Db4/Fb4/Ab4 4,
Bb4/Db5/Fb5 4`,
}

var songThree = songData{
	name: "Ghostlights in the Fog",
	lead: `Gb4 1, Ab4 1, Cb5 2,
Db5 1, Cb5 1, Gb4 2,
Ebb5 2, Db5 2,
Cb5 1, Ab4 1, Gb4 2,
Fb4 2, Gb4 2,
Ab4 1, Cb5 1, Db5 2,
Ebb5 2, Db5 2,
Gb4 4`,

	harmony: `Gb4/Bbb4/Db5 4,
Cb4/Ebb4/Gb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bbb4/Db5 4,

Gb4/Bbb4/Db5 4,
Cb4/Ebb4/Gb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bbb4/Db5 4,

Gb4/Bbb4/Db5 4,
Cb4/Ebb4/Gb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bbb4/Db5 4,

Gb4/Bbb4/Db5 4,
Cb4/Ebb4/Gb4 4,
Db4/Fb4/Ab4 4,
Gb4/Bbb4/Db5 4`,

	bass: `Gb2 4,
Cb2 4,
Db2 4,
Gb2 4,

Gb2 4,
Cb2 4,
Db2 4,
Gb2 4,

Gb2 4,
Cb2 4,
Db2 4,
Gb2 4,

Gb2 4,
Cb2 4,
Db2 4,
Gb2 4`,
}

var songFour = songData{
	name: "The Moonlit Maw",
	lead: `C#4 1, D#4 1, Fb4 2,
F#4 1, Fb4 1, D#4 2,
G#4 1, A4 1, Cb5 2,
C#5 2, A4 2,

C#4 1, D#4 1, F#4 2,
G#4 1, F#4 1, D#4 2,
Fb4 1, D#4 1, C#4 2,
Cb4 4`,

	harmony: `C#4/F#4/G#4 4,
D#4/A4/C#5 4,
Fb4/G#4/D#5 4,
C#4/F#4/G#4 4,

C#4/F#4/G#4 4,
D#4/A4/C#5 4,
Fb4/G#4/D#5 4,
C#4/F#4/G#4 4,

C#4/F#4/G#4 4,
D#4/A4/C#5 4,
Fb4/G#4/D#5 4,
C#4/F#4/G#4 4,

C#4/F#4/G#4 4,
D#4/A4/C#5 4,
Fb4/G#4/D#5 4,
C#4/F#4/G#4 4`,

	bass: `C#2 4,
D#2 4,
Fb2 4,
C#2 4,

C#2 4,
D#2 4,
Fb2 4,
C#2 4,

C#2 4,
D#2 4,
Fb2 4,
C#2 4,

C#2 4,
D#2 4,
Fb2 4,
C#2 4`,
}

var songFive = songData{
	name: "Iron Tides at Dawn",
	lead: `F#4 1, G#4 1, Bbb4 2,
Cb5 1, Bbb4 1, G#4 2,
Fb4 1, F#4 1, G#4 2,
C#5 2, G#4 2,

F#4 1, Bbb4 1, Cb5 2,
C#5 1, Cb5 1, Eb5 2,
Fb5 1, Eb5 1, C#5 2,
F#4 4`,

	harmony: `F#4/Bbb4/C#5 4,
G#4/Cb5/Fb5 4,
Bbb4/Eb5/F#5 4,
F#4/Bbb4/C#5 4,

F#4/Bbb4/C#5 4,
G#4/Cb5/Fb5 4,
Bbb4/Eb5/F#5 4,
F#4/Bbb4/C#5 4,

F#4/Bbb4/C#5 4,
G#4/Cb5/Fb5 4,
Bbb4/Eb5/F#5 4,
F#4/Bbb4/C#5 4,

F#4/Bbb4/C#5 4,
G#4/Cb5/Fb5 4,
Bbb4/Eb5/F#5 4,
F#4/Bbb4/C#5 4`,

	bass: `F#2 4,
G#2 4,
Bbb2 4,
F#2 4,

F#2 4,
G#2 4,
Bbb2 4,
F#2 4,

F#2 4,
G#2 4,
Bbb2 4,
F#2 4,

F#2 4,
G#2 4,
Bbb2 4,
F#2 4
`,
}
