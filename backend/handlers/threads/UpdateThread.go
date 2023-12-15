package threads

import (
	"backend/tutorial"
	"backend/utils"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"net/http"
)

// UpdateThread Handler for /api/v1/updateThread
func UpdateThread(w http.ResponseWriter, r *http.Request) {
	// Only POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get details from request body
	username := r.FormValue("username")
	threadId := r.FormValue("id")
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

	conn, err := pgx.Connect(ctx, "user=postgres dbname=cvwo-1 password=cs2102")
	if err != nil {
		log.Println("[ERROR] Unable to connect to database: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close(ctx)

	queries := tutorial.New(conn)

	// Create thread UUID for pg
	var threadUUID pgtype.UUID

	threadUUID.Scan(threadId)

	// Check if user is creator of thread
	isThreadCreator, err := queries.CheckThreadCreator(ctx, tutorial.CheckThreadCreatorParams{Creator: username,
		ID: threadUUID})

	if err != nil || !isThreadCreator {
		log.Println("[ERROR] Unable to check if user is creator of thread: ", err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Update the thread
	err = queries.UpdateThread(ctx, tutorial.UpdateThreadParams{ID: threadUUID, Title: title, Body: body, Creator: username})

	if err != nil {
		log.Println("[ERROR] Unable to update thread: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
