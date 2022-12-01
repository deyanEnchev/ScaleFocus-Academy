package main

import (
	"final/cmd"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	apiGroup := router.Group("/api")
	apiGroup.GET("/lists", )



	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}


	// router.Use(func(ctx *gin.Context) {
	// 	// This is a sample demonstration of how to attach middlewares in Gin
	// 	log.Println("Gin middleware was called")
	// 	ctx.Next()
	// })

	// // Add your handler (API endpoint) registrations here
	// router.GET("/api", func(ctx *gin.Context) {
	// 	ctx.JSON(200, "Hello, World!")
	// })