package middleware

import (
	"github.com/anthonydenecheau/gopocservice/health/delivery/renderings"
	"github.com/labstack/echo/v4"
	"github.com/satori/go.uuid"
	"net/http"
)

const (
	RequestIDContextKey = "gopocservice_correlation_id"
)

type goMiddleware struct {
	// another stuff , may be needed by middleware
}

func (m *goMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}
func (m *goMiddleware) RequestIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		correlation := uuid.NewV4()
		c.Set(RequestIDContextKey, correlation)
		c.Response().Header().Set(RequestIDContextKey, correlation.String())
		c.Logger().Infof("Set %s %s", RequestIDContextKey, c.Get(RequestIDContextKey))
		return next(c)
	}
}
func HealthCheck(c echo.Context) error {
	if reqID, ok := c.Get(RequestIDContextKey).(uuid.UUID); ok {
		c.Logger().Infof("RequestID: %s", reqID.String())
	}
	resp := renderings.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}
func InitMiddleware() *goMiddleware {
	return &goMiddleware{}
}
