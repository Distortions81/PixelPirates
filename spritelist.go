package main

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var spriteNames []string = []string{
	"boat",
	"sun",
	"title",
}

func loadSprites() {
	spriteList = make(map[string]spriteItem)

	for _, name := range spriteNames {
		image, err := loadSprite(name+".png", false)
		if err == nil {
			spriteList[name] = spriteItem{Name: name, Size: point(image.Bounds().Max), image: image}
			doLog(true, "loading sprite '"+name+"'")
		} else {
			doLog(true, "loading sprite '"+name+"' failed.")
			os.Exit(1)
		}
	}
}

var spriteList map[string]spriteItem

type spriteItem struct {
	Name string
	Size point

	image *ebiten.Image
}
