package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const titleFadeTime = time.Minute * 2

func (g *Game) drawTitle(screen *ebiten.Image) {

	unix := time.Now().Unix()
	//Sky, water, horizon
	vector.DrawFilledRect(screen, 0, 0, dWinWidth, dWinHeight/2-(1),
		g.colors.day.sky, false)
	vector.DrawFilledRect(screen, 0, dWinHeight/2, dWinWidth, dWinHeight/2,
		g.colors.day.water, false)
	vector.DrawFilledRect(screen, 0, dWinHeight/2-(1), dWinWidth, 1,
		g.colors.day.horizon, false)

	//Draw boat
	op := &ebiten.DrawImageOptions{}
	offset := iPoint{}
	if unix%3 == 0 {
		offset.Y = 1
	}
	boatFrame := autoAnimatePingPong(boat2SP)
	op.GeoM.Translate(
		float64((dWinWidth/2)-(boatFrame.Bounds().Dx())/2+offset.X),
		float64((dWinHeight/1.8)-(boatFrame.Bounds().Dy())/2+offset.Y)+2)

	screen.DrawImage(boatFrame, op)

	//Draw re-color flag
	op = &ebiten.DrawImageOptions{}
	offset = iPoint{}
	if unix%3 == 0 {
		offset.Y = 1
	}
	boatFrame = autoAnimatePingPong(boat2SP_flag)
	op.GeoM.Translate(
		float64((dWinWidth/2)-(boatFrame.Bounds().Dx())/2+offset.X),
		float64((dWinHeight/1.8)-(boatFrame.Bounds().Dy())/2+offset.Y)+2)
	op.ColorScale.ScaleWithColor(color.RGBA{R: 255, G: 160, B: 0, A: 255})

	screen.DrawImage(boatFrame, op)

	//Sun
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(64, 24)
	screen.DrawImage(sunSP.image, op)

	//Text
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64((dWinWidth/2)-(titleSP.image.Bounds().Dx())/2),
		float64((dWinHeight/4)-(titleSP.image.Bounds().Dy())/2))
	op.ColorScale.ScaleAlpha(0.3)
	screen.DrawImage(titleSP.image, op)

	//Click message
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64((dWinWidth/2)-(clickStartSP.image.Bounds().Dx())/2),
		float64((dWinHeight/4)*3-(clickStartSP.image.Bounds().Dy())/2))
	op.ColorScale.ScaleAlpha(0.3)
	screen.DrawImage(clickStartSP.image, op)

	drawWaves(g, screen)
}
