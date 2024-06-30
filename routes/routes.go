package routes

import (
	"terraform-deployer-api/handlers"

	"github.com/gorilla/mux"
)

// SetupRoutes initializes the router and sets up the routes
func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Define routes
    router.HandleFunc("/deploy", handlers.DeployHandler).Methods("POST")
    
    return router
}
