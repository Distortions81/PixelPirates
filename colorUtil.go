package main

import (
	"image/color"
	"math"
	"strconv"
	"strings"
	"time"
)

var COLOR_WHITE = color.NRGBA{R: 255, G: 255, B: 255}

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

// lerpHSV linearly interpolates between two HSV values
func lerpHSV(h1, h2 hsv, t float64) hsv {
	h := h1.H + (h2.H-h1.H)*t
	// Ensure hue wraps around 360 degrees
	if h < 0 {
		h += 360
	} else if h > 360 {
		h -= 360
	}
	return hsv{
		H: h,
		S: h1.S + (h2.S-h1.S)*t,
		V: h1.V + (h2.V-h1.V)*t,
	}
}

// GetFadeColorHSV returns the current color based on wall time, interpolating in HSV
func getFadeColor(start, end color.RGBA, duration time.Duration) color.RGBA {
	// Convert RGB colors to HSV
	startHSV := rgbToHSV(start)
	endHSV := rgbToHSV(end)

	// Get the current time
	currentTime := time.Now()

	// Calculate the number of seconds since the Unix epoch
	secondsSinceEpoch := currentTime.Unix()

	// Calculate how far we are within the fade cycle
	t := float64(secondsSinceEpoch%int64(duration.Seconds())) / duration.Seconds()

	// Make the fade go back and forth
	if t > 0.5 {
		t = 1 - t
	}
	t = 2 * t // Scale for full forward and back transition

	// Interpolate in HSV space
	currentHSV := lerpHSV(startHSV, endHSV, t)

	// Convert the interpolated HSV back to RGB (uint8)
	outColor := hsvToRGB(currentHSV)
	outColor.A = 255
	return outColor
}

// rgbToHSV converts an RGB color to HSV
func rgbToHSV(c color.RGBA) hsv {
	r := float64(c.R) / 255
	g := float64(c.G) / 255
	b := float64(c.B) / 255

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	delta := max - min

	h := 0.0
	if delta != 0 {
		switch max {
		case r:
			h = math.Mod((g-b)/delta, 6.0)
		case g:
			h = (b-r)/delta + 2
		case b:
			h = (r-g)/delta + 4
		}
		h *= 60
		if h < 0 {
			h += 360
		}
	}

	s := 0.0
	if max != 0 {
		s = delta / max
	}

	v := max

	return hsv{H: h, S: s, V: v}
}

// hsvToRGB converts an HSV color back to RGB (uint8)
func hsvToRGB(hsv hsv) color.RGBA {
	h := hsv.H
	s := hsv.S
	v := hsv.V

	c := v * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := v - c

	var r, g, b float64
	switch {
	case h >= 0 && h < 60:
		r, g, b = c, x, 0
	case h >= 60 && h < 120:
		r, g, b = x, c, 0
	case h >= 120 && h < 180:
		r, g, b = 0, c, x
	case h >= 180 && h < 240:
		r, g, b = 0, x, c
	case h >= 240 && h < 300:
		r, g, b = x, 0, c
	case h >= 300 && h < 360:
		r, g, b = c, 0, x
	}

	return color.RGBA{
		R: uint8((r + m) * 255),
		G: uint8((g + m) * 255),
		B: uint8((b + m) * 255),
		A: 255,
	}
}
