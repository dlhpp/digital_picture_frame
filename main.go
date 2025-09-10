package main

import (
	"net/http"

	"github.com/dlhpp/digital_picture_frame/internal"
	"github.com/dlhpp/digital_picture_frame/logging"
	"github.com/dlhpp/digital_picture_frame/yaml"
)

func main() {

	logging.SetLevel(3) // set to 1 for verbose, 5 for normal, 9 for very quiet

	commandLineFlags := internal.GetCommandLineArgs()

	yamlConfig := yaml.OpenYamlFile("config.yaml")

	store := internal.GetImageStore(yamlConfig, commandLineFlags)

	internal.SetupHttpHandlers(store)

	logging.Log("main: launch set to: ", 5, commandLineFlags.Launch)
	if commandLineFlags.Launch {
		// internal.LaunchDefaultBrowser(yamlConfig)
		// internal.LaunchBrowser(yamlConfig)
		// internal.LaunchBrowser02(yamlConfig)
		// internal.LaunchBrowser03(yamlConfig)
	}

	url := yaml.GetString(yamlConfig, "url", "localhost:81")
	logging.Log("main: listening:", 5, "url", url)
	if err := http.ListenAndServe(url, nil); err != nil {
		panic("main: Server failed to start: " + err.Error())
	}
}
