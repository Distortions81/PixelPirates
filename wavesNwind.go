package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type waveData struct {
	pos    iPoint
	length int

	start time.Time
	life  time.Duration
}

const maxWaves = 20

var (
	waves    []waveData
	numWaves int
)

func drawWaves(g *Game, screen *ebiten.Image) {
	testColor := color.NRGBA{R: 255, G: 255, B: 255, A: 64}

	for _, wave := range waves {
		vector.DrawFilledRect(screen, float32(wave.pos.X)-g.boatPos.X, float32(wave.pos.Y), float32(wave.length), 1, testColor, false)
	}
}

func (g Game) makeWave() {
	if numWaves > maxWaves {
		for i := numWaves - 1; i >= 0; i-- {
			if time.Since(waves[i].start) > waves[i].life {
				// Remove the element at index i
				waves = append(waves[:i], waves[i+1:]...)
				numWaves--
			}
		}
	} else {
		newWave := waveData{pos: iPoint{X: int(rand.Float64()*dWinWidth + float64(g.boatPos.X)),
			Y: int(rand.Float64()*dWinHeight/2) + dWinHeight/2}, length: 1 + int(rand.Float64()*5),
			start: time.Now(), life: time.Millisecond * time.Duration(100+(rand.Float64()*200))}
		waves = append(waves, newWave)
		numWaves++
	}
}
