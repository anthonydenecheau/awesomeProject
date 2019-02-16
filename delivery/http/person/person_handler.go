package person

import (
	request "github.com/anthonydenecheau/gopocservice/delivery"
	model "github.com/anthonydenecheau/gopocservice/model"
	repository "github.com/anthonydenecheau/gopocservice/repository"
	"github.com/labstack/echo/v4"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"strconv"
)

func GetPersonEndpoint(c echo.Context) error {

	if reqID, ok := c.Get(request.RequestIDContextKey).(uuid.UUID); ok {
		c.Logger().Infof("RequestID: %s", reqID.String())
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Person `id` = ", id, " not found")
	}

	var p *model.Ws_dog_eleveur
	p, err = repository.GetPerson().GetPeopleById(int64(id))
	if err != nil {
		log.Println("Db error {}", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Person `id` = ", id, " not found")
	}
	return c.JSON(http.StatusOK, p)
}
