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
	maxScale              = 24
	sampleRate            = 48000

	dataDir    = "data/"
	spritesDir = dataDir + "sprites/"
	txtDir     = dataDir + "txt/"
)

var (
	wasmMode                                  bool
	nomusic, qtest, qlive, qisland, debugMode *bool
	fullscreen                                *bool
)

func main() {
	dump := flag.Bool("dumpMusic", false, "Dump songs out as WAV and quit.")
	qtest = flag.Bool("qtest", false, "skip title screen")
	qisland = flag.Bool("qisland", false, "go directly to welcome island")
	qlive = flag.Bool("qlive", false, "live reload textures (slow)")
	nomusic = flag.Bool("nomusic", false, "disable music")
	debugMode = flag.Bool("debug", false, "debug info")

	fullscreen := flag.Bool("fullscreen", false, "fullscreen mode.")
	flag.Parse()

	startLog()
	logDaemon()
	doLog(true, true, "Game res: %v,%v (%vx) : (%v, %v)", dWinWidth, dWinHeight, magScale, dWinWidth*magScale, dWinHeight*magScale)

	if *dump {
		//dumpMusic()
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
	ebiten.SetWindowSizeLimits(dWinWidth, dWinHeight, dWinWidth*maxScale, dWinHeight*maxScale)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetScreenClearedEveryFrame(false)
	ebiten.SetFullscreen(*fullscreen)
	ebiten.SetWindowTitle("Pixel Pirates")

	loadSprites()

	if *qlive {
		go func() {
			for {
				loadSprites()
				time.Sleep(time.Second * 1)
				doLog(true, true, "Reloading textures.")
			}
		}()
	}

	if err := ebiten.RunGameWithOptions(newGame(), &ebiten.RunGameOptions{GraphicsLibrary: ebiten.GraphicsLibraryOpenGL}); err != nil {
		return
	}
}

func newGame() *Game {

	gMode := GAME_TITLE

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

	initNoise(g)
	initSprites(g)
	initIslands(g)

	g.audioContext = audio.NewContext(sampleRate)
	g.cloudChunks = map[int]*cloudData{}
	g.worldGradImg = ebiten.NewImage(1, dWinHeight)
	g.worldGradDirty = true

	if *qisland {
		g.visiting = &islands[0]
		g.gameMode = GAME_ISLAND
	} else if *qtest {
		g.gameMode = GAME_PLAY
	}

	go func() {
		time.Sleep(time.Second)
		playMusicPlaylist(g, g.gameMode, gameModePlaylists[g.gameMode])
	}()

	g.startFade(g.gameMode, time.Second*2, false, COLOR_BLACK, FADE_IN)

	g.lastUpdate = time.Now()
	return g

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return dWinWidth, dWinHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.frameNumber++
	startTime := time.Now()

	switch g.gameMode {
	case GAME_TITLE:
		g.drawTitle(screen)
	case GAME_PLAY:
		g.drawGame(screen)
	case GAME_ISLAND:
		g.drawIsland(screen)
	default:
		screen.Fill(COLOR_BLACK)
		ebitenutil.DebugPrint(screen, "Inavlid Game Mode")
		return
	}
	if g.modeTransition {
		g.drawFade(screen)
	}

	if *debugMode {
		if g.frameNumber%60 == 0 {
			renderTime := time.Since(startTime).Microseconds()
			displayTime := time.Since(g.displayStamp).Microseconds()

			g.debugBuf = fmt.Sprintf("Render: %4du, Display: %0.2fms, %%%0.2f, FPS: %0.2f",
				renderTime,
				float64(displayTime)/1000,
				float64(renderTime)/float64(displayTime)*100,
				ebiten.ActualFPS())
		}

		ebitenutil.DebugPrintAt(screen, g.debugBuf, 0, dWinHeight-16)
		g.displayStamp = time.Now()
	}

}
