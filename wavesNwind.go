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
	maxWaves          = 4000
	spawnPerFrame     = maxWaves / 60
	minWaveLifeMS     = 100
	maxWaveLifeRandMS = 500
)

var (
	waves    []waveData
	numWaves int
)

func drawWaves(g *Game, screen *ebiten.Image) {

	for _, wave := range waves {
		//Lower alpha for waves that are farther away
		alpha := uint8(10 * (1 + wave.logVal*persVal))
		waveColor := color.NRGBA{R: 255, G: 255, B: 255, A: alpha}

		bPos := g.boatPos.X * wave.logVal

		//Inverse X for correct direction, then modulo1(wavex + boatx / winwidth), then * winWidth to re-expand
		var x float32 = dWinWidth - float32(math.Mod(wave.linVal+bPos/dWinWidth, 1)*dWinWidth)
		//Start at horizon, add logVal * half winHeight
		var y float32 = float32(wave.logVal * (dWinHeight / 2))
		//Width is based on logVal
		var width float32 = float32(1 + (wave.logVal * persVal))

		vector.DrawFilledRect(screen, x, (dWinHeight/2)+y, width, 1, waveColor, false)
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
			logVal: logDist(rand.Float64()),
			linVal: rand.Float64(),
			start:  time.Now(),
			life:   time.Millisecond * time.Duration(minWaveLifeMS+(rand.Float64()*maxWaveLifeRandMS))}
		waves = append(waves, newWave)
		numWaves++
	}
}

func logDist(uniform float64) float64 {
	biased := math.Abs(math.Log(1 - uniform))
	return math.Min(biased/persVal, 1.0)
}
