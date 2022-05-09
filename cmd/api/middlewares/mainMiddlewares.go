package middlewares

import "github.com/labstack/echo/v4"

func SetMainMiddlewares(e *echo.Echo) {
	e.Use(serverHeader)
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set(echo.HeaderServer, "ScaleFocusTODO")

		return next(ctx)
	}
}