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
	c.AddFunc(*cronFlag, func() {
		// Silently ignore errors and hope for success in the next try
		_ = changeWallpaper(*wallpaperURL)
	})

	c.Start()

	if len(c.Entries()) == 0 {
		fmt.Println("Invalid cron expression. For help visit https://crontab.guru")
		return
	}

	// Block forever until app is stopped
	select {}
}

func changeWallpaper(url string) error {
	return wallpaper.SetFromURL(url)
}
