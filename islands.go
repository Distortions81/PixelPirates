package main

const islandChunkSize = dWinWidth

var islandChunks map[int]islandChunkData

type islandData struct {
	name, desc string
	pos        int

	spriteName string
	sprite     *spriteItem
}

type islandChunkData struct {
	islands []islandData
}

func getIslands(pos int) []islandData {
	var islandsFound []islandData

	minpos, maxpos := pos-dWinWidth, pos+(dWinWidth*2)
	for x := minpos; x < maxpos; x++ {
		islandsFound = append(islandsFound, islandChunks[x/islandChunkSize].islands...)
	}

	return islandsFound
}

var islands []islandData = []islandData{
	{name: "Welcome island", desc: "Learn the basics here!", pos: -dWinWidthHalf, spriteName: "island1"},
}
