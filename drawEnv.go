package main

import (
	"image"
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
	persVal      = 8 //Used for perspective
	skyPersVal   = 5 //AirWaves perspective
	distParallax = 0.01

	//Waves
	// This helps prevent the waves from getting into visually noticable cycles.
	spawnPerFrame = 66
	//Performance safeguard
	maxCollisions = spawnPerFrame / 4

	//Sea
	maxWaves          = 600
	minWaveLifeMS     = 100
	maxWaveLifeRandMS = 500
	waveAlpha         = 25
	waveDistAlpha     = 1.5

	//Air
	windSpeed            = 7 //mph-like
	maxAirWaves          = 10
	minAirWaveLifeMS     = 2000
	maxAirWaveLifeRandMS = 4000

	//Sun reflect -- disabled
	sunReflectHeight = 10.0
	sunReflectAlpha  = 0.8
	sunX             = 64.0

	//Water gradient
	waterStartColor   = 175
	waterHueShift     = 25
	waterBrightStart  = 0.9
	waterDarkenDivide = 3
	waterSaturate     = 0.8

	//Sky gradient
	skyStartColor   = 220
	skyHueShift     = -40
	skyBrightStart  = 1.0
	skyDarkenDivide = 2.5
	skySaturate     = 0.5

	cloudReflectAlpha   = 0.5
	cloudReflectStretch = 1.5
)

func (g *Game) updateWorldGrad() {
	g.worldGradImg.Clear()

	var y float32
	for y = 0; y < dWinHeightHalf; y++ {
		//Water
		color := color.RGBA{}
		vVal := (float64(y) / dWinHeightHalf)
		color = hsvToRGB(hsv{
			H: waterStartColor + (vVal * waterHueShift),
			S: waterSaturate,
			V: waterBrightStart - math.Min(vVal/waterDarkenDivide, 1.0)})
		vector.DrawFilledRect(g.worldGradImg, 0, dWinHeightHalf+y,
			1, 1, color, false)

		//Sky
		color = hsvToRGB(hsv{
			H: skyStartColor + (vVal * skyHueShift),
			S: skySaturate,
			V: skyBrightStart - math.Min(((1.0-vVal)/skyDarkenDivide), 1.0)})
		vector.DrawFilledRect(g.worldGradImg, 0, y, 1, 1, color, false)

	}

	//Horizon
	vector.DrawFilledRect(g.worldGradImg, 0, dWinHeightHalf-(1), dWinWidth, 1,
		g.envColors.day.horizon, false)
}

func drawSun(g *Game, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.sunSP.image.Bounds().Dx())+sunX, 24)
	screen.DrawImage(g.sunSP.image, op)
}

func drawWorldGrad(g *Game, screen *ebiten.Image) {
	//Draw world grads (cached)
	if g.worldGradDirty {
		g.worldGradDirty = false
		g.updateWorldGrad()
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(dWinWidth, 1)
	screen.DrawImage(g.worldGradImg, op)
}

func drawSunReflect(g *Game, screen *ebiten.Image) {
	subRect := image.Rectangle{
		Min: image.Point{X: 0, Y: g.sunSP.blurred.Bounds().Dy() / 2.0},
		Max: image.Point{X: g.sunSP.blurred.Bounds().Dx(), Y: g.sunSP.blurred.Bounds().Dy()},
	}
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	sub := g.sunSP.blurred.SubImage(subRect)
	op.GeoM.Reset()
	xScale, yScale := 1.0, sunReflectHeight
	op.GeoM.Scale(xScale, yScale)
	op.GeoM.Translate((float64(sub.Bounds().Dx())/xScale)+sunX, dWinHeightHalf)
	op.Blend = ebiten.BlendLighter
	op.ColorScale.ScaleAlpha(sunReflectAlpha)
	screen.DrawImage(sub.(*ebiten.Image), op)

}

func drawAir(g *Game, screen *ebiten.Image) {

	for y, line := range g.airWaveLines {
		for _, wave := range line.waves {
			// Lower alpha for waves that are farther away
			alpha := uint8(math32.Min(30+(float32(y)*2.5), 255))
			waveColor := color.NRGBA{R: 255, G: 255, B: 255, A: alpha}

			//Expand wave.x 2x to screen, add boat pos.x, multiply by Y for parallax.
			timeOff := float64(time.Now().UnixMilli()) / (1 / (windSpeed / 300.0))
			preMod := float64(wave.x*2) + g.boatPos.X*(float64(y+1)*distParallax) + timeOff
			//Modulo to wrap around the screen
			modPos := math.Mod(preMod, dWinWidth)

			//Fix negative coordinates
			if modPos < 0 {
				modPos += dWinWidth
			}

			//Wave width based on distance
			width := float32(y) / 11.0
			width = max(width, 2)

			vector.DrawFilledRect(screen, float32(dWinWidth-modPos), float32(dWinHeightHalf-y), width, 1, waveColor, false)
		}
	}
}

func drawWaves(g *Game, screen *ebiten.Image) {
	for y, line := range g.wavesLines {
		for _, wave := range line.waves {
			// Lower alpha for waves that are farther away
			alpha := uint8(math32.Min(waveAlpha+(float32(y)*waveDistAlpha), 255))
			waveColor := color.NRGBA{R: 255, G: 255, B: 255, A: alpha}

			//Expand wave.x 2x to screen, add boat pos.x, multiply by Y for parallax.
			preMod := float64((wave.x)*2) + (g.boatPos.X)*(float64(y+1)*distParallax)
			//Modulo to wrap around the screen
			modPos := math.Mod(preMod, dWinWidth)

			//Fix negative coordinates
			if modPos < 0 {
				modPos += dWinWidth
			}

			//Wave width based on distance
			width := float32(y) / 11.0
			width = max(width, 2)

			vector.DrawFilledRect(screen, float32(dWinWidth-modPos), float32(dWinHeightHalf+y), width, 1, waveColor, false)
		}
	}
}

func (g *Game) makeWave() {
	if g.numWaves > 0 {
		for l, line := range g.wavesLines {
			for w, wave := range line.waves {
				if time.Since(wave.start) > wave.life {
					// Remove the element at index i
					g.wavesLines[l].waves = append(g.wavesLines[l].waves[:w], g.wavesLines[l].waves[w+1:]...)
					g.numWaves--
					g.wavesLines[l].count--
					break
				}
			}
		}
	}
	spawns := 0
	for spawns < spawnPerFrame && g.numWaves < maxWaves {
		y := int(logDistWave(rand.Float64()) * dWinHeightHalf)
		y = min(y, dWinHeightHalf-1)
		y = max(y, 0)

		var newWave = waveData{
			x:     rand.Intn(dWinWidth / 2),
			start: time.Now(),
			life:  time.Millisecond * time.Duration(minWaveLifeMS+(rand.Float64()*maxWaveLifeRandMS)),
		}

		g.wavesLines[y].waves = append(g.wavesLines[y].waves, newWave)
		g.wavesLines[y].count++
		g.numWaves++
		spawns++
	}
}

func (g *Game) makeAirWave() {
	if g.numAirWaves > 0 {
		for l, line := range g.airWaveLines {
			for w, wave := range line.waves {
				if time.Since(wave.start) > wave.life {
					// Remove the element at index i
					g.airWaveLines[l].waves = append(g.airWaveLines[l].waves[:w], g.airWaveLines[l].waves[w+1:]...)
					g.numAirWaves--
					g.airWaveLines[l].count--
					break
				}
			}
		}
	}
	spawns := 0
	for spawns < spawnPerFrame && g.numAirWaves < maxAirWaves {
		y := int(logDistAirWave(rand.Float64()) * dWinHeightHalf)
		y = min(y, dWinHeightHalf-1)
		y = max(y, 0)

		var newWave waveData

		newWave = waveData{
			x:     rand.Intn(dWinWidth / 2),
			start: time.Now(),
			life:  time.Millisecond * time.Duration(minAirWaveLifeMS+(rand.Float64()*maxAirWaveLifeRandMS)),
		}

		g.airWaveLines[y].waves = append(g.airWaveLines[y].waves, newWave)
		g.airWaveLines[y].count++
		g.numAirWaves++
		spawns++
	}

}

func logDistWave(uniform float64) float64 {
	biased := math.Abs(math.Log(float64(1 - uniform)))
	return math.Min(biased/persVal, 1.0)
}

func logDistAirWave(uniform float64) float64 {
	biased := math.Abs(math.Log(1 - uniform))
	return math.Min(biased/skyPersVal, 1.0)
}
