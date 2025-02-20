package main

import (
	"image"
	"image/png"
	"io/fs"
	"os"

	"github.com/anthonynsimon/bild/blur"
	"github.com/hajimehoshi/ebiten/v2"
)

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
