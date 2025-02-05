package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	for {
		time.Sleep(2 * time.Second) // Simulate periodic updates
		msg := fmt.Sprintf("Server Time: %s", time.Now().Format("15:04:05"))
		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			break
		}
	}
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		// slog.Error("Error loading .env file", "err", err)
		log.Fatal(err)

	}
	// router := chi.NewRouter()
	// router.Handle("/*", public())
	// router.Get("/", handlers.Process(handlers.HandleHome))
	// fmt.Println("Hello, world!")
	listenAddr := os.Getenv("LISTEN_ADDR")
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Static("/public", "public")
	e.GET("/ws", wsHandler)
	// handlers.SetupRoutes(e)
	// slog.Info("Starting server...", "listenAddr", listenAddr)
	// if err := http.ListenAndServe(listenAddr, router); err != nil {
	// 	slog.Error("Error starting server", "err", err)
	// 	log.Fatal(err)
	// }
	// nats := NatsHandler{}
	// go nats.Subscribe()

	e.Logger.Fatal(e.Start(listenAddr))
}
