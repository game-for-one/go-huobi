package client

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/game-for-one/go-huobi/internal"
	"github.com/game-for-one/go-huobi/internal/requestbuilder"
	"github.com/game-for-one/go-huobi/pkg/model"
	"github.com/game-for-one/go-huobi/pkg/model/stablecoin"
	"github.com/valyala/fasthttp"
)

// Responsible to operate wallet
type StableCoinClient struct {
	httpCli           *fasthttp.Client
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

// Initializer
func (p *StableCoinClient) Init(accessKey string, secretKey string, host string, httpCli *fasthttp.Client) *StableCoinClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	p.httpCli = httpCli
	return p
}

// Get stable coin exchange rate
func (p *StableCoinClient) GetExchangeRate(currency string, amount string, exchangeType string) (*stablecoin.GetExchangeRateResponse, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("currency", currency)
	request.AddParam("amount", amount)
	request.AddParam("type", exchangeType)

	url := p.privateUrlBuilder.Build("GET", "/v1/stable-coin/quote", request)
	getResp, getErr := internal.HttpGet(p.httpCli, url)
	if getErr != nil {
		return nil, getErr
	}

	result := stablecoin.GetExchangeRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {
		return &result, nil
	}
	return nil, errors.New(getResp)
}

// Exchange stable coin
func (p *StableCoinClient) ExchangeStableCoin(quoteId string) (*stablecoin.ExchangeStableCoinResponse, error) {
	postBody := fmt.Sprintf("{ \"quote-id\": \"%s\"}", quoteId)

	url := p.privateUrlBuilder.Build("POST", "/v1/stable-coin/exchange", nil)
	postResp, postErr := internal.HttpPost(p.httpCli, url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := stablecoin.ExchangeStableCoinResponse{}
	jsonErr := json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if result.Status == "ok" && result.Data != nil {
		return &result, nil
	}
	return nil, errors.New(postResp)
}
