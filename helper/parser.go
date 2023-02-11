package helper

import (
	"strconv"
	"time"

	"github.com/marcioecom/arbitrage-bot/models"
)

func ParseBinanceTicker(bticker models.BinanceTicker) models.Ticker {
	return models.Ticker{
		Symbol:      bticker.Symbol,
		Ask:         parseToFloat64(bticker.AskPrice),
		AskVolume:   parseToFloat64(bticker.AskQty),
		Bid:         parseToFloat64(bticker.BidPrice),
		BidVolume:   parseToFloat64(bticker.BidQty),
		High:        parseToFloat64(bticker.High),
		Low:         parseToFloat64(bticker.Low),
		Average:     parseToFloat64(bticker.WeightedAvgPrice),
		BaseVolume:  parseToFloat64(bticker.BaseVolume),
		QuoteVolume: parseToFloat64(bticker.QuoteVolume),
		Open:        parseToFloat64(bticker.Open),
		Last:        parseToFloat64(bticker.Close),
		Change:      parseToFloat64(bticker.Change),
		Percentage:  parseToFloat64(bticker.Percentage),
		Amount:      parseToFloat64(bticker.CloseQty),
		OpenTime:    time.Unix(0, bticker.OpenTime*1000000).UTC(),
		CloseTime:   time.Unix(0, bticker.CloseTime*1000000).UTC(),
		FirstID:     bticker.FirstID,
		LastID:      bticker.LastID,
		Count:       bticker.Count,
		Timestamp:   time.Unix(0, bticker.Time*1000000).UTC(),
	}
}

func parseToFloat64(value string) float64 {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return f
}
