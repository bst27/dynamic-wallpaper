package main

import (
	"github.com/reujab/wallpaper"
)

func main() {
	err := changeWallpaper()
	if err != nil {
		panic(err)
	}
}

func changeWallpaper() error {
	// Works fine with linux but does not revert on windows
	return wallpaper.SetFromURL("https://source.unsplash.com/random/1920x1080")
}
