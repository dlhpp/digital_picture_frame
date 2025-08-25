package main

import (
	"log/slog"
	"net/http"

	"github.com/dlhpp/digital_picture_frame/internal"
)

func main() {
	internal.SetupLogging()

	commandLineFlags := internal.SetupCommandLineArgs()

	slog.Info("main: commandLineFlags", "commandLineFlags", commandLineFlags)

	store := internal.GetImageStore(commandLineFlags)

	internal.SetupHttpHandlers(store)

	internal.LaunchDefaultBrowser(commandLineFlags)

	// host := "localhost:81"
	// slog.Info("main: listening:", "host", host)
	// if err := http.ListenAndServe(host, nil); err != nil {
	// 	panic("main: Server failed to start: " + err.Error())
	// }

	host := commandLineFlags.Url
	slog.Info("main: listening:", "host", host)
	if err := http.ListenAndServe(host, nil); err != nil {
		panic("main: Server failed to start: " + err.Error())
	}
}
