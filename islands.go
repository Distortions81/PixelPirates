package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	islandChunkSize = dWinWidthHalf
	checkChunks     = 4
)

var islandChunks map[int]*islandChunkData

type islandData struct {
	name, desc string
	pos        int

	spriteName string
	sprite     *spriteItem
}

type islandChunkData struct {
	islands []islandData
}

func init() {
	islandChunks = map[int]*islandChunkData{}
	for i, island := range islands {
		islandChunkPos := island.pos / islandChunkSize
		if islandChunks[islandChunkPos] == nil {
			islandChunks[islandChunkPos] = &islandChunkData{}
		}
		fmt.Printf("Storing island: #%v '%v' in block %v.\n", i+1, island.name, islandChunkPos)
		islandChunks[islandChunkPos].islands = append(islandChunks[islandChunkPos].islands, island)
	}
}

func drawIslands(g *Game, screen *ebiten.Image) {
	paralaxPos := g.boatPos.X * float64(islandY/dWinWidth)

	islands := getIslands(int(paralaxPos))

	for _, island := range islands {
		islandPos := dWinWidth - (paralaxPos + float64(island.pos))

		//Island
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(
			islandPos,
			dWinHeightHalf-float64(island1SP.image.Bounds().Dy())+islandY)
		screen.DrawImage(island1SP.image, op)

		//Visit sign
		spriteSize := float64(island1SP.image.Bounds().Dx())
		if paralaxPos > islandPos-(spriteSize*2) && paralaxPos < islandPos+spriteSize {
			op.GeoM.Translate(-3, -25)
			op.ColorScale.ScaleAlpha(0.2)
			screen.DrawImage(visitSP.image, op)
		}

		//Island refection
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, -(1 / islandRefectionShrink))
		op.ColorScale.ScaleAlpha(islandReflectionAlpha)
		op.GeoM.Translate(
			islandPos,
			dWinHeightHalf+float64(islandY+island1SP.image.Bounds().Dy()-5)/islandRefectionShrink)
		screen.DrawImage(island1SP.blurred, op)

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

var islands []islandData = []islandData{
	{name: "Welcome island", desc: "Learn the basics here!", pos: dWinWidthHalf, spriteName: "island1"},
}
