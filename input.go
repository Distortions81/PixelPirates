package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Ebiten input handler
func (g *Game) Update() error {

	pressedKeys := inpututil.AppendPressedKeys(nil)

	if g.gameMode == GAME_TITLE {
		if pressedKeys != nil {
			g.gameMode = GAME_PLAY
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
