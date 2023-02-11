package ws

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/marcioecom/arbitrage-bot/models"
	"go.uber.org/zap"
)

func New(endpoint string) IWebsocket {
	return &Websocket{
		EndPoint: endpoint,
	}
}

func (w *Websocket) Init() error {
	dialer := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
	}

	conn, res, err := dialer.Dial(w.EndPoint, nil)
	if err != nil {
		zap.L().Error("error dialing", zap.Int("statusCode", res.StatusCode), zap.Error(err))
		return err
	}
	w.conn = conn

	zap.L().Info("connected", zap.String("status", res.Status))
	return nil
}

func (w *Websocket) Close() error {
	if err := w.conn.Close(); err != nil {
		zap.L().Error("error closing connection", zap.Error(err))
		return err
	}
	return nil
}

func (w *Websocket) ListenMessages(handler models.HandlerTicker) {
	go func() {
		for {
			_, msg, err := w.conn.ReadMessage()
			if err != nil {
				zap.L().Error("error reading message", zap.Error(err))
				continue
			}

			var bt models.BinanceTicker
			if err := json.Unmarshal(msg, &bt); err != nil {
				zap.L().Error("error unmarshalling message", zap.Error(err))
				continue
			}

			handler(bt)
		}
	}()
}
