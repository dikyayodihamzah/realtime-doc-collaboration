package model

import (
	"log"

	"github.com/gorilla/websocket"
)

// Connection represents a single WebSocket connection
type Connection struct {
	WS   *websocket.Conn
	Send chan []byte
}

// readPump reads messages from the WebSocket and sends them to the hub
func (c *Connection) ReadPump(h *Hub) {
	defer func() {
		h.Unregister <- c
		c.WS.Close()
	}()
	for {
		_, message, err := c.WS.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}
		// Send the message to the hub for broadcasting
		h.Broadcast <- message
	}
}

// writePump sends messages to the WebSocket from the hub
func (c *Connection) WritePump() {
	defer c.WS.Close()
	for message := range c.Send {
		if err := c.WS.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("WebSocket write error:", err)
			break
		}
	}
}
