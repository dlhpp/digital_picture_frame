package internal

import (
	"flag"
)

func SetupCommandLineArgs() *FlagSettings {

	fullscreen := flag.Bool("fullscreen", false, "Open browser in fullscreen mode")
	screensize := flag.String("screensize", "--start-fullscreen", "Command for opening fullscreen.  Options: --start-fullscreen (Chrome), --kiosk (Chrome, Firefox)")
	random := flag.Bool("random", true, "Randomize the order of images")
	url := flag.String("url", "localhost:81", "URL for the initial container page.  This page defines '/next' as the link to get images.")

	flag.Parse()

	return &FlagSettings{
		Fullscreen: *fullscreen,
		Screensize: *screensize,
		Random:     *random,
		Url:        *url,
	}
}
