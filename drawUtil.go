package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) doFade(screen *ebiten.Image, dur time.Duration, fadeFrom color.NRGBA, fadeIn bool) {
	sinceStart := time.Since(g.fadeStart)
	if sinceStart < dur {
		durMS := float64(dur.Milliseconds())
		sinceMS := float64(sinceStart.Milliseconds())

		var amount uint8
		if !fadeIn {
			amount = uint8((sinceMS / durMS) * 255.0)
		} else {
			amount = 255 - uint8((sinceMS/durMS)*255.0)
		}
		fadeColor := color.NRGBA{R: fadeFrom.R, G: fadeFrom.G, B: fadeFrom.B, A: amount}
		vector.DrawFilledRect(screen, 0, 0, dWinWidth, dWinWidth, fadeColor, false)
	}
}
