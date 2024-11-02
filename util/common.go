package util

import (
	"runtime"
)

func GetFunctionName() string {
	counter, _, _, success := runtime.Caller(1)

	if !success {
		return ""
	}

	return runtime.FuncForPC(counter).Name()
}
