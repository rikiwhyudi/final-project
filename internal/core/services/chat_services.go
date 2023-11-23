package services

import (
	"chatting-app/internal/core/models"
	"chatting-app/internal/core/ports"
	"chatting-app/pkg/websockets"

	"log"
	"strings"

	"github.com/gorilla/websocket"
)

type webSocketServiceImpl struct {
	webSocketManager *websockets.WebSocketManager
}

func NewWebSocketService(webSocketManager *websockets.WebSocketManager) ports.WebSocketService {
	return &webSocketServiceImpl{webSocketManager}
}

func (s *webSocketServiceImpl) WebSocketConnection(conn *websocket.Conn, username string) {

	currentConn := models.WebSocketConnection{Conn: conn, Username: username}
	s.webSocketManager.Connections = append(s.webSocketManager.Connections, &currentConn)

	go s.WebSocketService(&currentConn, s.webSocketManager.Connections)

}

func (s *webSocketServiceImpl) WebSocketService(currentConn *models.WebSocketConnection, connection []*models.WebSocketConnection) {

	const (
		MESSAGE_NEW_USER = "New User"
		MESSAGE_CHAT     = "Chat"
		MESSAGE_LEAVE    = "Leave"
	)

	websockets.GetWebSocketManager().BroadcastMessage(s.webSocketManager.Connections, currentConn, MESSAGE_NEW_USER, "")

	for {
		payload := models.SocketPayload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				websockets.GetWebSocketManager().BroadcastMessage(s.webSocketManager.Connections, currentConn, MESSAGE_LEAVE, "")
				websockets.GetWebSocketManager().EjectConnection(currentConn)
				return
			}

			log.Println("ERROR", err.Error())
			continue
		}

		websockets.GetWebSocketManager().BroadcastMessage(s.webSocketManager.Connections, currentConn, MESSAGE_CHAT, payload.Message)
	}
}
