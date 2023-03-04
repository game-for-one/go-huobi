package main

import (
	"github.com/game_for_one/go-huobi/config"
	"github.com/game_for_one/go-huobi/logging/applogger"
	"github.com/game_for_one/go-huobi/pkg/client"
)

func getSystemStatus() {
	client := new(client.CommonClient).Init(config.Host)
	resp, err := client.GetSystemStatus()
	if err != nil {
		applogger.Error("Get system status error: %s", err)
	} else {
		applogger.Info("Get system status %s", resp)
	}
}

func main() {
	getSystemStatus()
}