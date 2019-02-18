package breeder

import (
	models "github.com/anthonydenecheau/gopocservice/breeder"
	render "github.com/anthonydenecheau/gopocservice/breeder/delivery/renderings"
	breederUcase "github.com/anthonydenecheau/gopocservice/breeder/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	logMiddleware "github.com/labstack/gommon/log"
	"log"
	"net/http"
	"strconv"
)

type HttpBreederHandler struct {
	BUsecase breederUcase.BreederUsecase
}

type authenticationMiddleware struct {
	tokenUsers map[string]string
}

// Initialize it somewhere
func (amw *authenticationMiddleware) Populate() {
	amw.tokenUsers["POyIiqsN6gQxde7zxuX5"] = "scc_expos"
	amw.tokenUsers["000000"] = "agria"
	amw.tokenUsers["111111"] = "cga"
}

func check(amw *authenticationMiddleware, key string, c echo.Context) (bool, error) {
	token := c.Request().Header.Get("X-SCC-authentification")
	if user, found := amw.tokenUsers[token]; found {
		log.Printf("Authenticated user %s\n", user)
		return true, nil
	} else {
		return false, nil
	}
}
func (b *HttpBreederHandler) GetByID(c echo.Context) error {

	/* [TODO]
	if reqID, ok := c.Get(RequestIDContextKey).(uuid.UUID); ok {
		c.Logger().Infof("RequestID: %s", reqID.String())
	}
	*/

	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	r, err := b.BUsecase.GetByID(id)

	if err != nil {
		return c.JSON(getStatusCode(err), err.Error())
	}

	var rBreeder = render.BreederResponse{Id: r.Id, Firstname: r.Prenom, Lastname: r.Nom, Address: nil}

	return c.JSON(http.StatusOK, rBreeder)
}
func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}

	//logrus.Error(err)
	switch err {
	case models.INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	case models.NOT_FOUND_ERROR:
		return http.StatusNotFound
	case models.CONFLIT_ERROR:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
func NewBreederHttpHandler(e *echo.Echo, us breederUcase.BreederUsecase) {
	handler := &HttpBreederHandler{
		BUsecase: us,
	}

	e.Logger.SetLevel(logMiddleware.INFO)

	// Middleware
	//[TODO] e.Pre(requestIDMiddleware)
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	tokens := authenticationMiddleware{make(map[string]string)}
	tokens.Populate()
	v1 := e.Group("/api/v1")
	v1.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:X-SCC-authentification",
		Validator: func(key string, c echo.Context) (bool, error) {
			token := c.Request().Header.Get("X-SCC-authentification")
			if user, found := tokens.tokenUsers[token]; found {
				c.Logger().Infof("Authenticated user %s", user)
				return true, nil
			} else {
				return false, nil
			}
		},
	}))

	// endpoints attachés aux API
	v1.GET("/people/:id", handler.GetByID)

}
