package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const titleFadeTime = time.Minute * 2

func (g *Game) drawTitle(screen *ebiten.Image) {

	unix := time.Now().Unix()
	//Sky, water, horizon
	vector.DrawFilledRect(screen, 0, 0, dWinWidth, dWinHeight/2-(magScale),
		GetFadeColor(g.colors.day.sky, g.colors.evening.sky, titleFadeTime), false)
	vector.DrawFilledRect(screen, 0, dWinHeight/2, dWinWidth, dWinHeight/2,
		GetFadeColor(g.colors.day.water, g.colors.evening.water, titleFadeTime), false)
	vector.DrawFilledRect(screen, 0, dWinHeight/2-(magScale), dWinWidth, magScale,
		GetFadeColor(g.colors.day.horizon, g.colors.evening.horizon, titleFadeTime), false)

	//Draw boat
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(magScale, magScale)
	offset := point{}
	if unix%3 == 0 {
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

	//Text
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(magScale, magScale)
	op.GeoM.Translate(
		float64((dWinWidth/2)-(titleSP.Bounds().Dx()*magScale)/2),
		float64((dWinHeight/4)-(titleSP.Bounds().Dy()*magScale)/2))
	screen.DrawImage(titleSP, op)

}
