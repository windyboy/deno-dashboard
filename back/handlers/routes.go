package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func StaticServe(e *echo.Echo) {
	// Serve static files from the public directory
	e.Static("/static", "public")
	e.Static("/assets", "public/assets")
	e.Static("/css", "public/css")
	e.Static("/js", "public/js")

	slog.Info("Static file server configured",
		"paths", []string{"/static", "/assets", "/css", "/js"},
		"directory", "public")
}

func WebSocketServer(e *echo.Echo) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins in development
		},
	}

	e.GET("/ws", func(c echo.Context) error {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}
		defer ws.Close()

		// WebSocket connection handler
		for {
			// Read message
			messageType, msg, err := ws.ReadMessage()
			if err != nil {
				slog.Error("Error reading message", "error", err)
				break
			}

			// Echo the message back
			if err := ws.WriteMessage(messageType, msg); err != nil {
				slog.Error("Error writing message", "error", err)
				break
			}
		}

		return nil
	})

	slog.Info("WebSocket server configured", "path", "/ws")
}
