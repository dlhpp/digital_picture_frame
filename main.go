package main

import (
	"fmt"
	"net/http"

	"github.com/dlhpp/digital_picture_frame/internal"
	"github.com/dlhpp/digital_picture_frame/logging"
	"github.com/dlhpp/digital_picture_frame/yaml"
)

func main() {

	logging.SetLevel(3)

	commandLineFlags := internal.SetupCommandLineArgs()

	yamlConfig := yaml.OpenYamlFile("yamlConfigComplex.yaml")

	logging.Log("main", 5, fmt.Sprintf("commandLineFlags = %+v", commandLineFlags))

	store := internal.GetImageStore(yamlConfig, commandLineFlags)

	internal.SetupHttpHandlers(store)

	internal.LaunchDefaultBrowser(commandLineFlags)

	host := commandLineFlags.Url
	logging.Log("main: listening:", 5, "host", host)
	if err := http.ListenAndServe(host, nil); err != nil {
		panic("main: Server failed to start: " + err.Error())
	}
}
