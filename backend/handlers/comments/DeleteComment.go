package comments

import (
	"backend/tutorial"
	"backend/utils"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"net/http"
)

// DeleteComment Handler for /api/v1/deleteComment
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	// Only POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get details from request body
	commentId := r.FormValue("id")
	username := r.FormValue("username")

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

	// Create comment UUID for pg
	var commentUUID pgtype.UUID

	commentUUID.Scan(commentId)

	// Check if the user is the creator of the comment
	isCreator, err := queries.CheckCommentCreator(ctx, tutorial.CheckCommentCreatorParams{Creator: username,
		ID: commentUUID})

	if err != nil || !isCreator {
		log.Println("[ERROR] Unable to verify creator: ", err, isCreator)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Delete the comment
	err = queries.DeleteComment(ctx, tutorial.DeleteCommentParams{
		ID:      commentUUID,
		Creator: username,
	})

	if err != nil {
		log.Println("[ERROR] Unable to delete comment: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
