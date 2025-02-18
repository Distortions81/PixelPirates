package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
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

func (g *Game) drawGame(screen *ebiten.Image) {

	g.makeWave()
	g.makeAirWave()

	drawWorldGrad(g, screen)
	drawSun(g, screen)
	drawClouds(g, screen)
	drawWaves(g, screen)
	drawIslands(g, screen)
	drawAir(g, screen)
	drawBoat(g, screen)

}

func drawBoat(g *Game, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	offset := iPoint{}
	if time.Now().Unix()%3 == 0 {
		offset.Y = 1
	}
	dir := directionFromCoords(g.oldBoatPos.X-g.boatPos.X, g.oldBoatPos.Y-g.boatPos.Y)
	dirName := "still"
	if dir == DIR_WEST || dir == DIR_NORTH_WEST || dir == DIR_SOUTH_WEST {
		dirName = "sail-rev"
	}
	if dir == DIR_EAST || dir == DIR_NORTH_EAST || dir == DIR_SOUTH_EAST {
		dirName = "sail"
	}
	boatFrame := autoAnimatePingPong(g.boat2SP, 0, dirName)
	if g.gameMode == GAME_PLAY {
		op.GeoM.Translate(
			float64((dWinWidth/4)-(boatFrame.Bounds().Dx())/2+offset.X),
			float64(dWinHeight/6.0)*4.5-float64(boatFrame.Bounds().Dy())/2+float64(offset.Y)+float64(g.boatPos.Y))
	} else {
		op.GeoM.Translate(
			float64((dWinWidth/2)-(boatFrame.Bounds().Dx())/2+offset.X),
			float64((dWinHeight/3)*2.0-(boatFrame.Bounds().Dy())/2+offset.Y)+float64(g.boatPos.Y))
	}

	screen.DrawImage(boatFrame, op)
}
