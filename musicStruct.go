package main

import "time"

type playlistData []songData

type songData struct {
	name                    string
	ins                     []insData
	bpm                     int
	reverb, delay, feedback float32

	notes  []ScheduledNote
	length time.Duration
}

type insData struct {
	name, data     string
	id             int
	volume, square float64
	waveform       int
	/*
		Attack: The time it takes for a sound to go from silence to its full volume when a key is first pressed.
		Decay: The time it takes for the sound to drop from its peak volume to the sustain level.
		Sustain: The constant volume level maintained while a key is held down.
		Release: The time it takes for the sound to fade from the sustain level to silence when the key is released.
	*/
	attack, decay, sustain, release float64

	chords []chordData
}

type chordData struct {
	start time.Time
	freq  []float64
	dur   time.Duration
}

// ScheduledNote wraps a Note with a "start time" so we know when to play it.
type ScheduledNote struct {
	Start     time.Duration
	Frequency []float64
	Duration  time.Duration
	InstrName string

	played   bool
	ins      *insData
	volume   float32
	waveform int
}
