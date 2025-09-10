package internal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/dlhpp/digital_picture_frame/yaml"
)

// BrowserConfig defines browser-specific commands and flags
type BrowserConfig struct {
	Executable string   // Browser executable name
	Args       []string // Arguments for full-screen/kiosk mode
	MacScript  string   // AppleScript for macOS (if needed)
}

var (
	// Browser configurations for each supported browser
	browserConfigs = map[string]map[string]BrowserConfig{
		"windows": {
			"chrome": {
				Executable: "chrome",
				Args:       []string{"--new-window", "--kiosk"}, // Chrome supports kiosk mode
			},
			"firefox": {
				Executable: "firefox",
				Args:       []string{"--new-window", "--full-screen"}, // Firefox uses --full-screen
			},
		},
		"linux": {
			"chromium": {
				Executable: "chromium-browser",
				Args:       []string{"--new-window", "--kiosk"}, // Chromium supports kiosk mode
			},
			"firefox": {
				Executable: "firefox",
				Args:       []string{"--new-window", "--full-screen"},
			},
			"midori": {
				Executable: "midori",
				Args:       []string{"--fullscreen"}, // Midori uses --fullscreen
			},
			"epiphany": {
				Executable: "epiphany",
				Args:       []string{"--new-window", "--fullscreen"}, // Epiphany uses --fullscreen
			},
		},
		"darwin": {
			"firefox": {
				Executable: "firefox",
				Args:       []string{"--new-window", "--full-screen"},
			},
			"safari": {
				Executable: "open",
				Args:       []string{"-a", "Safari"}, // Use 'open' for Safari
				MacScript: `
                    tell application "Safari"
                        activate
                        set URL of document 1 to "%s"
                        tell application "System Events"
                            keystroke "t" using {command down} -- New tab
                            keystroke "f" using {command down, control down} -- Full-screen
                        end tell
                    end tell
                `,
			},
		},
	}
)

func LaunchBrowser03(yamlConfig *map[string]any) {
	// Command-line flags
	// browser := flag.String("browser", "", "Browser to use (e.g., chrome, firefox, chromium, midori, epiphany, safari)")
	// url := flag.String("url", "https://example.com", "URL to open")
	// mode := flag.String("mode", "kiosk", "Mode: fullscreen or kiosk")
	// flag.Parse()
	// if *browser == "" {
	// 	fmt.Println("Error: --browser flag is required")
	// 	flag.Usage()
	// 	os.Exit(1)
	// }

	browser := "chrome"
	mode := "fullscreen"
	url := yaml.GetString(yamlConfig, "url", "localhost:81")

	// Determine the OS
	osType := runtime.GOOS
	configs, exists := browserConfigs[osType]
	if !exists {
		fmt.Printf("Error: Unsupported operating system: %s\n", osType)
		os.Exit(1)
	}
	// config, exists := configs[*browser]
	config, exists := configs["chrome"]
	if !exists {
		// fmt.Printf("Error: Browser %s is not supported on %s\n", *browser, osType)
		fmt.Printf("Error: Browser %s is not supported on %s\n", browser, osType)
		os.Exit(1)
	}
	// Adjust arguments based on mode
	args := make([]string, len(config.Args))
	copy(args, config.Args)
	// if *mode == "fullscreen" {
	// Replace --kiosk with --start-fullscreen or equivalent where applicable
	for i, arg := range args {
		if arg == "--kiosk" {
			args[i] = "--start-fullscreen"
		}
	}
	// }

	// // Special handling for Safari on macOS
	// if osType == "darwin" && *browser == "safari" {
	// 	if *mode == "kiosk" {
	// 		fmt.Println("Warning: Safari does not support kiosk mode; using full-screen mode")
	// 	}
	// 	err := openSafari(*url, config.MacScript)
	// 	if err != nil {
	// 		fmt.Printf("Error opening Safari: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	return
	// }

	// General case: execute browser command
	// args = append(args, *url)
	args = append(args, url)
	cmd := exec.Command(config.Executable, args...)

	// Set PATH to ensure browsers can be found
	// DLH: this makes no sense on Windows or macOS, only Linux?
	// cmd.Env = append(os.Environ(), "PATH=/usr/bin:/usr/local/bin:/bin")

	err := cmd.Start()
	if err != nil {
		// fmt.Printf("Error opening %s: %v\n", *browser, err)
		fmt.Printf("Error opening %s: %v\n", browser, err)
		os.Exit(1)
	}

	// fmt.Printf("Opened %s in %s mode with URL: %s\n", *browser, *mode, *url)
	fmt.Printf("Opened %s in %s mode with URL: %s\n", browser, mode, url)
}

// // openSafari uses AppleScript to open Safari in full-screen mode
// func openSafari(url, script string) error {
// 	if script == "" {
// 		return fmt.Errorf("no AppleScript defined for Safari")
// 	}
// 	// Format the AppleScript with the URL
// 	script = fmt.Sprintf(script, url)
// 	// Execute AppleScript using osascript
// 	cmd := exec.Command("osascript", "-e", script)
// 	return cmd.Run()
// }
