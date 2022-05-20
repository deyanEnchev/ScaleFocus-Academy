package middlewares

import (
	"final/cmd/api/authentication"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetMainMiddlewares(e *echo.Echo) {
	e.Use(serverHeader)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.Use(middleware.BasicAuth(authentication.Authentication))
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		c.Set("Weather-API", "https://api.openweathermap.org/data/2.5/weather")
		return next(c)
	}
}
