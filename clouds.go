package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	cloudExpire  = 60 * 5
	checkRecycle = 60 * 10
	cloudChunkX  = dWinHeightHalf
	cloudChunkY  = dWinHeightHalf
)

type cloudData struct {
	id             int
	image, blurImg *ebiten.Image
	lastUsed       uint64
}

var (
	cloudChunks    map[int]*cloudData
	recycledChunks []*cloudData
	cloudsDirty    bool
	chunkIDTop     int
)

func drawCloudsNew(g *Game, screen *ebiten.Image) {

	draws := 0

	x := 0
	numChunks := dWinWidth/cloudChunkX + 1
	if g.boatPos.X <= 0 {
		x = -1
	}
	for ; x < numChunks+1; x++ {

		pos := int((g.boatPos.X) * float64(cloudY/dWinWidth))
		chunkNum := (pos / cloudChunkX) - x
		cloud := cloudChunks[chunkNum]

		if cloud == nil {
			var newCloud *cloudData

			//New or recycle chunk
			nc := len(recycledChunks)
			if nc == 0 {
				newCloud = &cloudData{id: chunkIDTop}
				chunkIDTop++
				newCloud.image = ebiten.NewImage(cloudChunkX, cloudChunkY)
				doLog(true, true, "Created new cloud chunk: %v (%v,%v)",
					newCloud.id, cloudChunkX, cloudChunkY)
			} else {
				newCloud = recycledChunks[0]
				doLog(true, true, "Reused cloud chunk: %v (%v,%v)",
					newCloud.id, cloudChunkX, cloudChunkY)

				if nc > 1 {
					recycledChunks = recycledChunks[1:]
				} else {
					recycledChunks = []*cloudData{}
				}
			}

			renderCloudChunk(chunkNum, newCloud)
			cloudChunks[chunkNum] = newCloud
			cloud = newCloud

		} else if cloudsDirty {
			//Rerender everything
			cloudsDirty = false
			for cn, fCloud := range cloudChunks {
				renderCloudChunk(cn, fCloud)
			}
		}
		//Just show the cached chunk
		cloud.lastUsed = frameNumber

		op := &ebiten.DrawImageOptions{}
		xtrans := float64(((x - numChunks) * cloudChunkX) + pos%cloudChunkX)
		op.GeoM.Translate(-xtrans, 0)
		screen.DrawImage(cloud.image, op)

		draws++

		//Reflection for water
		op.GeoM.Scale(1, -cloudReflectStretch)
		op.GeoM.Translate(0, dWinHeight*(cloudReflectStretch/1.25))
		op.ColorScale.ScaleAlpha(cloudReflectAlpha)
		screen.DrawImage(cloud.image, op)
	}

	//Get rid of old cloud chunks
	if frameNumber%checkRecycle == 0 {
		for xc, xCloud := range cloudChunks {
			if frameNumber-xCloud.lastUsed > cloudExpire {
				doLog(true, true, "Recycled chunk: %v", xCloud.id)
				recycledChunks = append(recycledChunks, xCloud)
				delete(cloudChunks, xc)
			}
		}
	}
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
	cloud.image.WritePixels(cBuf)
	if *debugMode {
		buf := fmt.Sprintf("%v: %v", chunkNum, cloud.id)
		ebitenutil.DebugPrint(cloud.image, buf)
	}
}
