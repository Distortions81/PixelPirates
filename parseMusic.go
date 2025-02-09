package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

func playMusicPlaylist(g *Game, gameMode int, songList []songData) {
	if *nomusic {
		return
	}
	if len(songList) == 0 {
		return
	}
	for {
		for _, song := range songList {
			if g.gameMode != gameMode {
				return
			}
			doLog(true, false, "Playing %v for %v.", song.name, song.length.String())
			playSong(g, &song)
		}
		doLog(true, true, "Restarting playlist...")
		time.Sleep(time.Second * 5)
	}
}

func playSong(g *Game, song *songData) {
	// 3) "Play" them in order
	startTime := time.Now()
	numIns := len(song.ins)

	for {
		for s, sn := range song.notes {
			if sn.played {
				continue
			}
			// How long until we reach sn.Start since the beginning?
			// This can be negative if we've already passed that time, so we clamp at 0.
			waitUntil := sn.Start - time.Since(startTime)
			if waitUntil > 0 {
				continue
			}

			// Now "play" the note
			fmt.Printf("[%s] Playing freq=%f for %v\n",
				sn.InstrName, sn.Frequency, sn.Duration)

			var notes []audioData
			for _, freq := range sn.Frequency {
				noteWave := make(audioData, int(float64(sampleRate)*sn.Duration.Seconds()))
				if freq > 0 {
					noteWave = generateWave(freq, sn.Duration, sn.waveform)
				} else if freq < 0 {
					noteWave = generateNoise(sn.Duration)
				}
				notes = append(notes, noteWave)
			}
			output := mixWaves(notes...)
			output = applyADSR(output, sn.ins, sn.volume*(1.0/float32(numIns)))
			song.notes[s].played = true
			go playWave(g, true, output)
		}
	}
}

func parseSong(song *songData) {
	beatDuration := time.Minute / time.Duration(song.bpm)

	// 1) Build scheduled notes for all instruments
	var scheduled []ScheduledNote

	var songLength time.Duration
	for i, ins := range song.ins {
		var elapsed time.Duration

		chords := strings.Split(ins.data, ",")
		for _, chord := range chords {
			note, duration := parseNote(chord)
			noteDuration := time.Duration(beatDuration.Seconds() * duration * float64(time.Second))
			chordNotes := strings.Split(note, "/")

			newChord := chordData{dur: noteDuration}
			if len(chordNotes) > 0 {
				for _, cnote := range chordNotes {
					freq := calculateFrequency(cnote)
					newChord.freq = append(newChord.freq, freq)
				}
			} else {
				newChord.freq = []float64{calculateFrequency(note)}
			}

			scheduled = append(scheduled, ScheduledNote{
				Start:     elapsed,
				Frequency: newChord.freq,
				Duration:  time.Duration(noteDuration),
				InstrName: ins.name,
				ins:       &ins,
				waveform:  ins.waveform,
				volume:    float32(ins.volume),
			})
			// Advance elapsed by the duration of this note
			elapsed += noteDuration
			if elapsed > songLength {
				songLength = elapsed
			}

			song.ins[i].chords = append(song.ins[i].chords, newChord)
		}

		// 2) Sort all scheduled notes by their start time
		sort.Slice(scheduled, func(i, j int) bool {
			return scheduled[i].Start < scheduled[j].Start
		})
	}

	song.notes = scheduled
	song.length = songLength
}

// parseNote parses a note string like "A4 1" into the note ("A4") and duration (1).
func parseNote(noteStr string) (string, float64) {
	parts := strings.Fields(noteStr)
	if len(parts) != 2 {
		return "", 0
	}
	note := parts[0]
	var duration float64
	_, err := fmt.Sscanf(parts[1], "%f", &duration)
	if err != nil {
		return "", 0
	}
	return note, duration
}

// calculateFrequency supports note names such as "A4", "C#4", "Db4", etc.
func calculateFrequency(note string) float64 {
	if note == "NN" {
		return 0 // rest
	}
	if note == "WN" {
		return -1 // white noise indicator
	}
	if len(note) < 2 {
		return 0
	}
	noteLetter := rune(note[0])
	var accidental rune
	var octave int
	var err error

	if len(note) >= 3 && (note[1] == '#' || note[1] == 'b') {
		accidental = rune(note[1])
		octave, err = strconv.Atoi(note[2:])
		if err != nil {
			return 0
		}
	} else {
		accidental = 0
		octave, err = strconv.Atoi(note[1:])
		if err != nil {
			return 0
		}
	}

	noteMap := map[rune]int{
		'C': 0,
		'D': 2,
		'E': 4,
		'F': 5,
		'G': 7,
		'A': 9,
		'B': 11,
	}
	base, ok := noteMap[noteLetter]
	if !ok {
		return 0
	}
	semitone := base
	if accidental == '#' {
		semitone++
	} else if accidental == 'b' {
		semitone--
	}
	offset := (octave-4)*12 + (semitone - 9)
	frequency := 440.0 * math.Pow(2, float64(offset)/12.0)
	return frequency
}
