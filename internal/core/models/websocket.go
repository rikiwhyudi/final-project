package models

import "github.com/gorilla/websocket"

type WebSocketConnection struct {
	*websocket.Conn
	Username string
}
