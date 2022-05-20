package router

import (
	"final/cmd/api"
	"final/cmd/api/handlers"
	"final/cmd/api/middlewares"
	"final/cmd/persistence"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	router := echo.New()

	// Set DB.
	repo := persistence.NewRepository(persistence.ConnectToDb())
	a := handlers.API{Storage: repo}

	// Create groups.
	apiGroup := router.Group("/api")

	// Set all middlewares.
	middlewares.SetMainMiddlewares(router)

	// Set main routes.
	api.MainGroup(apiGroup, a)

	return router
}
