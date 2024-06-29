package main

import (
	"log"
	"net/http"
	"terraform-deployer-api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/deploy", handlers.DeployHandler).Methods("POST")

	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}