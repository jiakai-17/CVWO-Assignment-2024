package threads

import (
	"backend/database"
	"backend/tutorial"
	"backend/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

// CreateThread Handler for /api/v1/createThread
func CreateThread(w http.ResponseWriter, r *http.Request) {
	// Only POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get details from request body
	username := r.FormValue("username")
	title := r.FormValue("title")
	body := r.FormValue("body")

	// Get JWT token from request header
	token := r.Header.Get("Authorization")

	// Remove "Bearer " from token
	token = token[7:]

	log.Println("[DEBUG] Token: ", token)

	// Verify token
	verifiedUsername, err := utils.VerifyJWT(token)

	if err != nil || verifiedUsername != username {
		log.Println("[ERROR] Unable to verify JWT token: ", err, verifiedUsername, username)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	queries := tutorial.New(conn)

	// Create the thread
	thread, err := queries.CreateThread(ctx, tutorial.CreateThreadParams{Creator: username, Title: title, Body: body})

	if err != nil {
		log.Println("[ERROR] Unable to create thread: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return comment as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(thread)

	if jsonErr != nil {
		log.Println("[ERROR] Unable to encode thread as JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
