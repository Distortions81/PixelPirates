package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	cloudExpire = time.Second * 30
	cloudChunkX = dWinHeightHalf
	cloudChunkY = dWinHeightHalf
)

type cloudData struct {
	image, blurImg *ebiten.Image
	wasInit        bool
	lastUsed       time.Time
}

var cloudChunks map[int]*cloudData
var cloudsDirty bool

func drawCloudsNew(g *Game, screen *ebiten.Image) {
	pos := int(g.boatPos.X * float64(cloudY/dWinWidth))
	chunkNum := pos / cloudChunkX
	cloud := cloudChunks[chunkNum]

	if cloud == nil || !cloud.wasInit {
		newCloud := &cloudData{}
		cloudChunks[chunkNum] = newCloud
		newCloud.image = ebiten.NewImage(cloudChunkX, cloudChunkY)
		newCloud.blurImg = ebiten.NewImage(
			cloudChunkX/cloudBlurAmountX,
			cloudChunkY/cloudBlurAmountY)

		newCloud.wasInit = true
		fmt.Printf("Created new cloud chunk: %v (%v,%v)\n",
			chunkNum, cloudChunkX, cloudChunkY)

		renderCloudChunk(chunkNum*cloudChunkX, newCloud)
	} else if cloudsDirty {
		//Rerender everything
		cloudsDirty = false
		for c, fCloud := range cloudChunks {
			if fCloud.wasInit {
				renderCloudChunk(c, fCloud)
			}
		}
	}
	//Just show the cached chunk
	cloudChunks[chunkNum].lastUsed = time.Now()
	screen.DrawImage(cloudChunks[chunkNum].image, nil)

	//Get rid of old cloud chunks
	for xc, xCloud := range cloudChunks {
		if time.Since(xCloud.lastUsed) > cloudExpire {
			fmt.Printf("Deleted chunk: %v\n", xc)
			delete(cloudChunks, xc)
		}
	}
}

func drawCloudsReflectNew(screen *ebiten.Image) {

}

func renderCloudChunk(chunkNum int, cloud *cloudData) {
	var cBuf []byte
	for y := 0; y < cloudChunkY; y++ {
		for x := 0; x < cloudChunkX; x++ {
			v := noiseMap(float32(x)+float32(chunkNum), float32((y-10)*4.0), 0)
			vi := byte(v / 5 * 255)
			cBuf = append(cBuf, []byte{vi, vi, vi, vi}...)
		}
	}
	cloud.image.Clear()
	cloud.image.WritePixels(cBuf)

	//reflection
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1.0/cloudBlurAmountX, 1.0/cloudBlurAmountY)
	op.Filter = ebiten.FilterLinear
	cloud.blurImg.Clear()
	cloud.blurImg.DrawImage(cloudbuf, op)

}

func drawClouds(g *Game, screen *ebiten.Image) {
	xpos := g.boatPos.X * float64(cloudY/dWinWidth)
	if int(xpos) != lastCloudPos {
		lastCloudPos = int(xpos)
		var cBuf []byte
		for y := 0; y < dWinHeightHalf; y++ {
			for x := 0; x < dWinWidth; x++ {
				v := noiseMap(float32(x)+float32(xpos), float32((y-10)*4.0), 0)
				vi := byte(v / 5 * 255)
				cBuf = append(cBuf, []byte{vi, vi, vi, vi}...)
			}
		}
		cloudbuf.WritePixels(cBuf)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1.0/cloudBlurAmountX, 1.0/cloudBlurAmountY)
		op.Filter = ebiten.FilterLinear
		cloudblur.Clear()
		cloudblur.DrawImage(cloudbuf, op)
	}
	drawCloudsReflect(screen)
}

func drawCloudsReflect(screen *ebiten.Image) {
	//Cloud reflection
	screen.DrawImage(cloudbuf, nil)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(cloudBlurAmountX, -cloudBlurAmountY/cloudBlurStrech)
	op.GeoM.Translate(0, dWinHeight*cloudBlurAmountY)
	op.ColorScale.ScaleAlpha(cloudReflectAlpha)
	//op.Blend = ebiten.BlendLighter
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(cloudblur, op)
}
