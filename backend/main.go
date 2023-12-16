package main

import (
	"backend/debug"
	"backend/handlers/comments"
	"backend/handlers/threads"
	"backend/handlers/user"
	"backend/utils"
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
)

// Helper function to pass the pgx.Conn object to the handler
func wrapDbConnection(
	handler func(http.ResponseWriter, *http.Request, *pgx.Conn),
	conn *pgx.Conn) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		handler(writer, request, conn)
	}
}

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

	// Connect to database
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=postgres dbname=cvwo-1 password=cs2102")
	if err != nil {
		log.Fatal("[ERROR] Unable to connect to database: ", err)
		return
	}
	defer conn.Close(ctx)

	// Routes
	// Authentication (debug)
	if IS_DEBUG {
		http.HandleFunc("/api/v1/getToken", debug.GetToken)
		http.HandleFunc("/api/v1/verifyToken", debug.VerifyToken)
	}

	// Authentication
	http.HandleFunc("/api/v1/createAccount", wrapDbConnection(user.CreateUser, conn))
	http.HandleFunc("/api/v1/login", wrapDbConnection(user.LoginUser, conn))

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

	// Start server
	utils.Log("main", "Listening on port 9090...", nil)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
