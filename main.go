package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/marcioecom/arbitrage-bot/helper"
	"github.com/marcioecom/arbitrage-bot/models"
	"github.com/marcioecom/arbitrage-bot/ws"
	"go.uber.org/zap"
)

func main() {
	helper.InitLogger()

	wsclient := ws.New("wss://stream.binance.com:9443/ws/btcusdt@ticker")

	if err := wsclient.Init(); err != nil {
		zap.L().Fatal("error initializing ws", zap.Error(err))
	}

	wsclient.ListenMessages(func(bt models.BinanceTicker) error {
		ticker := helper.ParseBinanceTicker(bt)
		zap.L().Info("handler", zap.Any("ticker", ticker))
		return nil
	})

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	zap.L().Info("stopping service...")
	go func() {
		time.Sleep(time.Second * 5)
		zap.L().Panic("stop timeout, killed")
	}()

	if err := wsclient.Close(); err != nil {
		zap.L().Fatal("error closing ws", zap.Error(err))
	}
}
