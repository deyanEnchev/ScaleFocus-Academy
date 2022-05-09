package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func SetCookieMiddlewares(g *echo.Group) {
	g.Use(checkCookie)
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("sessionID")

		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return ctx.String(http.StatusUnauthorized, "you don't have any cookie")
			}
			return err
		}

		if cookie.Value == "some_string" {
			return next(ctx)
		}
		return ctx.String(http.StatusUnauthorized, "you dont have the right cookie, cookie")
	}
}
