package comments

import (
	"backend/tutorial"
	"backend/utils"
	"context"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"net/http"
)

// DeleteComment godoc
// @Summary Handles comment deletion requests
// @Description Deletes a comment
// @Tags comment
// @Param id path string true "Comment UUID"
// @Success 200
// @Failure 401 "Invalid JWT token"
// @Failure 403 "User is not the creator of the comment"
// @Failure 500
// @Router /comment/{id} [delete]
func DeleteComment(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	// Only DELETE
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get details from request body
	commentId := mux.Vars(r)["id"]
	utils.Log("deleteComment", "[DEBUG] Comment ID: "+commentId, nil)

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
		return
	}
	username := verifiedUsername

	// Connect to database
	ctx := context.Background()
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
