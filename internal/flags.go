package internal

import (
	"flag"
)

func SetupCommandLineArgs() *FlagSettings {

	browser := flag.Bool("browser", true, "auto-open default web browser")
	fullscreen := flag.Bool("fullscreen", false, "If automatically opening a browser window, open it in fullscreen mode")
	random := flag.Bool("random", true, "Randomize the order of images")
	url := flag.String("url", "localhost:81", "Randomize the order of images")

	flag.Parse()

	return &FlagSettings{
		Browser:    *browser,
		Fullscreen: *fullscreen,
		Random:     *random,
		Url:        *url,
	}
}
