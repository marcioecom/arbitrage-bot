package exchange

import "github.com/marcioecom/arbitrage-bot/models"

func New() IExchange {
	return &Exchange{}
}

func (x *Exchange) FetchOrderBook(symbol string) (models.OrderBook, error) {
	panic("unimplemented")
}

func (x *Exchange) GetTicker(symbol string) (models.Ticker, error) {
	panic("unimplemented")
}
