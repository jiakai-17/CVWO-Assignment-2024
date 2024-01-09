package main

import (
	"backend/debug"
	"backend/handlers/comments"
	"backend/handlers/threads"
	"backend/handlers/user"
	"backend/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var BASE_PATH = "/api/v1/"

// TODO: Remove this in prod, extract to env variable
var IS_DEBUG = true

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

	r := mux.NewRouter()

	// Routes
	// Authentication (debug)
	if IS_DEBUG {
		http.HandleFunc("/api/v1/getToken", debug.GetToken)
		http.HandleFunc("/api/v1/verifyToken", debug.VerifyToken)
	}

	// Authentication
	http.HandleFunc(BASE_PATH+"user/create", user.CreateUser)
	http.HandleFunc(BASE_PATH+"user/login", user.LoginUser)

	// Comments
	r.HandleFunc(BASE_PATH+"thread/{thread_id}/comments", comments.GetComment).Methods("GET")
	http.HandleFunc(BASE_PATH+"comment/create", comments.CreateComment)
	r.HandleFunc(BASE_PATH+"comment/{id}", comments.UpdateComment).Methods("PUT")
	r.HandleFunc(BASE_PATH+"comment/{id}", comments.DeleteComment).Methods("DELETE")

	// Threads
	//r.HandleFunc(BASE_PATH+"threads", threads.GetThreads).Methods("GET")
	r.HandleFunc(BASE_PATH+"thread/{id}", threads.GetThread).Methods("GET")
	http.HandleFunc(BASE_PATH+"thread/create", threads.CreateThread)
	r.HandleFunc(BASE_PATH+"thread/{id}", threads.UpdateThread).Methods("PUT")
	r.HandleFunc(BASE_PATH+"thread/{id}", threads.DeleteThread).Methods("DELETE")

	// Search Threads
	http.HandleFunc(BASE_PATH+"thread", threads.SearchThread)

	// Start server
	http.Handle("/", r)
	utils.Log("main", "Listening on port 9090...", nil)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
