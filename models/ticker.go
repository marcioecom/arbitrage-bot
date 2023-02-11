package models

import "time"

type HandlerTicker = func(BinanceTicker) error

type Ticker struct {
	FirstID     int64
	LastID      int64
	Count       int64
	Ask         float64
	AskVolume   float64
	Bid         float64
	BidVolume   float64
	High        float64
	Low         float64
	Average     float64
	BaseVolume  float64
	QuoteVolume float64
	Open        float64
	Last        float64
	Change      float64
	Percentage  float64
	Amount      float64
	Symbol      string
	OpenTime    time.Time
	CloseTime   time.Time
	Timestamp   time.Time
}
