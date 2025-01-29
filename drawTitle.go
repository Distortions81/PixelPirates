package main

import (
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) drawTitle(screen *ebiten.Image) {

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

	//Clouds -- TODO: render chunks and cache them
	xpos := g.boatPos.X * float64(islandVert/dWinWidth)
	if int(xpos) != lastCloudPos {
		lastCloudPos = int(xpos)
		var cBuf []byte
		for y := 0; y < dWinHeightHalf; y++ {
			for x := 0; x < dWinWidth; x++ {
				v := noiseMap(float32(x*2.0)+float32(xpos), float32(y-10*2.0), 0)
				vi := byte(v / 5 * 255)
				cBuf = append(cBuf, []byte{vi, vi, vi, vi}...)
			}
		}
		cloudbuf.WritePixels(cBuf)
	}
	screen.DrawImage(cloudbuf, nil)

	drawWaves(g, screen)
	drawAir(g, screen)

	// Island
	op := &ebiten.DrawImageOptions{}
	islandPos := dWinWidthHalf / 8
	op.GeoM.Translate(
		float64(islandPos),
		dWinHeightHalf-float64(island1SP.image.Bounds().Dy())+islandVert)
	screen.DrawImage(island1SP.image, op)

	//Island refection
	op.GeoM.Reset()
	op.GeoM.Scale(1, -(1 / refectionShrink))
	op.ColorScale.ScaleAlpha(refectionAlpha)
	op.GeoM.Translate(
		float64(islandPos),
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
		float64((dWinWidth/2)-(boatFrame.Bounds().Dx())/2),
		float64((dWinHeight/6)*4-(boatFrame.Bounds().Dy())/2+offset.Y)+float64(g.boatPos.Y))

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
		float64((dWinHeight/6)*5-(clickStartSP.image.Bounds().Dy())/2))
	op.ColorScale.ScaleAlpha(0.3)
	screen.DrawImage(clickStartSP.image, op)

	if g.gameMode == GAME_FADEOUT {
		fadeDur := time.Millisecond * 500
		g.doFade(screen, fadeDur, color.NRGBA{R: 255, G: 255, B: 255}, false)
		if time.Since(g.fadeStart) > fadeDur {
			g.fadeStart = time.Now()
			g.gameMode = GAME_PLAY
		}
	}

	g.doFade(screen, time.Millisecond*500, color.NRGBA{R: 255, G: 255, B: 255}, true)
}
