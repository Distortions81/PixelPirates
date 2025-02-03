package main

import (
	"fmt"
	"image"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var spriteList map[string]*spriteItem = map[string]*spriteItem{
	"boat2":         {Path: "boats/"},
	"boat2-flag":    {Path: "boats/"},
	"sun":           {Path: "world/"},
	"island1":       {Path: "world/", doReflect: true},
	"title":         {Path: "title/"},
	"clickstart":    {Path: "title/"},
	"visit":         {Path: "ui/"},
	"testScene1":    {Path: "islands/"},
	"island-scene1": {Path: "islands/"},
}

var (
	testScene1SP, visitSP, boat2SP, boat2SP_flag, sunSP, island1SP, titleSP, clickStartSP *spriteItem
)

func initSprites() {
	titleSP = spriteList["title"]
	clickStartSP = spriteList["clickstart"]

	sunSP = spriteList["sun"]
	boat2SP = spriteList["boat2"]
	boat2SP_flag = spriteList["boat2-flag"]
	visitSP = spriteList["visit"]
	island1SP = spriteList["island1"]

	testScene1SP = spriteList["island-scene1"]

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

	fKey := ani.animation.SortedFrames[frame]
	fRect := ani.animation.Frames[fKey].Frame
	rect := image.Rectangle{Min: image.Point{X: fRect.X, Y: fRect.Y}, Max: image.Point{X: fRect.X + fRect.W, Y: fRect.Y + fRect.H}}
	subFrame := ani.image.SubImage(rect).(*ebiten.Image)
	return subFrame
}

func autoAnimate(ani *spriteItem, offset int) *ebiten.Image {
	firstFrame := ani.animation.SortedFrames[0]
	speed := ani.animation.Frames[firstFrame].Duration
	time := time.Now().UnixMilli() / int64(speed)
	frameNum := time % (ani.animation.NumFrames)

	return getAniFrame(frameNum, ani, offset)
}

func autoAnimatePingPong(ani *spriteItem, offset int) *ebiten.Image {

	period := 2*ani.animation.NumFrames - 2
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
		frameNum = int64(period - framePosition)
	}

	return getAniFrame(frameNum, ani, offset)
}
