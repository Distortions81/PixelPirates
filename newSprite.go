package main

import (
	"image"
	"image/png"
	"io/fs"
	"os"
	"strings"

	"github.com/anthonynsimon/bild/blur"
	"github.com/hajimehoshi/ebiten/v2"
)

type islandInfoData struct {
	Name, Desc string

	Distance,
	Level int
}

const (
	infoJsonFile    = "info.json"
	mainSpriteName  = "world.png"
	spriteSheetName = "spritesheet.png"
	spriteSheetJson = "spritesheet.json"
)

// Load sprites
func loadImage(name string, unmanaged bool, doBlur bool) (*ebiten.Image, *ebiten.Image, error) {

	//Open file
	var (
		err     error
		pngData fs.File
	)

	if wasmMode {
		pngData, err = efs.Open(name)
	} else {
		pngData, err = os.Open(name)
	}
	if err != nil {
		doLog(true, false, "loadSprite: Open: %v", err)
		return nil, nil, err
	}

	//Decode png
	m, err := png.Decode(pngData)
	if err != nil {
		doLog(true, false, "loadSprite: Decode: %v", err)
		return nil, nil, err
	}

	//Create image
	var (
		img, blurImg *ebiten.Image
		newBlur      image.Image
	)
	if doBlur {
		newBlur = blur.Box(m, islandReflectionBlur)
	}

	if unmanaged {
		img = ebiten.NewImageFromImageWithOptions(m, &ebiten.NewImageFromImageOptions{Unmanaged: true})
		if doBlur {
			blurImg = ebiten.NewImageFromImageWithOptions(newBlur, &ebiten.NewImageFromImageOptions{Unmanaged: true})
		}
	} else {
		img = ebiten.NewImageFromImage(m)
		if doBlur {
			blurImg = ebiten.NewImageFromImage(newBlur)
		}
	}

	return img, blurImg, nil
}

func scanIslandsFolder() {
	var dir []os.DirEntry
	var err error
	dirPath := dataDir + spritesDir + islandsDir

	doLog(true, true, "scanIslandsFolder: Scanning.")

	if wasmMode {
		dir, err = efs.ReadDir(dirPath)
	} else {
		dir, err = os.ReadDir(dirPath)
	}
	if err != nil {
		doLog(true, false, "scanIslandsFolder: readDir: %v", err)
		return
	}

	var islandsList []string
	for _, item := range dir {
		if item.IsDir() {
			islandsList = append(islandsList, item.Name())
		}
	}
	doLog(true, true, "Islands found: %v", strings.Join(islandsList, ", "))

	for _, island := range islandsList {
		_, err := os.ReadFile(dirPath + "/" + island + "/" + infoJsonFile)
		if err != nil {
			doLog(true, false, "Island '%v' has no %v file.", island, infoJsonFile)
			return
		}
	}
}
