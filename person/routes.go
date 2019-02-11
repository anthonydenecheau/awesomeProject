package person

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"PeopleList",
		"GET",
		"/",
		GetPeopleEndpoint,
	},
	Route{
		"PeopleCreate",
		"POST",
		"/",
		CreatePersonEndpoint,
	},
	Route{
		"PeopleShow",
		"GET",
		"/{id}",
		GetPersonEndpoint,
	},
	Route{
		"PeopleDelete",
		"DELETE",
		"/{id}",
		DeletePersonEndpoint,
	},
}

func NewRoutes() Routes {
	return routes
}
