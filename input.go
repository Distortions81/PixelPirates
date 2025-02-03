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

	boatYSpeed = 60.0 * 1000
	boatXSpeed = 10 * 1000

	playerSpeed = 1
)

// Ebiten input handler
func (g *Game) Update() error {

	defer func() {
		lastUpdate = time.Now()
		g.clampBoatPos()
	}()

	pressedKeys := inpututil.AppendPressedKeys(nil)

	if g.gameMode == GAME_TITLE {
		if pressedKeys != nil ||
			inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) ||
			inpututil.IsMouseButtonJustPressed(ebiten.MouseButton1) {
			g.startFade(GAME_PLAY, time.Second, true, COLOR_WHITE, FADE_CROSSFADE)
		}
		return nil
	} else if g.gameMode == GAME_PLAY {

		vspeed := float64(time.Since(lastUpdate).Microseconds()) / boatYSpeed
		hspeed := float64(time.Since(lastUpdate).Microseconds()) / boatXSpeed

		xs := hspeed
		for _, key := range pressedKeys {
			if key == ebiten.KeyShiftLeft || key == ebiten.KeyShiftRight {
				xs = hspeed * 50
			}
		}
		for _, key := range pressedKeys {
			if key == ebiten.KeyW ||
				key == ebiten.KeyArrowUp {
				g.boatPos.Y -= vspeed
			}
			if key == ebiten.KeyA ||
				key == ebiten.KeyArrowLeft {
				g.boatPos.X -= xs
			}
			if key == ebiten.KeyS ||
				key == ebiten.KeyArrowDown {
				g.boatPos.Y += vspeed
			}
			if key == ebiten.KeyD ||
				key == ebiten.KeyArrowRight {
				g.boatPos.X += xs
			}
			if key == ebiten.KeyE {
				if g.canVisit != nil {
					g.visiting = g.canVisit
					g.startFade(GAME_ISLAND, time.Second, true, COLOR_WHITE, FADE_CROSSFADE)
				}
				return nil
			}
		}
	} else if g.gameMode == GAME_ISLAND {
		ps := float64(playerSpeed)
		for _, key := range pressedKeys {
			if key == ebiten.KeyShiftLeft || key == ebiten.KeyShiftRight {
				ps = playerSpeed * 50
			}
		}
		for _, key := range pressedKeys {
			if key == ebiten.KeyE {
				g.startFade(GAME_PLAY, time.Second, true, COLOR_WHITE, FADE_CROSSFADE)
			}
			if key == ebiten.KeyW ||
				key == ebiten.KeyArrowUp {
				g.playerPos.Y -= ps
			}
			if key == ebiten.KeyA ||
				key == ebiten.KeyArrowLeft {
				g.playerPos.X -= ps
			}
			if key == ebiten.KeyS ||
				key == ebiten.KeyArrowDown {
				g.playerPos.Y += ps
			}
			if key == ebiten.KeyD ||
				key == ebiten.KeyArrowRight {
				g.playerPos.X += ps
			}
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
