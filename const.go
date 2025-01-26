package main

const (
	dataDir    = "data/"
	spritesDir = dataDir + "sprites/"
	txtDir     = dataDir + "txt/"
)

const (
	magScale = 8
)

const (
	GAME_TITLE = iota
	GAME_MENU
	GAME_JOIN
	GAME_LOBBY
	GAME_PLAY
	GAME_DEAD
	GAME_QUIT
)
