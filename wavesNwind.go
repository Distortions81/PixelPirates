package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type waveData struct {
	logVal, linVal float64
	start          time.Time
	life           time.Duration
}

const (
	persVal           = 10 //Used for perspective
	skyPersVal        = 5  //Used for perspective (airwaves)
	colorVal          = 10 //Used for perspective (waves)
	maxWaves          = 4000
	spawnPerFrame     = maxWaves / 60
	minWaveLifeMS     = 100
	maxWaveLifeRandMS = 500

	maxAirWaves          = 10
	minAirWaveLifeMS     = 2000
	maxAirWaveLifeRandMS = 4000
)

var (
	waves                 []waveData
	airwave               []waveData
	numWaves, numAirWaves int
)

func drawAir(g *Game, screen *ebiten.Image) {

	for _, wave := range airwave {
		//Dim alpha with time
		lifeLeft := float64(wave.life-time.Since(wave.start)) / float64(wave.life) * 2
		//Lower alpha for waves that are farther away
		alpha := uint8(math.Min(math.Abs(255-lifeLeft*255), 255.0))
		waveColor := color.NRGBA{R: 255, G: 255, B: 255, A: 85 - alpha/3}

		bPos := g.boatPos.X * wave.logVal

		// Calculate the raw modulo
		rawMod := wave.linVal + bPos/dWinWidth
		modVal := math.Mod(rawMod, 1)

		// Ensure the modulo result is positive
		if modVal < 0 {
			modVal += 1
		}

		// Inverse X for correct direction, then apply the positive modulo, then re-expand
		x := dWinWidth - float32(modVal*dWinWidth)

		// Start at horizon, add logVal * half winHeight
		y := float32(wave.logVal * dWinHeightHalf)

		// Width is based on logVal
		width := float32(1 + (wave.logVal * persVal))

		vector.DrawFilledRect(screen, x, dWinHeightHalf-y+1, width, 1, waveColor, false)
	}
}

func drawWaves(g *Game, screen *ebiten.Image) {
	for _, wave := range waves {
		// Lower alpha for waves that are farther away
		alpha := uint8(15 * (1 + wave.logVal*colorVal))
		waveColor := color.NRGBA{R: 255, G: 255, B: 255, A: alpha}

		bPos := g.boatPos.X * wave.logVal

		// Calculate the raw modulo
		rawMod := wave.linVal + bPos/dWinWidth
		modVal := math.Mod(rawMod, 1)

		// Ensure the modulo result is positive
		if modVal < 0 {
			modVal += 1
		}

		// Inverse X for correct direction, then apply the positive modulo, then re-expand
		x := dWinWidth - float32(modVal*dWinWidth)

		// Start at horizon, add logVal * half winHeight
		y := float32(wave.logVal * dWinHeightHalf)

		// Width is based on logVal
		width := float32(1 + (wave.logVal * persVal))

		vector.DrawFilledRect(screen, x, dWinHeightHalf+y, width, 1, waveColor, false)
	}
}

func (g Game) makeWave() {
	for i := numWaves - 1; i >= 0; i-- {
		if time.Since(waves[i].start) > waves[i].life {
			// Remove the element at index i
			waves = append(waves[:i], waves[i+1:]...)
			numWaves--
		}
	}
	for z := 0; z < spawnPerFrame && numWaves < maxWaves; z++ {
		newWave := waveData{
			logVal: logDistWave(rand.Float64()),
			linVal: rand.Float64(),
			start:  time.Now(),
			life:   time.Millisecond * time.Duration(minWaveLifeMS+(rand.Float64()*maxWaveLifeRandMS))}
		waves = append(waves, newWave)
		numWaves++
	}
}

func (g Game) makeAirWave() {
	for i := numAirWaves - 1; i >= 0; i-- {
		if time.Since(airwave[i].start) > airwave[i].life {
			// Remove the element at index i
			airwave = append(airwave[:i], airwave[i+1:]...)
			numAirWaves--
		}
	}
	for z := 0; z < spawnPerFrame && numAirWaves < maxAirWaves; z++ {
		newWave := waveData{
			logVal: logDistAirwave(rand.Float64()),
			linVal: rand.Float64(),
			start:  time.Now(),
			life:   time.Millisecond * time.Duration(minAirWaveLifeMS+(rand.Float64()*maxAirWaveLifeRandMS))}
		airwave = append(airwave, newWave)
		numAirWaves++
	}
}

func logDistWave(uniform float64) float64 {
	biased := math.Abs(math.Log(1 - uniform))
	return math.Min(biased/persVal, 1.0)
}

func logDistAirwave(uniform float64) float64 {
	biased := math.Abs(math.Log(1 - uniform))
	return math.Min(biased/skyPersVal, 1.0)
}
