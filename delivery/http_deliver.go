package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/satori/go.uuid"
)

const (
	RequestIDContextKey = "gopocservice_correlation_id"
)

func RequestIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		correlation := uuid.NewV4()
		c.Set(RequestIDContextKey, correlation)
		c.Response().Header().Set(RequestIDContextKey, correlation.String())
		c.Logger().Infof("Set %s %s", RequestIDContextKey, c.Get(RequestIDContextKey))
		return next(c)
	})
}
