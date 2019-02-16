package middleware

import (
	request "github.com/anthonydenecheau/gopocservice/delivery"
	"github.com/anthonydenecheau/gopocservice/delivery/renderings"
	"github.com/labstack/echo/v4"
	"github.com/satori/go.uuid"
	"net/http"
)

func HealthCheck(c echo.Context) error {
	if reqID, ok := c.Get(request.RequestIDContextKey).(uuid.UUID); ok {
		c.Logger().Infof("RequestID: %s", reqID.String())
	}
	resp := renderings.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}
