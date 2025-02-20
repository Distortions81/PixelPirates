package main

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	islandChunkSize = dWinWidthHalf
	checkChunks     = 4
)

type islandData struct {
	name, desc string
	pos        int
	spawn      fPoint

	oceanSprite *spriteItem
	visitSprite *spriteItem
	objects     []*spriteItem

	collision map[iPoint]bool
}

type islandChunkData struct {
	islands []islandData
}

func drawIslands(g *Game, screen *ebiten.Image) {

	paralaxPos := g.boatPos.X * (islandY * distParallax)

	islands := getIslands(g, int(paralaxPos))
	drewSign := false

	for i, island := range islands {

		islandPosX := -(paralaxPos + float64(-island.pos))
		islandPosY := dWinHeightHalf - float64(island.oceanSprite.image.Bounds().Dy()) + islandY

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
			g.canVisit = &islands[i]
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
		g.canVisit = nil
	}
}

func getIslands(g *Game, pos int) []islandData {
	var islandsFound []islandData

	cPos := pos / islandChunkSize

	for x := cPos - checkChunks; x < cPos+checkChunks; x++ {
		if g.islandChunks[x] == nil {
			continue
		}
		islandsFound = append(islandsFound, g.islandChunks[x].islands...)
	}

	return islandsFound
}

func findSpawns() fPoint {
	for i, island := range islands {
		for _, item := range island.objects {
			if strings.Contains(item.Name, "spawn") {
				name := item.animation.sortedFrames[0]
				frame := item.animation.Frames[name]
				newSpawn := fPoint{X: float64(frame.SpriteSourceSize.X), Y: float64(frame.SpriteSourceSize.Y)}
				islands[i].spawn = newSpawn
				doLog(true, false, "Found Spawn for: %v at %v,%v", island.name, newSpawn.X, newSpawn.Y)
				return newSpawn
			}
		}
	}

	return fPoint{}
}
