package ports

import (
	"chatting-app/internal/core/models"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WebSocketService interface {
	WebSocketConnection(conn *websocket.Conn, username string)
	WebSocketService(currentConn *models.WebSocketConnection, connection []*models.WebSocketConnection)
}

type WebSocketHandlers interface {
	HomeHandler(c echo.Context) error
	WebSocketHandler(c echo.Context) error
}
