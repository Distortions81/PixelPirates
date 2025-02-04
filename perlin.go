package main

import (
	"math/rand"

	"github.com/aquilax/go-perlin"
)

var MapSeed int64 = 0

func initNoise() {
	MapSeed = rand.Int63()

	for p := range noiseLayers {
		noiseLayers[p].randomSeed = MapSeed - noiseLayers[p].seedOffset
		noiseLayers[p].randomSource = rand.NewSource(noiseLayers[p].randomSeed)
		noiseLayers[p].perlinNoise = perlin.NewPerlinRandSource(float64(noiseLayers[p].alpha), float64(noiseLayers[p].beta), noiseLayers[p].n, noiseLayers[p].randomSource)
		doLog(true, true, "init noise layer: %v: %v", p, MapSeed)
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
const numResourceTypes = 1

var noiseLayers = [numResourceTypes]noiseLayerData{
	{name: "Clouds1",
		seedOffset: 5147,
		scale:      64,
		alpha:      2,
		beta:       2.0,
		n:          4,
		contrast:   0.8,
		brightness: 0,
		maxValue:   5,
		minValue:   0,

		resourceMultiplier: 1,
		redMulti:           1,
		blueMulti:          1,
		greenMulti:         1,
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
