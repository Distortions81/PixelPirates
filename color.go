package main

import (
	"image/color"
	"strconv"
	"strings"
)

func hexToRGB(hex string) color.RGBA {
	// Remove the '#' if present
	hex = strings.TrimPrefix(hex, "#")

	// Parse the hex string to its R, G, B components
	if len(hex) != 6 {
		return color.RGBA{}
	}

	r, err := strconv.ParseInt(hex[0:2], 16, 32)
	if err != nil {
		return color.RGBA{}
	}

	g, err := strconv.ParseInt(hex[2:4], 16, 32)
	if err != nil {
		return color.RGBA{}
	}

	b, err := strconv.ParseInt(hex[4:6], 16, 32)
	if err != nil {
		return color.RGBA{}
	}

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}
