package models

type OBEntry struct {
	Price  float64
	Amount float64
}

type OrderBook struct {
	LastUpdateID int64
	Bids         []OBEntry
	Asks         []OBEntry
}
