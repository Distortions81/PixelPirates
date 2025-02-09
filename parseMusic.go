package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	WAVE_SINE = iota
	WAVE_SQUARE
	WAVE_TRIANGLE
	WAVE_SAW
)

const deadNotesGiveup = 10

var sampleCache map[string]*audioData

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
			doLog(true, false, "Playing %v for %v.", song.name, song.length.Round(time.Second).String())
			playSong(g, &song)
			time.Sleep(time.Millisecond * 500)
		}
		time.Sleep(time.Millisecond * 500)
		doLog(true, true, "Restarting playlist...")
	}
}

const interval = time.Millisecond * (1000 / 32)

func playSong(g *Game, song *songData) {
	sampleCache = map[string]*audioData{}

	startTime := time.Now()

	numNotes := len(song.notes)
	var notes []ScheduledNote = make([]ScheduledNote, numNotes)
	copy(notes, song.notes)

	loops := 0
	lastTime := time.Now()

	for {
		if g.stopMusic {
			g.stopMusic = false
			break
		}

		if numNotes <= 1 {
			break
		}
		time.Sleep(interval - time.Since(lastTime))
		lastTime = time.Now()
		deadCount := 0
		for z := numNotes - 1; z > 0; z-- {
			loops++

			sn := notes[z]
			var freqs []int
			for _, freq := range sn.Frequency {
				freqs = append(freqs, int(freq))
			}
			buf := fmt.Sprintf("%v %v %v", freqs, sn.Duration.Milliseconds(), sn.ins.id)

			//Sleep until next note signature
			waitUntil := sn.Start - time.Since(startTime)
			if waitUntil > interval { //Too far into future
				deadCount++
				if deadCount > deadNotesGiveup {
					//At this point, all the notes are in the future.
					break
				}
				continue
				//Within reason, not too old
				//Otherwise skip the note
			} else if waitUntil <= interval {
				sc := sampleCache[buf]
				if sc != nil {
					playWave(g, true, *sc, true)
					//fmt.Printf("Used cached: %v\n", buf)
				} else {
					var notes []audioData
					for _, freq := range sn.Frequency {
						var noteWave audioData
						if freq > 0 {
							noteWave = generateWave(freq, sn.Duration, sn.waveform)
						} else if freq < 0 {
							noteWave = generateNoise(sn.Duration)
						} else {
							continue
						}
						//fmt.Printf("[%s] Playing wave=%v freq=%f for %v\n", sn.waveform, sn.InstrName, sn.Frequency, sn.Duration)

						notes = append(notes, noteWave)
					}

					var output audioData
					numNotes := len(notes)
					if numNotes > 1 {
						output = mixWaves(numNotes, int(float64(sampleRate)*sn.Duration.Seconds()), notes...)
					} else if numNotes == 1 {
						output = notes[0]
					} else {
						continue
					}

					output = applyADSR(output, sn.ins, sn.volume)
					sampleCache[buf] = cloneAudioData(output)

					//playWave(g, true, applyReverb(output, sn.volume, song.delay, song.feedback), true)
					playWave(g, true, output, true)
					fmt.Printf("RENDERED: %v\n", buf)
				}
			}

			//Delete note from list
			if numNotes > 1 {
				notes = append(notes[:z], notes[z+1:]...)
				numNotes--
			} else {
				notes = []ScheduledNote{}
				numNotes = 0
			}

		}

		took := time.Since(lastTime)
		if loops > 0 && took > time.Millisecond*2 {
			doLog(true, true, " render took: %v", took)
		}
	}
	sampleCache = map[string]*audioData{}
	doLog(true, true, "%v loops, cleared cached %v samples.", loops, len(sampleCache))
}

func cloneAudioData(data audioData) *audioData {
	cloned := make(audioData, len(data))
	copy(cloned, data)
	return &cloned
}

func parseSong(song *songData) {
	beatDuration := time.Minute / time.Duration(song.bpm)

	// 1) Build scheduled notes for all instruments
	var scheduled []ScheduledNote

	var songLength time.Duration
	for i, ins := range song.ins {
		song.ins[i].id = i
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
				ins:       &song.ins[i],
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
			return scheduled[i].Start > scheduled[j].Start
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
