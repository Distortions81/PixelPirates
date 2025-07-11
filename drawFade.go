package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	FADE_CROSSFADE = iota
	FADE_IN
	FADE_OUT
)

type fadeData struct {
	fadeToMode int

	fadeStarted   time.Time
	fadeType      int
	duration      time.Duration
	fadeDirection bool
	stopMusic     bool

	color color.NRGBA
}

func startGame(g *Game) {

	initSprites(g)
	loadSprites()

	g.cloudChunks = map[int]*cloudData{}
	g.worldGradImg = ebiten.NewImage(1, dWinHeight)
	g.worldGradDirty = true

	if wasmMode || *nomusic {
		return
	}

}

func (g *Game) startFade(toMode int, duration time.Duration, stopMusic bool, color color.NRGBA, fadeType int) {
	if g.modeTransition {
		return
	}

	fadeStart := time.Now()
	if fadeType == FADE_IN {
		fadeStart = fadeStart.Add(-(duration / 2))
	}
	g.fade = fadeData{
		fadeToMode: toMode, fadeStarted: fadeStart, fadeType: fadeType,
		duration: duration, color: color, stopMusic: stopMusic}

	g.modeTransition = true

	if duration == 0 {
		modeChange(g)
		return
	}
}

func (g *Game) drawFade(screen *ebiten.Image) {
	sinceStart := time.Since(g.fade.fadeStarted)
	if sinceStart > g.fade.duration {
		g.modeTransition = false
	}

	durMS := float64(g.fade.duration.Milliseconds())
	sinceMS := float64(sinceStart.Milliseconds())

	var value float64
	if g.fade.fadeType != FADE_OUT {
		value = min((sinceMS/durMS), 1.0) * 2
	} else {
		value = min((sinceMS / durMS), 1.0)
	}

	var amount uint8
	if sinceStart < g.fade.duration/2 || g.fade.fadeType == FADE_OUT {
		//Fade out
		g.fade.fadeDirection = true
		amount = uint8(value * 255.0)
	} else {
		//Fade in
		if g.fade.fadeDirection {
			g.fade.fadeDirection = false
			modeChange(g)
		}
		amount = uint8(254 - (value * 255.0))
	}
	fadeColor := color.NRGBA{R: g.fade.color.R, G: g.fade.color.G, B: g.fade.color.B, A: amount}
	vector.DrawFilledRect(screen, 0, 0, dWinWidth, dWinWidth, fadeColor, false)

}

func modeChange(g *Game) {

	if g.fade.stopMusic {
		g.stopMusic = true
	}

	oldMode := g.gameMode
	g.gameMode = g.fade.fadeToMode

	//Mode deinits
	if oldMode == GAME_TITLE {
		g.clickStartSP.image.Deallocate()
		g.clickStartSP.image = nil
		g.titleSP.image.Deallocate()
		g.titleSP.image = nil
		doLog(true, true, "Deallocated title screen assets.")

	} else if oldMode == GAME_ISLAND && g.gameMode != GAME_ROOM {
		g.inIsland.spriteSheet.image.Deallocate()
		g.inIsland.spriteSheet.image = nil
		g.inIsland = nil
		doLog(true, true, "Deallocated island spriteSheet.")
	} else if oldMode == GAME_ROOM {
		g.playPos = g.inRoom.oldPos
		g.inRoom = nil
	}

	//Mode inits
	if g.gameMode == GAME_PLAY {
		initNoise(g)
		scanIslandsFolder()
		initIslands(g)
	} else if g.gameMode == GAME_ISLAND {
		scanIslandsFolder()
	} else if g.gameMode == GAME_TITLE {
		initNoise(g)
	} else if g.gameMode == GAME_ROOM {
		g.inRoom.oldPos = g.playPos
	}
	if *debugMode {
		doLog(true, true, "Mode: %v --> %v", modeNames[oldMode], modeNames[g.gameMode])
	}

	go func(g *Game) {
		time.Sleep(time.Second)
		if g.audioContext == nil {
			g.audioContext = audio.NewContext(sampleRate)
		}
		playMusicPlaylist(g, g.gameMode, gameModePlaylists[g.gameMode])
	}(g)
}
