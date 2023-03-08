package internal

import (
	"fmt"

	"github.com/game-for-one/go-huobi/logging/perflogger"
	"github.com/valyala/fasthttp"
)

var HttpClient = &fasthttp.Client{}

func HttpGet(url string) (string, error) {
	logger := perflogger.GetInstance()
	logger.Start()

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	req.Header.SetMethod("GET")
	req.Header.SetRequestURI(url)

	err := HttpClient.Do(req, res)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)
	if err != nil {
		fmt.Printf("Error making request: %e", err)
		return "", err
	}

	logger.StopAndLog("GET", url)

	return string(res.Body()), err
}

func HttpPost(url string, body string) (string, error) {
	logger := perflogger.GetInstance()
	logger.Start()

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.Header.SetRequestURI(url)
	req.SetBodyString(body)

	err := HttpClient.Do(req, res)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)
	if err != nil {
		fmt.Printf("Error making request: %e", err)
		return "", err
	}

	logger.StopAndLog("POST", url)

	return string(res.Body()), err
}
