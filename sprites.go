package main

import (
	"image"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// For some currently hardcoded sprites
var spriteList map[string]*spriteItem = map[string]*spriteItem{
	//Title
	"title":      {Path: "title/"},
	"clickstart": {Path: "title/"},

	//Game & Title
	"sun":     {Path: "world/"},
	"island1": {Path: "world/", doReflect: true, onDemand: true},
	"boat2":   {Path: "boats/"},

	//Islands
	"island-scene1":  {Path: "islands/", onDemand: true},
	"default-player": {Path: "characters/", onDemand: true},
}

func initSprites(g *Game) {
	g.titleSP = spriteList["title"]
	g.clickStartSP = spriteList["clickstart"]

	g.sunSP = spriteList["sun"]
	g.boat2SP = spriteList["boat2"]
}

func loadSprites() {

	for name, sprite := range spriteList {
		loadSprite(name, sprite, false)
	}
}

func loadSprite(name string, sprite *spriteItem, demanded bool) {
	var image, blurImg *ebiten.Image
	var err error

	if !sprite.onDemand || demanded {
		image, blurImg, err = loadImage(sprite.Path+name, false, sprite.doReflect)
		doLog(true, true, "loading sprite '"+name+"'")
	}
	if err == nil {
		spriteList[name].image = image
		spriteList[name].blurred = blurImg
		spriteList[name].Name = name
	} else {
		doLog(true, false, "loading sprite '"+name+"' failed.")
	}

	aniData, err := loadAnimationData(sprite.Path + name)
	if err == nil {
		spriteList[name].animation = aniData
	}
}

type spriteItem struct {
	Name, Path string

	image, blurred *ebiten.Image
	doReflect      bool
	onDemand       bool

	animation *animationData
}

func getAniFrame(frame int64, ani *spriteItem, offset int) *ebiten.Image {
	numFrames := int64(len(ani.animation.Frames))
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
	return getAniFrame(frameNum, ani, offset)
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

	return getAniFrame(frameNum+int64(frameRange.start), ani, offset)
}
