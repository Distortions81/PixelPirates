package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	islandChunkSize = dWinWidthHalf
	checkChunks     = 4
)

var islandChunks map[int]*islandChunkData

type islandData struct {
	name, desc string
	pos        int
	spawn      fPoint

	spriteName string
	visitName  string

	sprite      *spriteItem
	visitSprite *spriteItem
}

type islandChunkData struct {
	islands []islandData
}

var islands []islandData = []islandData{
	{name: "Welcome island", desc: "Learn the basics here!", pos: dWinWidth, spriteName: "island1", visitName: "island-scene1"},
}

func initIslands() {
	islandChunks = map[int]*islandChunkData{}

	for i, island := range islands {
		islandChunkPos := island.pos / islandChunkSize
		if islandChunks[islandChunkPos] == nil {
			islandChunks[islandChunkPos] = &islandChunkData{}
		}
		islands[i].sprite = spriteList[island.spriteName]
		vsp := spriteList[island.visitName]
		islands[i].visitSprite = vsp
		islands[i].spawn = fPoint{X: float64(vsp.image.Bounds().Dx()) / 2, Y: float64(vsp.image.Bounds().Dy())}

		fmt.Printf("Spawn: %v,%v -- ", islands[i].spawn.X, islands[i].spawn.Y)
		fmt.Printf("Storing island: #%v '%v' in block %v.\n", i+1, island.name, islandChunkPos)

		islandChunks[islandChunkPos].islands = append(islandChunks[islandChunkPos].islands, islands[i])
	}
}

func drawIslands(g *Game, screen *ebiten.Image) {

	paralaxPos := g.boatPos.X * (islandY * distParallax)

	islands := getIslands(int(paralaxPos))
	drewSign := false

	for i, island := range islands {
		islandPosX := -(paralaxPos + float64(-island.pos))
		islandPosY := dWinHeightHalf - float64(island.sprite.image.Bounds().Dy()) + islandY

		//Island
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(
			islandPosX,
			islandPosY,
		)
		screen.DrawImage(island.sprite.image, op)

		//Visit sign
		spriteSize := float64(island.sprite.image.Bounds().Dx())
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
		screen.DrawImage(island.sprite.blurred, op)
	}
	//Clear target
	if !drewSign {
		g.canVisit = nil
	}
}

func getIslands(pos int) []islandData {
	var islandsFound []islandData

	cPos := pos / islandChunkSize

	for x := cPos - checkChunks; x < cPos+checkChunks; x++ {
		if islandChunks[x] == nil {
			continue
		}
		islandsFound = append(islandsFound, islandChunks[x].islands...)
	}

	return islandsFound
}

func (g *Game) drawIsland(screen *ebiten.Image) {
	screen.Clear()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-g.playPos.X, -g.playPos.Y)
	screen.DrawImage(g.visiting.visitSprite.image, op)
	buf := fmt.Sprintf("Test Island scene, E to Exit. %v,%v", g.playPos.X, g.playPos.Y)
	ebitenutil.DebugPrint(screen, buf)

	op = &ebiten.DrawImageOptions{}
	ani := defPlayerSP
	fKey := ani.animation.sortedFrames[0]
	fRect := ani.animation.Frames[fKey].Frame

	charX, charY := float64(fRect.W/2), float64(fRect.H/2)
	op.GeoM.Translate(dWinWidthHalf-charX, dWinHeightHalf-charY)

	faceDir := directionFromCoords(g.oldPlayPos.X-g.playPos.X, g.oldPlayPos.Y-g.playPos.Y)
	var dirName string
	var playerImg *ebiten.Image
	if faceDir < 0 {
		dirName = "idle"
		lface := g.playerFacing
		if lface < 0 {
			lface = 4
		}
		playerImg = getAniFrame(int64(faceFix[lface]), defPlayerSP, 0)
	} else {
		dirName = fmt.Sprintf("%v move", moveFix[faceDir])
		playerImg = autoAnimate(defPlayerSP, 0, dirName)
	}

	screen.DrawImage(playerImg, op)
}

var moveFix [8]int = [8]int{12, 2, 3, 4, 6, 8, 9, 10}
var faceFix [8]int = [8]int{4, 3, 2, 1, 0, 7, 6, 5}
