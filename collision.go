package main

import (
	"image"
	"image/color"
)

func getPixelData(img image.Image) [][]color.Color {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	pixels := make([][]color.Color, height)
	for y := 0; y < height; y++ {
		pixels[y] = make([]color.Color, width)
		for x := 0; x < width; x++ {
			pixels[y][x] = img.At(x, y)
		}
	}
	return pixels
}

func checkPixelCollision(img1Pixels [][]color.Color, img2Pixels [][]color.Color, x1, y1, x2, y2 int) bool {
	for y := 0; y < len(img1Pixels); y++ {
		for x := 0; x < len(img1Pixels[0]); x++ {
			if x+x1 >= 0 && x+x1 < len(img2Pixels[0]) && y+y1 >= 0 && y+y1 < len(img2Pixels) {
				_, _, _, a1 := img1Pixels[y][x].RGBA()
				_, _, _, a2 := img2Pixels[y+y1][x+x1].RGBA()
				if a1 > 0 && a2 > 0 {
					return true
				}
			}
		}
	}
	return false
}
