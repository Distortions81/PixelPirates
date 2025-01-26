package main

import "github.com/hajimehoshi/ebiten/v2"

var spriteNames []string = []string{
	"boat",
}

func loadSprites() {
	spriteList = make(map[string]spriteItem)

	for _, name := range spriteNames {
		image, err := loadSprite(name+".png", false)
		if err == nil {
			spriteList[name] = spriteItem{Name: name, Size: point(image.Bounds().Max), image: image}
		} else {
			doLog(true, "loading sprite '"+name+"' failed.")
		}
	}
}

var spriteList map[string]spriteItem

type spriteItem struct {
	Name string
	Size point

	image *ebiten.Image
}
