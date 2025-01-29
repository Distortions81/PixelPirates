package main

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	//Object reflect
	refectionShrink     = 0.4
	refectionBlurAmount = 2
	refectionAlpha      = 0.15

	//Sun reflect -- disabled
	sunReflectHeight = 10.0
	sunReflectAlpha  = 0.8
	sunX             = 64.0

	//Water gradient
	waterStartColor   = 175
	waterHueShift     = 25
	waterBrightStart  = 0.9
	waterDarkenDivide = 3
	waterSaturate     = 0.8

	//Sky gradient
	skyStartColor   = 220
	skyHueShift     = -40
	skyBrightStart  = 1.0
	skyDarkenDivide = 2.5
	skySaturate     = 0.5

	//Island settings
	islandVert  = 6.0
	islandStart = -dWinWidthHalf

	cloudBlurAmountX  = 32
	cloudBlurAmountY  = 1
	cloudBlurStrech   = 1
	cloudReflectAlpha = 0.5
)

var (
	cloudbuf, cloudblur *ebiten.Image
	lastCloudPos        int = -1000
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
		color = HSVToRGB(HSV{
			H: skyStartColor + (vVal * skyHueShift),
			S: skySaturate,
			V: skyBrightStart - math.Min(((1.0-vVal)/skyDarkenDivide), 1.0)})
		vector.DrawFilledRect(screen, 0, y, dWinWidth, 1, color, false)
	}

	//Clouds -- TODO: render chunks and cache them
	xpos := g.boatPos.X * float64(islandVert/dWinWidth)
	if int(xpos) != lastCloudPos {
		lastCloudPos = int(xpos)
		var cBuf []byte
		for y := 0; y < dWinHeightHalf; y++ {
			for x := 0; x < dWinWidth; x++ {
				v := noiseMap(float32(x)+float32(xpos), float32((y-10)*4.0), 0)
				vi := byte(v / 5 * 255)
				cBuf = append(cBuf, []byte{vi, vi, vi, vi}...)
			}
		}
		cloudbuf.WritePixels(cBuf)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1.0/cloudBlurAmountX, 1.0/cloudBlurAmountY)
		op.Filter = ebiten.FilterLinear
		cloudblur.Clear()
		cloudblur.DrawImage(cloudbuf, op)
	}
	//Cloud reflection
	screen.DrawImage(cloudbuf, nil)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(cloudBlurAmountX, -cloudBlurAmountY/cloudBlurStrech)
	op.GeoM.Translate(0, dWinHeight*cloudBlurAmountY)
	op.ColorScale.ScaleAlpha(cloudReflectAlpha)
	//op.Blend = ebiten.BlendLighter
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(cloudblur, op)

	//Sun reflect -- Disabled until objects can block it
	/*
		subRect := image.Rectangle{
			Min: image.Point{X: 0, Y: sunSP.blurred.Bounds().Dy() / 2.0},
			Max: image.Point{X: sunSP.blurred.Bounds().Dx(), Y: sunSP.blurred.Bounds().Dy()},
		}
		op := &ebiten.DrawImageOptions{}
		op.Filter = ebiten.FilterLinear
		sub := sunSP.blurred.SubImage(subRect)
		op.GeoM.Reset()
		xScale, yScale := 1.0, sunReflectHeight
		op.GeoM.Scale(xScale, yScale)
		op.GeoM.Translate((float64(sub.Bounds().Dx())/xScale)+sunX, dWinHeightHalf)
		op.Blend = ebiten.BlendLighter
		op.ColorScale.ScaleAlpha(sunReflectAlpha)
		screen.DrawImage(sub.(*ebiten.Image), op)
	*/

	drawWaves(g, screen)
	drawAir(g, screen)

	// Island
	op = &ebiten.DrawImageOptions{}
	islandPos := g.boatPos.X*float64(islandVert/dWinWidth) - islandStart
	op.GeoM.Translate(
		dWinWidth-float64(islandPos),
		dWinHeightHalf-float64(island1SP.image.Bounds().Dy())+islandVert)
	screen.DrawImage(island1SP.image, op)

	//Island refection
	op.GeoM.Reset()
	op.GeoM.Scale(1, -(1 / refectionShrink))
	op.ColorScale.ScaleAlpha(refectionAlpha)
	op.GeoM.Translate(
		dWinWidth-float64(islandPos),
		dWinHeightHalf+float64(islandVert+island1SP.image.Bounds().Dy()-5)/refectionShrink)
	screen.DrawImage(island1SP.blurred, op)

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
	op.GeoM.Translate(float64(sunSP.image.Bounds().Dx())+sunX, 24)
	screen.DrawImage(sunSP.image, op)

	g.doFade(screen, time.Millisecond*500, color.NRGBA{R: 255, G: 255, B: 255}, true)
	buf := fmt.Sprintf("%4.0f,%3.0f (%3d,%3d)", g.boatPos.X, g.boatPos.Y, numWaves, numAirWaves)

	if *debug {
		ebitenutil.DebugPrint(screen, buf)
	}
}
