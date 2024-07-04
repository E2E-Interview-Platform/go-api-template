package helpers

import (
	"fmt"
	"runtime"
)

func GetCallerInfo() string {
	pc, file, line, ok := runtime.Caller(2)
	var funcName string
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}

	file = GetSuffixJoinedAfterSplit(file, "/", 2)
	funcName = GetSuffixJoinedAfterSplit(funcName, "/", 1)
	funcPath := fmt.Sprintf("%s:%d %s", file, line, funcName)

	return funcPath
}
