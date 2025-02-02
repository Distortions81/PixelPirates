package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) startFade(toMode int, duration time.Duration, color color.NRGBA) {
	if g.modeTransition {
		return
	}
	g.fade = fadeData{
		fadeToMode: toMode, fadeStarted: time.Now(),
		duration: duration, color: color}
	g.modeTransition = true
}

func (g *Game) drawFade(screen *ebiten.Image) {
	sinceStart := time.Since(g.fade.fadeStarted)
	if sinceStart > g.fade.duration {
		g.modeTransition = false
	}

	durMS := float64(g.fade.duration.Milliseconds())
	sinceMS := float64(sinceStart.Milliseconds())

	value := min((sinceMS/durMS), 1.0) * 2

	var amount uint8
	if sinceStart < g.fade.duration/2 {
		//Fade out
		g.fade.fadeDirection = true
		amount = uint8(value * 255.0)
	} else {
		//Fade in
		if g.fade.fadeDirection {
			g.fade.fadeDirection = false
			g.gameMode = g.fade.fadeToMode
		}
		amount = uint8(254 - (value * 255.0))
	}
	fadeColor := color.NRGBA{R: g.fade.color.R, G: g.fade.color.G, B: g.fade.color.B, A: amount}
	vector.DrawFilledRect(screen, 0, 0, dWinWidth, dWinWidth, fadeColor, false)
}
