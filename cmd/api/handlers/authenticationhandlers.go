package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Login(ctx echo.Context) error {
	username := ctx.QueryParam("username")
	password := ctx.QueryParam("password")

	//check username and password against DB after hashing the password
	if username == "jack" && password == "1234" {
		cookie := new(http.Cookie)
		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)
		ctx.SetCookie(cookie)

		return ctx.String(http.StatusOK, "You were logged in!")
	}
	return ctx.String(http.StatusUnauthorized, "Your username or passowrd were wrong.")
}
