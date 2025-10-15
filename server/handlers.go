package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Ошибка при повышении до WebSocket: %+v", err)
		return
	}
	defer conn.Close()

	log.Println("✅ Клиент успешно подключился по WebSocket")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка при чтении сообщения:", err)
			break
		}
		log.Printf("📥 Получено сообщение: %s", string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("Ошибка при отправке сообщения:", err)
			break
		}
	}
}
