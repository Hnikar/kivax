package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
	room *Room
}
type Room struct {
	ID        string
	clients   map[*Client]bool
	broadcast chan []byte
	join      chan *Client
	leave     chan *Client
}

type Hub struct {
	rooms        map[string]*Room
	registerRoom chan *Room
}

func NewHub() *Hub {
	return &Hub{
		rooms:        make(map[string]*Room),
		registerRoom: make(chan *Room),
	}
}

func (h *Hub) Run() {
	for {
		select {

		case room := <-h.registerRoom:
			h.rooms[room.ID] = room
		}
	}
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			log.Printf("Клиент присоединился к комнате %s", r.ID)

		case client := <-r.leave:
			if _, ok := r.clients[client]; ok {
				delete(r.clients, client)
				close(client.send)
				log.Printf("Клиент покинул комнату %s", r.ID)
			}

		case message := <-r.broadcast:
			for client := range r.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(r.clients, client)
				}
			}
		}
	}
}

func (c *Client) readPump() {
	defer func() {
		c.room.leave <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		c.room.broadcast <- message
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}
