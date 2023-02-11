package strategy

import (
	"github.com/marcioecom/arbitrage-bot/exchange"
	"github.com/marcioecom/arbitrage-bot/helper"
	"github.com/marcioecom/arbitrage-bot/models"
	"github.com/marcioecom/arbitrage-bot/ws"
	"go.uber.org/zap"
)

func New(ws ws.IWebsocket, ex exchange.IExchange) IStrategy {
	return &Strategy{
		ws:       ws,
		exchange: ex,
	}
}

func (s *Strategy) Start() {
	s.ws.ListenMessages(func(bt models.BinanceTicker) {
		ticker := helper.ParseBinanceTicker(bt)
		zap.L().Info("handler", zap.Any("ticker", ticker))
	})
}

func (s *Strategy) Stop() {
	panic("unimplemented")
}
