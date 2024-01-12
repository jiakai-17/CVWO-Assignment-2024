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
		w.Write([]byte("Method not allowed"))
		return
	}

	type ThreadCreate struct {
		Title string   `json:"title"`
		Body  string   `json:"body"`
		Tags  []string `json:"tags"`
	}

	// Get details from request body
	var threadCreate ThreadCreate

	err := json.NewDecoder(r.Body).Decode(&threadCreate)

	if err != nil {
		log.Println("[ERROR] Unable to decode JSON: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Malformed JSON"))
		return
	}

	title := threadCreate.Title
	body := threadCreate.Body

	if len(threadCreate.Tags) > 3 {
		log.Println("[ERROR] Too many tags")
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		w.Write([]byte("Input too large"))
		return
	}

	if title == "" || body == "" {
		log.Println("[ERROR] Title or body is empty")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Title or body is empty"))
		return
	}

	// Get JWT token from request header
	token := r.Header.Get("Authorization")

	// Remove "Bearer " from token
	token = token[7:]

	log.Println("[DEBUG] Token: ", token)

	// Verify token
	verifiedUsername, err := utils.VerifyJWT(token)

	if err != nil {
		log.Println("[ERROR] Unable to verify JWT token: ", err, verifiedUsername)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	defer conn.Close(ctx)
	queries := tutorial.New(conn)

	// Create the thread
	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Println("[ERROR] Unable to begin transaction: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	defer tx.Rollback(ctx)
	qtx := queries.WithTx(tx)
	thread, err := qtx.CreateThread(ctx, tutorial.CreateThreadParams{Creator: verifiedUsername, Title: title, Body: body})

	//thread, err := queries.CreateThread(ctx, tutorial.CreateThreadParams{Creator: verifiedUsername, Title: title, Body: body})

	if err != nil {
		log.Println("[ERROR] Unable to create thread: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	var threadId = thread.ID

	// Create tags
	err = qtx.AddNewTags(ctx, threadCreate.Tags)
	if err != nil {
		log.Println("[ERROR] Unable to create tags: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Add tags to thread
	err = qtx.AddThreadTags(ctx, tutorial.AddThreadTagsParams{ThreadID: threadId, Tagarray: threadCreate.Tags})

	if err != nil {
		log.Println("[ERROR] Unable to add tags to thread: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	err = tx.Commit(ctx)

	if err != nil {
		log.Println("[ERROR] Unable to commit transaction: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	updatedThread, _ := queries.GetThreadDetails(ctx, threadId)

	// Return thread as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(updatedThread)

	if jsonErr != nil {
		log.Println("[ERROR] Unable to encode thread as JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	return
}
