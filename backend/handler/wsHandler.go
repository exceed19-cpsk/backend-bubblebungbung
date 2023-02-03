package handler

import (
	"github.com/exceed19-cpsk/backend-bubblebungbung/service"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// ServeWs handles websocket requests from the peer.
func ServeWs(hub *service.WsHub, w http.ResponseWriter, r *http.Request, u websocket.Upgrader) {
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &service.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
}
