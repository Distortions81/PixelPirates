package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	dWinWidth, dWinHeight = 1280 / magScale, 720 / magScale
	magScale              = 4
	sampleRate            = 48000
	verbose               = false
)

var (
	WASMMode     bool
	fxtest       *bool
	audioContext *audio.Context
)

func main() {
	fmt.Printf("Game res: %v,%v (%vx) : (%v, %v)\n", dWinWidth, dWinHeight, magScale, dWinWidth*magScale, dWinHeight*magScale)
	dump := flag.Bool("dumpMusic", false, "Dump songs out as WAV and quit.")
	fxtest = flag.Bool("fxtest", false, "test sound effects.")
	flag.Parse()

	audioContext = audio.NewContext(sampleRate)

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
	ebiten.SetWindowSize(dWinWidth*magScale, dWinHeight*magScale)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Pixel Pirates")

	loadSprites()

	if err := ebiten.RunGameWithOptions(newGame(), &ebiten.RunGameOptions{GraphicsLibrary: ebiten.GraphicsLibraryOpenGL}); err != nil {
		return
	}
}

var (
	boat1SP, boat2SP, boat2SP_flag, sunSP, titleSP, clickStartSP *spriteItem
)

func newGame() *Game {

	boat1SP = spriteList["boat1"]
	boat2SP = spriteList["boat2"]
	boat2SP_flag = spriteList["boat2-flag"]
	sunSP = spriteList["sun"]
	titleSP = spriteList["title"]
	clickStartSP = spriteList["clickstart"]

	g := &Game{
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

	if *fxtest {
		go PlayFx(g)
	} else {
		go PlayTitleMusic(g)
	}
	return g

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
