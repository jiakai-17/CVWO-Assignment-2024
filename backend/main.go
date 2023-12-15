package main

import (
	"backend/debug"
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

	log.Fatal(http.ListenAndServe(":9090", nil))
}
