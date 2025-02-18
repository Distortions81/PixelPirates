package main

import (
	"fmt"
	"math"
	"os"
	"path"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var playerImg *ebiten.Image

const (
	islandChunkSize = dWinWidthHalf
	checkChunks     = 4
)

type islandData struct {
	name, desc string
	pos        int
	spawn      fPoint

	spriteName string
	visitName  string

	sprite      *spriteItem
	visitSprite *spriteItem
	objects     []*spriteItem

	collision map[iPoint]bool
}

type islandChunkData struct {
	islands []islandData
}

var islands []islandData = []islandData{
	{name: "Welcome island", desc: "Learn the basics here!",
		pos: dWinWidth, spriteName: "island1", visitName: "island-scene1"},
}

func visitIsland(g *Game) {
	if g.canVisit == nil {
		return
	}
	g.visiting = g.canVisit

	loadSprite(islandsDir+g.visiting.visitName+"/"+g.visiting.visitName, g.visiting.visitSprite, true)

	//Load objects
	for _, obj := range g.visiting.objects {
		obj.image, _, _ = loadImage(obj.Fullpath, true, false)
		aniData, err := loadAnimationData(obj.Fullpath)
		if err == nil && aniData != nil {
			obj.animation = aniData
			doLog(true, true, "loaded animation: %v", obj.Fullpath)
		}
		doLog(true, true, "loaded sprite: %v", obj.Fullpath)
	}

	makeCollisionMaps(g)
	fixPos := findSpawns()

	frameRange := g.defPlayerSP.animation.animations["idle"]
	name := g.defPlayerSP.animation.sortedFrames[frameRange.start]
	frame := g.defPlayerSP.animation.Frames[name]

	fixPos.X -= (dWinWidth / 2)
	fixPos.Y -= (dWinHeight / 2)

	fixPos.X += float64(frame.SpriteSourceSize.W / 2)
	fixPos.Y += float64(frame.SpriteSourceSize.H / 2)
	g.playPos = fixPos
}

func initIslands(g *Game) {
	g.islandChunks = map[int]*islandChunkData{}

	for i, island := range islands {
		fPath := islandsDir + island.visitName
		islandChunkPos := island.pos / islandChunkSize
		if g.islandChunks[islandChunkPos] == nil {
			g.islandChunks[islandChunkPos] = &islandChunkData{}
			var islandDir []os.DirEntry
			var err error
			if !wasmMode {
				islandDir, err = os.ReadDir(dataDir + spritesDir + fPath)
			} else {
				islandDir, err = efs.ReadDir(dataDir + spritesDir + fPath)
			}
			if err != nil {
				doLog(true, false, "initIslands: readSprites: %v", err.Error())
				return
			}
			islands[i].visitSprite = &spriteItem{onDemand: true, unmanged: true}
			islands[i].sprite = &spriteItem{doReflect: true}
			for _, item := range islandDir {

				if strings.HasSuffix(item.Name(), ".png") {
					fileName := path.Base(item.Name())
					trimName := strings.TrimSuffix(fileName, ".png")

					if strings.EqualFold(trimName, island.spriteName) {
					} else if strings.EqualFold(trimName, island.visitName) {
						loadSprite(fPath+"/"+trimName, islands[i].visitSprite, false)
					} else {
						newSprite := &spriteItem{Name: trimName, onDemand: true, unmanged: true, Fullpath: dataDir + spritesDir + fPath + "/" + trimName}
						loadSprite(fPath+"/"+trimName, newSprite, false)
						islands[i].objects = append(islands[i].objects, newSprite)
					}
				}
			}
		}

		doLog(true, true, "Storing island: #%v '%v' in block %v.", i+1, island.name, islandChunkPos)

		g.islandChunks[islandChunkPos].islands = append(g.islandChunks[islandChunkPos].islands, islands[i])
	}
}

func drawIslands(g *Game, screen *ebiten.Image) {

	paralaxPos := g.boatPos.X * (islandY * distParallax)

	islands := getIslands(g, int(paralaxPos))
	drewSign := false

	for i, island := range islands {
		if island.sprite.image == nil {
			loadSprite(islandsDir+island.visitName+"/"+island.spriteName, island.sprite, true)
		}
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

func (g *Game) drawIsland(screen *ebiten.Image) {

	if g.visiting == nil {
		screen.Clear()
		ebitenutil.DebugPrint(screen, "Invalid g.visiting.")
		return
	}
	if g.visiting.visitSprite.image == nil {
		screen.Clear()
		ebitenutil.DebugPrint(screen, "Invalid visitSprite.")
		return
	}
	//Draw island ground
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-g.playPos.X, -g.playPos.Y)
	screen.DrawImage(g.visiting.visitSprite.image, op)

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
		playerImg = getAniFrame(int64(faceFix[lface]), g.defPlayerSP, 0)
	} else {
		dirName = fmt.Sprintf("%v move", moveFix[faceDir])
		playerImg = autoAnimate(g.defPlayerSP, 0, dirName)
	}
	screen.DrawImage(playerImg, op)

	for _, obj := range g.visiting.objects {
		op := &ebiten.DrawImageOptions{}
		name := obj.animation.sortedFrames[0]
		frame := obj.animation.Frames[name]

		//TODO: Replace with sprite values
		offsety := 0.0
		if strings.Contains(obj.Name, "shore") {
			fraction := float64(time.Now().UnixMilli()%10000) / 10000.0
			offsety = math.Sin(2*math.Pi*fraction)*25 + 50
		}

		op.GeoM.Translate(
			float64(frame.SpriteSourceSize.X-int(g.playPos.X)),
			float64(frame.SpriteSourceSize.Y-int(g.playPos.Y))+offsety)

		if strings.Contains(obj.Name, "collision") ||
			strings.Contains(obj.Name, "spawn") {
			if *debugMode {
				op.ColorScale.ScaleAlpha(0.15)
			} else {
				continue
			}
		}
		screen.DrawImage(obj.image, op)
	}

	buf := fmt.Sprintf("Test Island scene, E to Exit. %0.0f,%0.0f", g.playPos.X, g.playPos.Y)
	ebitenutil.DebugPrint(screen, buf)
}

// TODO: Update sprite tags instead,
var (
	moveFix [9]int = [9]int{12, 12, 2, 3, 4, 6, 8, 9, 10}
	faceFix [9]int = [9]int{0, 4, 3, 2, 1, 0, 7, 6, 5}
)

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
