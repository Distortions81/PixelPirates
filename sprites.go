package main

import (
	"fmt"
	"image"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var spriteList map[string]*spriteItem = map[string]*spriteItem{
	"title":      {Path: "title/"},
	"clickstart": {Path: "title/"},

	"sun":     {Path: "world/"},
	"island1": {Path: "world/", doReflect: true},

	"default-player": {Path: "characters/"},

	"boat2": {Path: "boats/"},

	"testScene1":    {Path: "islands/"},
	"island-scene1": {Path: "islands/"},
}

var (
	defPlayerSP, testScene1SP, boat2SP, sunSP, island1SP, titleSP, clickStartSP *spriteItem
)

func initSprites() {
	titleSP = spriteList["title"]
	clickStartSP = spriteList["clickstart"]

	sunSP = spriteList["sun"]
	boat2SP = spriteList["boat2"]
	island1SP = spriteList["island1"]

	testScene1SP = spriteList["island-scene1"]
	defPlayerSP = spriteList["default-player"]

}

func loadSprites() {

	for name, sprite := range spriteList {
		image, blurImg, err := loadSprite(sprite.Path+name, false, sprite.doReflect)
		if err == nil {
			doLog(true, "loading sprite '"+name+"'")
			spriteList[name].image = image
			spriteList[name].blurred = blurImg
			spriteList[name].Name = name
		} else {
			doLog(true, "loading sprite '"+name+"' failed.")
		}

		aniData, err := loadAnimationData(sprite.Path + name)
		if err == nil {
			spriteList[name].animation = aniData
		}
	}
}

type spriteItem struct {
	Name, Path string

	image     *ebiten.Image
	blurred   *ebiten.Image
	doReflect bool

	animation *animationData
	pingDir   bool
}

func getAniFrame(frame int64, ani *spriteItem, offset int) *ebiten.Image {
	numFrames := int64(len(ani.animation.Frames))
	if frame < 0 || frame >= numFrames {
		fmt.Printf("%v: invalid frame number: %v\n", ani.Name, frame)
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
		fmt.Printf("** %v: %v: NO FRAMES: %v -> %v\n\n", tag, ani.Name, frameRange.start, frameRange.end)
		return nil
	}
	frameNum := (time % numFrames) + int64(frameRange.start)
	return getAniFrame(frameNum, ani, offset)
}

func autoAnimatePingPong(ani *spriteItem, offset int, tag string) *ebiten.Image {

	frameRange := ani.animation.animations[tag]
	numFrames := int64(frameRange.end - frameRange.start)
	if numFrames <= 0 {
		fmt.Printf("** %v: NO FRAMES: %v -> %v\n\n", ani.Name, frameRange.start, frameRange.end)
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
