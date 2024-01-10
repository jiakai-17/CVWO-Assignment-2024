package threads

import (
	"backend/database"
	"backend/tutorial"
	"backend/utils"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"net/http"
)

// UpdateThread Handler for /api/v1/updateThread
func UpdateThread(w http.ResponseWriter, r *http.Request) {
	// Only POST
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	type ThreadUpdate struct {
		Title string   `json:"title"`
		Body  string   `json:"body"`
		Tags  []string `json:"tags"`
	}

	// Get details from request
	vars := mux.Vars(r)
	threadId := vars["id"]

	var threadUpdate ThreadUpdate
	err := json.NewDecoder(r.Body).Decode(&threadUpdate)

	if err != nil {
		log.Println("[ERROR] Unable to decode JSON: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Malformed JSON"))
		return
	}

	title := threadUpdate.Title
	body := threadUpdate.Body
	tags := threadUpdate.Tags

	if len(tags) > 3 {
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

	//username := r.FormValue("username")
	//threadId := r.FormValue("id")
	//title := r.FormValue("title")
	//body := r.FormValue("body")

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
	queries := tutorial.New(conn)

	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Println("[ERROR] Unable to begin transaction: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	defer tx.Rollback(ctx)
	qtx := queries.WithTx(tx)

	// Create thread UUID for pg
	var threadUUID pgtype.UUID

	threadUUID.Scan(threadId)

	// Check if user is creator of thread
	isThreadCreator, err := qtx.CheckThreadCreator(ctx, tutorial.CheckThreadCreatorParams{Creator: verifiedUsername,
		ID: threadUUID})

	if err != nil || !isThreadCreator {
		log.Println("[ERROR] Unable to check if user is creator of thread: ", err)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Forbidden"))
		return
	}

	// Update the thread
	err = qtx.UpdateThread(ctx, tutorial.UpdateThreadParams{ID: threadUUID, Title: title, Body: body,
		Creator: verifiedUsername})

	// Update the tags
	err = qtx.DeleteThreadTags(ctx, threadUUID)

	if err != nil {
		log.Println("[ERROR] Unable to delete thread tags: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	err = qtx.DeleteUnusedTags(ctx)

	if err != nil {
		log.Println("[ERROR] Unable to delete unused tags: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	err = qtx.AddNewTags(ctx, tags)

	if err != nil {
		log.Println("[ERROR] Unable to add new tags: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	err = qtx.AddThreadTags(ctx, tutorial.AddThreadTagsParams{ThreadID: threadUUID, Tagarray: tags})

	if err != nil {
		log.Println("[ERROR] Unable to add thread tags: ", err)
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

	return
}
