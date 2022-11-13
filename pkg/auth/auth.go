package auth

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TokenHeaderName is the header name to set authentication token.
const TokenHeaderName = "X-Auth-Token"

// Middleware middleware check for an authorized token.
func Middleware(authorizedTokens []string) echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup:  fmt.Sprintf("header:%s", TokenHeaderName),
		AuthScheme: "",
		Validator: func(key string, c echo.Context) (bool, error) {
			for _, token := range authorizedTokens {
				if key == token {
					return true, nil
				}
			}
			return false, nil
		},
	})
}
