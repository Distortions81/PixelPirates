package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Ebiten input handler
func (g *Game) Update() error {

	pressedKeys := inpututil.AppendPressedKeys(nil)

	if g.gameMode == GAME_TITLE {
		if pressedKeys != nil ||
			inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) ||
			inpututil.IsMouseButtonJustPressed(ebiten.MouseButton1) {
			g.stopMusic = true
			g.gameMode = GAME_PLAY
			time.Sleep(time.Millisecond * 500)
			go PlayGameMusic(g)
		}
		return nil
	}

	for _, key := range pressedKeys {
		if key == ebiten.KeyW ||
			key == ebiten.KeyArrowUp {
		}
		if key == ebiten.KeyA ||
			key == ebiten.KeyArrowLeft {
		}
		if key == ebiten.KeyS ||
			key == ebiten.KeyArrowDown {
		}
		if key == ebiten.KeyD ||
			key == ebiten.KeyArrowRight {
		}
	}
	return nil
}
