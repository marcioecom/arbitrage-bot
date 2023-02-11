package ws

import "github.com/marcioecom/arbitrage-bot/models"

type IWebsocket interface {
	Init() error
	Close() error
	ListenMessages(handler models.HandlerTicker)
}
