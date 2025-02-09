package main

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	MaxBoatY = 25
	MinBoatY = -35

	boatYSpeed = 60 * 1000
	boatXSpeed = 10 * 1000

	playerSpeed = 11 * 1000
	turboSpeed  = 10
)

const (
	DIR_NORTH = iota
	DIR_NORTH_EAST
	DIR_EAST
	DIR_SOUTH_EAST
	DIR_SOUTH
	DIR_SOUTH_WEST
	DIR_WEST
	DIR_NORTH_WEST
)

// Ebiten input handler
func (g *Game) Update() error {

	defer func() {
		g.lastUpdate = time.Now()
		g.clampBoatPos()
	}()

	pressedKeys := inpututil.AppendPressedKeys(nil)

	if g.gameMode == GAME_TITLE {
		if pressedKeys != nil ||
			inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) ||
			inpututil.IsMouseButtonJustPressed(ebiten.MouseButton1) {
			if g.modeTransition {
				return nil
			}

			g.startFade(GAME_PLAY, time.Second, true, COLOR_WHITE, FADE_CROSSFADE)
		}
		return nil
	} else if g.gameMode == GAME_PLAY {
		g.oldBoatPos = g.boatPos

		xBase := float64(time.Since(g.lastUpdate).Microseconds()) / boatYSpeed
		yBase := float64(time.Since(g.lastUpdate).Microseconds()) / boatXSpeed

		xSpeed, ySpeed := yBase, xBase
		for _, key := range pressedKeys {
			if key == ebiten.KeyShiftLeft || key == ebiten.KeyShiftRight {
				xSpeed = yBase * turboSpeed
				ySpeed = xBase * turboSpeed
			}
		}
		for _, key := range pressedKeys {
			if key == ebiten.KeyW ||
				key == ebiten.KeyArrowUp {
				g.boatPos.Y -= ySpeed
			}
			if key == ebiten.KeyA ||
				key == ebiten.KeyArrowLeft {
				g.boatPos.X -= xSpeed
			}
			if key == ebiten.KeyS ||
				key == ebiten.KeyArrowDown {
				g.boatPos.Y += ySpeed
			}
			if key == ebiten.KeyD ||
				key == ebiten.KeyArrowRight {
				g.boatPos.X += xSpeed
			}
			if key == ebiten.KeyE {
				if g.canVisit != nil && !g.modeTransition {
					visitIsland(g)
					g.startFade(GAME_ISLAND, time.Second, true, COLOR_WHITE, FADE_CROSSFADE)
				}
				return nil
			}
		}
	} else if g.gameMode == GAME_ISLAND {

		if g.visiting == nil || g.visiting.visitSprite.image == nil {
			return nil
		}

		pBase := float64(time.Since(g.lastUpdate).Microseconds()) / playerSpeed

		g.oldPlayPos = g.playPos

		pSpeed := float64(pBase)
		for _, key := range pressedKeys {
			if key == ebiten.KeyShiftLeft || key == ebiten.KeyShiftRight {
				pSpeed *= turboSpeed
			}
		}

		sceneX, sceneY := float64(g.visiting.visitSprite.image.Bounds().Dx()), float64(g.visiting.visitSprite.image.Bounds().Dy())
		sceneX, sceneY = sceneX-dWinWidth, sceneY-dWinHeight
		oldPos := g.playPos
		for _, key := range pressedKeys {
			if key == ebiten.KeyE {
				g.startFade(GAME_PLAY, time.Second, true, COLOR_WHITE, FADE_CROSSFADE)
				return nil
			}
			if key == ebiten.KeyW ||
				key == ebiten.KeyArrowUp {
				g.playPos.Y -= pSpeed
			}
			if key == ebiten.KeyA ||
				key == ebiten.KeyArrowLeft {
				g.playPos.X -= pSpeed
			}
			if key == ebiten.KeyS ||
				key == ebiten.KeyArrowDown {
				g.playPos.Y += pSpeed
			}
			if key == ebiten.KeyD ||
				key == ebiten.KeyArrowRight {
				g.playPos.X += pSpeed
			}
		}
		face := directionFromCoords(oldPos.X-g.playPos.X, oldPos.Y-g.playPos.Y)
		if face >= 0 {
			g.playerFacing = face
		}
		g.playPos = clampPos(fPoint{X: 0, Y: 0}, g.playPos, fPoint{X: sceneX, Y: sceneY})
	}

	return nil
}

func clampPos(low, pos, max fPoint) fPoint {
	if pos.Y > max.Y {
		pos.Y = max.Y
	}
	if pos.Y < low.Y {
		pos.Y = low.Y
	}

	if pos.X > max.X {
		pos.X = max.X
	}
	if pos.X < low.X {
		pos.X = low.X
	}
	return pos
}

func (g *Game) clampBoatPos() {
	if g.boatPos.Y > MaxBoatY {
		g.boatPos.Y = MaxBoatY
	}
	if g.boatPos.Y < MinBoatY {
		g.boatPos.Y = MinBoatY
	}
}

func directionFromCoords(x, y float64) int {
	x = -x
	switch {
	case y > 0 && x == 0:
		return DIR_NORTH // north
	case x > 0 && y > 0:
		return DIR_NORTH_EAST // north-east
	case x > 0 && y == 0:
		return DIR_EAST // east
	case x > 0 && y < 0:
		return DIR_SOUTH_EAST // south-east
	case x == 0 && y < 0:
		return DIR_SOUTH // south
	case x < 0 && y < 0:
		return DIR_SOUTH_WEST // south-west
	case x < 0 && y == 0:
		return DIR_WEST // west
	case x < 0 && y > 0:
		return DIR_NORTH_WEST // north-west
	default:
		// x == 0 && y == 0 → no movement
		// or any unhandled case
		return DIR_SOUTH //default south
	}
}

var diag = math.Sqrt(2) / 2

func applyDirection(x, y float64, direction int, speed float64) (float64, float64) {
	// 45° diagonal movement factor (cos 45°, sin 45°)

	diagonal := diag * speed

	switch direction {
	case 0:
		// north
		y += speed
	case 1:
		// north-east
		x += diagonal
		y += diagonal
	case 2:
		// east
		x += speed
	case 3:
		// south-east
		x += diagonal
		y -= diagonal
	case 4:
		// south
		y -= speed
	case 5:
		// south-west
		x -= diagonal
		y -= diagonal
	case 6:
		// west
		x -= speed
	case 7:
		// north-west
		x -= diagonal
		y += diagonal
	default:
		// If the direction is invalid, do nothing
	}

	return x, y
}
