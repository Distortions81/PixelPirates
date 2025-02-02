package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	wasmMode     bool
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
		dumpMusic()
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
		envColors: colorData{
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

	g.startFade(g.gameMode, time.Second*2, false, COLOR_BLACK, FADE_IN)

	lastUpdate = time.Now()
	return g

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return dWinWidth, dWinHeight
}

func (g *Game) Draw(screen *ebiten.Image) {

	switch g.gameMode {
	case GAME_TITLE:
		g.drawTitle(screen)
	case GAME_PLAY:
		g.drawGame(screen)
	case GAME_ISLAND:
		screen.Fill(COLOR_BLACK)
		g.drawIsland(screen)
	default:
		screen.Fill(COLOR_BLACK)
		ebitenutil.DebugPrint(screen, "Inavlid Game Mode")
		return
	}

	if g.modeTransition {
		g.drawFade(screen)
	}
}
