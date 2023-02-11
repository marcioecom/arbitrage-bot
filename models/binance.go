package models

type BinanceTicker struct {
	Event            string `json:"e"`
	Time             int64  `json:"E"`
	Symbol           string `json:"s"`
	Change           string `json:"p"`
	Percentage       string `json:"P"`
	WeightedAvgPrice string `json:"w"`
	PrevClosePrice   string `json:"x"`
	Close            string `json:"c"`
	CloseQty         string `json:"Q"`
	BidPrice         string `json:"b"`
	BidQty           string `json:"B"`
	AskPrice         string `json:"a"`
	AskQty           string `json:"A"`
	Open             string `json:"o"`
	High             string `json:"h"`
	Low              string `json:"l"`
	BaseVolume       string `json:"v"`
	QuoteVolume      string `json:"q"`
	OpenTime         int64  `json:"O"`
	CloseTime        int64  `json:"C"`
	FirstID          int64  `json:"F"`
	LastID           int64  `json:"L"`
	Count            int64  `json:"n"`
}
