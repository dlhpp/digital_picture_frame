package internal

import (
	"flag"
)

func SetupCommandLineArgs() *FlagSettings {

	kiosk := flag.Bool("kiosk", false, "Open browser in kiosk mode")
	random := flag.Bool("random", true, "Randomize the order of images")
	url := flag.String("url", "localhost:81", "URL for the initial container page.  This page defines '/next' as the link to get images.")

	flag.Parse()

	return &FlagSettings{
		Kiosk:  *kiosk,
		Random: *random,
		Url:    *url,
	}
}
