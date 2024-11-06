package config

import "github.com/dikyayodihamzah/realtime-doc-collaboration/pkg/model"

// NewHub initializes a new Hub
func NewHub(doc *model.Document) *model.Hub {
	return &model.Hub{
		Connections: make(map[*model.Connection]bool),
		Broadcast:   make(chan []byte),
		Register:    make(chan *model.Connection),
		Unregister:  make(chan *model.Connection),
		Doc:         doc,
	}
}
