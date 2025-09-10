package internal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/dlhpp/digital_picture_frame/logging"
	"github.com/dlhpp/digital_picture_frame/utils"
	"github.com/dlhpp/digital_picture_frame/yaml"
)

// BrowserType represents different browsers
type BrowserType string

const (
	Chrome   BrowserType = "chrome"
	Firefox  BrowserType = "firefox"
	Safari   BrowserType = "safari"
	Chromium BrowserType = "chromium"
	Midori   BrowserType = "midori"
	Epiphany BrowserType = "epiphany"
)

// BrowserLauncher handles launching browsers in kiosk/fullscreen mode
type BrowserLauncher struct {
	URL     string
	Browser BrowserType
}

// NewBrowserLauncher creates a new browser launcher instance
func NewBrowserLauncher(url string, browser BrowserType) *BrowserLauncher {
	return &BrowserLauncher{
		URL:     url,
		Browser: browser,
	}
}

// Launch opens the browser in kiosk/fullscreen mode
func (bl *BrowserLauncher) Launch() error {
	switch runtime.GOOS {
	case "windows":
		return bl.launchWindows()
	case "linux":
		return bl.launchLinux()
	case "darwin":
		return bl.launchMacOS()
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

// launchWindows handles Windows-specific browser launching
func (bl *BrowserLauncher) launchWindows() error {
	var cmd *exec.Cmd

	switch bl.Browser {
	case Chrome:
		cmd = exec.Command("chrome.exe",
			"--kiosk",
			"--disable-infobars",
			"--disable-session-crashed-bubble",
			"--disable-restore-session-state",
			"--disable-features=TranslateUI",
			"--disable-web-security",
			"--disable-features=VizDisplayCompositor",
			bl.URL,
		)
	case Firefox:
		cmd = exec.Command("firefox.exe",
			"-kiosk",
			bl.URL,
		)
	default:
		return fmt.Errorf("unsupported browser for Windows: %s", bl.Browser)
	}

	return bl.executeCommand(cmd)
}

// launchLinux handles Linux/Raspberry Pi-specific browser launching
func (bl *BrowserLauncher) launchLinux() error {
	var cmd *exec.Cmd

	switch bl.Browser {
	case Chrome:
		// Try google-chrome first, then chrome
		if bl.commandExists("google-chrome") {
			cmd = exec.Command("google-chrome",
				"--kiosk",
				"--no-first-run",
				"--disable-infobars",
				"--disable-session-crashed-bubble",
				"--disable-restore-session-state",
				"--disable-features=TranslateUI",
				"--disable-web-security",
				"--disable-dev-shm-usage",
				"--no-sandbox",
				bl.URL,
			)
		} else if bl.commandExists("chrome") {
			cmd = exec.Command("chrome",
				"--kiosk",
				"--no-first-run",
				"--disable-infobars",
				"--disable-session-crashed-bubble",
				"--disable-restore-session-state",
				"--disable-features=TranslateUI",
				"--disable-web-security",
				"--disable-dev-shm-usage",
				"--no-sandbox",
				bl.URL,
			)
		} else {
			return fmt.Errorf("chrome not found in PATH")
		}
	case Chromium:
		cmd = exec.Command("chromium-browser",
			"--kiosk",
			"--no-first-run",
			"--disable-infobars",
			"--disable-session-crashed-bubble",
			"--disable-restore-session-state",
			"--disable-features=TranslateUI",
			"--disable-web-security",
			"--disable-dev-shm-usage",
			"--no-sandbox",
			bl.URL,
		)
	case Firefox:
		cmd = exec.Command("firefox",
			"-kiosk",
			bl.URL,
		)
	case Midori:
		cmd = exec.Command("midori",
			"-e", "Fullscreen",
			"-a", bl.URL,
		)
	case Epiphany:
		cmd = exec.Command("epiphany",
			"--application-mode",
			bl.URL,
		)
	default:
		return fmt.Errorf("unsupported browser for Linux: %s", bl.Browser)
	}

	return bl.executeCommand(cmd)
}

// launchMacOS handles macOS-specific browser launching
func (bl *BrowserLauncher) launchMacOS() error {
	var cmd *exec.Cmd

	switch bl.Browser {
	case Firefox:
		cmd = exec.Command("open", "-a", "Firefox", "--args", "-kiosk", bl.URL)
	case Safari:
		// Safari doesn't have a built-in kiosk mode, so we'll open it normally
		// and use AppleScript to make it fullscreen
		cmd = exec.Command("osascript", "-e",
			fmt.Sprintf(`
				tell application "Safari"
					activate
					open location "%s"
					delay 2
					tell application "System Events"
						keystroke "f" using {control down, command down}
					end tell
				end tell
			`, bl.URL))
	case Chrome:
		// Try Google Chrome
		cmd = exec.Command("open", "-a", "Google Chrome", "--args",
			"--kiosk",
			"--disable-infobars",
			"--disable-session-crashed-bubble",
			"--disable-restore-session-state",
			"--disable-features=TranslateUI",
			bl.URL,
		)
	default:
		return fmt.Errorf("unsupported browser for macOS: %s", bl.Browser)
	}

	return bl.executeCommand(cmd)
}

// executeCommand executes the given command and handles errors
func (bl *BrowserLauncher) executeCommand(cmd *exec.Cmd) error {
	if cmd == nil {
		return fmt.Errorf("no command to execute")
	}

	// Set environment variables if needed
	cmd.Env = os.Environ()

	// For Linux, set DISPLAY if not set
	if runtime.GOOS == "linux" {
		display := os.Getenv("DISPLAY")
		if display == "" {
			cmd.Env = append(cmd.Env, "DISPLAY=:0")
		}
	}

	fmt.Printf("Executing: %s %s\n", cmd.Path, strings.Join(cmd.Args[1:], " "))

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start browser: %w", err)
	}

	return nil
}

// commandExists checks if a command exists in the system PATH
func (bl *BrowserLauncher) commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// GetSupportedBrowsers returns a list of browsers supported on the current OS
func GetSupportedBrowsers() []BrowserType {
	switch runtime.GOOS {
	case "windows":
		return []BrowserType{Chrome, Firefox}
	case "linux":
		return []BrowserType{Epiphany, Midori, Firefox, Chromium, Chrome}
	case "darwin":
		return []BrowserType{Firefox, Safari, Chrome}
	default:
		return []BrowserType{}
	}
}

// DetectAvailableBrowsers detects which browsers are actually installed
func DetectAvailableBrowsers() []BrowserType {
	var available []BrowserType
	supported := GetSupportedBrowsers()
	logging.Log("DetectAvailableBrowsers", 5, utils.DescribeVariable("supported", supported))

	for _, browser := range supported {
		if isInstalled(browser) {
			available = append(available, browser)
		}
	}

	return available
}

// isInstalled checks if a browser is installed on the system
func isInstalled(browser BrowserType) bool {
	var commands []string

	switch runtime.GOOS {
	case "windows":
		switch browser {
		case Chrome:
			commands = []string{"chrome.exe", "chrome"}
		case Firefox:
			commands = []string{"firefox.exe", "firefox"}
		}
	case "linux":
		switch browser {
		case Epiphany:
			commands = []string{"epiphany"}
		case Midori:
			commands = []string{"midori"}
		case Firefox:
			commands = []string{"firefox"}
		case Chromium:
			commands = []string{"chromium-browser", "chromium"}
		case Chrome:
			commands = []string{"google-chrome", "chrome"}
		}
	case "darwin":
		switch browser {
		case Firefox:
			// Check if Firefox.app exists
			if _, err := os.Stat("/Applications/Firefox.app"); err == nil {
				return true
			}
			commands = []string{"firefox"}
		case Safari:
			// Safari is built into macOS
			return true
		case Chrome:
			// Check if Google Chrome.app exists
			if _, err := os.Stat("/Applications/Google Chrome.app"); err == nil {
				return true
			}
			commands = []string{"google-chrome"}
		}
	}

	for _, cmd := range commands {
		if _, err := exec.LookPath(cmd); err == nil { // DLH: this failed on my windows!
			logging.Log("isInstalled", 59, "browser install FOUND! "+string(browser))
			return true
		}
	}

	logging.Log("isInstalled", 5, "browser install not found: "+string(browser))

	// return false
	return true // DLH: temporarily return true for testing on Windows
}

func LaunchBrowser02(yamlConfig *map[string]any) {
	logging.Log("LaunchBrowser02", 1, "entering", utils.DescribeVariable("yamlConfig", yamlConfig))

	url := yaml.GetString(yamlConfig, "url", "localhost:81")

	// browser.OpenURL("http://" + url) // open default browser using: github.com/pkg/browser

	// Detect available browsers
	available := DetectAvailableBrowsers()
	logging.Log("LaunchBrowser02", 1, fmt.Sprintf("Available browsers on %s: %v\n", runtime.GOOS, available))

	if len(available) == 0 {
		panic("LaunchBrowser02: No supported browsers found!")
	}

	// Use the first available browser
	browser := available[0]
	logging.Log("LaunchBrowser02", 1, fmt.Sprintf("Using browser: %s\n", browser))

	// Create launcher and launch browser
	launcher := NewBrowserLauncher(url, browser)
	err := launcher.Launch()
	if err != nil {
		fmt.Printf("LaunchBrowser02: Error launching browser: %v\n", err)
		os.Exit(1)
	}

	logging.Log("LaunchBrowser02", 5, "url="+url, ", Browser launched successfully!")

}
