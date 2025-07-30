package main

import (
	"flag"
	"fmt"
	"github.com/reujab/wallpaper"
	"github.com/robfig/cron/v3"
	"os"
)

func main() {
	var cronFlag = flag.String("cron", "* * * * *", "Cron expression which defines when to change wallpaper")
	var wallpaperURL = flag.String("url", "https://picsum.photos/1920/1080", "URL which returns a wallpaper image")
	flag.Parse()

	args := os.Args
	if len(args) == 2 && args[1] == "help" {
		flag.PrintDefaults()
		os.Exit(0)
	}

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
