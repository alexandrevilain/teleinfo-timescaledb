package teleinfogrp

import (
	"net/http"

	"github.com/alexandrevilain/teleinfo-timescaledb/pkg/teleinfo"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Service *teleinfo.Service
}

func (h Handlers) Create(c echo.Context) error {
	f := &teleinfo.RawFrame{}
	err := c.Bind(f)
	if err != nil {
		return err
	}

	err = h.Service.Create(c.Request().Context(), f)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "OK")
}
