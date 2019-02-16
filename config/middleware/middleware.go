package middleware

import (
	prehandle "github.com/anthonydenecheau/gopocservice/delivery"
	"github.com/anthonydenecheau/gopocservice/delivery/renderings"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	logMiddleware "github.com/labstack/gommon/log"
)

func CreateRoutesGeneric(router *echo.Echo) *echo.Echo {

	// endpoints attachés au routeur
	for _, route := range routesMiddleware {
		log.Println("Add route {}", route.Name)
		router.Logger.Infof("Add route %s", route.Name)
		router.Add(route.Method, route.Pattern, route.HandlerFunc)
	}

	return router
}
func CreateRoutesApi(router *echo.Group) *echo.Group {

	// endpoints attachés aux API
	for _, route := range routesApi {
		log.Println("Add route {}", route.Name)
		//router.Infof("Add route %s", route.Name)
		router.Add(route.Method, route.Pattern, route.HandlerFunc)
	}

	return router
}
func NewRouter() *echo.Echo {

	log.Println("Create router ... ")
	r := echo.New()

	r.Logger.SetLevel(logMiddleware.INFO)

	// Middleware
	r.Pre(prehandle.RequestIDMiddleware)
	r.Pre(middleware.RemoveTrailingSlash())

	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	r.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}
		if he, ok := err.(*echo.HTTPError); ok {
			c.JSON(he.Code, renderings.Error{
				Status:  he.Code,
				Message: he.Error(),
			})
		}
	}

	//CORS
	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// endpoints attachés au routeur
	CreateRoutesGeneric(r)

	tokens := authenticationMiddleware{make(map[string]string)}
	tokens.Populate()
	v1 := r.Group("/api/v1")
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
	CreateRoutesApi(v1)

	return r

}
