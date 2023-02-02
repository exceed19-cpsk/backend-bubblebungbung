package handler

import (
	"fmt"
	"github.com/exceed19-cpsk/backend-bubblebungbung/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageHandler struct {
	wsHub *service.WsHub
}

type MessageRequestBody struct {
	Message string
}

func NewMessageHandler(wsHub *service.WsHub) *MessageHandler {
	return &MessageHandler{
		wsHub: wsHub,
	}
}

func (m MessageHandler) SendMessage(g *gin.Context) {
	var requestBody MessageRequestBody

	if err := g.BindJSON(&requestBody); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "incorrect body format",
		})
	}

	fmt.Println(requestBody.Message)
	m.wsHub.Broadcast <- []byte(requestBody.Message)
}
