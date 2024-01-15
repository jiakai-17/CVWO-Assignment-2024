package comments

import (
	"backend/internal/database"
	"backend/internal/utils"
	"context"
	"errors"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
)

// DeleteComment godoc
// @Summary Handles comment deletion requests
// @Description Deletes a comment
// @Tags comment
// @Param id path string true "Comment UUID"
// @Security ApiKeyAuth
// @Success 200
// @Failure 401 "Invalid JWT token"
// @Failure 403 "No permission to delete comment"
// @Failure 405 "Method not allowed"
// @Failure 500 "Internal server error"
// @Router /comment/{id} [delete]
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	// Only DELETE
	if r.Method != http.MethodDelete {
		utils.Log("DeleteComment", "Method not allowed", errors.New("method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Method not allowed"))
		if err != nil {
			utils.Log("DeleteComment", "Unable to write response", err)
		}
		return
	}

	// Get commentId from request
	commentId := mux.Vars(r)["id"]

	// Get and verify JWT token from request header
	token := r.Header.Get("Authorization")[7:]

	verifiedUsername, err := utils.VerifyJWT(token)

	if err != nil {
		utils.Log("DeleteComment", "Unable to verify JWT token", err)
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Invalid JWT token"))
		if err != nil {
			utils.Log("DeleteComment", "Unable to write response", err)
		}
		return
	}

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	defer database.CloseConnection(conn)
	queries := database.New(conn)

	// Create comment UUID for pg
	var pgCommentId pgtype.UUID
	err = pgCommentId.Scan(commentId)
	if err != nil {
		utils.Log("DeleteComment", "Unable to scan commentId", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("DeleteComment", "Unable to write response", err)
		}
		return
	}

	// Check if the user is the creator of the comment
	isCreator, err := queries.CheckCommentCreator(ctx, database.CheckCommentCreatorParams{
		Creator: verifiedUsername,
		ID:      pgCommentId})

	if err != nil || !isCreator {
		utils.Log("DeleteComment", "User is not the creator of the comment", err)
		w.WriteHeader(http.StatusForbidden)
		_, err := w.Write([]byte("No permission to delete comment"))
		if err != nil {
			utils.Log("DeleteComment", "Unable to write response", err)
		}
		return
	}

	// Delete the comment
	err = queries.DeleteComment(ctx, database.DeleteCommentParams{
		ID:      pgCommentId,
		Creator: verifiedUsername,
	})

	if err != nil {
		utils.Log("DeleteComment", "Unable to delete comment", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("DeleteComment", "Unable to write response", err)
		}
		return
	}

	utils.Log("DeleteComment", "Comment "+commentId+" deleted by "+verifiedUsername, nil)

	return
}
