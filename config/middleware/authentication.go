package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
)

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
