package market

import (
	"github.com/game-for-one/go-huobi/pkg/model/base"
)

type SubscribeLast24hCandlestickResponse struct {
	base.WebSocketResponseBase
	Data *Candlestick
	Tick *Candlestick
}
