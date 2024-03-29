package comments

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/utils"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
	"strings"
)

// UpdateComment godoc
// @Summary Handles comment update requests
// @Description Updates a comment
// @Tags comment
// @Param id path string true "Comment UUID"
// @Param data body models.UpdateCommentRequest true "Comment data"
// @Security ApiKeyAuth
// @Success 200
// @Failure 400 "Invalid data"
// @Failure 401 "Invalid JWT token"
// @Failure 403 "No permission to update comment"
// @Failure 405 "Method not allowed"
// @Failure 500 "Internal server error"
// @Router /comment/{id} [put]
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	// Only PUT
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Method not allowed"))
		if err != nil {
			utils.Log("UpdateComment", "Unable to write response", err)
		}
		return
	}

	// Get details from request
	var commentUpdate models.UpdateCommentRequest

	err := json.NewDecoder(r.Body).Decode(&commentUpdate)

	if err != nil {
		utils.Log("UpdateComment", "Unable to decode JSON", err)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Invalid data"))
		if err != nil {
			utils.Log("UpdateComment", "Unable to write response", err)
		}
		return
	}

	commentId := mux.Vars(r)["id"]
	body := strings.TrimSpace(commentUpdate.Body)

	// Ensure comment body is not empty and is not too long
	if len(body) == 0 || len(body) > 3000 {
		utils.Log("UpdateComment", "Invalid comment body", nil)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Invalid data"))
		if err != nil {
			utils.Log("UpdateComment", "Unable to write response", err)
		}
		return
	}

	// Get and verify JWT token from request header
	token := r.Header.Get("Authorization")[7:]
	verifiedUsername, err := utils.VerifyJWT(token)

	if err != nil {
		utils.Log("UpdateComment", "Unable to verify JWT token", err)
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Invalid JWT token"))
		if err != nil {
			utils.Log("UpdateComment", "Unable to write response", err)
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
		utils.Log("UpdateComment", "Unable to scan commentId", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("UpdateComment", "Unable to write response", err)
		}
		return
	}

	// Check if the user is the creator of the comment
	isCreator, err := queries.CheckCommentCreator(ctx, database.CheckCommentCreatorParams{
		Creator: verifiedUsername,
		ID:      pgCommentId})

	if err != nil || !isCreator {
		utils.Log("UpdateComment", "User is not the creator of the comment", err)
		w.WriteHeader(http.StatusForbidden)
		_, err := w.Write([]byte("No permission to update comment"))
		if err != nil {
			utils.Log("UpdateComment", "Unable to write response", err)
		}
		return
	}

	// Update the comment
	err = queries.UpdateComment(ctx, database.UpdateCommentParams{
		Body:    body,
		Creator: verifiedUsername,
		ID:      pgCommentId,
	})

	if err != nil {
		utils.Log("UpdateComment", "Unable to update comment", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("UpdateComment", "Unable to write response", err)
		}
		return
	}

	utils.Log("UpdateComment", "Comment "+commentId+" updated by "+verifiedUsername, nil)

	return
}
