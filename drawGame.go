package main

import (
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	colorLogVal = 3 //Used for color
	cAmnt       = 80.0

	waterStartColor   = 175
	waterHueShift     = 25
	waterBrightStart  = 0.9
	waterDarkenDivide = 3
	waterSaturate     = 0.8

	skyStartColor   = 220
	skyHueShift     = -40
	skyBrightStart  = 1.0
	skyDarkenDivide = 2.5
	skySaturate     = 0.5

	islandVert  = 6.0
	islandStart = -dWinWidthHalf
)

func (g *Game) drawGame(screen *ebiten.Image) {

	unix := time.Now().Unix()

	//Horizon
	vector.DrawFilledRect(screen, 0, dWinHeightHalf-(1), dWinWidth, 1,
		g.colors.day.horizon, false)

	var y float32
	for y = 0; y < dWinHeightHalf; y++ {
		//Water
		color := color.RGBA{}
		vVal := (float64(y) / dWinHeightHalf)
		color = HSVToRGB(HSV{
			H: waterStartColor + (vVal * waterHueShift),
			S: waterSaturate,
			V: waterBrightStart - math.Min(vVal/waterDarkenDivide, 1.0)})
		vector.DrawFilledRect(screen, 0, dWinHeightHalf+y,
			dWinWidth, 1, color, false)

		//Sky
		//Water
		color = HSVToRGB(HSV{
			H: skyStartColor + (vVal * skyHueShift),
			S: skySaturate,
			V: skyBrightStart - math.Min(((1.0-vVal)/skyDarkenDivide), 1.0)})
		vector.DrawFilledRect(screen, 0, y, dWinWidth, 1, color, false)
	}

	drawWaves(g, screen)

	// Island
	op := &ebiten.DrawImageOptions{}
	islandPos := g.boatPos.X*float64(islandVert/dWinWidth) - islandStart
	op.GeoM.Translate(dWinWidth-float64(islandPos),
		dWinHeightHalf-float64(island1SP.image.Bounds().Dy())+islandVert)
	screen.DrawImage(island1SP.image, op)

	//Draw boat
	op = &ebiten.DrawImageOptions{}
	offset := iPoint{}
	if unix%3 == 0 {
		offset.Y = 1
	}
	boatFrame := autoAnimatePingPong(boat2SP)
	op.GeoM.Translate(
		float64((dWinWidth/4)-(boatFrame.Bounds().Dx())/2+offset.X),
		float64((dWinHeight/6)*4.5-(boatFrame.Bounds().Dy())/2+offset.Y)+float64(g.boatPos.Y))

	screen.DrawImage(boatFrame, op)

	//Sun
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(64, 24)
	screen.DrawImage(sunSP.image, op)

	g.doFade(screen, time.Millisecond*500, color.NRGBA{R: 255, G: 255, B: 255}, true)
}
