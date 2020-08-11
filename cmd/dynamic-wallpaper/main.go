package main

import (
	"fmt"
	"github.com/reujab/wallpaper"
	"time"
)

func main() {
	background, err := wallpaper.Get()

	if err != nil {
		panic(err)
	}

	fmt.Println("Current wallpaper:", background)

	// Works fine with linux but does not revert on windows
	err = wallpaper.SetFromURL("https://source.unsplash.com/random/1920x1080")
	if err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)

	wallpaper.SetFromFile(background)
}
