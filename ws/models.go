package ws

import "github.com/gorilla/websocket"

type Websocket struct {
	EndPoint string
	conn     *websocket.Conn
}
