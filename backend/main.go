package main

import (
	"backend/debug"
	"backend/handlers/user"
	"log"

	"net/http"
)

// TODO: Remove this in prod, extract to env variable
var IS_DEBUG = true

func main() {
	// Routes
	// Authentication (debug)
	if IS_DEBUG {
		http.HandleFunc("/api/v1/getToken", debug.GetToken)
		http.HandleFunc("/api/v1/verifyToken", debug.VerifyToken)
	}

	// Authentication
	http.HandleFunc("/api/v1/createAccount", user.CreateUser)
	http.HandleFunc("/api/v1/login", user.LoginUser)

	log.Fatal(http.ListenAndServe(":9090", nil))
}
