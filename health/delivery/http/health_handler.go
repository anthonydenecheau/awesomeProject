package http

import (
	render "github.com/anthonydenecheau/gopocservice/health/delivery/renderings"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HttpHealthHandler struct {
}

func (b *HttpHealthHandler) GetByID(c echo.Context) error {
	resp := render.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}
func NewHealthHttpHandler(e *echo.Echo) {
	handler := &HttpHealthHandler{}
	// endpoints attachés au serveur
	e.GET("/health", handler.GetByID)
}
