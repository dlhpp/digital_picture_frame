package yaml

import (
	"fmt"
	"testing"

	"github.com/dlhpp/digital_picture_frame/logging"
	"github.com/dlhpp/digital_picture_frame/utils"
	"github.com/stretchr/testify/assert"
)

func Test_Get_various_flavors(t *testing.T) {

	testCases := []struct {
		resultType    int // 1=>string, 2=>int, 3=>map, 4=>intArray, 5=>stringArray
		path          string
		expectedValue any
		defaultValue  any
	}{
		// ----------------------- string -----------------------
		{1, "user.address.city", "Metropolis", "--not found--"},
		{1, "user.address.country", "Fictionland", "--not found--"},
		{1, "database.host", "localhost", "--not found--"},
		{1, "deeply.nested.structure.key", "Some nested string", "--not found--"},
		{1, "logging.level", "DEBUG", "--not found--"},
		{1, "multiple.0", "item1", "--not found--"},
		{1, "multiple.1", "item2", "--not found--"},

		// ----------------------- string NOT found -----------------------
		{1, "deeply.nested.structure.keyNotFound", "--not found--", "--not found--"},
		{1, "logging[\"level\"]", "--not found--", "--not found--"},
		{1, "multiple[0]", "--not found--", "--not found--"},
		{1, "multiple[1]", "--not found--", "--not found--"},

		// ----------------------- int -----------------------
		{2, "nested_numbers.multiple.0", 5, -1},
		{2, "nested_numbers.multiple.1", 7, -2},
		{2, "nested_numbers.multiple.2", 11, -3},

		// ----------------------- int NOT found -----------------------
		{2, "nested_numbers.multiple.3", -17, -17},

		// ----------------------- maps and arrays -----------------------
		{3, "logging", map[string]any{"file": "app.log", "level": "DEBUG"}, map[string]any{}},     // will be map with two fields
		{4, "nested_numbers.multiple", []int{5, 7, 11}, []int{11, 13, 15}},                        // array of int
		{5, "multiple", []string{"item1", "item2", "item3"}, []string{"item4", "item5", "item6"}}, // will be array of strings
	}

	cfg := OpenYamlFile("yamlConfigComplex.yaml")

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", idx), func(t *testing.T) {

			var expected, got any

			expected = tc.expectedValue

			// 1=>string, 2=>int, 3=>map, 4=>intArray, 5=>stringArray
			switch tc.resultType {
			case 1: // string
				got = GetString(cfg, tc.path, tc.defaultValue.(string))

			case 2: // int
				got = GetInt(cfg, tc.path, tc.defaultValue.(int))

			case 3: // map
				got = Get(cfg, tc.path, tc.defaultValue.(map[string]any))
				logging.Log("Test_Get_various_flavors", 3, utils.DescribeVariable("got", got))

			case 4: // intArray
				// got = Get(cfg, tc.path, tc.defaultValue.([]int))
				got = GetIntArray(cfg, tc.path, tc.defaultValue.([]int))
				logging.Log("Test_Get_various_flavors", 3, utils.DescribeVariable("got", got))

			case 5: // stringArray
				got = GetStringArray(cfg, tc.path, tc.defaultValue.([]string))
				logging.Log("Test_Get_various_flavors", 3, utils.DescribeVariable("got", got))

			default:
				t.Errorf("path=%s, unknown resultType=%d \n", tc.path, tc.resultType)
				return
			}

			logging.Log("Test_Get_various_flavors", 3, "DONE!", fmt.Sprintf("path=%s, expected=%+v, got=%+v \n", tc.path, expected, got))

			assert := assert.New(t)
			assert.Equal(expected, got, "expected value != got value \n")
		})
	}

}

func L(n string, v any) string {
	return fmt.Sprintf("%s = %+v", n, v)
}

func NoTest_BrowserConfig(t *testing.T) {

	logging.SetLevel(5)
	cfg := OpenYamlFile("config_browser.yaml")
	logging.Log("Test_BrowserConfig", 3, L("cfg", cfg))

	testCases := []struct {
		resultType    int // 1=>string, 2=>int, 3=>map, 4=>intArray, 5=>stringArray
		path          string
		expectedValue any
		defaultValue  any
	}{
		// ----------------------- "executable" string -----------------------
		{1, "windows.chrome.executable", "cmd", "--not found--"},
		{1, "windows.firefox.executable", "cmd", "--not found--"},

		{1, "linux.chromium.executable", "chromium-browser", "--not found--"},
		{1, "linux.firefox.executable", "firefox", "--not found--"},
		{1, "linux.epiphany.executable", "epiphany", "--not found--"},
		{1, "linux.midori.executable", "midori", "--not found--"},

		{1, "darwin.firefox.executable", "firefox", "--not found--"},
		{1, "darwin.safari.executable", "open", "--not found--"},

		// ----------------------- "args" []string -----------------------
		{5, "windows.chrome.args", []string{"/c", "start", "chrome", "--new-window", "--full-screen"}, []string{"--not found--"}},
		{5, "windows.firefox.args", []string{"/c", "start", "firefox", "--new-window", "--full-screen"}, []string{"--not found--"}},

		{5, "linux.chromium.args", []string{"--new-window", "--full-screen"}, []string{"--not found--"}},
		{5, "linux.firefox.args", []string{"--new-window", "--full-screen"}, []string{"--not found--"}},
		{5, "linux.epiphany.args", []string{"--new-window", "--fullscreen"}, []string{"--not found--"}},
		{5, "linux.midori.args", []string{"--fullscreen"}, []string{"--not found--"}},

		{5, "darwin.firefox.args", []string{"--new-window", "--full-screen"}, []string{"--not found--"}},
		{5, "darwin.safari.args", []string{"-a", "Safari"}, []string{"--not found--"}},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", idx), func(t *testing.T) {

			var expected, got any

			expected = tc.expectedValue

			// 1=>string, 2=>int, 3=>map, 4=>intArray, 5=>stringArray
			switch tc.resultType {
			case 1: // string
				got = GetString(cfg, tc.path, tc.defaultValue.(string))

			case 2: // int
				got = GetInt(cfg, tc.path, tc.defaultValue.(int))

			case 3: // map
				got = Get(cfg, tc.path, tc.defaultValue.(map[string]any))
				logging.Log("Test_BrowserConfig", 3, utils.DescribeVariable("got", got))

			case 4: // intArray
				// got = Get(cfg, tc.path, tc.defaultValue.([]int))
				got = GetIntArray(cfg, tc.path, tc.defaultValue.([]int))
				logging.Log("Test_BrowserConfig", 3, utils.DescribeVariable("got", got))

			case 5: // stringArray
				got = GetStringArray(cfg, tc.path, tc.defaultValue.([]string))
				logging.Log("Test_BrowserConfig", 3, utils.DescribeVariable("got", got))

			default:
				t.Errorf("path=%s, unknown resultType=%d \n", tc.path, tc.resultType)
				return
			}

			assert := assert.New(t)
			assert.Equal(expected, got, L("path", tc.path))
		})
	}
}
