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

	err = wallpaper.SetFromURL("https://picsum.photos/1920/1080")
	if err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)

	wallpaper.SetFromFile(background)
}
