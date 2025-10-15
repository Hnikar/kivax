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

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request, roomId string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	room, ok := hub.rooms[roomId]
	if !ok {
		room = &Room{
			ID:        roomId,
			clients:   make(map[*Client]bool),
			broadcast: make(chan []byte),
			join:      make(chan *Client),
			leave:     make(chan *Client),
		}
		hub.registerRoom <- room
		go room.Run()
	}

	client := &Client{
		conn: conn,
		send: make(chan []byte, 256),
		room: room,
	}

	client.room.join <- client

	go client.writePump()
	go client.readPump()

	log.Printf("Клиент подключился к комнате %s. Всего клиентов: %d", client.room.ID, len(client.room.clients))
}
