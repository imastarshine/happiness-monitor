package main

import (
	"fmt"
	"os"
	"time"
)

var logFilename = ""

func appenToFile(message string) {
	file, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != err {
		fmtError("an error occurred for file: " + logFilename + " with err: " + err.Error())
	}
	defer file.Close()

	_, wErr := file.WriteString(message + "\n")
	if wErr != nil {
		fmtError("an error occurred for writing string in file: " + logFilename + " with err: " + wErr.Error())
	}
}

func fmtError(logStr string) {
	now := time.Now()
	timeString := now.Format("2006-01-02 15:04:05.000")

	logString := fmt.Sprintf("%s UTC+00 [LOG] :: %s", timeString, logStr)
	fmt.Println(logString)
}

func loggerLog(logStr string) {
	now := time.Now()
	timeString := now.Format("2006-01-02 15:04:05.000")

	logString := fmt.Sprintf("%s UTC+00 [LOG] :: %s", timeString, logStr)
	fmt.Println(logString)
	appenToFile(logString)
}

func loggerError(logStr string) {
	now := time.Now()
	timeString := now.Format("2006-01-02 15:04:05.000")

	logString := fmt.Sprintf("%s UTC+00 [ERROR] :: %s", timeString, logStr)
	fmt.Println(logString)
}

func loggerMain() {
	err := os.MkdirAll("logs", 0o755)
	if err != nil {
		fmtError("an error occurred on making dir logs: " + err.Error())
	}
	now := time.Now()
	logFilename = "logs/" + now.Format("2006-01-02") + ".log"
}
