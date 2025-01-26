package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var vflop, hflop bool

func (g *Game) drawTitle(screen *ebiten.Image) {

	time := time.Now().Unix()
	//Sky, water, horizon
	vector.DrawFilledRect(screen, 0, 0, dWinWidth, dWinHeight/2, g.colors.sky, false)
	vector.DrawFilledRect(screen, 0, dWinHeight/2, dWinWidth, dWinHeight/2, g.colors.water, false)
	vector.DrawFilledRect(screen, 0, dWinHeight/2, dWinWidth, magScale, g.colors.horizon, false)

	//Draw boat
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(magScale, magScale)
	offset := point{}
	if time%3 == 0 {
		offset.Y = 1
	}
	op.GeoM.Translate(
		float64((dWinWidth/2)-(boatSP.Bounds().Dx()*magScale)/2+offset.X*magScale),
		float64((dWinHeight/2)-(boatSP.Bounds().Dy()*magScale)/2+offset.Y*magScale)+magScale*2)
	screen.DrawImage(boatSP, op)

	//Sun
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(magScale, magScale)
	op.GeoM.Translate(32*magScale, 8*magScale)
	screen.DrawImage(sunSP, op)

}
