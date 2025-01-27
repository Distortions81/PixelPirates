package main

import (
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var spriteList map[string]*spriteItem = map[string]*spriteItem{
	"boat1":      {Path: "boats/"},
	"sun":        {Path: "world/"},
	"title":      {Path: "title/"},
	"clickstart": {Path: "title/"},
}

func loadSprites() {

	for name, sprite := range spriteList {
		image, err := loadSprite(sprite.Path+name, false)
		if err == nil {
			doLog(true, "loading sprite '"+name+"'")
			spriteList[name].image = image
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
	animation *animationData
}

func getAniFrame(frame int64, ani *spriteItem) *ebiten.Image {
	numFrames := int64(len(ani.animation.Frames))
	if frame < 0 || frame > numFrames {
		return nil
	}

	fKey := ani.animation.SortedFrames[frame]
	fRect := ani.animation.Frames[fKey].Frame
	rect := image.Rectangle{Min: image.Point{X: fRect.X, Y: fRect.Y}, Max: image.Point{X: fRect.X + fRect.W, Y: fRect.Y + fRect.H}}
	subFrame := ani.image.SubImage(rect).(*ebiten.Image)
	return subFrame
}

func autoAnimate(ani *spriteItem) *ebiten.Image {
	firstFrame := boatSP.animation.SortedFrames[0]
	speed := boatSP.animation.Frames[firstFrame].Duration
	time := time.Now().UnixMilli() / int64(speed)
	frameNum := time % (boatSP.animation.NumFrames)
	return getAniFrame(frameNum, boatSP)
}
