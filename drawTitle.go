package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) drawTitle(screen *ebiten.Image) {
	g.makeWave()
	g.makeAirWave()

	drawWorldGrad(g, screen)
	drawSun(screen)
	drawCloudsNew(g, screen)
	drawWaves(g, screen)
	drawIslands(g, screen)
	drawAir(g, screen)
	drawBoat(g, screen)

	if !g.modeTransition {
		//Text
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(
			float64((dWinWidth/2)-(titleSP.image.Bounds().Dx())/2),
			float64((dWinHeight/4)-(titleSP.image.Bounds().Dy())/2))
		//op.ColorScale.ScaleAlpha(0.8)
		screen.DrawImage(titleSP.image, op)

		//Click message
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(
			float64((dWinWidth/2)-(clickStartSP.image.Bounds().Dx())/2),
			float64((dWinHeight/6)*5-(clickStartSP.image.Bounds().Dy())/2))
		op.ColorScale.ScaleAlpha(0.3)
		screen.DrawImage(clickStartSP.image, op)
	}
}
