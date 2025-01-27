package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	dWinWidth, dWinHeight = 1280, 720
)

var WASMMode bool

func main() {
	dump := flag.Bool("dumpMusic", false, "Dump songs out as WAV and quit.")
	flag.Parse()

	if *dump {
		DumpMusic()
		fmt.Println("Music dump complete.")
		return
	}
	/*
		go func() {
			http.ListenAndServe("localhost:6060", nil)
		}()
	*/

	ebiten.SetTPS(ebiten.SyncWithFPS)
	ebiten.SetWindowSize(dWinWidth, dWinHeight)
	ebiten.SetWindowTitle("Pixel Pirates")

	loadSprites()

	if err := ebiten.RunGameWithOptions(newGame(), &ebiten.RunGameOptions{GraphicsLibrary: ebiten.GraphicsLibraryOpenGL}); err != nil {
		return
	}
}

var (
	boatSP, sunSP, titleSP, clickStartSP *ebiten.Image
)

func newGame() *Game {

	boatSP = spriteList["boat"].image
	sunSP = spriteList["sun"].image
	titleSP = spriteList["title"].image
	clickStartSP = spriteList["clickstart"].image
	go PlayMusic()

	return &Game{
		gameMode: GAME_TITLE,
		colors: colorData{
			day: colors{
				water:   hexToRGB("009688"),
				horizon: hexToRGB("4fc3f7"),
				sky:     hexToRGB("5b6ee1"),
			},
			evening: colors{
				water:   hexToRGB("006064"),
				horizon: hexToRGB("303f9f"),
				sky:     hexToRGB("1a237e"),
			},
		},
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return dWinWidth, dWinHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.gameMode == GAME_TITLE {
		g.drawTitle(screen)
	} else if g.gameMode == GAME_PLAY {
		g.drawGame(screen)
	}
}
