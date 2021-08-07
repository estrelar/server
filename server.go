package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(*http.Request) bool { return true },
}

func main() {
	setRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setRoutes() {
	http.HandleFunc("/ws", websocketHandler)
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("[ERROR	] Connection -> %v", err)
		return
	}

	handleWebsocketConnection(conn)
}

func handleWebsocketConnection(c *websocket.Conn) {
	messageType, p, err := c.ReadMessage()

	switch messageType {
	case websocket.TextMessage:
		log.Printf("[Text	] %v", p)
	case websocket.BinaryMessage:
		log.Printf("[Binary	] %v", p)
	}

	err = c.WriteMessage(websocket.TextMessage, []byte("ueba"))
	if err != nil {
		log.Println(err)
	}
}
