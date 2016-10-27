package xlog

import (
        "fmt"
        "log"
        "strings"
)

// _, filename, _, _ := runtime.Caller(0)
func Error(filepath string, args ...interface{}) {
        log.Printf(generateOutputFormat("ERROR", filepath, len(args)), args...)
}

func Info(filepath string, args ...interface{}) {
        log.Printf(generateOutputFormat("INFO", filepath, len(args)), args...)
}

func Debug(filepath string, args ...interface{}) {
        log.Printf(generateOutputFormat("DEBUG", filepath, len(args)), args...)
}

func Warn(filepath string, args ...interface{}) {
        log.Printf(generateOutputFormat("WARN", filepath, len(args)), args...)
}

func generateOutputFormat(level string, path string, length int) string {
        result := fmt.Sprintf("[%s][%s]", level, path)
        result += strings.Repeat(" %v", length)

        return result
}
