package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"image"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed data

var efs embed.FS

const loadEmbedSprites = true

// Load sprites
func loadSprite(name string, unmanaged bool) (*ebiten.Image, error) {

	if loadEmbedSprites {

		//Open file
		pngData, err := efs.Open(spritesDir + name + ".png")
		if err != nil {
			doLog(true, "loadSprite: Embedded: %v", err)
			return nil, err
		}

		//Decode png
		m, _, err := image.Decode(pngData)
		if err != nil {
			doLog(true, "loadSprite: Embedded: %v", err)
			return nil, err
		}

		//Create image
		var img *ebiten.Image
		if unmanaged {
			img = ebiten.NewImageFromImageWithOptions(m, &ebiten.NewImageFromImageOptions{Unmanaged: true})
		} else {
			img = ebiten.NewImageFromImage(m)
		}
		return img, nil

	} else {
		img, _, err := ebitenutil.NewImageFromFile(dataDir + spritesDir + name)
		if err != nil {
			doLog(true, "loadSprite: File: %v", err)
		}
		return img, err
	}
}

func loadAnimationData(name string) (*animationData, error) {
	if loadEmbedSprites {
		jdata, err := efs.Open(spritesDir + name + ".json")
		if err != nil {
			return nil, err
		}
		buf, err := io.ReadAll(jdata)
		if err != nil {
			doLog(true, "loadAnimationData: Embedded: %v", err)
			return nil, err
		}

		aniJSON, err := decodeAniJSON(buf)
		if err != nil {
			doLog(true, "loadAnimationData: Embedded: %v", err)
			return nil, err
		}

		return &aniJSON, nil
	} else {
		buf, err := os.ReadFile(spritesDir + name + ".json")
		if err != nil {
			return nil, err
		}

		aniJSON, err := decodeAniJSON(buf)
		if err != nil {
			doLog(true, "loadAnimationData: Embedded: %v", err)
			return nil, err
		}

		return &aniJSON, nil
	}
}

func decodeAniJSON(data []byte) (animationData, error) {

	var root animationData
	err := json.Unmarshal(data, &root)
	if err != nil {
		doLog(true, "decodeAniJSON: %v", err)
		return animationData{}, err
	}

	// Extract and sort frame names based on the numerical part.
	sorted, err := getSortedFrameNames(root.Frames)
	if err != nil {
		log.Fatalf("Error sorting frame names: %v", err)
	}
	root.SortedFrames = sorted

	if verbose {
		fmt.Println("Frames:")
		for _, fKey := range root.SortedFrames {
			frameData := root.Frames[fKey]
			fmt.Printf("Frame Name: %s\n", fKey)
			fmt.Printf("  Position: (%d, %d)\n", frameData.Frame.X, frameData.Frame.Y)
			fmt.Printf("  Size: %dx%d\n", frameData.Frame.W, frameData.Frame.H)
			fmt.Printf("  Rotated: %t\n", frameData.Rotated)
			fmt.Printf("  Trimmed: %t\n", frameData.Trimmed)
			fmt.Printf("  Sprite Source Size: (%d, %d, %dx%d)\n", frameData.SpriteSourceSize.X, frameData.SpriteSourceSize.Y, frameData.SpriteSourceSize.W, frameData.SpriteSourceSize.H)
			fmt.Printf("  Source Size: %dx%d\n", frameData.SourceSize.W, frameData.SourceSize.H)
			fmt.Printf("  Duration: %dms\n", frameData.Duration)
			fmt.Println()
		}
	} else {
		fmt.Printf("Loaded animation for: %v\n", root.Meta.Image)
	}

	return root, nil
}

// animationData represents the top-level structure of the JSON.
type animationData struct {
	Frames       map[string]aniFrame `json:"frames"`
	SortedFrames []string            `json:"-"`
	Meta         aniMeta             `json:"meta"`
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
	App       string        `json:"app"`
	Version   string        `json:"version"`
	Image     string        `json:"image"`
	Format    string        `json:"format"`
	Size      aniSize       `json:"size"`
	Scale     string        `json:"scale"`
	FrameTags []interface{} `json:"frameTags"`
	Layers    []aniLayer    `json:"layers"`
	Slices    []interface{} `json:"slices"`
}

// aniLayer represents each layer in the "layers" array within "meta".
type aniLayer struct {
	Name      string `json:"name"`
	Opacity   int    `json:"opacity"`
	BlendMode string `json:"blendMode"`
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
