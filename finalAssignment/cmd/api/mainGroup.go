package api

import (
	"final/cmd/api/handlers"

	"github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Group, api handlers.API) {

	e.GET("/lists", api.GetLists)
	e.GET("/lists/:id/tasks", api.GetTasks)
	e.GET("/weather", handlers.GetWeather)
	e.GET("/list/export", api.ExportLists)

	e.POST("/lists/:id/tasks", api.AddTask)
	e.POST("/lists", api.AddList)
	e.PATCH("/tasks/:id", api.ToggleTask)
	e.DELETE("/tasks/:id", api.DeleteTask)
	e.DELETE("/lists/:id", api.DeleteList)
}
