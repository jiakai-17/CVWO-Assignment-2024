package server

import (
	"backend/internal/router"
	"backend/internal/utils"
	"log"
	"net/http"
)

// StartServer Starts the server
func StartServer() {
	// Initialise JWT secret
	utils.InitJwtSecret()

	// Start server
	http.Handle("/", router.SetupRouter())

	utils.Log("main", "Listening on port 9090...", nil)

	log.Fatal(http.ListenAndServe(":9090", nil))
}
