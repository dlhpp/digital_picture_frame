package internal

// Can the default browser be launched in full-screen mode?
// Yes, Go can launch the default browser in full-screen mode,
// but the approach depends on the operating system and browser,
// as full-screen behavior is controlled by browser-specific
// flags or system commands. The os/exec package can be used
// to pass appropriate command-line arguments to the browser,
// but direct full-screen support varies.
//
// Below is an updated example that attempts to open the default
// browser in full-screen mode where possible:
//
// see:   https://github.com/pkg/browser
// this might be useful for interacting with the browser.

/*
Key Points:

Browser-Specific Flags:
	- Google Chrome: Use --start-fullscreen for full-screen mode or --kiosk for a more locked-down kiosk mode.
	- Firefox: Use -kiosk for full-screen kiosk mode.
	- Microsoft Edge: Supports --start-fullscreen like Chrome.
	- 	- Safari: No direct command-line flag for full-screen; it may require AppleScript or manual user action.

Operating System Notes:
	- Windows: The start command doesn't natively support full-screen, so you must specify the browser (e.g., chrome, msedge, firefox) and its flags. Replace "chrome" with the desired browser executable name.
	- macOS: The open -a command launches a specific browser with --args for flags. Safari is trickier since it lacks full-screen flags.
	- Linux: Directly invoke the browser (e.g., google-chrome or firefox) with flags, asруса

Limitations:
	- The default browser isn't directly controlled; you must specify a browser (e.g., Chrome, Firefox) to use its flags.
	- Some browsers or environments may not support full-screen flags (e.g., Safari on macOS).
	- If the specified browser isn't installed, the command will fail. You may need to check for browser availability or fall back to xdg-open on Linux.

Error Handling:
	- Add error handling for cases where the browser or command isn't found.
	- Example: Check if err != nil and log the issue.

Alternative Approach:
	- For more control, you could use a library like github.com/pkg/browser to open URLs in the default browser, but it doesn't support full-screen flags directly. You'd still need to use browser-specific commands for full-screen.

Testing:
	- Test with your target browser, as behavior varies (e.g., Chrome's --start-fullscreen vs. Firefox's -kiosk).
	- Ensure the browser is installed and the executable name matches (e.g., google-chrome vs. chromium on Linux).
	- If you need a specific browser or OS setup, let me know, and I can tailor the solution further!
*/

import (
	"log"
	"os/exec"
	"runtime"
)

func openBrowser(url string) error {
	log.Println("openBrowser: entering")
	var cmd string
	var args []string

	log.Println("openBrowser: runtime.GOOS = " + runtime.GOOS)

	switch runtime.GOOS {
	case "windows":
		// Windows: Use 'start' with browser-specific flags
		// Note: 'start' doesn't directly support full-screen, but we can try launching a specific browser
		cmd = "cmd"
		args = []string{"/c", "start", "chrome", url} // For Chrome
		// args = []string{"/c", "start", "chrome", "--start-fullscreen", url} // For Chrome
		// args = []string{"/c", "start", "chrome", "--kiosk", url} // works well - hard to close or escape.

		// For Edge: args = []string{"/c", "start", "msedge", "--start-fullscreen", url}
		// For Firefox: args = []string{"/c", "start", "firefox", "-kiosk", url}
	case "darwin": // macOS
		// macOS: Use 'open' with browser-specific flags
		cmd = "open"
		args = []string{"-a", "Google Chrome", "--args", "--start-fullscreen", url} // For Chrome
		// For Safari: args = []string{"-a", "Safari", url} // Safari doesn't support full-screen flag directly
		// For Firefox: args = []string{"-a", "Firefox", "--args", "-kiosk", url}
	default: // Linux, BSD, etc.
		// Linux: Use 'xdg-open' or specific browser with flags
		// cmd = "google-chrome"
		// args = []string{"--start-fullscreen", url} // For Chrome
		// For Firefox: cmd = "firefox"; args = []string{"-kiosk", url}

		cmd = "chromium"
		// args = []string{"--start-fullscreen", url} // For Chrome
		// args = []string{"--kiosk", url} // For Chrome
	}

	log.Println("openBrowser: cmd = ", cmd, ", args = ", args)
	return exec.Command(cmd, args...).Start()
}

func LaunchDefaultBrowser() {
	log.Println("LaunchDefaultBrowser: entering")
	err := openBrowser("http://localhost:81")
	// err := openBrowser("https://www.youtube.com/watch?v=ML_MyBtPh8A&t=322s")
	// err := openBrowser("https://www.youtube.com/watch?v=3NUp_RJ9JL0")
	// err := openBrowser("http://rpi2a:82/images/car.png")
	// err := openBrowser("https://example.com")
	if err != nil {
		panic(err)
	}
}
