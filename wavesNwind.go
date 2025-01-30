package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/chewxy/math32"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type waveLine struct {
	waves []waveData
	count int
}

type waveData struct {
	x     int
	start time.Time
	life  time.Duration
}

const (
	persVal           = 10 //Used for perspective
	skyPersVal        = 5  //Used for perspective (airwaves)
	colorVal          = 10 //Used for perspective (waves)
	maxWaves          = 2000
	spawnPerFrame     = 66
	minWaveLifeMS     = 100
	maxWaveLifeRandMS = 500

	maxAirWaves          = 10
	minAirWaveLifeMS     = 2000
	maxAirWaveLifeRandMS = 4000
)

var (
	wavesLines            [dWinHeightHalf]waveLine
	airWaveLines          [dWinHeightHalf]waveLine
	numWaves, numAirWaves int
)

func drawAir(g *Game, screen *ebiten.Image) {

	/*
		for _, wave := range airwave {
			//Fade in and out
			lifeLeft := float64(wave.life-time.Since(wave.start)) / float64(wave.life) * 2
			alpha := uint8(math.Min(math.Abs(255-lifeLeft*255), 255.0))
			waveColor := color.NRGBA{R: 255, G: 255, B: 255, A: 85 - alpha/3}

			bPos := g.boatPos.X * wave.logVal

			// Calculate the raw modulo
			rawMod := wave.linVal + bPos/dWinWidth
			modVal := math.Mod(rawMod, 1)

			// Ensure the modulo result is positive
			if modVal < 0 {
				modVal += 1
			}

			// Inverse X for correct direction, then apply the positive modulo, then re-expand
			x := dWinWidth - float32(modVal*dWinWidth)

			// Start at horizon, add logVal * half winHeight
			y := float32(wave.logVal * dWinHeightHalf)

			// Width is based on logVal
			width := float32(1 + (wave.logVal * persVal))

			vector.DrawFilledRect(screen, x, dWinHeightHalf-y+1, width, 1, waveColor, false)
		}
	*/
}

func drawWaves(g *Game, screen *ebiten.Image) {
	for y, line := range wavesLines {
		for _, wave := range line.waves {
			// Lower alpha for waves that are farther away
			alpha := uint8(math32.Min(64+(float32(y)*3.0), 255))
			waveColor := color.NRGBA{R: 255, G: 255, B: 255, A: alpha}

			bPos := float64(wave.x) //g.boatPos.X*float64(y)

			// Width is based on logVal
			width := float32(y) / 11.0
			width = max(width, 2)

			vector.DrawFilledRect(screen, float32(bPos), float32(dWinHeightHalf+y), width, 1, waveColor, false)
		}
	}
}

func (g Game) makeWave() {
	if numWaves > 0 {
		for l, line := range wavesLines {
			for w, wave := range line.waves {
				if time.Since(wave.start) > wave.life {
					// Remove the element at index i
					wavesLines[l].waves = append(wavesLines[l].waves[:w], wavesLines[l].waves[w+1:]...)
					numWaves--
					wavesLines[l].count--
					break
				}
			}
		}
	}
	spawns := 0
	collisions := 0
spawn:
	for spawns < spawnPerFrame && numWaves < maxWaves && collisions < 15 {
		y := int(logDistWave(rand.Float64()) * dWinHeightHalf)
		y = min(y, dWinHeightHalf-1)
		y = max(y, 0)

		var newWave waveData

		newWave = waveData{
			x:     rand.Intn(dWinWidth),
			start: time.Now(),
			life:  time.Millisecond * time.Duration(minWaveLifeMS+(rand.Float64()*maxWaveLifeRandMS)),
		}
		for _, check := range wavesLines[y].waves {
			if check.x == newWave.x {
				collisions++
				goto spawn
			}
		}

		wavesLines[y].waves = append(wavesLines[y].waves, newWave)
		wavesLines[y].count++
		numWaves++
		spawns++
	}
	//fmt.Printf("C: %v\n", collisions)
}

func (g Game) makeAirWave() {
	/*
		for i := numAirWaves - 1; i >= 0; i-- {
			if time.Since(airwave[i].start) > airwave[i].life {
				// Remove the element at index i
				airwave = append(airwave[:i], airwave[i+1:]...)
				numAirWaves--
			}
		}
		for z := 0; z < spawnPerFrame && numAirWaves < maxAirWaves; z++ {
			newWave := waveData{
				logVal: logDistAirwave(rand.Float64()),
				linVal: rand.Float64(),
				start:  time.Now(),
				life:   time.Millisecond * time.Duration(minAirWaveLifeMS+(rand.Float64()*maxAirWaveLifeRandMS))}
			airwave = append(airwave, newWave)
			numAirWaves++
		}
	*/
}

func logDistWave(uniform float64) float64 {
	biased := math.Abs(math.Log(float64(1 - uniform)))
	return math.Min(biased/persVal, 1.0)
}

func logDistAirwave(uniform float64) float64 {
	biased := math.Abs(math.Log(1 - uniform))
	return math.Min(biased/skyPersVal, 1.0)
}
