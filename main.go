package main

import (
	"log"
	"net/http"
	"terraform-deployer-api/routes"
)

func main() {
    // Set up the router
    router := routes.SetupRoutes()

    // Start the HTTP server
    log.Println("Starting server on :8080...")
    err := http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}