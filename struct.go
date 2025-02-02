package main

import (
	"image/color"
	"time"
)

// Game mode
const (
	GAME_TITLE = iota
	GAME_PLAY
	GAME_ISLAND
	GAME_MAX //Don't use
)

const (
	FADE_CROSSFADE = iota
	FADE_IN
	FADE_OUT
)

type fadeData struct {
	fadeToMode int

	fadeStarted   time.Time
	duration      time.Duration
	fadeDirection bool
	stopMusic     bool

	color color.NRGBA
}

type Game struct {
	gameMode       int
	modeTransition bool
	fade           fadeData

	stopMusic bool

	boatPos   fPoint
	envColors colorData

	visiting, canVisit *islandData
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
