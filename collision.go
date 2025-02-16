package main

import (
	"math"
	"strings"
)

func BresenhamLine(a, b iPoint) []iPoint {
	var points []iPoint

	dx := int(math.Abs(float64(b.X - a.X)))
	sx := 1
	if a.X > b.X {
		sx = -1
	}

	dy := -int(math.Abs(float64(b.Y - a.Y)))
	sy := 1
	if a.Y > b.Y {
		sy = -1
	}

	err := dx + dy
	x, y := a.X, a.Y

	for {
		// Add the current point.
		points = append(points, iPoint{X: x, Y: y})

		// Break if we’ve reached the end point.
		if x == b.X && y == b.Y {
			break
		}

		// e2 is twice the 'error' term.
		e2 := 2 * err

		// Move in x if the error is larger in that dimension.
		if e2 >= dy {
			err += dy
			x += sx
		}
		// Move in y if the error is larger in that dimension.
		if e2 <= dx {
			err += dx
			y += sy
		}
	}

	return points
}

func checkPixelCollision(g *Game) fPoint {

	var prev iPoint = g.oldPlayPos.QuantizePoint()
	//Check all pixels in source to destination
	points := BresenhamLine(g.oldPlayPos.QuantizePoint(), g.playPos.QuantizePoint())
	points = SortLinePoints(points, g.oldPlayPos.QuantizePoint(), g.playPos.QuantizePoint())

	//Only search pixels around the feet that collide
	for _, p := range points {
		for _, pixel := range g.defCollision.collision {
			if g.visiting.collision[iPoint{X: dWinWidthHalf - (pWidth / 2) + int(p.X) + pixel.X, Y: dWinHeightHalf - (pHeight / 2) + int(p.Y) + pixel.Y}] {
				return fPoint{X: float64(prev.X), Y: float64(prev.Y)}
			}
		}
		prev = p
	}
	return fPoint{}
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
