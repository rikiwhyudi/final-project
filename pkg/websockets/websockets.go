package websockets

import (
	"chatting-app/internal/core/models"
)

type WebSocketManager struct {
	Connections []*models.WebSocketConnection
}

var instance *WebSocketManager

func GetWebSocketManager() *WebSocketManager {
	if instance == nil {
		instance = &WebSocketManager{}
	}
	return instance
}

func (m *WebSocketManager) EjectConnection(connToRemove *models.WebSocketConnection) {

	var updatedConnections []*models.WebSocketConnection
	for _, conn := range m.Connections {
		if conn != connToRemove {
			updatedConnections = append(updatedConnections, conn)
		}
	}

	m.Connections = updatedConnections
}

func (m *WebSocketManager) BroadcastMessage(connections []*models.WebSocketConnection, currentConn *models.WebSocketConnection, kind, message string) {
	for _, eachConn := range connections {
		if eachConn == currentConn {
			continue
		}

		eachConn.WriteJSON(models.SocketResponse{
			From:    currentConn.Username,
			Type:    kind,
			Message: message,
		})
	}
}
