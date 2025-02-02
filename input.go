package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var lastUpdate time.Time

const (
	MaxBoatY = 25
	MinBoatY = -35

	//larger numbers are slower
	vspeed = 60.0 * 1000
	xspeed = 10 * 1000
)

// Ebiten input handler
func (g *Game) Update() error {

	defer func() {
		lastUpdate = time.Now()
		g.clampBoatPos()
	}()

	pressedKeys := inpututil.AppendPressedKeys(nil)

	if g.gameMode == GAME_FADEOUT {
		return nil
	} else if g.gameMode == GAME_TITLE {
		if pressedKeys != nil ||
			inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) ||
			inpututil.IsMouseButtonJustPressed(ebiten.MouseButton1) {
			g.stopMusic = true
			g.gameMode = GAME_FADEOUT
			g.fadeStart = time.Now()
			initNoise()
			go PlayGameMusic(g)
		}
		return nil
	}

	vspeed := float64(time.Since(lastUpdate).Microseconds()) / vspeed
	hspeed := float64(time.Since(lastUpdate).Microseconds()) / xspeed

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

func (g *Game) clampBoatPos() {
	if g.boatPos.Y > MaxBoatY {
		g.boatPos.Y = MaxBoatY
	}
	if g.boatPos.Y < MinBoatY {
		g.boatPos.Y = MinBoatY
	}
}
