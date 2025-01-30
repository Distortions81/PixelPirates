package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	//Island settings
	islandY      = 6.0
	islandStartX = -dWinWidthHalf

	//Object reflect
	islandRefectionShrink = 0.4
	islandReflectionBlur  = 2
	islandReflectionAlpha = 0.15
)

func (g *Game) drawGame(screen *ebiten.Image) {

	drawWorldGrad(g, screen)
	drawSun(screen)
	drawClouds(g, screen)
	drawWaves(g, screen)
	drawIsland(g, screen)
	drawAir(g, screen)
	drawBoat(g, screen)

	g.doFade(screen, time.Millisecond*500, color.NRGBA{R: 255, G: 255, B: 255}, true)
	if *debug {
		buf := fmt.Sprintf("boat: %4.0f,%3.0f w: %3d, a: %3d", g.boatPos.X, g.boatPos.Y, numWaves, numAirWaves)
		ebitenutil.DebugPrint(screen, buf)
	}
}

func drawIsland(g *Game, screen *ebiten.Image) {
	// Island
	op := &ebiten.DrawImageOptions{}
	islandPos := g.boatPos.X*float64(islandY/dWinWidth) - islandStartX
	op.GeoM.Translate(
		dWinWidth-float64(islandPos),
		dWinHeightHalf-float64(island1SP.image.Bounds().Dy())+islandY)
	screen.DrawImage(island1SP.image, op)

	//Island refection
	op.GeoM.Reset()
	op.GeoM.Scale(1, -(1 / islandRefectionShrink))
	op.ColorScale.ScaleAlpha(islandReflectionAlpha)
	op.GeoM.Translate(
		dWinWidth-float64(islandPos),
		dWinHeightHalf+float64(islandY+island1SP.image.Bounds().Dy()-5)/islandRefectionShrink)
	screen.DrawImage(island1SP.blurred, op)

}

func drawBoat(g *Game, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	offset := iPoint{}
	if time.Now().Unix()%3 == 0 {
		offset.Y = 1
	}
	boatFrame := autoAnimatePingPong(boat2SP, 0)
	op.GeoM.Translate(
		float64((dWinWidth/4)-(boatFrame.Bounds().Dx())/2+offset.X),
		float64((dWinHeight/6)*4.5-(boatFrame.Bounds().Dy())/2+offset.Y)+float64(g.boatPos.Y))

	screen.DrawImage(boatFrame, op)
}
