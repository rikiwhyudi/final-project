package handlers

import (
	"chatting-app/internal/core/ports"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type webSocketHandlerImpl struct {
	webSocketService ports.WebSocketService
}

func NewWebSocketHandler(webSocketService ports.WebSocketService) ports.WebSocketHandlers {
	return &webSocketHandlerImpl{webSocketService}
}

func (h *webSocketHandlerImpl) HomeHandler(c echo.Context) error {
	content, err := os.ReadFile("view/index.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Could not open requested file")
	}

	return c.HTMLBlob(http.StatusOK, content)
}

func (h *webSocketHandlerImpl) WebSocketHandler(c echo.Context) error {
	currentGorillaConn, err := websocket.Upgrade(c.Response(), c.Request(), c.Response().Header(), 1024, 1024)
	if err != nil {
		return c.String(http.StatusBadRequest, "Could not open WebSocket connection")
	}

	username := c.Request().URL.Query().Get("username")
	h.webSocketService.WebSocketConnection(currentGorillaConn, username)

	return nil

}
