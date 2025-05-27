package handlers

import (
	"context"
	"errors"
	gen "fullGen/tintin"
)

func (s *Server) TintinsIDPut(ctx context.Context, req *gen.Tintin, params gen.TintinsIDPutParams) (gen.TintinsIDPutRes, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.storage[params.ID]; !ok {
		return &gen.TintinsIDPutNotFound{}, errors.New("not found")
	}

	req.ID = params.ID
	s.storage[params.ID] = *req

	tintin := s.storage[params.ID]

	return &tintin, nil
}
