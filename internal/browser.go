package internal

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/dlhpp/digital_picture_frame/logging"
	"github.com/dlhpp/digital_picture_frame/utils"
	"github.com/dlhpp/digital_picture_frame/yaml"

	"github.com/pkg/browser"
)

func LaunchBrowser(yamlConfig *map[string]any) error {
	launchBrowser := (yaml.Get(yamlConfig, "launch", true)).(bool)
	logging.Log("LaunchBrowser", 5, fmt.Sprintf("launchBrowser = %t", launchBrowser))

	browser := yaml.GetString(yamlConfig, "browser", "chrome")
	url := "http://" + yaml.GetString(yamlConfig, "host", "localhost:81")

	if !launchBrowser {
		return nil
	} else if browser == "default" {
		err := LaunchDefaultBrowser(yamlConfig)
		logging.Log("LaunchBrowser", 5, fmt.Sprintf("err(1) = %+v", err))
		return err
	} else {
		osType := runtime.GOOS
		logging.Log("LaunchBrowser", 5, fmt.Sprintf("osType = %s, browser = %s", osType, browser))

		browserConfig := yaml.OpenYamlFile("static/config/config_browser.yaml")
		cmd := yaml.GetString(browserConfig, fmt.Sprintf("%s.%s.executable", osType, browser), "")
		args := yaml.GetStringArray(browserConfig, fmt.Sprintf("%s.%s.args", osType, browser), []string{})
		args = append(args, url)
		logging.Log("LaunchBrowser", 5, fmt.Sprintf("cmd=%s, args=%+v", cmd, args))

		err := exec.Command(cmd, args...).Start()
		errMsg := utils.Tern(err == nil, "Browser launched successfully", fmt.Sprintf("Error launching browser: %+v", err))
		logging.Log("LaunchBrowser", 5, fmt.Sprintf("%+v", errMsg))
		return err
	}
}

func LaunchDefaultBrowser(yamlConfig *map[string]any) error {
	host := yaml.GetString(yamlConfig, "host", "localhost:81")
	logging.Log("LaunchDefaultBrowser", 5, "host="+host)
	err := browser.OpenURL("http://" + host)
	logging.Log("LaunchDefaultBrowser", 5, fmt.Sprintf("err = %+v", err))
	return err
}
