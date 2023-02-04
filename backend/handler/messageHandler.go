package handler

import (
	"encoding/json"
	"github.com/exceed19-cpsk/backend-bubblebungbung/config"
	"net/http"
	"strconv"

	"github.com/exceed19-cpsk/backend-bubblebungbung/service"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	wsHub     *service.WsHub
	appConfig config.Config
}

type MessageRequestBody struct {
	Message string
	Color   string
}

func NewMessageHandler(wsHub *service.WsHub, appConfig config.Config) *MessageHandler {
	return &MessageHandler{
		wsHub:     wsHub,
		appConfig: appConfig,
	}
}

func (m MessageHandler) SendMessagePost(g *gin.Context) {
	var requestBody MessageRequestBody

	if err := g.BindJSON(&requestBody); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "incorrect body format",
		})
		return
	}
	requestBody.Color = "blue"
	if !m.isMessageValid(requestBody.Message) {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "message size exceed the limit",
		})
	}
	message, err := json.Marshal(requestBody)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "error encoding message",
		})
		return
	}

	m.wsHub.Broadcast <- message
}

func (m MessageHandler) SendMessagePut(g *gin.Context) {
	var requestBody MessageRequestBody

	if err := g.BindJSON(&requestBody); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "incorrect body format",
		})
		return
	}
	requestBody.Color = "green"
	if !m.isMessageValid(requestBody.Message) {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "message size exceed the limit",
		})
	}
	message, err := json.Marshal(requestBody)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "error encoding message",
		})
		return
	}

	m.wsHub.Broadcast <- message
}

func (m MessageHandler) SendMessageDelete(g *gin.Context) {
	var requestBody MessageRequestBody

	if err := g.BindJSON(&requestBody); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "incorrect body format",
		})
		return
	}
	requestBody.Color = "red"
	if !m.isMessageValid(requestBody.Message) {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "message size exceed the limit",
		})
	}
	message, err := json.Marshal(requestBody)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "error encoding message",
		})
		return
	}

	m.wsHub.Broadcast <- message
}

func (m MessageHandler) SendMessageGet(g *gin.Context) {
	var requestBody MessageRequestBody
	paramMessage := g.Param("message")
	requestBody.Message = paramMessage
	requestBody.Color = "yellow"
	if !m.isMessageValid(requestBody.Message) {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "message size exceed the limit",
		})
	}

	message, err := json.Marshal(requestBody)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "error encoding message",
		})
		return
	}

	m.wsHub.Broadcast <- message
}

func (m MessageHandler) SendQueryMessageGet(g *gin.Context) {
	var requestBody MessageRequestBody
	paramMessage := g.Query("message")
	requestBody.Message = paramMessage
	requestBody.Color = "purple"
	if !m.isMessageValid(requestBody.Message) {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "message size exceed the limit",
		})
	}

	message, err := json.Marshal(requestBody)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": "error encoding message",
		})
		return
	}

	m.wsHub.Broadcast <- message
}

func (m MessageHandler) isMessageValid(s string) bool {
	messageSize, err := strconv.Atoi(m.appConfig.MESSAGE_SIZE)
	if err != nil {
		return false
	}
	return len(s) <= messageSize
}
