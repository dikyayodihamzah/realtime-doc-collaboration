package wssrv

import (
	"log"
	"net/http"

	"github.com/dikyayodihamzah/realtime-doc-collaboration/pkg/config"
	"github.com/dikyayodihamzah/realtime-doc-collaboration/pkg/model"
)

type WebSocketService interface {
	Serve(w http.ResponseWriter, r *http.Request)
}

type webSocketService struct {
	Hub *model.Hub
}

func New(h *model.Hub) WebSocketService {
	return &webSocketService{
		Hub: h,
	}
}

// Serve upgrades the HTTP connection to WebSocket and registers it
func (s *webSocketService) Serve(w http.ResponseWriter, r *http.Request) {
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
	s.Hub.Register <- client

	// Start write and read pumps for the connection
	go client.WritePump()
	go client.ReadPump(s.Hub)
}
