package main

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var islands []islandData

const (
	islandChunkSize = dWinWidthHalf
	checkChunks     = 4

	infoJsonFile    = "info"
	oceanSpriteFile = "ocean"
	spriteSheetName = "spritesheet"
	spriteSheetJson = "spritesheet"
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

type islandInfoData struct {
	Comment, Name, Desc string

	Pos,
	Level int
}

func initIslands(g *Game) {
	g.islandChunks = map[int]*islandChunkData{}

	for i, island := range islands {
		islandChunkPos := island.pos / islandChunkSize
		if g.islandChunks[islandChunkPos] == nil {
			g.islandChunks[islandChunkPos] = &islandChunkData{}
		}

		doLog(true, true, "Storing island: #%v '%v' in block %v.", i+1, island.name, islandChunkPos)

		g.islandChunks[islandChunkPos].islands = append(g.islandChunks[islandChunkPos].islands, islands[i])
	}
}

func writeInfoJson(path string, island islandInfoData) error {

	if wasmMode {
		return nil
	}

	data, err := json.MarshalIndent(island, "", "  ")
	if err != nil {
		doLog(true, false, "writeInfoJson: jsonMarshal: %v", err)
		return err
	}

	err = os.WriteFile(path, data, 0755)
	if err != nil {
		doLog(true, false, "writeInfoJson: WriteFile: %v", err)
		return err
	}

	return nil
}

func readInfoJson(path string) (islandInfoData, error) {
	var fileData []byte
	var err error

	fpath := dataDir + spritesDir + islandsDir + path

	if wasmMode {
		fileData, err = efs.ReadFile(fpath + "/info.json")
	} else {
		fileData, err = os.ReadFile(fpath + "/info.json")
	}
	if err != nil {
		doLog(true, false, "readInfoJson: readFile: %v", err)
		return islandInfoData{}, err
	}

	var info islandInfoData
	err = json.Unmarshal(fileData, &info)
	if err != nil {
		doLog(true, false, "decodeAniJSON: %v", err)
		return islandInfoData{}, err
	}

	return info, nil
}

func scanIslandsFolder() error {
	var dir []os.DirEntry
	var err error
	dirPath := dataDir + spritesDir + islandsDir
	islands = []islandData{}

	doLog(true, true, "scanIslandsFolder: Scanning.")

	if wasmMode {
		dir, err = efs.ReadDir(dirPath)
	} else {
		dir, err = os.ReadDir(dirPath)
	}
	if err != nil {
		doLog(true, false, "scanIslandsFolder: readDir: %v", err)
		return err
	}

	var islandFolders []string
	for _, item := range dir {
		if item.IsDir() {
			islandFolders = append(islandFolders, item.Name())
		}
	}
	doLog(true, true, "Islands found: %v", strings.Join(islandFolders, ", "))

	var islandsAdded []string
	for _, island := range islandFolders {
		infoPath := dirPath + "/" + island + "/"
		_, err := os.ReadFile(infoPath + infoJsonFile + ".json")
		if err != nil {
			doLog(true, false, "Island '%v' has no %v file.", island, infoJsonFile)
			newInfo := islandInfoData{
				Comment: "Once complete, rename this file to info.json",
				Name:    island, Desc: "In-game description", Pos: 320}
			writeInfoJson(infoPath+"info-example.json", newInfo)

			return err
		}
		info, err := readInfoJson(island)
		if err != nil {
			doLog(true, false, "scanIslandsFolder: %v file for %v is invalid.", infoJsonFile, island)
			return nil
		}
		islands = append(islands,
			islandData{
				name: info.Name,
				desc: info.Desc,
				pos:  info.Pos,
				oceanSprite: &spriteItem{doReflect: true, onDemand: true,
					Fullpath: dataDir + spritesDir + islandsDir + island + "/" + oceanSpriteFile}})
		islandsAdded = append(islandsAdded, info.Name)
	}

	doLog(true, true, "Islands added: %v", strings.Join(islandsAdded, ", "))

	return nil
}

func drawIslands(g *Game, screen *ebiten.Image) {

	paralaxPos := g.boatPos.X * (islandY * distParallax)

	islands := getIslands(g, int(paralaxPos))
	drewSign := false

	for i, island := range islands {

		if island.oceanSprite.image == nil {
			loadSprite(island.oceanSprite.Fullpath, island.oceanSprite, true)
		}
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
