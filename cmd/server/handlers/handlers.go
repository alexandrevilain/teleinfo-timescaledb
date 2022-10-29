package handlers

import (
	"github.com/alexandrevilain/teleinfo-timescaledb/cmd/server/handlers/v1/teleinfogrp"
	"github.com/alexandrevilain/teleinfo-timescaledb/pkg/teleinfo"
	"github.com/go-logr/logr"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func AddRoutes(log logr.Logger, db *sqlx.DB, e *echo.Echo) error {
	tiStore, err := teleinfo.NewStore(log, db)
	if err != nil {
		return err
	}
	tiGrp := teleinfogrp.Handlers{
		Service: teleinfo.NewService(log, tiStore),
	}

	v1 := e.Group("/v1")
	v1.POST("/frames", tiGrp.Create)

	return nil
}
