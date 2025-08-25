package main

import (
	"dlh/slideshow/internal"
	"log/slog"
	"net/http"
)

func main() {
	internal.SetupLogging()

	slog.Info("main: Begin ")

	internal.LaunchDefaultBrowser()

	store := internal.GetImageStore()

	internal.SetupHttpHandlers(store)

	host := "localhost:81"
	slog.Info("main: listening:", "host", host)
	if err := http.ListenAndServe(host, nil); err != nil {
		panic("main: Server failed to start: " + err.Error())
	}

	slog.Info("main: Program has quit and we'll never reach this point.")
}
