package main

import (
	"github.com/game-for-one/go-huobi/config"
	"github.com/game-for-one/go-huobi/logging/applogger"
	"github.com/game-for-one/go-huobi/pkg/client"
	"github.com/valyala/fasthttp"
)

func getSystemStatus() {
	client := new(client.CommonClient).Init(config.Host, &fasthttp.Client{})
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
