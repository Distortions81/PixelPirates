package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) drawTitle(screen *ebiten.Image) {

	drawWorldGrad(g, screen)
	drawSun(screen)
	drawCloudsNew(g, screen)
	drawWaves(g, screen)
	drawIsland(g, screen)
	drawAir(g, screen)
	drawBoat(g, screen)

	//Text
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64((dWinWidth/2)-(titleSP.image.Bounds().Dx())/2),
		float64((dWinHeight/4)-(titleSP.image.Bounds().Dy())/2))
	op.ColorScale.ScaleAlpha(0.3)
	screen.DrawImage(titleSP.image, op)

	//Click message
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64((dWinWidth/2)-(clickStartSP.image.Bounds().Dx())/2),
		float64((dWinHeight/6)*5-(clickStartSP.image.Bounds().Dy())/2))
	op.ColorScale.ScaleAlpha(0.3)
	screen.DrawImage(clickStartSP.image, op)

	if g.gameMode == GAME_FADEOUT {
		fadeDur := time.Millisecond * 500
		g.doFade(screen, fadeDur, color.NRGBA{R: 255, G: 255, B: 255}, false)
		if time.Since(g.fadeStart) > fadeDur {
			g.fadeStart = time.Now()
			g.boatPos = fPoint{X: 0, Y: 0}
			g.gameMode = GAME_PLAY
		}
	}

	g.doFade(screen, time.Millisecond*500, color.NRGBA{R: 255, G: 255, B: 255}, true)
}
