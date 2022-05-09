package api

import (
	"final/cmd/api/handlers"

	"github.com/labstack/echo/v4"
)

func MainGroup(router *echo.Echo) {
	router.GET("/login", handlers.Login)
	router.GET("/api", handlers.ApiHandler)

	router.GET("/lists", handlers.GetLists)

	router.GET("/lists/:id/tasks", handlers.GetTasks)

	router.POST("/lists/:id/tasks", handlers.AddTask)
	router.POST("/lists", handlers.AddList)
	router.PATCH("/tasks/:id", handlers.ToggleTask)
	router.DELETE("/tasks/:id", handlers.DeleteTask)
	router.DELETE("/lists/:id", handlers.DeleteList)
}