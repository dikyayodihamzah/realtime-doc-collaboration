package model

// Hub manages all active WebSocket connections
type Hub struct {
	Connections map[*Connection]bool
	Broadcast   chan []byte
	Register    chan *Connection
	Unregister  chan *Connection
	Doc         *Document
}
