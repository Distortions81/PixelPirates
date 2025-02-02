package main

import (
	"image/color"
	"time"
)

const (
	GAME_TITLE = iota
	GAME_FADEOUT
	GAME_PLAY
)

type Game struct {
	fadeStart time.Time
	gameMode  int
	stopMusic bool
	boatPos   fPoint

	colors colorData
}

// hsv represents a color in hsv space
type hsv struct {
	H, S, V float64
}

type colorData struct {
	day, evening colors
}

type colors struct {
	sky, water, horizon color.RGBA
}

type iPoint struct {
	X, Y int
}

type fPoint struct {
	X, Y float64
}
