package main

import (
	"fmt"
	"net/http"

	"github.com/exceed19-cpsk/backend-bubblebungbung/config"
	"github.com/exceed19-cpsk/backend-bubblebungbung/handler"
	"github.com/exceed19-cpsk/backend-bubblebungbung/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var appConfig config.Config

func init() {
	appConfig = config.Load()
}

func initUpgrader() websocket.Upgrader {
	var upgrader = websocket.Upgrader{
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return upgrader
}

func main() {
	server := gin.Default()
	api := server.Group(appConfig.PROXY_URL)
	hub := service.NewWsHub()
	upgrader := initUpgrader()

	go hub.Run()

	messageHandler := handler.NewMessageHandler(hub, appConfig)
	api.GET("/ws", func(c *gin.Context) {
		handler.ServeWs(hub, c.Writer, c.Request, upgrader)
	})
	api.GET("/message/:message", messageHandler.SendMessageGet)
	api.GET("/message", messageHandler.SendQueryMessageGet)
	api.POST("/message", messageHandler.SendMessagePost)
	api.PUT("/message", messageHandler.SendMessagePut)
	api.DELETE("/message", messageHandler.SendMessageDelete)
	server.Run(fmt.Sprint(":", appConfig.LISTENING_PORT))
}
