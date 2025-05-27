package handlers

import (
	"context"
	"errors"
	gen "fullGen/tintin"
)

func (s *Server) TintinsGet(ctx context.Context) ([]gen.Tintin, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tintins := make([]gen.Tintin, 0, len(s.storage))
	for _, t := range s.storage {
		tintins = append(tintins, t)
	}
	return tintins, nil
}

func (s *Server) TintinsIDGet(ctx context.Context, params gen.TintinsIDGetParams) (gen.TintinsIDGetRes, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tintin, ok := s.storage[params.ID]
	if !ok {
		return &gen.TintinsIDGetNotFound{}, errors.New("not found")
	}

	return &tintin, nil
}
