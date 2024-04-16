package logger

import (
	"fmt"
	"runtime"
)

func Info(msg ...interface{}) {
	output := fmt.Sprintf("\033[1;32m[INFO]\033[0m%s: %s", getCallerName(1), strConstructor(msg...))
	fmt.Println(output)
}

func Error(msg ...interface{}) {
	output := fmt.Sprintf("\033[1;31m[ERROR]\033[0m%s: %s", getCallerName(1), strConstructor(msg...))
	fmt.Println(output)
}

func Debug(msg ...interface{}) {
	output := fmt.Sprintf("\033[1;34m[DEBUG]\033[0m%s: %s", getCallerName(1), strConstructor(msg...))
	fmt.Println(output)
}

func Warn(msg ...interface{}) {
	output := fmt.Sprintf("\033[1;33m[WARN]\033[0m%s: %s", getCallerName(1), strConstructor(msg...))
	fmt.Println(output)
}

func strConstructor(msg ...interface{}) string {
	var output string
	for _, v := range msg {
		output += fmt.Sprintf("%v", v)
	}
	return output
}

func getCallerName(times int) string {
	pc, _, _, _ := runtime.Caller(times + 1)
	return runtime.FuncForPC(pc).Name()
}
