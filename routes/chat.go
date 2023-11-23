package routes

import (
	"chatting-app/internal/adapters/handlers"
	"chatting-app/internal/core/services"
	"chatting-app/pkg/websockets"

	"github.com/labstack/echo/v4"
)

func ChatRoutes(e *echo.Group) {

	webSocketManager := websockets.GetWebSocketManager()
	webSocketService := services.NewWebSocketService(webSocketManager)

	h := handlers.NewWebSocketHandler(webSocketService)

	e.GET("/", h.HomeHandler)
	e.GET("/ws", h.WebSocketHandler)
}
