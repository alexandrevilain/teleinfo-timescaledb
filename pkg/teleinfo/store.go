package teleinfo

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/jmoiron/sqlx"
)

var (
	createHyperTableStmt = `
	CREATE TABLE IF NOT EXISTS teleinfo_data (
		time TIMESTAMPTZ NOT NULL,
		instant_intensity INTEGER,
		max_intensity INTEGER,
		total_energy INTEGER,
	);
	SELECT create_hypertable('teleinfo_data', 'time', if_not_exists => TRUE);
	`
)

type Store struct {
	log logr.Logger
	db  *sqlx.DB
}

func NewStore(log logr.Logger, db *sqlx.DB) (*Store, error) {
	s := &Store{log: log, db: db}
	return s, s.init()
}

func (s *Store) init() error {
	_, err := s.db.Exec(createHyperTableStmt)
	return err
}

func (s *Store) Create(ctx context.Context, frame *Frame) error {
	_, err := s.db.Exec("INSERT INTO teleinfo_data (time, instant_intensity, max_intensity,  total_energy) VALUES (now(), $1, $2, $3)",
		frame.InstantIntensity,
		frame.MaxIntensity,
		frame.TotalEnergy,
	)
	return err
}
