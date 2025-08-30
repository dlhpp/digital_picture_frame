package main

import (
	"fmt"
	"net/http"

	"github.com/dlhpp/digital_picture_frame/internal"
)

func main() {

	internal.DLHLogSetLevel(1)

	commandLineFlags := internal.SetupCommandLineArgs()

	internal.DLHLog("main", 5, fmt.Sprintf("commandLineFlags = %+v", commandLineFlags))

	store := internal.GetImageStore(commandLineFlags)

	internal.SetupHttpHandlers(store)

	internal.LaunchDefaultBrowser(commandLineFlags)

	host := commandLineFlags.Url
	internal.DLHLog("main: listening:", 5, "host", host)
	if err := http.ListenAndServe(host, nil); err != nil {
		panic("main: Server failed to start: " + err.Error())
	}
}
