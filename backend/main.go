package main

import (
	"fmt"
	"github.com/exceed19-cpsk/backend-bubblebungbung/config"
	"github.com/exceed19-cpsk/backend-bubblebungbung/handler"
	"github.com/exceed19-cpsk/backend-bubblebungbung/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/exp/slices"
	"log"
	"net/http"
)

var appConfig config.Config

func init() {
	appConfig = config.Load()
}

func initUpgrader() websocket.Upgrader {
	var upgrader = websocket.Upgrader{
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			log.Println(origin, "trying to connect to websocket")
			if len(appConfig.ALLOW_ORIGINS) == 0 {
				return true
			}
			return slices.Contains(appConfig.ALLOW_ORIGINS, origin)
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

	messageHandler := handler.NewMessageHandler(hub)

	api.GET("/ws", func(c *gin.Context) {
		handler.ServeWs(hub, c.Writer, c.Request, upgrader)
	})
	api.POST("/message", messageHandler.SendMessage)
	server.Run(fmt.Sprint(":", appConfig.LISTENING_PORT))
}
