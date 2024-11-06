package hubsrv

import "github.com/dikyayodihamzah/realtime-doc-collaboration/pkg/model"

// run handles the registration, unregistration, and broadcasting of messages
func Run(h *model.Hub) {
	for {
		select {
		case conn := <-h.Register:
			h.Connections[conn] = true
			// Send the current document content to the new connection
			h.Doc.Mu.Lock()
			conn.Send <- []byte(h.Doc.Content)
			h.Doc.Mu.Unlock()
		case conn := <-h.Unregister:
			if _, ok := h.Connections[conn]; ok {
				delete(h.Connections, conn)
				close(conn.Send)
			}
		case message := <-h.Broadcast:
			// Update document content with the new message
			h.Doc.Mu.Lock()
			h.Doc.Content = string(message)
			h.Doc.Mu.Unlock()
			// Broadcast to all clients
			for conn := range h.Connections {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(h.Connections, conn)
				}
			}
		}
	}
}
