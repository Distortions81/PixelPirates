package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const shoreMovement = 25

func (g *Game) drawIsland(screen *ebiten.Image) {

	if g.inIsland == nil {
		screen.Clear()
		ebitenutil.DebugPrint(screen, "Invalid g.inIsland.")
		return
	}
	if g.inIsland.spriteSheet.image == nil {
		screen.Clear()
		ebitenutil.DebugPrint(screen, "Invalid spriteSheet.")
		return
	}

	op := &ebiten.DrawImageOptions{}
	var ground *ebiten.Image
	if g.inRoom == nil {
		//Draw island ground
		op.GeoM.Translate(-g.playPos.X, -g.playPos.Y)
		ground = getLayerFromName("ground", g.inIsland.spriteSheet)
	} else {
		//Draw room
		screen.Clear()
		ground = getLayerFromName(g.inRoom.room, g.inIsland.spriteSheet)
		room := g.inIsland.spriteSheet.animation.layers[g.inRoom.room]
		xpos, ypos :=
			float64(room.SpriteSourceSize.X-int(g.playPos.X)),
			float64(room.SpriteSourceSize.Y-int(g.playPos.Y))
		op.GeoM.Translate(xpos, ypos)
	}
	screen.DrawImage(ground, op)

	//Draw player
	op = &ebiten.DrawImageOptions{}
	ani := g.defPlayerSP
	fKey := ani.animation.sortedFrames[0]
	fRect := ani.animation.Frames[fKey].Frame

	charX, charY := float64(fRect.W/2), float64(fRect.H/2)
	op.GeoM.Translate(dWinWidthHalf-charX, dWinHeightHalf-charY)

	faceDir := directionFromCoords(g.oldPlayPos.X-g.playPos.X, g.oldPlayPos.Y-g.playPos.Y)
	var (
		dirName string
	)
	if faceDir == DIR_NONE {
		dirName = "idle"
		lface := g.playerFacing
		playerImg = getFrameNumber(int64(faceFix[lface]), g.defPlayerSP, 0)
	} else {
		dirName = fmt.Sprintf("%v move", moveFix[faceDir])
		playerImg = autoAnimate(g.defPlayerSP, 0, dirName)
	}
	screen.DrawImage(playerImg, op)

	//Draw layers
	for layerName, layer := range g.inIsland.spriteSheet.animation.layers {
		op := &ebiten.DrawImageOptions{}

		if g.inRoom != nil && !strings.HasPrefix(layerName, "building") {
			continue
		} else {
			if layerName == "ground" || strings.HasPrefix(layerName, "building") {
				continue
			}
		}

		//TODO: Replace with sprite values
		offsety := 0.0
		if layerName == "water" {
			fraction := float64(time.Now().UnixMilli()%10000) / 10000.0
			offsety = (math.Sin(2*math.Pi*fraction) * shoreMovement) + shoreMovement
		}

		xpos, ypos :=
			float64(layer.SpriteSourceSize.X-int(g.playPos.X)),
			float64(layer.SpriteSourceSize.Y-int(g.playPos.Y))+offsety

		if strings.HasPrefix(layerName, "door") {
			if g.availRoom == nil || g.availRoom.door != layerName {
				continue
			} else {
				op.ColorScale.ScaleAlpha(0.5)
				ebitenutil.DebugPrintAt(screen, "E to enter", int(xpos-20), int(ypos-20))
			}
		}

		op.GeoM.Translate(xpos, ypos)

		if layerName == "edges" || layerName == "spawn" {
			if *debugMode {
				op.ColorScale.ScaleAlpha(0.15)
			} else {
				continue
			}
		}

		subImg := getLayer(layer, g.inIsland.spriteSheet)
		screen.DrawImage(subImg, op)
	}

	buf := fmt.Sprintf("Test Island scene, E to Exit. %0.0f,%0.0f", g.playPos.X, g.playPos.Y)
	ebitenutil.DebugPrint(screen, buf)
}

func drawIslands(g *Game, screen *ebiten.Image) {

	paralaxPos := g.boatPos.X * (islandY * distParallax)

	iList := getIslands(g, int(paralaxPos))
	drewSign := false

	for i, island := range iList {
		if island.oceanSprite.image == nil {
			loadSprite(island.oceanSprite.Fullpath, island.oceanSprite, true)
		}
		islands[i].oceanSeen = time.Now()

		islandPosX := -(paralaxPos + float64(-island.pos))
		islandPosY := dWinHeightHalf - float64(island.oceanSprite.image.Bounds().Dy()) + islandY

		if paralaxPos < 0 || islandPosX > dWinWidth {
			continue //prevent overdraw
		}

		//Island
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(
			islandPosX,
			islandPosY,
		)
		screen.DrawImage(island.oceanSprite.image, op)

		//Visit sign
		spriteSize := float64(island.oceanSprite.image.Bounds().Dx())
		if !drewSign && islandPosX > 0 && islandPosX < spriteSize {
			ebitenutil.DebugPrintAt(screen, island.name+"\nE: Visit", int(islandPosX)+10, int(islandPosY)-32)
			drewSign = true
			g.availIsland = &islands[i]
		}

		//Island refection
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, -(1 / islandRefectionShrink))
		op.ColorScale.ScaleAlpha(islandReflectionAlpha)
		op.GeoM.Translate(
			islandPosX,
			islandPosY*1.5,
		)
		screen.DrawImage(island.oceanSprite.blurred, op)
	}
	//Clear target
	if !drewSign {
		g.availIsland = nil
	}
}
