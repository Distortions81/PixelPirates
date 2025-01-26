package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const dWinWidth, dWinHeight = 1280, 720

func main() {
	ebiten.SetTPS(ebiten.SyncWithFPS)
	ebiten.SetWindowSize(dWinWidth, dWinHeight)
	ebiten.SetWindowTitle("Pixel Pirates")

	if err := ebiten.RunGameWithOptions(newGame(), &ebiten.RunGameOptions{GraphicsLibrary: ebiten.GraphicsLibraryOpenGL}); err != nil {
		return
	}
}

type Game struct {
}

func newGame() *Game {
	return &Game{}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return dWinWidth, dWinHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
}

// Ebiten input handler
func (g *Game) Update() error {
	return nil
}
