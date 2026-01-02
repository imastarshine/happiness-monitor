package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Site struct {
	URL  string
	Name string
}

type NetError struct {
	FindTemplate string
	Text         string
}

// TODO: добавить основные ошибки в короткий формат, для удобства. Сделать через массив и RegEXP
// TODO добавить считалку сколько времени занял запрос

func padEnd(s string, length int, char string) string {
	if len(s) >= length {
		return s
	}
	return s + strings.Repeat(char, length-len(s))
}

func monitor() {
	for {
		sites := []Site{
			{URL: "https://google.com", Name: "Google"},
			{URL: "https://github.com", Name: "Github"},
			{URL: "https://ya.ru", Name: "Yandex"},
			{URL: "https://vk.ru", Name: "VK"},
		}

		if !wifiIsOk() {
			loggerError("No Wi-Fi Connection!")
			dbInsertInformation("localhost", "No Wifi", 0)
			time.Sleep(60 * time.Second)
		}

		for _, site := range sites {
			url := site.URL
			name := site.Name

			status, _, elapsedTime, checkErr := checkNetwork(url)
			elapsedString := "(" + strconv.FormatInt(elapsedTime, 10) + "ms)"
			finalStatus := status

			if checkErr == nil {
				line := padEnd(name, 15, ".") + " " + padEnd(elapsedString, 10, ".") + " " + padEnd(url, 40, ".") + " " + status
				loggerLog(line)
			} else {
				finalStatus = checkErr.Error()
				line := padEnd(name, 15, ".") + " " + padEnd(elapsedString, 10, ".") + " " + padEnd(url, 40, ".") + " ERROR: " + finalStatus
				loggerError(line)
			}

			dbInsertInformation(url, finalStatus, elapsedTime)
		}
		loggerLog(strings.Repeat("-", 75))
		time.Sleep(60 * time.Second)
	}
}

func main() {
	fmt.Println("Starting program...")
	// initialization
	loggerMain()
	dbMain()
	dbClearOldInformation()

	monitor()
	//for {
	//	time.Sleep(1 * time.Hour)
	//}
}
