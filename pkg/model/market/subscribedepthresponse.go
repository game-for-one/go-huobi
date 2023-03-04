package market

import (
	"github.com/game-for-one/go-huobi/pkg/model/base"
)

type SubscribeDepthResponse struct {
	base.WebSocketResponseBase
	Data *Depth
	Tick *Depth
}
