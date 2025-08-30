package internal

import (
	"fmt"
	"log"
)

var GlobalLevel int = 5 // 1 ==> most verbose, 9 ==> least verbose

func DLHLog(method string, level int, args ...any) {

	if level < GlobalLevel {
		return
	}

	concatenatedArgs := ""
	for _, singleArg := range args {
		if concatenatedArgs != "" {
			concatenatedArgs += ", "
		}
		concatenatedArgs += fmt.Sprintf("%+v", singleArg)
	}

	// later i'll implement ability for different output formats
	// along with named logging levels.
	log.Printf("%s: %+v", method, concatenatedArgs)
}

func DLHLogSetLevel(l int) {
	GlobalLevel = l
}

func DLHLogGetLevel(l int) int {
	return GlobalLevel
}

// func main() {
// 	DLHLog("main", 1, "one", 2, 3.0, true, []string{"a", "b", "c"}, map[string]int{"x": 1, "y": 2})
// 	DLHLog("main", 2, "one", 2, 3.0, true, []string{"a", "b", "c"}, map[string]int{"x": 1, "y": 2})
// 	DLHLog("main", 3, "one", 2, 3.0, true, []string{"a", "b", "c"}, map[string]int{"x": 1, "y": 2})
// 	DLHLog("main", 4, "one", 2, 3.0, true, []string{"a", "b", "c"}, map[string]int{"x": 1, "y": 2})
// 	DLHLog("main", 5, "one", 2, 3.0, true, []string{"a", "b", "c"}, map[string]int{"x": 1, "y": 2})
// 	DLHLog("main", 6, "one", 2, 3.0, true, []string{"a", "b", "c"}, map[string]int{"x": 1, "y": 2})
// 	DLHLog("main", 7, "one", 2, 3.0, true, []string{"a", "b", "c"}, map[string]int{"x": 1, "y": 2})
// 	DLHLog("main", 8, "one", 2, 3.0, true, []string{"a", "b", "c"}, map[string]int{"x": 1, "y": 2})
// 	DLHLog("main", 9, "one", 2, 3.0, true, []string{"a", "b", "c"}, map[string]int{"x": 1, "y": 2})
// }
