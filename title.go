package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var vflop, hflop bool

func (g *Game) drawTitle(screen *ebiten.Image) {

	time := time.Now().Unix()
	vector.DrawFilledRect(screen, 0, 0, dWinWidth, dWinHeight/2, g.colors.sky, false)
	vector.DrawFilledRect(screen, 0, dWinHeight/2, dWinWidth, dWinHeight/2, g.colors.water, false)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(magScale, magScale)
	offset := point{}
	if time%3 == 0 {
		offset.Y = 1
	}
	op.GeoM.Translate(
		float64((dWinWidth/2)-(boatSP.Bounds().Dx()*magScale)/2+offset.X*magScale),
		float64((dWinHeight/2)-(boatSP.Bounds().Dy()*magScale)/2+offset.Y*magScale))
	screen.DrawImage(boatSP, op)

}
