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

func openBrowser(yamlConfig *map[string]any) error {
	logging.Log("openBrowser", 5, "entering, yamlConfig = ", yamlConfig)
	url := "http://" + yaml.GetString(yamlConfig, "url", "localhost:81")
	var cmd string
	var kiosk string = ""
	var args []string
	logging.Log("openBrowser", 5, "runtime.GOOS = "+runtime.GOOS)

	// Fullscreen Options: --start-fullscreen (Chrome), --kiosk (Chrome, Firefox)")
	switch runtime.GOOS {
	case "windows":
		// For Edge: args = []string{"/c", "start", "msedge", "--start-fullscreen", url}
		// For Firefox: args = []string{"/c", "start", "firefox", "-kiosk", url}

		cmd = "cmd"
		args = []string{"/c", "start", "chrome", kiosk, url}

	case "darwin": // macOS
		// macOS: Use 'open' with browser-specific flags
		// For Safari: args = []string{"-a", "Safari", url} // Safari doesn't support full-screen flag directly
		// For Firefox: args = []string{"-a", "Firefox", "--args", "-kiosk", url}

		cmd = "open"
		args = []string{"-a", "Google Chrome", "--args", kiosk, url} // For Chrome

	default: // Linux, BSD, etc.
		// Linux: Use 'xdg-open' or specific browser with flags
		// cmd = "google-chrome"
		// For Firefox: cmd = "firefox"; args = []string{"-kiosk", url}

		cmd = "chromium"
		args = []string{kiosk, url} // For Chrome
	}

	logging.Log("openBrowser", 5, fmt.Sprintf("cmd = %s, args = %s", cmd, args))
	return exec.Command(cmd, args...).Start()
}

func LaunchBrowser(yamlConfig *map[string]any) {
	logging.Log("LaunchBrowser", 5, "entering, yamlConfig = ", yamlConfig)
	err := openBrowser(yamlConfig)
	if err != nil {
		panic(err)
	}
}

func LaunchDefaultBrowser(yamlConfig *map[string]any) {
	logging.Log("LaunchDefaultBrowser", 1, "entering", utils.DescribeVariable("yamlConfig", yamlConfig))

	url := yaml.GetString(yamlConfig, "url", "localhost:81")
	browser.OpenURL("http://" + url)

	// logging.Log("LaunchDefaultBrowser", 3, "url="+url, ", launchCommand="+launchCommand)
	logging.Log("LaunchDefaultBrowser", 3, "url="+url)
}
