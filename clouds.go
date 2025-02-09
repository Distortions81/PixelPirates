package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	cloudExpire  = 60 * 5.0
	checkRecycle = 60 * 10.0
	cloudChunkX  = dWinHeightHalf
	cloudChunkY  = dWinHeightHalf
)

type cloudData struct {
	id             int
	image, blurImg *ebiten.Image
	lastUsed       uint64
}

func drawCloudsNew(g *Game, screen *ebiten.Image) {

	draws := 0

	var numChunks int = int(math.Round((dWinWidth / cloudChunkX) + 1))
	for x := 0; x < numChunks+1; x++ {

		pos := int((g.boatPos.X) * float64(cloudY/dWinWidth))
		chunkNum := (pos / cloudChunkX) - x
		cloud := g.cloudChunks[chunkNum]

		if cloud == nil {
			var newCloud *cloudData

			//New or recycle chunk
			nc := len(g.recycledChunks)
			if nc == 0 {
				newCloud = &cloudData{id: g.chunkIDTop}
				g.chunkIDTop++
				newCloud.image = ebiten.NewImage(cloudChunkX, cloudChunkY)
				doLog(true, true, "Created new cloud chunk: %v (%v,%v)",
					newCloud.id, cloudChunkX, cloudChunkY)
			} else {
				newCloud = g.recycledChunks[0]
				doLog(true, true, "Reused cloud chunk: %v (%v,%v)",
					newCloud.id, cloudChunkX, cloudChunkY)

				if nc > 1 {
					g.recycledChunks = g.recycledChunks[1:]
				} else {
					g.recycledChunks = []*cloudData{}
				}
			}

			renderCloudChunk(chunkNum, newCloud)
			g.cloudChunks[chunkNum] = newCloud
			cloud = newCloud

		} else if g.cloudsDirty {
			//Rerender everything
			g.cloudsDirty = false
			for cn, fCloud := range g.cloudChunks {
				renderCloudChunk(cn, fCloud)
			}

			doLog(true, true, "Rerendered all clouds.")
		}
		//Just show the cached chunk
		cloud.lastUsed = g.frameNumber

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
	if g.frameNumber%checkRecycle == 0 {
		for xc, xCloud := range g.cloudChunks {
			if g.frameNumber-xCloud.lastUsed > cloudExpire {
				doLog(true, true, "Recycled chunk: %v", xCloud.id)
				g.recycledChunks = append(g.recycledChunks, xCloud)
				delete(g.cloudChunks, xc)
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
	/*
		if *debugMode {
			buf := fmt.Sprintf("%v: %v", chunkNum, cloud.id)
			ebitenutil.DebugPrint(cloud.image, buf)
		}
	*/
}
