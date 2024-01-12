package threads

import (
	"backend/database"
	"backend/tutorial"
	"backend/utils"
	"context"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"net/http"
)

// DeleteThread Handler for /api/v1/deleteThread
func DeleteThread(w http.ResponseWriter, r *http.Request) {
	// Only DELETE
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	// Get details from request body
	vars := mux.Vars(r)
	threadId := vars["id"]

	//username := r.FormValue("username")
	//threadId := r.FormValue("id")

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

	// Create thread UUID for pg
	var threadUUID pgtype.UUID

	threadUUID.Scan(threadId)

	// Check if user is creator of thread
	isThreadCreator, err := queries.CheckThreadCreator(ctx, tutorial.CheckThreadCreatorParams{Creator: verifiedUsername,
		ID: threadUUID})

	if err != nil || !isThreadCreator {
		log.Println("[ERROR] Unable to check if user is creator of thread: ", err)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Forbidden"))
		return
	}

	// Update the thread
	err = queries.DeleteThread(ctx, tutorial.DeleteThreadParams{
		ID:      threadUUID,
		Creator: verifiedUsername,
	})

	if err != nil {
		log.Println("[ERROR] Unable to delete thread: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	return
}
