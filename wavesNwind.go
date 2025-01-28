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
	maxWaves          = 200
	minWaveLifeMS     = 1000
	maxWaveLifeRandMS = 200
)

var (
	waves    []waveData
	numWaves int
)

func drawWaves(g *Game, screen *ebiten.Image) {
	testColor := color.NRGBA{R: 255, G: 255, B: 255, A: 64}

	for _, wave := range waves {
		var x, y, width float32 = dWinWidth - float32(math.Mod(wave.linVal+g.boatPos.X/dWinWidth, 1)*dWinWidth), float32((dWinHeight / 2) + (wave.logVal * (dWinHeight / 2))), float32(1 + (wave.logVal * logVal))
		vector.DrawFilledRect(screen, x, y, width, 1, testColor, false)
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
	if numWaves < maxWaves {
		newWave := waveData{
			logVal: logDist(rand.Float64()),
			linVal: rand.Float64(),
			start:  time.Now(),
			life:   time.Millisecond * time.Duration(minWaveLifeMS+(rand.Float64()*maxWaveLifeRandMS))}
		waves = append(waves, newWave)
		numWaves++
	}
}

const logVal = 5.0

func logDist(uniform float64) float64 {
	biased := math.Abs(math.Log(1 - uniform))
	return math.Min(biased/logVal, 1.0)
}
