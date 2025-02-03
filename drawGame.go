package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	//Island settings
	islandY      = 18.0
	cloudY       = 6.0
	islandStartX = -dWinWidthHalf

	//Object reflect
	islandRefectionShrink = 0.4
	islandReflectionBlur  = 2
	islandReflectionAlpha = 0.15
)

var (
	displayStamp time.Time
	frameNumber  uint64

	debugBuf string
)

func (g *Game) drawGame(screen *ebiten.Image) {
	frameNumber++
	startTime := time.Now()

	g.makeWave()
	g.makeAirWave()

	drawWorldGrad(g, screen)
	drawSun(screen)
	drawCloudsNew(g, screen)
	drawWaves(g, screen)
	drawIslands(g, screen)
	drawAir(g, screen)
	drawBoat(g, screen)

	if *debug {

		if frameNumber%60 == 0 {
			renderTime := time.Since(startTime).Microseconds()
			displayTime := time.Since(displayStamp).Microseconds()

			debugBuf = fmt.Sprintf("Render: %4du, Display: %0.2fms, %%%0.2f, FPS: %3d",
				renderTime,
				float64(displayTime)/1000,
				float64(renderTime)/float64(displayTime)*100,
				int(1000/(float64(displayTime)/1000)))
		}

		buf := fmt.Sprintf("x: %d", int(g.boatPos.X))
		ebitenutil.DebugPrint(screen, buf)
		ebitenutil.DebugPrintAt(screen, debugBuf, 0, dWinHeight-16)
		displayStamp = time.Now()
	}
}

func drawBoat(g *Game, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	offset := iPoint{}
	if time.Now().Unix()%3 == 0 {
		offset.Y = 1
	}
	boatFrame := autoAnimatePingPong(boat2SP, 0, "sail")
	if g.gameMode == GAME_PLAY {
		op.GeoM.Translate(
			float64((dWinWidth/4)-(boatFrame.Bounds().Dx())/2+offset.X),
			float64((dWinHeight/6)*4.5-(boatFrame.Bounds().Dy())/2+offset.Y)+float64(g.boatPos.Y))
	} else {
		op.GeoM.Translate(
			float64((dWinWidth/2)-(boatFrame.Bounds().Dx())/2+offset.X),
			float64((dWinHeight/3)*2.0-(boatFrame.Bounds().Dy())/2+offset.Y)+float64(g.boatPos.Y))
	}

	screen.DrawImage(boatFrame, op)
}
