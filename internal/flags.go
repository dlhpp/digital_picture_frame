package internal

import (
	"flag"

	"github.com/dlhpp/digital_picture_frame/logging"
	"github.com/dlhpp/digital_picture_frame/utils"
)

func GetCommandLineArgs() *FlagSettings {

	browser := flag.String("browser", "chrome", "Required when launch=true.  Options: default, chrome, firefox, epiphany, midori, safari.")
	launch := flag.Bool("launch", true, "Automatically launch browser.  If false then manually open browser to http://localhost:81")
	random := flag.Bool("random", true, "Randomize the order of images")

	flag.Parse()

	result := &FlagSettings{
		Browser: *browser,
		Launch:  *launch,
		Random:  *random,
		Rest:    flag.Args(),
	}

	logging.Log("GetCommandLineArgs", 3, utils.DescribeVariable("result", result))

	return result
}
