package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

// Game mode
const (
	GAME_BOOT = iota
	GAME_START
	GAME_TITLE
	GAME_PLAY
	GAME_ISLAND
	GAME_MAX //Don't use
)

var modeNames [GAME_MAX]string = [GAME_MAX]string{
	"BOOT",
	"START",
	"TITLE",
	"PLAY",
	"ISLAND",
}

type Game struct {
	gameMode int
	//Fade
	modeTransition bool

	//Draw
	displayStamp time.Time
	frameNumber  uint64

	debugBuf string

	//Input
	lastUpdate time.Time

	//Audio
	audioContext *audio.Context
	stopMusic    bool

	fade fadeData

	//Ocean specific
	boatPos, oldBoatPos fPoint
	//Waves
	wavesLines            [dWinHeightHalf]waveLine
	airWaveLines          [dWinHeightHalf]waveLine
	numWaves, numAirWaves int

	//Visit-Island specific
	islandChunks        map[int]*islandChunkData
	playPos, oldPlayPos fPoint

	playerFacing int

	visiting, canVisit *islandData

	//Clouds
	cloudChunks    map[int]*cloudData
	recycledChunks []*cloudData
	cloudsDirty    bool
	chunkIDTop     int

	//Sky & water colors
	envColors      colorData
	worldGradImg   *ebiten.Image
	worldGradDirty bool

	//Hardcoded sprites
	defPlayerSP, defCollision, boat2SP, sunSP, titleSP, clickStartSP *spriteItem
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
