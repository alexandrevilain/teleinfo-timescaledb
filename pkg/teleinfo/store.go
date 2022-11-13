package teleinfo

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"gorm.io/gorm"
)

type Store struct {
	log logr.Logger
	db  *gorm.DB
}

func NewStore(log logr.Logger, db *gorm.DB) (*Store, error) {
	s := &Store{log: log, db: db}
	return s, s.init()
}

func (s *Store) init() error {
	err := s.db.AutoMigrate(&Frame{})
	if err != nil {
		return fmt.Errorf("can't auto migrate Frame model: %w", err)
	}
	return s.db.Exec("SELECT create_hypertable(?, 'time', if_not_exists => TRUE)", Frame{}.TableName()).Error
}

func (s *Store) Create(ctx context.Context, frame *Frame) error {
	frame.Time = time.Now().UTC() // Always override data
	return s.db.Create(frame).Error
}
