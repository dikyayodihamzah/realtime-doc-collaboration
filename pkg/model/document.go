package model

import "sync"

// Document represents the shared document state
type Document struct {
	Content string
	Mu      sync.Mutex // To protect content updates
}
