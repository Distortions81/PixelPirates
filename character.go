package main

import "github.com/hajimehoshi/ebiten/v2"

// TODO: Update sprite tags instead,
var (
	playerImg *ebiten.Image
	moveFix   [9]int = [9]int{12, 12, 2, 3, 4, 6, 8, 9, 10}
	faceFix   [9]int = [9]int{0, 4, 3, 2, 1, 0, 7, 6, 5}
)

func loadDefaultChar(g *Game) {

	if g.defPlayerSP.image == nil {
		loadSprite(g.defPlayerSP.Fullpath, g.defPlayerSP, true)
	}
	if g.defCollision.image == nil {
		loadSprite(g.defCollision.Fullpath, g.defCollision, true)
	}

	if len(g.defCollision.collision) == 0 {
		//Save player size
		img := getFrameNumber(0, g.defPlayerSP, 0)
		pWidth = img.Bounds().Dx()
		pHeight = img.Bounds().Dy()

		saveCollisionList(g)
	}
}
