package strategy

import (
	"sync"

	"github.com/marcioecom/arbitrage-bot/exchange"
	"github.com/marcioecom/arbitrage-bot/models"
	"github.com/marcioecom/arbitrage-bot/ws"
)

type Strategy struct {
	obMutex    sync.RWMutex
	orderbooks map[string]models.OrderBook
	ws         ws.IWebsocket
	exchange   exchange.IExchange
}
