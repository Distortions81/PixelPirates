package main

import "github.com/hajimehoshi/ebiten/v2"

// TODO: Update sprite tags instead,
var (
	playerImg *ebiten.Image
	moveFix   [9]int = [9]int{12, 12, 2, 3, 4, 6, 8, 9, 10}
	faceFix   [9]int = [9]int{0, 4, 3, 2, 1, 0, 7, 6, 5}
)
