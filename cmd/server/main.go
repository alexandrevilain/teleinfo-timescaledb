package main

import (
	"fmt"
	"os"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/pflag"

	"github.com/alexandrevilain/teleinfo-timescaledb/cmd/server/handlers"
	"github.com/alexandrevilain/teleinfo-timescaledb/pkg/database"
	"github.com/alexandrevilain/teleinfo-timescaledb/pkg/log"
)

func main() {
	err := run()
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
}

func run() error {
	logger, err := log.CreateLogger()
	if err != nil {
		return err
	}

	k := koanf.New(".")

	f := pflag.NewFlagSet("", pflag.ContinueOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}
	f.StringSlice("config", []string{"config.yaml"}, "path to one or more .yaml config files")
	f.Parse(os.Args[1:])

	configFiles, _ := f.GetStringSlice("config")
	for _, c := range configFiles {
		if err := k.Load(file.Provider(c), yaml.Parser()); err != nil {
			return fmt.Errorf("error loading file: %v", err)
		}
	}

	dbConfig := database.Config{}
	k.Unmarshal("db", &dbConfig)

	db, err := database.Open(dbConfig)
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}

	logger.Info("Starting server")

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogError:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request", "URI", v.URI, "status", v.Status, "error", v.Error)
			return nil
		},
	}))

	if err := handlers.AddRoutes(logger, db, e); err != nil {
		return err
	}

	return e.Start(k.String("listenAdress"))
}
