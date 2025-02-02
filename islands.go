package main

import "github.com/hajimehoshi/ebiten/v2"

const islandChunkSize = dWinWidth

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
	for _, island := range islands {
		islandChunkPos := island.pos / islandChunkSize
		if islandChunks[islandChunkPos] == nil {
			islandChunks[islandChunkPos] = &islandChunkData{}
		}
		islandChunks[islandChunkPos].islands = append(islandChunks[islandChunkPos].islands, island)
	}
}

func drawIslands(g *Game, screen *ebiten.Image) {
	paralaxPos := g.boatPos.X * float64(islandY/dWinWidth)

	islands := getIslands(int(paralaxPos))

	for _, island := range islands {
		islandPos := (paralaxPos + float64(island.pos))
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(
			dWinWidth-islandPos,
			dWinHeightHalf-float64(island1SP.image.Bounds().Dy())+islandY)
		screen.DrawImage(island1SP.image, op)

		//Island refection
		op.GeoM.Reset()
		op.GeoM.Scale(1, -(1 / islandRefectionShrink))
		op.ColorScale.ScaleAlpha(islandReflectionAlpha)
		op.GeoM.Translate(
			dWinWidth-islandPos,
			dWinHeightHalf+float64(islandY+island1SP.image.Bounds().Dy()-5)/islandRefectionShrink)
		screen.DrawImage(island1SP.blurred, op)
	}
}

func getIslands(pos int) []islandData {
	var islandsFound []islandData

	minpos, maxpos := pos-dWinWidth, pos+(dWinWidth*2)
	for x := minpos; x < maxpos; x++ {
		if islandChunks[x/islandChunkSize] == nil {
			continue
		}
		islandsFound = append(islandsFound, islandChunks[x/islandChunkSize].islands...)
	}

	return islandsFound
}

var islands []islandData = []islandData{
	{name: "Welcome island", desc: "Learn the basics here!", pos: dWinWidthHalf, spriteName: "island1"},
}
