package main

import (
	"flag"
	"fmt"
	"github.com/reujab/wallpaper"
	"github.com/robfig/cron/v3"
)

func main() {
	var cronFlag = flag.String("cron", "* * * * *", "Cron expression which defines when to change wallpaper")
	var wallpaperURL = flag.String("url", "https://source.unsplash.com/random/1920x1080", "URL which returns a wallpaper image")
	flag.Parse()

	c := cron.New()
	_, err := c.AddFunc(*cronFlag, func() {
		// Silently ignore errors and hope for success in the next try
		_ = changeWallpaper(*wallpaperURL)
	})

	if err != nil {
		fmt.Println("Invalid cron expression.", "For help visit https://crontab.guru")
		fmt.Println(err)
		return
	}

	c.Start()

	// Block forever until app is stopped
	select {}
}

func changeWallpaper(url string) error {
	return wallpaper.SetFromURL(url)
}
