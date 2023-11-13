package middleware

import (
	"github.com/Stanxxy/stan-go-web/internal/context"
	"github.com/labstack/echo/v4"
)

func AppContext(cc *context.AppContext) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc.Context = c
			return h(cc)
		}
	}
}
