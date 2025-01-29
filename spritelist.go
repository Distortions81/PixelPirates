package main

import (
	"fmt"
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var spriteList map[string]*spriteItem = map[string]*spriteItem{
	"boat1":      {Path: "boats/"},
	"boat2":      {Path: "boats/"},
	"boat2-flag": {Path: "boats/"},
	"sun":        {Path: "world/", doReflect: true},
	"island1":    {Path: "world/", doReflect: true},
	"title":      {Path: "title/"},
	"clickstart": {Path: "title/"},
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

func getAniFrame(frame int64, ani *spriteItem) *ebiten.Image {
	numFrames := int64(len(ani.animation.Frames))
	if frame < 0 || frame >= numFrames {
		fmt.Printf("%v: invalid frame number: %v\n", ani.Name, frame)
		if frame >= numFrames {
			frame = numFrames - 1
		} else if frame < 0 {
			frame = 0
		}
	}

	fKey := ani.animation.SortedFrames[frame]
	fRect := ani.animation.Frames[fKey].Frame
	rect := image.Rectangle{Min: image.Point{X: fRect.X, Y: fRect.Y}, Max: image.Point{X: fRect.X + fRect.W, Y: fRect.Y + fRect.H}}
	subFrame := ani.image.SubImage(rect).(*ebiten.Image)
	return subFrame
}

func autoAnimate(ani *spriteItem) *ebiten.Image {
	firstFrame := ani.animation.SortedFrames[0]
	speed := ani.animation.Frames[firstFrame].Duration
	time := time.Now().UnixMilli() / int64(speed)
	frameNum := time % (ani.animation.NumFrames)
	return getAniFrame(frameNum, ani)
}

func autoAnimatePingPong(ani *spriteItem) *ebiten.Image {

	period := 2*ani.animation.NumFrames - 1
	firstFrame := ani.animation.SortedFrames[0]
	speed := ani.animation.Frames[firstFrame].Duration
	time := time.Now().UnixMilli() / int64(speed)
	framePosition := time % period

	var frameNum int64
	if framePosition < int64(ani.animation.NumFrames) {
		// Forward direction
		frameNum = int64(framePosition)
	} else {
		// Backward direction
		frameNum = int64(period - framePosition - 1)
	}

	return getAniFrame(frameNum, ani)
}
