package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	fmt.Println(config.Port)
	fmt.Println(config.DataSourceName)

	fmt.Printf("API running at localhost:%d", config.Port)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", config.Port),
		router.New()),
	)
}
