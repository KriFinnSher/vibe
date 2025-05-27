package handlers

import (
	"sync"

	gen "fullGen/tintin"
)

type Server struct {
	mu      sync.RWMutex
	storage map[string]gen.Tintin
}

func NewServer() *Server {
	return &Server{
		storage: make(map[string]gen.Tintin),
	}
}
