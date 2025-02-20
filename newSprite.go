package main

import (
	"encoding/json"
	"image"
	"image/png"
	"io/fs"
	"os"
	"strings"

	"github.com/anthonynsimon/bild/blur"
	"github.com/hajimehoshi/ebiten/v2"
)

var islands []islandData

type islandInfoData struct {
	Comment, Name, Desc string

	Pos,
	Level int
}

const (
	infoJsonFile    = "info.json"
	mainSpriteName  = "world.png"
	spriteSheetName = "spritesheet.png"
	spriteSheetJson = "spritesheet.json"
)

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
		_, err := os.ReadFile(infoPath + infoJsonFile)
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
			})
		islandsAdded = append(islandsAdded, info.Name)
	}

	doLog(true, true, "Islands added: %v", strings.Join(islandsAdded, ", "))

	return nil
}

// Load sprites
func loadImage(path string, unmanaged bool, doBlur bool) (*ebiten.Image, *ebiten.Image, error) {

	//Open file
	var (
		err     error
		pngData fs.File
	)

	if wasmMode {
		pngData, err = efs.Open(path + ".png")
	} else {
		pngData, err = os.Open(path + ".png")
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
