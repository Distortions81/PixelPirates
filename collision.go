package main

import (
	"strings"
)

func checkPixelCollision(g *Game) bool {
	visiting := g.visiting
	img := getAniFrame(0, g.defPlayerSP, 0)
	if img == nil || visiting == nil {
		return false
	}
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img := img.At(x, y)
			_, _, _, alpha := img.RGBA()
			if alpha < 254 {
				continue
			}
			if g.visiting.collision[iPoint{X: dWinWidthHalf - (width / 2) + int(g.playPos.X) + x, Y: dWinHeightHalf - (height / 2) + int(g.playPos.Y) + y}] {
				return true
			}
		}
	}
	return false
}

func findSpawns() {
	for i, island := range islands {
		for _, item := range island.objects {
			if strings.Contains(item.Name, "spawn") {
				name := item.animation.sortedFrames[0]
				frame := item.animation.Frames[name]
				newSpawn := fPoint{X: float64(frame.SpriteSourceSize.X), Y: float64(frame.SpriteSourceSize.Y)}
				islands[i].spawn = newSpawn
				doLog(true, false, "Found Spawn for: %v at %v,%v", island.name, newSpawn.X, newSpawn.Y)
				break
			}
		}
	}
}

func makeCollisionMaps() {
	for i, island := range islands {
		for _, item := range island.objects {
			if strings.Contains(item.Name, "collision") {
				islands[i].collision = map[iPoint]bool{}
				count := 0
				for x := 0; x < item.image.Bounds().Dx(); x++ {
					for y := 0; y < item.image.Bounds().Dy(); y++ {
						pixel := item.image.At(x, y)
						_, _, _, alpha := pixel.RGBA()
						if alpha > 128 {
							islands[i].collision[iPoint{X: x, Y: y}] = true
							count++
						}
					}
				}
				doLog(true, false, "Got collision map for island: %v (%v pixels)", island.name, count)
			}
		}
	}
}
