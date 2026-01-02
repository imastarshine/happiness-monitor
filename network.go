package main

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"time"
)

func wifiIsOk() bool {
	result, err := exec.Command("termux-wifi-connectioninfo").Output()
	if err != nil {
		loggerError("Something went wrong on wifi checking: " + err.Error())
	}

	var data struct {
		IP string `json:"ip"`
	}

	json.Unmarshal(result, &data)
	if data.IP == "" || data.IP == "0.0.0.0" {
		return false
	}
	return true
}

func checkNetwork(url string) (string, int, int64, error) {
	startTime := time.Now()
	response, err := http.Get(url)
	endTime := time.Since(startTime).Milliseconds()
	if err != nil {
		return "", 0, endTime, err
	}
	defer response.Body.Close()
	return response.Status, response.StatusCode, endTime, nil
}
