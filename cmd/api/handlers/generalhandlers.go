package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ApiHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello, World!")
}