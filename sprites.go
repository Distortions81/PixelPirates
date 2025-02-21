package main

import (
	"image"
	"math"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var pWidth, pHeight int

// For some currently hardcoded sprites
// Sprites with an animation json auto load as unmanged
var spriteList map[string]*spriteItem = map[string]*spriteItem{
	//Title
	"title":      {Path: "title/", onDemand: true},
	"clickstart": {Path: "title/", onDemand: true},

	//Game & Title
	"sun":     {Path: "world/"},
	"island1": {Path: "world/", doReflect: true, onDemand: true},
	"boat2":   {Path: "boats/"},

	"default-player":           {Path: "characters/", onDemand: true},
	"default-player-collision": {Path: "characters/", onDemand: true},
}

func initSprites(g *Game) {
	g.titleSP = spriteList["title"]
	g.clickStartSP = spriteList["clickstart"]

	g.defCollision = spriteList["default-player-collision"]
	g.defPlayerSP = spriteList["default-player"]

	g.sunSP = spriteList["sun"]
	g.boat2SP = spriteList["boat2"]

}

func loadSprites() {
	for name, sprite := range spriteList {
		path := dataDir + spritesDir + sprite.Path + name
		loadSprite(path, sprite, false)
		spriteList[name].Fullpath = path
	}
}

func loadSprite(name string, sprite *spriteItem, demanded bool) {
	var image, blurImg *ebiten.Image
	var err error

	if sprite == nil {
		doLog(true, false, "loadSprite: nil spriteItem: %v", name)
		return
	}

	if !sprite.onDemand || demanded {
		if sprite.image != nil {
			return
		}
		aniData, err := loadAnimationData(name)
		sprite.animation = aniData

		image, blurImg, err = loadImage(name, sprite.unmanged, sprite.doReflect)
		if err != nil {
			doLog(true, false, "loadImage Failed: %v", err)
			return
		}
		doLog(true, true, "loadSprite '"+name+"'")
	}
	if err == nil {
		sprite.image = image
		sprite.blurred = blurImg
	}
}

type spriteItem struct {
	Name, Path, Fullpath string

	image, blurred *ebiten.Image
	doReflect      bool
	onDemand       bool
	unmanged       bool
	collision      []iPoint

	animation *animationData
}

func getFrameNumber(frame int64, ani *spriteItem, offset int) *ebiten.Image {
	numFrames := ani.animation.numFrames
	if frame < 0 || frame >= numFrames {
		doLog(true, false, "%v: invalid frame number: %v", ani.Name, frame)
		if frame >= numFrames {
			frame = numFrames - 1
		} else if frame < 0 {
			frame = 0
		}
	}

	//Frame offset
	if offset != 0 {
		frame = int64(math.Mod(float64(frame+int64(offset)), float64(numFrames)))
	}

	fKey := ani.animation.sortedFrames[frame]
	fRect := ani.animation.Frames[fKey].Frame
	rect := image.Rectangle{Min: image.Point{X: fRect.X, Y: fRect.Y}, Max: image.Point{X: fRect.X + fRect.W, Y: fRect.Y + fRect.H}}
	subFrame := ani.image.SubImage(rect).(*ebiten.Image)
	return subFrame
}

func getLayerFromName(layerName string, ani *spriteItem) *ebiten.Image {
	layer := ani.animation.layers[strings.ToLower(layerName)]
	if layer == nil {
		doLog(true, false, "getLayerFromName: layer: %v not found.", layerName)
		return nil
	}
	layerRect := layer.Frame
	rect := image.Rectangle{Min: image.Point{
		X: layerRect.X, Y: layerRect.Y},
		Max: image.Point{X: layerRect.X + layerRect.W, Y: layerRect.Y + layerRect.H}}
	subFrame := ani.image.SubImage(rect).(*ebiten.Image)
	return subFrame
}

func getLayer(layer *aniFrame, ani *spriteItem) *ebiten.Image {
	layerRect := layer.Frame
	rect := image.Rectangle{Min: image.Point{
		X: layerRect.X, Y: layerRect.Y},
		Max: image.Point{X: layerRect.X + layerRect.W, Y: layerRect.Y + layerRect.H}}
	subFrame := ani.image.SubImage(rect).(*ebiten.Image)
	return subFrame
}

func autoAnimate(ani *spriteItem, offset int, tag string) *ebiten.Image {
	frameRange := ani.animation.animations[tag]
	numFrames := int64(frameRange.end-frameRange.start) + 1

	firstFrame := ani.animation.sortedFrames[frameRange.start]
	speed := ani.animation.Frames[firstFrame].Duration
	time := time.Now().UnixMilli() / int64(speed)

	if numFrames <= 0 {
		doLog(true, false, "** %v: %v: NO FRAMES: %v -> %v", tag, ani.Name, frameRange.start, frameRange.end)
		return nil
	}
	frameNum := (time % numFrames) + int64(frameRange.start)
	return getFrameNumber(frameNum, ani, offset)
}

func autoAnimatePingPong(ani *spriteItem, offset int, tag string) *ebiten.Image {

	frameRange := ani.animation.animations[tag]
	numFrames := int64(frameRange.end - frameRange.start)
	if numFrames <= 0 {
		doLog(true, false, "** %v: NO FRAMES: %v -> %v", ani.Name, frameRange.start, frameRange.end)
		return nil
	}

	period := 2*numFrames - 2
	firstFrame := ani.animation.sortedFrames[0]
	speed := ani.animation.Frames[firstFrame].Duration

	time := time.Now().UnixMilli() / int64(speed)
	framePosition := time % period

	var frameNum int64
	if framePosition < int64(numFrames) {
		// Forward direction
		frameNum = int64(framePosition)
	} else {
		// Backward direction
		frameNum = int64(period - framePosition)
	}

	return getFrameNumber(frameNum+int64(frameRange.start), ani, offset)
}
