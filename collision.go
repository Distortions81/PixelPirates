package main

import (
	"math"
)

func calculateDistance(origin, target iPoint) float64 {
	// Calculate differences as float64 to ensure accuracy
	dx := float64(target.X - origin.X)
	dy := float64(target.Y - origin.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

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

func checkPixelCollision(prev, new fPoint, g *Game) fPoint {

	//Check all pixels in source to destination
	points := BresenhamLine(g.oldPlayPos.ToInt(), new.ToInt())
	points = SortLinePoints(points, g.oldPlayPos.ToInt(), new.ToInt())

	//Only search pixels around the feet that collide
	for _, p := range points {
		for _, pixel := range g.defCollision.collision {
			if g.inIsland.collision[iPoint{X: dWinWidthHalf - (pWidth / 2) + int(p.X) + pixel.X, Y: dWinHeightHalf - (pHeight / 2) + int(p.Y) + pixel.Y}] {
				return fPoint{X: float64(prev.X), Y: float64(prev.Y)}
			}
		}
		prev = p.ToFloat()
	}
	return fPoint{}
}

func makeCollisionMaps(g *Game) {
	if g.inIsland == nil {
		return
	}

	if len(g.inIsland.collision) > 0 {
		return
	}

	edges := g.inIsland.spriteSheet.animation.layers["edges"]
	g.inIsland.collision = map[iPoint]bool{}
	count := 0

	for x := 0; x < edges.SourceSize.W; x++ {
		for y := 0; y < edges.SourceSize.H; y++ {
			sx, sy :=
				x+edges.Frame.X-edges.SpriteSourceSize.X,
				y+edges.Frame.Y-edges.SpriteSourceSize.Y
			//Check for trimmed area
			if sx < edges.Frame.X || sy < edges.Frame.Y {
				continue
			}
			//Check for trimmed area
			if sx > edges.Frame.X+edges.Frame.W || sy > edges.Frame.Y+edges.Frame.H {
				continue
			}
			pixel := g.inIsland.spriteSheet.image.At(sx, sy)
			_, _, _, alpha := pixel.RGBA()
			if alpha > 254 {
				g.inIsland.collision[iPoint{X: x, Y: y}] = true
				count++
			}
		}
	}
	doLog(true, false, "Parsed collision map for island: %v (%v points)", g.inIsland.name, count)
}
