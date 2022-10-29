package teleinfo

import (
	"context"

	"github.com/go-logr/logr"
)

type Service struct {
	store *Store
}

func NewService(log logr.Logger, store *Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) Create(ctx context.Context, f *RawFrame) error {
	frame := NewFrameFromRaw(f)
	return s.store.Create(ctx, frame)
}
