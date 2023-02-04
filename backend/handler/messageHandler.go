package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/exceed19-cpsk/backend-bubblebungbung/service"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	wsHub *service.WsHub
}

type MessageRequestBody struct {
	Message string
	Color   string
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
		return
	}

	message, err := json.Marshal(requestBody)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "error encoding message",
		})
		return
	}

	fmt.Println(message)
	m.wsHub.Broadcast <- message
}
