package main

import (
	"fmt"
	"math/rand"

	"github.com/aquilax/go-perlin"
)

var MapSeed int64 = 0

/* Init random seeds for the perlin noise layers */
func initNoise() {
	if MapSeed == 0 {
		MapSeed = rand.Int63()
	}

	for p := range noiseLayers {
		noiseLayers[p].randomSeed = MapSeed - noiseLayers[p].seedOffset
		noiseLayers[p].randomSource = rand.NewSource(noiseLayers[p].randomSeed)
		noiseLayers[p].perlinNoise = perlin.NewPerlinRandSource(float64(noiseLayers[p].alpha), float64(noiseLayers[p].beta), noiseLayers[p].n, noiseLayers[p].randomSource)
		fmt.Printf("init: %v\n", p)
	}
}

/* Get resource value at xy */
func noiseMap(x, y float32, p int) float32 {

	val := float32(noiseLayers[p].perlinNoise.Noise2D(
		float64(x/noiseLayers[p].scale),
		float64(y/noiseLayers[p].scale)))/float32(noiseLayers[p].contrast) + noiseLayers[p].brightness

	if val > noiseLayers[p].maxValue {
		return noiseLayers[p].maxValue
	} else if val < noiseLayers[p].minValue {
		return noiseLayers[p].minValue
	}

	return val
}

/* Resource layers */
const numResourceTypes = 7

var noiseLayers = [numResourceTypes]noiseLayerData{
	{name: "Ground",
		seedOffset: 5147,
		scale:      32,
		alpha:      2,
		beta:       2.0,
		n:          3,
		contrast:   2,
		brightness: 2,
		maxValue:   5,
		minValue:   -1,

		resourceMultiplier: 0,
		redMulti:           1,
		blueMulti:          1,
		greenMulti:         1,
	},

	/* Resources */
	{name: "Oil",
		seedOffset: 6812,
		scale:      256,
		alpha:      2,
		beta:       2.0,
		n:          3,

		contrast:   0.2,
		brightness: -2.2,
		maxValue:   5,
		minValue:   0,

		modGreen: true,

		resourceMultiplier: 1,
		greenMulti:         1,
	},
	{name: "Natural Gas",
		seedOffset: 240,
		scale:      128,
		alpha:      2,
		beta:       2.0,
		n:          3,

		contrast:   0.3,
		brightness: -1.5,
		maxValue:   5,
		minValue:   0,

		modRed:   true,
		modGreen: true,

		resourceMultiplier: 1,
		redMulti:           0.80,
		greenMulti:         1,
	},
	{name: "Coal",
		seedOffset: 7266,
		scale:      256,
		alpha:      2,
		beta:       2.0,
		n:          3,

		contrast:   0.3,
		brightness: -1.0,
		maxValue:   5,
		minValue:   0,

		modRed: true,

		redMulti: 1,
	},
	{name: "Iron Ore",
		seedOffset: 5324,
		scale:      256,
		alpha:      2,
		beta:       2.0,
		n:          3,

		contrast:   0.3,
		brightness: -1.0,
		maxValue:   5,
		minValue:   0,

		modRed:   true,
		modGreen: true,

		resourceMultiplier: 1,
		redMulti:           1,
		greenMulti:         0.5,
	},
	{name: "Copper Ore",
		seedOffset: 1544,
		scale:      256,
		alpha:      2,
		beta:       2.0,
		n:          3,

		contrast:   0.3,
		brightness: -1.0,
		maxValue:   5,
		minValue:   0,

		modGreen: true,
		modBlue:  true,

		resourceMultiplier: 1,
		greenMulti:         1,
		blueMulti:          1,
	},
	{name: "Stone Ore",
		seedOffset: 8175,
		scale:      256,
		alpha:      2,
		beta:       2.0,
		n:          3,

		contrast:   0.4,
		brightness: -0.75,
		maxValue:   5,
		minValue:   0,

		modRed:   true,
		modGreen: true,
		modBlue:  true,

		resourceMultiplier: 1,
		redMulti:           0.5,
		greenMulti:         0.5,
		blueMulti:          0.5,
	},
}

/* Perlin noise data */
type noiseLayerData struct {
	name       string
	seedOffset int64

	/* Perlin values */
	scale      float32
	alpha      float32
	beta       float32
	n          int32
	contrast   float32
	brightness float32
	maxValue   float32
	minValue   float32

	/* Output adjustments */
	modRed   bool
	modGreen bool
	modBlue  bool

	resourceMultiplier float64

	redMulti   float32
	greenMulti float32
	blueMulti  float32

	randomSource rand.Source
	randomSeed   int64
	perlinNoise  *perlin.Perlin
}
