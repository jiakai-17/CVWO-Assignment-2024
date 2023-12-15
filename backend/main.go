package main

import (
	"backend/debug"
	"backend/handlers/comments"
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

	// Comments
	http.HandleFunc("/api/v1/getComment", comments.GetComment)
	http.HandleFunc("/api/v1/createComment", comments.CreateComment)
	http.HandleFunc("/api/v1/updateComment", comments.UpdateComment)
	http.HandleFunc("/api/v1/deleteComment", comments.DeleteComment)

	log.Fatal(http.ListenAndServe(":9090", nil))
}
