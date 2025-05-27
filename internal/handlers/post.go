package handlers

import (
	"context"
	"errors"
	"fmt"
	gen "fullGen/tintin"
)

func (s *Server) TintinsPost(ctx context.Context, req *gen.Tintin) (*gen.Tintin, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if req.ID == "" {
		return nil, errors.New("id is required")
	}
	if _, exists := s.storage[req.ID]; exists {
		return nil, fmt.Errorf("tintin with id %s already exists", req.ID)
	}

	s.storage[req.ID] = *req
	return req, nil
}
