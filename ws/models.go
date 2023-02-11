package ws

import "github.com/gorilla/websocket"

type Websocket struct {
	endPoint string
	conn     *websocket.Conn
}
