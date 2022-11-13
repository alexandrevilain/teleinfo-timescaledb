package database

import (
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config is the required properties to use the database.
type Config struct {
	User         string
	Password     string
	Host         string
	Name         string
	MaxIdleConns int
	MaxOpenConns int
	DisableTLS   bool
	Timezone     string
}

// Open knows how to open a database connection based on the configuration.
func Open(cfg Config) (*gorm.DB, error) {
	sslMode := "require"
	if cfg.DisableTLS {
		sslMode = "disable"
	}

	timezone := "utc"
	if cfg.Timezone != "" {
		timezone = cfg.Timezone
	}

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", timezone)

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     cfg.Host,
		Path:     cfg.Name,
		RawQuery: q.Encode(),
	}

	db, err := gorm.Open(postgres.Open(u.String()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

	return db, nil
}
