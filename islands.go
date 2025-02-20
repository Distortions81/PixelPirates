package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var islands []islandData

const (
	islandChunkSize = dWinWidthHalf
	checkChunks     = 4

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
					Fullpath: dataDir + spritesDir + islandsDir + island + "/" + oceanSpriteFile},
				spriteSheet: &spriteItem{onDemand: true,
					Fullpath: dataDir + spritesDir + islandsDir + island + "/" + spriteSheetName},
			})
		islandsAdded = append(islandsAdded, info.Name)
	}

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

func findSpawns(g *Game) fPoint {

	if g.visiting == nil {
		return fPoint{}
	}
	spawn := g.visiting.spriteSheet.animation.layers["spawn"]
	if spawn == nil {
		doLog(true, false, "Island has no spawn layer.")
		return fPoint{}
	}
	newSpawn := fPoint{X: float64(spawn.SpriteSourceSize.X), Y: float64(spawn.SpriteSourceSize.Y)}
	doLog(true, false, "Found Spawn for: %v at %v,%v", g.visiting.name, newSpawn.X, newSpawn.Y)
	return newSpawn
}

func visitIsland(g *Game) {
	if g.canVisit == nil {
		return
	}
	doLog(true, true, "Visiting: %v", g.canVisit.name)
	if g.canVisit.spriteSheet.image == nil {
		loadSprite(g.canVisit.spriteSheet.Fullpath, g.canVisit.spriteSheet, true)
	}

	g.visiting = g.canVisit

	makeCollisionMaps(g)
	fixPos := findSpawns(g)

	fixPos.X -= (dWinWidth / 2)
	fixPos.Y -= (dWinHeight / 2)

	g.playPos = fixPos
}

const shoreMovement = 25

func (g *Game) drawIsland(screen *ebiten.Image) {

	if g.visiting == nil {
		screen.Clear()
		ebitenutil.DebugPrint(screen, "Invalid g.visiting.")
		return
	}
	if g.visiting.spriteSheet.image == nil {
		screen.Clear()
		ebitenutil.DebugPrint(screen, "Invalid visitSprite.")
		return
	}
	//Draw island ground
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-g.playPos.X, -g.playPos.Y)
	ground := getLayerFromName("ground", g.visiting.spriteSheet)
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

	for layerName, layer := range g.visiting.spriteSheet.animation.layers {
		op := &ebiten.DrawImageOptions{}

		if layerName == "ground" {
			continue
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
		op.GeoM.Translate(xpos, ypos)

		if layerName == "edges" || layerName == "spawn" {
			if *debugMode {
				op.ColorScale.ScaleAlpha(0.15)
			} else {
				continue
			}
		}

		subImg := getLayer(layer, g.visiting.spriteSheet)
		screen.DrawImage(subImg, op)
	}

	buf := fmt.Sprintf("Test Island scene, E to Exit. %0.0f,%0.0f", g.playPos.X, g.playPos.Y)
	ebitenutil.DebugPrint(screen, buf)
}
