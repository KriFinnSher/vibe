package handlers

import (
	"context"
	gen "fullGen/tintin"
	"github.com/go-faster/errors"
)

func (s *Server) TintinsIDDelete(ctx context.Context, params gen.TintinsIDDeleteParams) (gen.TintinsIDDeleteRes, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.storage[params.ID]; !ok {
		return &gen.TintinsIDDeleteNotFound{}, errors.New("not found")
	}
	delete(s.storage, params.ID)
	return &gen.TintinsIDDeleteNoContent{}, nil
}
