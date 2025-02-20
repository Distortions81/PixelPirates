package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//go:embed data/sprites/boats data/sprites/characters  data/sprites/islands data/sprites/title data/sprites/world
var efs embed.FS

func loadAnimationData(name string) (*animationData, error) {
	var data []byte
	var err error

	if wasmMode {
		data, err = efs.ReadFile(name + ".json")
	} else {
		data, err = os.ReadFile(name + ".json")
	}
	if err != nil {
		//doLog(true, false, "loadAnimationData: %v: %v", name, err)
		return nil, err
	}

	aniJSON, err := decodeAniJSON(data)
	if err != nil {
		doLog(true, false, "loadAnimationData: decodeAniJSON: %v: %v", name, err)
		return nil, err
	}

	return &aniJSON, nil
}

func decodeAniJSON(data []byte) (animationData, error) {

	var root animationData
	err := json.Unmarshal(data, &root)
	if err != nil {
		doLog(true, false, "decodeAniJSON: %v", err)
		return animationData{}, err
	}

	//Parse frame tags
	root.animations = map[string]frameRange{}

	var buf string
	for _, item := range root.Meta.FrameTags {
		if item.From+item.To == 0 {
			doLog(true, false, "Empty Animation: '%v', %v->%v", item.Name, item.From, item.To)
			continue
		}
		if *debugMode {
			if buf != "" {
				buf = buf + ", "
			}
			buf = buf + fmt.Sprintf("%v: %v->%v", item.Name, item.From, item.To)
		}
		root.animations[item.Name] = frameRange{start: item.From, end: item.To}
	}
	if buf != "" {
		doLog(true, true, "Parsing animation for: %v", root.Meta.Image)
		doLog(true, true, buf)
	}

	// Extract and sort frame names based on the numerical part.
	sorted, err := getSortedFrameNames(root.Frames)

	if err != nil {
		root.numFrames = 0
		for name, _ := range root.Frames {
			root.sortedFrames = append(root.sortedFrames, name)
			root.numFrames++
		}
	} else {
		root.sortedFrames = sorted
		root.numFrames = int64(len(sorted))
	}

	re := regexp.MustCompile(`\(([^)]+)\)`)

	root.layers = map[string]*aniFrame{}

	//Parse layers/frames
	for _, layer := range root.Meta.Layers {
		for fname, frame := range root.Frames {
			matches := re.FindStringSubmatch(fname)
			if len(matches) != 2 {
				continue
			}
			lName := strings.ToLower(matches[1])
			if strings.EqualFold(lName, layer.Name) {
				root.layers[lName] = &frame
				doLog(true, true, "found layer: %v", lName)
			}
		}
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
	layers       map[string]*aniFrame  `json:"-"`
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
