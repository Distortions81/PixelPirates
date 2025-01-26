package main

import "image/color"

type Game struct {
	gameMode        int
	boatPos, camPos point

	colors colorData
}

type colorData struct {
	sky, water, horizon color.RGBA
}

type point struct {
	X, Y int
}
