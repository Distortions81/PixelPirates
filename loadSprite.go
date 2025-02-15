package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io"
	"io/fs"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/anthonynsimon/bild/blur"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed data/sprites/boats data/sprites/characters data/sprites/islands data/sprites/title data/sprites/world

var efs embed.FS

// Load sprites
func loadImage(name string, unmanaged bool, doBlur bool) (*ebiten.Image, *ebiten.Image, error) {

	//Open file
	var (
		err     error
		pngData fs.File
	)

	if wasmMode {
		pngData, err = efs.Open(name + ".png")
	} else {
		pngData, err = os.Open(name + ".png")
	}
	if err != nil {
		doLog(true, false, "loadSprite: Open: %v", err)
		return nil, nil, err
	}

	//Decode png
	m, err := png.Decode(pngData)
	if err != nil {
		doLog(true, false, "loadSprite: Decode: %v", err)
		return nil, nil, err
	}

	//Create image
	var (
		img, blurImg *ebiten.Image
		newBlur      image.Image
	)
	if doBlur {
		newBlur = blur.Box(m, islandReflectionBlur)
	}

	if unmanaged {
		img = ebiten.NewImageFromImageWithOptions(m, &ebiten.NewImageFromImageOptions{Unmanaged: true})
		if doBlur {
			blurImg = ebiten.NewImageFromImageWithOptions(newBlur, &ebiten.NewImageFromImageOptions{Unmanaged: true})
		}
	} else {
		img = ebiten.NewImageFromImage(m)
		if doBlur {
			blurImg = ebiten.NewImageFromImage(newBlur)
		}
	}

	return img, blurImg, nil
}

func loadAnimationData(name string) (*animationData, error) {
	if wasmMode {
		jdata, err := efs.Open(dataDir + spritesDir + name + ".json")
		if err != nil {
			return nil, err
		}
		buf, err := io.ReadAll(jdata)
		if err != nil {
			doLog(true, false, "loadAnimationData: Embedded: %v", err)
			return nil, err
		}

		aniJSON, err := decodeAniJSON(buf)
		if err != nil {
			//doLog(true, false, "loadAnimationData: Embedded: %v", err)
			return nil, err
		}

		return &aniJSON, nil
	} else {
		buf, err := os.ReadFile(dataDir + spritesDir + name + ".json")
		if err != nil {
			return nil, err
		}

		aniJSON, err := decodeAniJSON(buf)
		if err != nil {
			//doLog(true, false, "loadAnimationData: Embedded: %v", err)
			return nil, err
		}

		return &aniJSON, nil
	}
}

func decodeAniJSON(data []byte) (animationData, error) {

	var root animationData
	err := json.Unmarshal(data, &root)
	if err != nil {
		doLog(true, false, "decodeAniJSON: %v", err)
		return animationData{}, err
	}

	root.animations = map[string]frameRange{}

	for _, item := range root.Meta.FrameTags {
		if item.From+item.To == 0 {
			doLog(true, false, "Empty Animation: '%v', %v->%v", item.Name, item.From, item.To)
			continue
		}
		if *debugMode {
			doLog(true, true, "Animation: %v, %v->%v", item.Name, item.From, item.To)
		}
		root.animations[item.Name] = frameRange{start: item.From, end: item.To}
	}

	// Extract and sort frame names based on the numerical part.
	sorted, err := getSortedFrameNames(root.Frames)
	if err != nil {
		//doLog(true, false, "Error sorting frame names: %v", err)
		return animationData{}, err
	}
	root.sortedFrames = sorted
	root.numFrames = int64(len(sorted))

	doLog(true, true, "Decoded animation for: %v", root.Meta.Image)
	if *debugMode {
		/*
			fmt.Println("Frames:")
			for _, fKey := range root.sortedFrames {
				frameData := root.Frames[fKey]
				doLog(true, "Frame Name: %s", fKey)
				doLog(true, "  Position: (%d, %d)", frameData.Frame.X, frameData.Frame.Y)
				doLog(true, "  Size: %dx%d", frameData.Frame.W, frameData.Frame.H)
				doLog(true, "  Rotated: %t", frameData.Rotated)
				doLog(true, "  Trimmed: %t", frameData.Trimmed)
				doLog(true, "  Sprite Source Size: (%d, %d, %dx%d)", frameData.SpriteSourceSize.X, frameData.SpriteSourceSize.Y, frameData.SpriteSourceSize.W, frameData.SpriteSourceSize.H)
				doLog(true, "  Source Size: %dx%d", frameData.SourceSize.W, frameData.SourceSize.H)
				doLog(true, "  Duration: %dms", frameData.Duration)
				fmt.Println()
			} */
	}

	return root, nil
}

// getSortedFrameNames extracts frame names and sorts them based on the numerical index.
// It assumes frame names follow the pattern "name <number>.extension", e.g., "boat 0.png".
func getSortedFrameNames(frames map[string]aniFrame) ([]string, error) {
	// Regular expression to extract the numerical part from the frame name.
	// This regex captures the number between the last space and the file extension.
	re := regexp.MustCompile(`\s*(\d+)\.\w+$`)

	// Create a slice to hold frame names and their corresponding indices.
	type FrameWithIndex struct {
		Name  string
		Index int
	}

	var frameList []FrameWithIndex

	for name := range frames {
		matches := re.FindStringSubmatch(name)
		if len(matches) < 2 {
			return nil, fmt.Errorf("frame name '%s' does not match expected pattern", name)
		}
		index, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, fmt.Errorf("invalid frame index in name '%s': %v", name, err)
		}
		frameList = append(frameList, FrameWithIndex{
			Name:  name,
			Index: index,
		})
	}

	// Sort the frames based on the extracted index.
	sort.Slice(frameList, func(i, j int) bool {
		return frameList[i].Index < frameList[j].Index
	})

	// Extract the sorted frame names.
	var sortedNames []string
	for _, frame := range frameList {
		sortedNames = append(sortedNames, frame.Name)
	}

	return sortedNames, nil
}

// animationData represents the top-level structure of the JSON.
type animationData struct {
	Frames map[string]aniFrame `json:"frames"`
	Meta   aniMeta             `json:"meta"`

	//Local
	sortedFrames []string              `json:"-"`
	numFrames    int64                 `json:"-"`
	animations   map[string]frameRange `json:"-"`
}

type frameRange struct {
	start, end int
}

// aniFrame represents each individual frame in the "frames" object.
type aniFrame struct {
	Frame            aniRect `json:"frame"`
	Rotated          bool    `json:"rotated"`
	Trimmed          bool    `json:"trimmed"`
	SpriteSourceSize aniRect `json:"spriteSourceSize"`
	SourceSize       aniSize `json:"sourceSize"`
	Duration         int     `json:"duration"`
}

// aniRect represents a rectangle with position and size.
type aniRect struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

// aniSize represents width and height.
type aniSize struct {
	W int `json:"w"`
	H int `json:"h"`
}

// aniMeta contains metadata about the frames and the source image.
type aniMeta struct {
	App       string         `json:"app"`
	Version   string         `json:"version"`
	Image     string         `json:"image"`
	Format    string         `json:"format"`
	Size      aniSize        `json:"size"`
	Scale     string         `json:"scale"`
	FrameTags []frameTagData `json:"frameTags"`
	Layers    []aniLayer     `json:"layers"`
	Slices    []interface{}  `json:"slices"`
}

type frameTagData struct {
	Name     string
	From, To int
}

// aniLayer represents each layer in the "layers" array within "meta".
type aniLayer struct {
	Name      string `json:"name"`
	Opacity   int    `json:"opacity"`
	BlendMode string `json:"blendMode"`
}
