package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	cloudExpire = time.Second * 30
	cloudChunkX = dWinHeightHalf
	cloudChunkY = dWinHeightHalf
)

type cloudData struct {
	image, blurImg *ebiten.Image
	lastUsed       time.Time
}

var cloudChunks map[int]*cloudData
var cloudsDirty bool

func drawCloudsNew(g *Game, screen *ebiten.Image) {
	numChunks := dWinWidth/cloudChunkX + 2
	dbuf := ""
	for x := 0; x < numChunks; x++ {
		pos := int((g.boatPos.X) * float64(cloudY/dWinWidth))
		chunkNum := (pos / cloudChunkX) - x
		cloud := cloudChunks[chunkNum]

		if cloud == nil {
			newCloud := &cloudData{}
			newCloud.image = ebiten.NewImage(cloudChunkX, cloudChunkY)
			fmt.Printf("Created new cloud chunk: %v (%v,%v)\n",
				chunkNum, cloudChunkX, cloudChunkY)

			renderCloudChunk(chunkNum, newCloud)
			cloudChunks[chunkNum] = newCloud
			cloud = newCloud

		} else if cloudsDirty {
			//Rerender everything
			cloudsDirty = false
			for _, fCloud := range cloudChunks {
				renderCloudChunk(chunkNum, fCloud)
			}
		}

		if cloud == nil {
			fmt.Printf("Fuck: %v\n", chunkNum)
			return
		}
		//Just show the cached chunk
		cloud.lastUsed = time.Now()

		op := &ebiten.DrawImageOptions{}
		xtrans := float64(((x - 4) * cloudChunkX) + pos%cloudChunkX)
		op.GeoM.Translate(-xtrans, 0)
		screen.DrawImage(cloud.image, op)
		dbuf = dbuf + fmt.Sprintf("%v, ", xtrans)
	}
	//fmt.Println(dbuf)

	//Get rid of old cloud chunks
	/*
		for xc, xCloud := range cloudChunks {
			if time.Since(xCloud.lastUsed) > cloudExpire {
				fmt.Printf("Deleted chunk: %v\n", xc)
				delete(cloudChunks, xc)
			}
		}
	*/
}

func drawCloudsReflectNew(screen *ebiten.Image) {

}

func renderCloudChunk(chunkNum int, cloud *cloudData) {
	var cBuf []byte
	offset := float32(chunkNum * cloudChunkX)
	for y := 0; y < cloudChunkY; y++ {
		for x := 0; x < cloudChunkX; x++ {
			v := noiseMap(float32(x)+offset, float32((y-10)*4.0), 0)
			vi := byte(v / 5 * 255)
			cBuf = append(cBuf, []byte{vi, vi, vi, vi}...)
		}
	}
	cloud.image.Clear()
	cloud.image.WritePixels(cBuf)
	buf := fmt.Sprintf("%v", chunkNum)
	ebitenutil.DebugPrint(cloud.image, buf)

	//reflection
	/*
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1.0/cloudBlurAmountX, 1.0/cloudBlurAmountY)
		op.Filter = ebiten.FilterLinear
		cloud.blurImg.Clear()
		cloud.blurImg.DrawImage(cloudbuf, op)
	*/

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
