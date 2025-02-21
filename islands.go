package main

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

var islands []islandData

const (
	chunksPerScreen = 4
	islandChunkSize = dWinWidth / chunksPerScreen
	checkChunks     = chunksPerScreen

	infoJsonFile    = "info"
	oceanSpriteFile = "ocean"
	spriteSheetName = "sprite-sheet"
)

type islandData struct {
	name, desc string
	pos        int
	spawn      fPoint

	oceanSprite *spriteItem
	oceanSeen   time.Time
	spriteSheet *spriteItem
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

	go cleanIslands()
}

func cleanIslands() {
	for {
		time.Sleep(time.Second)

		for i, island := range islands {
			if island.oceanSprite.image != nil {
				if !island.oceanSeen.IsZero() && time.Since(island.oceanSeen) > time.Second {
					islands[i].oceanSeen = time.Time{}
					islands[i].oceanSprite.image.Deallocate()
					islands[i].oceanSprite.blurred.Deallocate()

					islands[i].oceanSprite.image = nil
					islands[i].oceanSprite.blurred = nil
					doLog(true, true, "Deallocated island ocean sprite: %v", island.name)
				}
			}
		}
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
	if len(islands) > 0 {
		return nil
	}
	var dir []os.DirEntry
	var err error
	dirPath := dataDir + spritesDir + islandsDir
	islands = []islandData{}

	doLog(true, true, "scanIslandsFolder: Scanning.")

	if wasmMode {
		dir, err = efs.ReadDir(strings.TrimSuffix(dirPath, "/"))
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
		infoPath := dirPath + island + "/"
		fullPath := infoPath + infoJsonFile + ".json"
		var err error
		if wasmMode {
			_, err = efs.ReadFile(fullPath)
		} else {
			_, err = os.ReadFile(fullPath)
		}
		if err != nil {
			doLog(true, false, "Island '%v' has no %v file.", fullPath, infoJsonFile)
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
					Fullpath: dataDir + spritesDir + islandsDir + island + "/" + oceanSpriteFile},
				spriteSheet: &spriteItem{onDemand: true,
					Fullpath: dataDir + spritesDir + islandsDir + island + "/" + spriteSheetName},
			})
		islandsAdded = append(islandsAdded, info.Name)
	}

	return nil
}

func getIslands(g *Game, pos int) []islandData {
	var islandsFound []islandData

	cPos := pos / islandChunkSize

	for x := cPos - 1; x < cPos+(checkChunks+1); x++ {
		if g.islandChunks[x] == nil {
			continue
		}
		islandsFound = append(islandsFound, g.islandChunks[x].islands...)
	}

	return islandsFound
}

func findSpawns(g *Game) fPoint {

	if g.inIsland == nil {
		return fPoint{}
	}
	spawn := g.inIsland.spriteSheet.animation.layers["spawn"]
	if spawn == nil {
		doLog(true, false, "Island has no spawn layer.")
		return fPoint{}
	}
	newSpawn := fPoint{X: float64(spawn.SpriteSourceSize.X), Y: float64(spawn.SpriteSourceSize.Y)}
	doLog(true, false, "Found Spawn for: %v at %v,%v", g.inIsland.name, newSpawn.X, newSpawn.Y)
	return newSpawn
}

func gotoIsland(g *Game) {
	if g.availIsland == nil {
		return
	}
	doLog(true, true, "Going to island: %v", g.availIsland.name)
	if g.availIsland.spriteSheet.image == nil {
		loadSprite(g.availIsland.spriteSheet.Fullpath, g.availIsland.spriteSheet, true)
	}

	loadDefaultChar(g)

	g.inIsland = g.availIsland

	makeCollisionMaps(g)
	fixPos := findSpawns(g)

	fixPos.X -= (dWinWidth / 2)
	fixPos.Y -= (dWinHeight / 2)

	g.playPos = fixPos
}
