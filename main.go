package main

import (
	"flag"
	"fmt"
	"runtime"
	"sort"
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
	spritesDir = "sprites/"
	islandsDir = "islands/"
)

var (
	wasmMode                                  bool
	nomusic, qtest, qlive, qisland, debugMode *bool
	fullscreen                                *bool
)

func isWasm() bool {
	// In WASM, GOARCH is "wasm" and GOOS is "js"
	return runtime.GOARCH == "wasm" && runtime.GOOS == "js"
}

func main() {
	dump := flag.Bool("dumpMusic", false, "Dump songs out as WAV and quit.")
	qtest = flag.Bool("qtest", false, "skip title screen")
	qisland = flag.Bool("qisland", false, "go directly to welcome island")
	qlive = flag.Bool("qlive", false, "live reload textures (slow)")
	nomusic = flag.Bool("nomusic", false, "disable music")
	debugMode = flag.Bool("debug", false, "debug info")

	fullscreen := flag.Bool("fullscreen", false, "fullscreen mode.")
	flag.Parse()

	if isWasm() {
		wasmMode = true

		value := true
		ptr := &value
		debugMode = ptr
		//nomusic = ptr
	}

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
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetWindowTitle("Pixel Pirates")

	loadSprites()

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

	/*
		if *qlive {
			go func() {
				for {
					loadSprite(islandsDir+"island-scene1/island-scene1", islands[0].visitSprite, true)
					time.Sleep(time.Second * 1)
					doLog(true, true, "Reloading textures.")
				}
			}()
		}
	*/

	g.audioContext = audio.NewContext(sampleRate)
	g.cloudChunks = map[int]*cloudData{}
	g.worldGradImg = ebiten.NewImage(1, dWinHeight)
	g.worldGradDirty = true

	if *qisland {
		g.canVisit = &islands[0]
		visitIsland(g)
		g.gameMode = GAME_ISLAND
	} else if *qtest {
		g.gameMode = GAME_PLAY
	}

	if !wasmMode {
		go func(g *Game) {
			time.Sleep(time.Second)
			playMusicPlaylist(g, g.gameMode, gameModePlaylists[g.gameMode])
		}(g)
	}

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

	//Operations that can only happen after game start
	if g.frameNumber == 1 {

	}

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

func (point fPoint) QuantizePoint() iPoint {
	return iPoint{X: int(point.X), Y: int(point.Y)}
}

func savePlayerCollisionList(g *Game) {
	if len(g.defCollision.collision) > 0 {
		return
	}
	for x := 0; x < g.defCollision.image.Bounds().Dx(); x++ {
		for y := 0; y < g.defCollision.image.Bounds().Dy(); y++ {
			img := g.defCollision.image.At(x, y)
			_, _, _, alpha := img.RGBA()
			if alpha > 254 {
				g.defCollision.collision = append(g.defCollision.collision,
					iPoint{X: x, Y: y})
			}
		}
	}
}

func SortLinePoints(points []iPoint, a, b iPoint) []iPoint {
	// Direction vector for the line
	dx := b.X - a.X
	dy := b.Y - a.Y

	// We only need the dot product with (dx, dy) to establish ordering.
	// That is, for each point p = (px, py), we compute:
	//     t = (px - a.X)*dx + (py - a.Y)*dy
	// We do NOT need to divide by (dx^2 + dy^2) for sorting because
	// the ratio is monotonic and the denominator would be constant anyway.

	type paramPoint struct {
		pt   iPoint
		dist int
	}

	paramPoints := make([]paramPoint, len(points))
	for i, p := range points {
		px, py := p.X, p.Y
		// Dot product relative to (a.X,a.Y):
		t := (px-a.X)*dx + (py-a.Y)*dy
		paramPoints[i] = paramPoint{pt: p, dist: t}
	}

	// Sort by dist (the dot product value).
	sort.Slice(paramPoints, func(i, j int) bool {
		return paramPoints[i].dist < paramPoints[j].dist
	})

	// Extract the sorted points
	sortedPoints := make([]iPoint, len(points))
	for i, pp := range paramPoints {
		sortedPoints[i] = pp.pt
	}

	return sortedPoints
}
