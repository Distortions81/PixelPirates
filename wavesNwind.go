package main

import (
	"image/color"
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

var waves []waveData

func drawWaves(g *Game, screen *ebiten.Image) {
	testColor := color.NRGBA{R: 255, G: 255, B: 255, A: 127}

	for _, wave := range waves {
		vector.DrawFilledRect(screen, float32(wave.pos.X), float32(wave.pos.Y), float32(wave.length), 1, testColor, false)
	}
}

func makeWave() {
}
