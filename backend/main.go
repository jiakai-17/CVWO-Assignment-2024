package main

import (
	"backend/cmd/server"
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
	server.StartServer()
}
