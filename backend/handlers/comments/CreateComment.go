package comments

import (
	"backend/tutorial"
	"backend/utils"
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"net/http"
)

// CreateComment Handler for /api/v1/createComment
func CreateComment(w http.ResponseWriter, r *http.Request) {
	// Only POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get details from request body
	username := r.FormValue("username")
	thread := r.FormValue("thread")
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
	var threadID pgtype.UUID

	threadID.Scan(thread)

	// Create the comment
	params := tutorial.CreateCommentParams{
		Body:     body,
		Creator:  username,
		ThreadID: threadID,
	}

	comment, err := queries.CreateComment(ctx, params)

	if err != nil {
		log.Println("[ERROR] Unable to create comment: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return comment as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(comment)

	if jsonErr != nil {
		log.Println("[ERROR] Unable to encode comment as JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
