package exchange

import "github.com/marcioecom/arbitrage-bot/models"

type IExchange interface {
	// GetTicker returns the ticker for the given symbol
	GetTicker(symbol string) (models.Ticker, error)
	// GetOrderBook returns the order book for the given symbol
	FetchOrderBook(symbol string) (models.OrderBook, error)
}
