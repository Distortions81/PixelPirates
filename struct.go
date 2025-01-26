package main

import "image/color"

type Game struct {
	gameMode        int
	boatPos, camPos point

	colors colorData
}

// HSV represents a color in HSV space
type HSV struct {
	H, S, V float64
}

type colorData struct {
	day, evening colors
}

type colors struct {
	sky, water, horizon color.RGBA
}

type point struct {
	X, Y int
}
