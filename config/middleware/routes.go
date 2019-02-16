package middleware

import (
	"github.com/anthonydenecheau/gopocservice/delivery/http/middleware"
	"github.com/anthonydenecheau/gopocservice/delivery/http/person"
	"github.com/labstack/echo/v4"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc echo.HandlerFunc
}

type Routes []Route

var routesMiddleware = Routes{
	Route{
		"HealthCheck",
		"GET",
		"/health",
		middleware.HealthCheck,
	},
}
var routesApi = Routes{
	/*
		Route{
			"PeopleList",
			"GET",
			"/people",
			person.GetPeopleEndpoint,
		},
			Route{
				"PeopleCreate",
				"POST",
				"/people",
				CreatePersonEndpoint,
			},
		Route{
			"PeopleDelete",
			"DELETE",
			"/people/:id",
			person.DeletePersonEndpoint,
		},
	*/
	Route{
		"PeopleShow",
		"GET",
		"/people/:id",
		person.GetPersonEndpoint,
	},
}

func NewRoutes() Routes {
	return routesApi
}
