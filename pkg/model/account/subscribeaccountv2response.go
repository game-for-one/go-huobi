package account

import "github.com/game_for_one/go-huobi/pkg/model/base"

type SubscribeAccountV2Response struct {
	base.WebSocketV2ResponseBase
	Data *struct {
		Currency    string `json:"currency"`
		AccountId   int    `json:"accountId"`
		Balance     string `json:"balance"`
		Available   string `json:"available"`
		ChangeType  string `json:"changeType"`
		AccountType string `json:"accountType"`
		ChangeTime  int64  `json:"changeTime"`
	}
}
