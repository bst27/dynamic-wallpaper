package main

import (
	"github.com/reujab/wallpaper"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("* * * * *", func() {
		_ = changeWallpaper()
	})

	c.Start()

	// Block forever until app is stopped
	select {}
}

func changeWallpaper() error {
	return wallpaper.SetFromURL("https://source.unsplash.com/random/1920x1080")
}
