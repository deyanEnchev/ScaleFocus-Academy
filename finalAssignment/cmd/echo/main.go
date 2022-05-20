package main

import (
	"final/cmd"
	"final/cmd/router"
	"log"
	"net/http"
)

func main() {
	router := router.New()

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
