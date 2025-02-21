package main

import (
	"flag"
	"fmt"
	"image/color"
	"runtime"
	"sort"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	dWinWidth, dWinHeight = 1280 / magScale, 720 / magScale
	dWinHeightHalf        = dWinHeight / 2
	dWinWidthHalf         = dWinWidth / 2
	magScale              = 4
	maxScale              = 24
	sampleRate            = 44100

	dataDir    = "data/"
	spritesDir = "sprites/"
	islandsDir = "islands/"
)

var (
	wasmMode                                              bool
	qtest, qisland, qlive, nomusic, debugMode, fullscreen *bool
)

func isWasm() bool {
	// In WASM, GOARCH is "wasm" and GOOS is "js"
	return runtime.GOARCH == "wasm" && runtime.GOOS == "js"
}

func main() {
	qtest = flag.Bool("qtest", false, "skip title screen")
	qisland = flag.Bool("qisland", false, "go directly to welcome island")
	qlive = flag.Bool("qlive", false, "live reload textures (slow)")
	nomusic = flag.Bool("nomusic", false, "disable music")
	debugMode = flag.Bool("debug", false, "debug info")
	fullscreen := flag.Bool("fullscreen", false, "fullscreen mode.")
	flag.Parse()

	if isWasm() {
		wasmMode = true

		//tvalue := true
		//tptr := &tvalue
		fvalue := false
		fptr := &fvalue

		qtest = fptr
		qisland = fptr
		nomusic = fptr
		debugMode = fptr
		fullscreen = fptr
	}

	if *debugMode {
		startLog()
		logDaemon()
		doLog(true, true, "Game res: %v,%v (%vx) : (%v, %v)", dWinWidth, dWinHeight, magScale, dWinWidth*magScale, dWinHeight*magScale)
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

	if err := ebiten.RunGameWithOptions(newGame(), &ebiten.RunGameOptions{GraphicsLibrary: ebiten.GraphicsLibraryOpenGL}); err != nil {
		return
	}
}

func newGame() *Game {

	g := &Game{
		gameMode: GAME_BOOT,
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
	case GAME_BOOT:
		screen.Fill(color.White)
		g.gameMode = GAME_START
		startGame(g)

		if *qisland {
			g.fade.fadeToMode = GAME_ISLAND
			modeChange(g)
			g.canVisit = &islands[0]
			visitIsland(g)
		} else if *qtest {
			g.fade.fadeToMode = GAME_PLAY
			modeChange(g)
		} else {
			g.fade.fadeToMode = GAME_TITLE
			modeChange(g)
		}

	case GAME_START:
		return
	default:
		screen.Fill(COLOR_BLACK)
		ebitenutil.DebugPrint(screen, "Inavlid Game Mode")
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

func (point fPoint) ToInt() iPoint {
	return iPoint{X: int(point.X), Y: int(point.Y)}
}

func (point iPoint) ToFloat() fPoint {
	return fPoint{X: float64(point.X), Y: float64(point.Y)}
}

func saveCollisionList(g *Game) {
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
