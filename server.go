package main

import (
	"Gari/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":3000"

	log.Println("Running Server")
	fmt.Println("Running Server")
	routes.RegisterRoutes()

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Failed to start server")
		fmt.Println("Failed to start server")
	}
}
