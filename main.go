package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	dWinWidth, dWinHeight = 1280 / magScale, 720 / magScale
	dWinHeightHalf        = dWinHeight / 2
	dWinWidthHalf         = dWinWidth / 2
	magScale              = 4
	sampleRate            = 48000

	dataDir    = "data/"
	spritesDir = dataDir + "sprites/"
	txtDir     = dataDir + "txt/"
)

var (
	WASMMode     bool
	qtest, debug *bool
)

func main() {
	fmt.Printf("Game res: %v,%v (%vx) : (%v, %v)\n", dWinWidth, dWinHeight, magScale, dWinWidth*magScale, dWinHeight*magScale)
	dump := flag.Bool("dumpMusic", false, "Dump songs out as WAV and quit.")
	qtest = flag.Bool("qtest", false, "skip title screen")
	debug = flag.Bool("debug", false, "debug info")
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
	ebiten.SetVsyncEnabled(true)
	ebiten.SetScreenClearedEveryFrame(false)
	ebiten.SetWindowTitle("Pixel Pirates")

	cloudChunks = map[int]*cloudData{}
	initNoise()
	cloudbuf = ebiten.NewImage(dWinWidth, dWinHeightHalf)
	cloudblur = ebiten.NewImage(dWinWidth/cloudBlurAmountX, dWinHeightHalf/cloudBlurAmountY)
	worldGradImg = ebiten.NewImage(1, dWinHeight)
	loadSprites()

	if err := ebiten.RunGameWithOptions(newGame(), &ebiten.RunGameOptions{GraphicsLibrary: ebiten.GraphicsLibraryOpenGL}); err != nil {
		return
	}
}

func newGame() *Game {

	initSprites()

	gMode := GAME_TITLE
	if *qtest {
		gMode = GAME_PLAY
	}
	g := &Game{
		gameMode: gMode,
		colors: colorData{
			day: colors{
				water:   hexToRGB("00a0a7"),
				horizon: hexToRGB("cdffff"),
				sky:     hexToRGB("7c95dc"),
			},
			evening: colors{
				water:   hexToRGB("006064"),
				horizon: hexToRGB("303f9f"),
				sky:     hexToRGB("1a237e"),
			},
		},
	}

	lastUpdate = time.Now()
	go PlayTitleMusic(g)
	return g

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return dWinWidth, dWinHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.makeWave()
	g.makeAirWave()
	if g.gameMode == GAME_TITLE || g.gameMode == GAME_FADEOUT {
		g.drawTitle(screen)
	}
	if g.gameMode == GAME_PLAY {
		g.drawGame(screen)
	}
}
