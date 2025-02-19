package main

import (
	"os"
	"strings"
)

const infoJsonFile = "info.json"

func scanIslandsFolder() {
	var dir []os.DirEntry
	var err error
	dirPath := dataDir + spritesDir + islandsDir

	doLog(true, true, "scanIslandsFolder: Scanning.")

	if wasmMode {
		dir, err = efs.ReadDir(dirPath)
	} else {
		dir, err = os.ReadDir(dirPath)
	}
	if err != nil {
		doLog(true, false, "scanIslandsFolder: readDir: %v", err)
		return
	}

	var islandsList []string
	for _, item := range dir {
		if item.IsDir() {
			islandsList = append(islandsList, item.Name())
		}
	}
	doLog(true, true, "Islands found: %v", strings.Join(islandsList, ", "))

	for _, island := range islandsList {
		_, err := os.ReadFile(dirPath + "/" + island + "/" + infoJsonFile)
		if err != nil {
			doLog(true, false, "Island '%v' has no %v file.", island, infoJsonFile)
			return
		}
	}
}
