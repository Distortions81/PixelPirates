package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/chewxy/math32"
)

const maxVolume = 0.8

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
			startTime := time.Now()
			if *debugMode {
				doLog(true, true, "Rendering: '%v'", song.name)
			}
			output := playSong(song)

			if song.reverb > 0 {
				output = applyReverb(output, song.delay, song.feedback, song.reverb)
			}
			runtime.GC()
			doLog(true, true, "Render took %v -- Now Playing: %v. (%v sec)", time.Since(startTime).Round(time.Millisecond), song.name, len(output)/sampleRate)

			playWave(g, true, output)
		}
		doLog(true, true, "Restarting playlist...")
	}
}

func dumpMusic() {
	os.Mkdir("dump", 0755)
	for _, song := range titleSongList {
		if *debugMode {
			doLog(true, true, "Rendering: '%v'", song.name)
		}
		output := playSong(song)
		if song.reverb > 0 {
			output = applyReverb(output, song.delay, song.feedback, song.reverb)
		}
		saveMono16BitWav("dump/"+song.name+".wav", output)
	}
}

func playSong(song songData) audioData {
	if *nomusic {
		return nil
	}
	var (
		waves    []audioData
		waveLock sync.Mutex
		wg       sync.WaitGroup
	)
	for _, instrument := range song.ins {
		if instrument.volume <= 0 {
			continue
		}
		wg.Add(1)
		go func(ins insData) {
			insWave := generateFromText(&song, &ins)
			waveLock.Lock()
			waves = append(waves, insWave)
			waveLock.Unlock()
			wg.Done()
		}(instrument)
	}
	wg.Wait()
	return mixWaves(waves...)
}

func generateFromText(song *songData, ins *insData) audioData {
	beatDuration := time.Minute / time.Duration(song.bpm)
	var finalWave audioData
	if *debugMode {
		doLog(true, true, "Rendering instrument: %v", ins.name)
	}
	for _, noteStr := range strings.Split(ins.data, ",") {
		note, duration := parseNote(noteStr)
		if note == "" {
			continue
		}
		noteDuration := time.Duration(beatDuration.Seconds() * duration * float64(time.Second))
		// Check for chord
		chordNotes := strings.Split(note, "/")
		if len(chordNotes) > 1 {
			chordWave := playChord(chordNotes, noteDuration, ins)
			finalWave = append(finalWave, chordWave...)
		} else {
			freq := calculateFrequency(note)
			noteWave := playNote(freq, noteDuration, ins)
			finalWave = append(finalWave, noteWave...)
		}
	}
	return finalWave
}

// playNote now uses the new waveform options.
func playNote(freq float32, duration time.Duration, ins *insData) audioData {
	if freq == 0 {
		return make(audioData, int(float64(sampleRate)*duration.Seconds()))
	}
	var wave audioData
	if freq == -1 {
		wave = generateNoise(duration)
	} else {
		waveformType := ins.waveform
		if waveformType == "" {
			waveformType = "mix"
		}
		wave = generateWave(freq, duration, waveformType, ins.square)
	}
	wave = applyADSR(wave, ins)
	for i := range wave {
		wave[i] *= float32(ins.volume)
	}
	return wave
}

// Global constant for white noise dB offset.
const noiseDBOffset = -3.0 // Adjust this value as needed

func generateNoise(duration time.Duration) audioData {
	numSamples := int(float64(sampleRate) * duration.Seconds())
	wave := make(audioData, numSamples)
	compensation := math32.Pow(10, float32(noiseDBOffset)/20)

	for i := 0; i < numSamples; i++ {
		sample := float32(rand.Float64()*2.0 - 1.0)
		wave[i] = sample * compensation
		// Repeat the sample for smoothing.
		for x := 0; x < 8; x++ {
			i++
			if i < numSamples {
				wave[i] = sample * compensation
			}
		}
	}
	return wave
}

// generateWave accepts a waveform type and blend factor.
func generateWave(freq float32, duration time.Duration, waveform string, blend float32) audioData {
	samples := int(float64(sampleRate) * duration.Seconds())
	wave := make(audioData, samples)
	period := 1.0 / float64(freq)

	var dbOffset float32
	switch waveform {
	case "sine":
		dbOffset = 0.0
	case "square":
		dbOffset = -6.0
	case "triangle":
		dbOffset = 0.0
	case "sawtooth":
		dbOffset = -12.0
	case "mix":
		dbOffset = blend*(-2.0) + (1-blend)*0.0
	default:
		dbOffset = blend*(-2.0) + (1-blend)*0.0
	}
	compensation := math32.Pow(10, dbOffset/20)

	for i := 0; i < samples; i++ {
		t := float32(i) / float32(sampleRate)
		sineVal := math32.Sin(2 * math32.Pi * freq * t)

		var squareVal float32
		if sineVal < 0 {
			squareVal = -1.0
		} else {
			squareVal = 1.0
		}

		tmod := float32(math.Mod(float64(t), period))
		triangleVal := 4*math32.Abs(tmod/float32(period)-0.5) - 1.0
		sawtoothVal := 2*float32(math.Mod(float64(t), period))/float32(period) - 1.0

		var sample float32
		switch waveform {
		case "sine":
			sample = sineVal
		case "square":
			sample = squareVal
		case "triangle":
			sample = triangleVal
		case "sawtooth":
			sample = sawtoothVal
		case "mix":
			sample = blend*squareVal + (1-blend)*sineVal
		default:
			sample = blend*squareVal + (1-blend)*sineVal
		}
		wave[i] = sample * compensation
	}
	return wave
}

func playChord(chord []string, duration time.Duration, ins *insData) audioData {
	var waves []audioData
	for _, note := range chord {
		freq := calculateFrequency(note)
		var noteWave audioData
		if freq == 0 {
			// Generate silence: an array of zeros of the proper length.
			noteWave = make(audioData, int(float64(sampleRate)*duration.Seconds()))
		} else {
			noteWave = generateWave(freq, duration, ins.waveform, ins.square)
		}
		noteWave = applyADSR(noteWave, ins)
		for i := range noteWave {
			noteWave[i] *= float32(ins.volume)
		}
		waves = append(waves, noteWave)
	}
	chordWave := make(audioData, len(waves[0]))
	for i := range chordWave {
		var sum float32
		for _, w := range waves {
			sum += w[i]
		}
		chordWave[i] = sum / float32(len(waves))
	}
	return chordWave
}

// mixWaves using the maximum length.
func mixWaves(waves ...audioData) audioData {
	var maxLen int
	for _, w := range waves {
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	mixed := make(audioData, maxLen)
	for _, w := range waves {
		for i := 0; i < maxLen; i++ {
			if i < len(w) {
				mixed[i] += w[i]
			}
		}
	}
	numWaves := float32(len(waves))
	if numWaves > 1.0 {
		for i := 0; i < maxLen; i++ {
			mixed[i] /= numWaves
		}
	}

	var maxAmp float32
	for _, sample := range mixed {
		absVal := sample
		if absVal < 0 {
			absVal = -absVal
		}
		if absVal > maxAmp {
			maxAmp = absVal
		}
	}
	if maxAmp > 0 {
		scale := maxVolume / maxAmp
		for i := range mixed {
			mixed[i] *= scale
		}
	}
	return mixed
}

func playWave(g *Game, music bool, wave audioData) {
	soundData := make([]byte, len(wave)*2)
	var prevError float32

	for i, sample := range wave {
		shapedSample := sample + 0.5*prevError
		if shapedSample > 1.0 {
			shapedSample = 1.0
		} else if shapedSample < -1.0 {
			shapedSample = -1.0
		}
		intVal := int16(math32.Round(shapedSample * 32767))
		soundData[i*2] = byte(intVal)
		soundData[i*2+1] = byte(intVal >> 8)
		actual := float32(intVal) / 32767.0
		prevError = shapedSample - actual
	}
	player := g.audioContext.NewPlayerFromBytes(soundData)
	player.Play()

	for player.IsPlaying() {
		if music && g.stopMusic {
			g.stopMusic = false
			fadeout := 50
			for x := 1; x < fadeout; x++ {
				volume := 1 - float64(x)/float64(fadeout)
				player.SetVolume(volume)
				time.Sleep(time.Millisecond * 10)
			}
			player.Close()
			return
		}
		time.Sleep(time.Millisecond * 10)
	}
}

// calculateFrequency supports note names such as "A4", "C#4", "Db4", etc.
func calculateFrequency(note string) float32 {
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
	frequency := 440.0 * math32.Pow(2, float32(offset)/12.0)
	return frequency
}

// / applyADSR applies an ADSR envelope to the input waveform 'wave'
// and guarantees that the returned output is exactly the same length as 'wave'.
// If the sum of attack, decay, and release (in samples) is less than the total length,
// the remaining samples are assigned to sustain.
// If the sum is greater than the total length, the attack, decay, and release are scaled down proportionally.
func applyADSR(wave audioData, ins *insData) audioData {
	totalSamples := len(wave)

	// Calculate the nominal phase lengths.
	nomAttackSamples := int(float32(sampleRate) * ins.attack)
	nomDecaySamples := int(float32(sampleRate) * ins.decay)
	nomReleaseSamples := int(float32(sampleRate) * ins.release)

	// The sum of the non-sustain phases.
	phaseSum := nomAttackSamples + nomDecaySamples + nomReleaseSamples

	var attackSamples, decaySamples, releaseSamples, sustainSamples int

	if phaseSum > totalSamples {
		// Not enough samples for full phases; scale them down proportionally.
		scale := float32(totalSamples) / float32(phaseSum)
		attackSamples = int(float32(nomAttackSamples) * scale)
		decaySamples = int(float32(nomDecaySamples) * scale)
		releaseSamples = int(float32(nomReleaseSamples) * scale)
		sustainSamples = 0 // No room for sustain.
	} else {
		attackSamples = nomAttackSamples
		decaySamples = nomDecaySamples
		releaseSamples = nomReleaseSamples
		// All remaining samples become sustain.
		sustainSamples = totalSamples - phaseSum
	}

	// Create the envelope array.
	envelope := make([]float32, totalSamples)
	index := 0

	// Attack phase: ramp linearly from 0 to 1.
	for i := 0; i < attackSamples && index < totalSamples; i++ {
		envelope[index] = float32(i) / float32(attackSamples)
		index++
	}

	// Decay phase: ramp from 1 down to ins.sustain.
	for i := 0; i < decaySamples && index < totalSamples; i++ {
		t := float32(i) / float32(decaySamples)
		envelope[index] = 1 - (1-ins.sustain)*t
		index++
	}

	// Sustain phase: constant at ins.sustain.
	for i := 0; i < sustainSamples && index < totalSamples; i++ {
		envelope[index] = ins.sustain
		index++
	}

	// Release phase: ramp from ins.sustain down to 0.
	for i := 0; i < releaseSamples && index < totalSamples; i++ {
		t := float32(i) / float32(releaseSamples)
		if t > 1 {
			t = 1
		}
		envelope[index] = ins.sustain * (1 - t)
		index++
	}

	// If for any reason we have not filled all samples (due to rounding), fill the remainder with 0.
	for index < totalSamples {
		envelope[index] = 0
		index++
	}

	// Apply the envelope to each sample.
	output := make(audioData, totalSamples)
	for i := 0; i < totalSamples; i++ {
		output[i] = wave[i] * envelope[i]
	}

	return output
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
