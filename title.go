package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) drawTitle(screen *ebiten.Image) {

	vector.DrawFilledRect(screen, 0, 0, dWinWidth, dWinHeight/2, g.colors.sky, false)
	vector.DrawFilledRect(screen, 0, dWinHeight/2, dWinWidth, dWinHeight/2, g.colors.water, false)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(magScale, magScale)
	op.GeoM.Translate(
		float64((dWinWidth/2)-(boatSP.Bounds().Dx()*magScale)/2),
		float64((dWinHeight/2)-(boatSP.Bounds().Dy()*magScale)/2))
	screen.DrawImage(boatSP, op)

}
