package main

import (
	"backend/internal/router"
	"backend/internal/utils"
	"log"
	"net/http"
)

// @title           CVWO Forum Backend API
// @version         1.0
// @description     This is the backend API for the forum.

// @license.name  All Rights Reserved.

// @host      localhost:9090
// @BasePath  /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description The word "Bearer", followed by a space, and then the JWT token.
func main() {
	// Initialise JWT secret
	utils.InitJwtSecret()

	// Start server
	http.Handle("/", router.SetupRouter())

	utils.Log("main", "Listening on port 9090...", nil)

	log.Fatal(http.ListenAndServe(":9090", nil))
}
