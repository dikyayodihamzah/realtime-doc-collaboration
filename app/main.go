package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dikyayodihamzah/realtime-doc-collaboration/app/service/hubsrv"
	"github.com/dikyayodihamzah/realtime-doc-collaboration/pkg/config"
	"github.com/dikyayodihamzah/realtime-doc-collaboration/pkg/model"
)

// serveWebSocket upgrades the HTTP connection to WebSocket and registers it
func serveWebSocket(h *model.Hub, w http.ResponseWriter, r *http.Request) {
	upgrader := config.NewUpgrader()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}

	client := &model.Connection{
		WS:   conn,
		Send: make(chan []byte, 256),
	}
	h.Register <- client

	// Start write and read pumps for the connection
	go client.WritePump()
	go client.ReadPump(h)
}

// Main function to set up the server
func main() {
	doc := new(model.Document)
	hub := config.NewHub(doc)
	go hubsrv.Run(hub)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWebSocket(hub, w, r)
	})

	addr := ":8080"
	fmt.Println("Starting server on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
