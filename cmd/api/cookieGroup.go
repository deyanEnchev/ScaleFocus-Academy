package api

import (
	"final/cmd/api/handlers"

	"github.com/labstack/echo/v4"
)

func CookieGroup(g *echo.Group) {
	g.GET("/main", handlers.MainCookie)
}