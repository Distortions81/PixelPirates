package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) drawTitle(screen *ebiten.Image) {
	g.makeWave()
	//g.makeAirWave()

	drawWorldGrad(g, screen)
	drawSun(g, screen)
	//drawClouds(g, screen)
	drawWaves(g, screen)
	//drawIslands(g, screen)
	//drawAir(g, screen)
	drawBoat(g, screen)

	if !g.modeTransition {

		/* Load if needed */
		if g.titleSP.image == nil {
			loadSprite(g.titleSP.Fullpath, g.titleSP, true)
		}
		if g.clickStartSP.image == nil {
			loadSprite(g.clickStartSP.Fullpath, g.clickStartSP, true)
		}
		//Text
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(
			float64(float64(dWinWidth/2.0)-float64(g.titleSP.image.Bounds().Dx())/2.0),
			float64(float64(dWinHeight/4.0)-float64(g.titleSP.image.Bounds().Dy())/2.0))
		//op.ColorScale.ScaleAlpha(0.8)
		screen.DrawImage(g.titleSP.image, op)

		//Click message
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(
			float64((dWinWidth/2)-(g.clickStartSP.image.Bounds().Dx())/2),
			float64((dWinHeight/6)*5-(g.clickStartSP.image.Bounds().Dy())/2))
		op.ColorScale.ScaleAlpha(0.3)
		screen.DrawImage(g.clickStartSP.image, op)
	}
}
