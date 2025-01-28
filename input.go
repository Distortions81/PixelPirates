package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var lastUpdate time.Time

// Ebiten input handler
func (g *Game) Update() error {

	defer func() {
		lastUpdate = time.Now()
	}()

	pressedKeys := inpututil.AppendPressedKeys(nil)

	if g.gameMode == GAME_TITLE {
		if pressedKeys != nil ||
			inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) ||
			inpututil.IsMouseButtonJustPressed(ebiten.MouseButton1) {
			g.stopMusic = true
			g.gameMode = GAME_PLAY

			go PlayGameMusic(g)
		}
		return nil
	}

	vspeed := float32(time.Since(lastUpdate).Milliseconds()) / 120.0
	hspeed := float32(time.Since(lastUpdate).Milliseconds()) / 60.0

	for _, key := range pressedKeys {
		if key == ebiten.KeyW ||
			key == ebiten.KeyArrowUp {
			g.boatPos.Y -= vspeed
		}
		if key == ebiten.KeyA ||
			key == ebiten.KeyArrowLeft {
			g.boatPos.X -= hspeed
		}
		if key == ebiten.KeyS ||
			key == ebiten.KeyArrowDown {
			g.boatPos.Y += vspeed
		}
		if key == ebiten.KeyD ||
			key == ebiten.KeyArrowRight {
			g.boatPos.X += hspeed
		}
	}

	return nil
}
