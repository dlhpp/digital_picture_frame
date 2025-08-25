package internal

import (
	"runtime"
	"strconv"
)

// NOTHING HERE IS CURRENTLY USED BUT MAY BE USEFUL LATER

func GetCurrentFuncNameLineNumber() string {

	/*
		This line:    slog.Info("indexHandler: entering", "func", getCurrentFuncNameLineNumber())
		produces:     msg="indexHandler: entering" func=main.(*ImageStore).indexHandler:49

		Explanation:
		runtime.Caller(1):
		This function returns the:
			program counter (PC),
			file name,
			line number,
			boolean indicating success for the caller of the current function.

		Passing 1 as an argument means we're interested in the stack frame of the
		function that called getCurrentFuncName.

		runtime.FuncForPC(pc):
		This function takes a program counter (PC) and returns a *runtime.Func object,
		which represents the function at that PC.

		fn.Name():
		The Name() method of the *runtime.Func object returns the fully qualified
		name of the function (e.g., main.myExampleFunction).
	*/

	pc, _, lineNumber, ok := runtime.Caller(1)
	if !ok {
		return "unknown-1"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown-2"
	}
	return fn.Name() + ":" + strconv.Itoa(lineNumber)
}
