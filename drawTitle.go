package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) drawTitle(screen *ebiten.Image) {

	unix := time.Now().Unix()

	//Draw world grads
	if worldGradDirty {
		worldGradDirty = false
		g.drawWorldGrad()
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(dWinWidth, 1)
	screen.DrawImage(worldGradImg, op)

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
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(cloudBlurAmountX, -cloudBlurAmountY*cloudBlurStrech)
	op.GeoM.Translate(0, dWinHeight)
	op.ColorScale.ScaleAlpha(cloudReflectAlpha)
	//op.Blend = ebiten.BlendLighter
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(cloudblur, op)

	drawWaves(g, screen)
	drawAir(g, screen)

	// Island
	op = &ebiten.DrawImageOptions{}
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
	boatFrame := autoAnimatePingPong(boat2SP, 0)
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
