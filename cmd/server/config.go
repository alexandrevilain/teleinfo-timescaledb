package main

import (
	"github.com/alexandrevilain/teleinfo-timescaledb/pkg/database"
)

type Config struct {
	ListenAdress     string          `koanf:"listenAdress"`
	DB               database.Config `koanf:"db"`
	AuthorizedTokens []string        `koanf:"authorizedTokens"`
}
