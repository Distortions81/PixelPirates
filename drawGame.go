package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) drawGame(screen *ebiten.Image) {

	unix := time.Now().Unix()
	//Sky, water, horizon
	vector.DrawFilledRect(screen, 0, 0, dWinWidth, dWinHeight/2-1,
		GetFadeColor(g.colors.day.sky, g.colors.evening.sky, titleFadeTime), false)
	vector.DrawFilledRect(screen, 0, dWinHeight/2, dWinWidth, dWinHeight/2,
		GetFadeColor(g.colors.day.water, g.colors.evening.water, titleFadeTime), false)
	vector.DrawFilledRect(screen, 0, dWinHeight/2-(1), dWinWidth, 1,
		GetFadeColor(g.colors.day.horizon, g.colors.evening.horizon, titleFadeTime), false)

	//Draw boat
	op := &ebiten.DrawImageOptions{}
	offset := iPoint{}
	if unix%3 == 0 {
		offset.Y = 1
	}
	boatFrame := autoAnimatePingPong(boat1SP)
	op.GeoM.Translate(
		float64((dWinWidth/2)-(boatFrame.Bounds().Dx())/2+offset.X),
		float64((dWinHeight/2)-(boatFrame.Bounds().Dy())/2+offset.Y)+2)
	op.ColorScale.ScaleWithColor(
		GetFadeColor(
			color.RGBA{R: 255, G: 255, B: 255, A: 255},
			color.RGBA{R: 128, G: 128, B: 128, A: 255},
			titleFadeTime))
	screen.DrawImage(boatFrame, op)

	//Sun
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(32, 8)
	screen.DrawImage(sunSP.image, op)
}
