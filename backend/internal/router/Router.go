package router

import (
	"backend/internal/handlers/comments"
	"backend/internal/handlers/threads"
	"backend/internal/handlers/user"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var BASE_PATH = "/api/v1/"

// SetupRouter Sets up the router for the server
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Get env variables
	if os.Getenv("BASE_PATH") != "" {
		BASE_PATH = os.Getenv("BASE_PATH")
	}

	// Routes
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

	return r
}
