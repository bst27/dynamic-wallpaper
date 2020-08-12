# About
This app changes your wallpaper. You can give an URL to define where to get the wallpaper(s) from.
To define when to change your wallpaper you can use a cron expression.

Per default your wallpaper is changed every minute with a random image from [Unsplash](https://unsplash.com).

# Usage:
```
./dynamic-wallpaper -url "<URL>" -cron "<CRON>"
```

# Examples
| URL                                          | Info                                            | See                          |
|----------------------------------------------|-------------------------------------------------|------------------------------|
| https://source.unsplash.com/random/1920x1080 | Get random image with 1920x1080px from Unsplash | https://source.unsplash.com/ |
| https://loremflickr.com/1920/1080/dog        | Get random dog image with 1920x1080px           | https://loremflickr.com/     |

| Cron | Info |
|------|------|
| * * * * *  | Change wallpaper at every minute |
| 15 9 * * * | Change wallpaper at 09:15 every day |
| 0 6 * * 1  | Change wallpaper at 06:00 on monday |

To fiddle around with cron expressions you can use [crontab.guru](https://crontab.guru).

# Build
To build executables for multiple platforms you can use the build script at `scripts/build.sh`.