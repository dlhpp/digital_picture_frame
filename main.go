package main

import (
	"net/http"

	"github.com/dlhpp/digital_picture_frame/internal"
	"github.com/dlhpp/digital_picture_frame/logging"
	"github.com/dlhpp/digital_picture_frame/yaml"
)

func main() {

	yamlConfig := yaml.OpenYamlFile("config.yaml")

	logging.SetLevel(yaml.GetInt(yamlConfig, "loglevel", 5)) // set to 1 for verbose, 5 for normal, 9 for very quiet

	store := internal.GetImageStore(yamlConfig)

	internal.SetupHttpHandlers(store)

	internal.LaunchBrowser(yamlConfig)

	host := yaml.GetString(yamlConfig, "host", "localhost:81")
	logging.Log("main: listening:", 5, "host", host)
	if err := http.ListenAndServe(host, nil); err != nil {
		panic("main: Server failed to start: " + err.Error())
	}
}
