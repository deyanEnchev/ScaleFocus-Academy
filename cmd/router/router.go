package router

import (
	"final/cmd/api"
	"final/cmd/api/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	router := echo.New()

	// Create groups.
	adminGroup := router.Group("/admin")
	cookieGroup := router.Group("/cookie")

	// Set all middlewares.
	middlewares.SetMainMiddlewares(router)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetCookieMiddlewares(cookieGroup)

	// Set main routes.
	api.MainGroup(router)
	
	// Set group routes.
	api.AdminGroup(adminGroup)
	api.CookieGroup(cookieGroup)
	return router
}

