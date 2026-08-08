package main

import (
	"devbook-api/internal/config"
	"devbook-api/internal/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	fmt.Printf("API running at localhost:%d", config.Port)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", config.Port),
		router.New()),
	)
}
