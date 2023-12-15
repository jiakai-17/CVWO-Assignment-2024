package main

import (
	"backend/debug"
	"backend/handlers/comments"
	"backend/handlers/threads"
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

	// Threads
	http.HandleFunc("/api/v1/getThread", threads.GetThread)
	http.HandleFunc("/api/v1/createThread", threads.CreateThread)
	http.HandleFunc("/api/v1/updateThread", threads.DeleteThread)
	http.HandleFunc("/api/v1/deleteThread", threads.DeleteThread)

	// Search Threads
	http.HandleFunc("/api/v1/searchThread", threads.SearchThread)

	log.Fatal(http.ListenAndServe(":9090", nil))
}
