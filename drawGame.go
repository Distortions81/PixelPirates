package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const dayFadeTime = time.Minute * 3

func (g *Game) drawGame(screen *ebiten.Image) {

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
		float64((dWinWidth/4)-(boatFrame.Bounds().Dx())/2+offset.X)+float64(g.boatPos.X),
		float64((dWinHeight/6)*4.5-(boatFrame.Bounds().Dy())/2+offset.Y)+2+float64(g.boatPos.Y))

	screen.DrawImage(boatFrame, op)

	//Sun
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(64, 24)
	screen.DrawImage(sunSP.image, op)

	drawWaves(g, screen)

	g.doFade(screen, time.Millisecond*500, color.NRGBA{R: 255, G: 255, B: 255}, true)
}
