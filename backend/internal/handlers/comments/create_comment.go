package comments

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/utils"
	"context"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
	"strings"
)

// CreateComment godoc
// @Summary Handles comment creation requests
// @Description Creates a new comment for the given thread
// @Tags comment
// @Accept json
// @Produce json
// @Param data body models.CreateCommentRequest true "Comment data"
// @Security ApiKeyAuth
// @Success 200 {object} models.Comment
// @Failure 400 "Invalid data"
// @Failure 401 "Invalid JWT token"
// @Failure 405 "Method not allowed"
// @Failure 500 "Internal server error"
// @Router /comment/create [post]
func CreateComment(w http.ResponseWriter, r *http.Request) {
	// Only POST
	if r.Method != http.MethodPost {
		utils.Log("CreateComment", "Method not allowed", errors.New("method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Method not allowed"))
		if err != nil {
			utils.Log("CreateComment", "Unable to write response", err)
		}
		return
	}

	// Get thread ID and comment body from request
	var createCommentRequest models.CreateCommentRequest

	err := json.NewDecoder(r.Body).Decode(&createCommentRequest)

	if err != nil {
		utils.Log("CreateComment", "Unable to decode JSON", err)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Invalid data"))
		if err != nil {
			utils.Log("CreateComment", "Unable to write response", err)
		}
		return
	}

	threadId := createCommentRequest.ThreadId
	body := strings.TrimSpace(createCommentRequest.Body)

	// Ensure comment body is not empty and is not too long
	if len(body) == 0 || len(body) > 3000 {
		utils.Log("CreateComment", "Invalid comment body", errors.New("invalid comment body"))
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Invalid data"))
		if err != nil {
			utils.Log("CreateComment", "Unable to write response", err)
		}
		return
	}

	// Get and verify JWT token from request header
	token := r.Header.Get("Authorization")[7:]

	verifiedUsername, err := utils.VerifyJWT(token)

	if err != nil {
		utils.Log("CreateComment", "Unable to verify JWT token", err)
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Invalid JWT token"))
		if err != nil {
			utils.Log("CreateComment", "Unable to write response", err)
		}
		return
	}

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	defer database.CloseConnection(conn)
	queries := database.New(conn)

	// Format threadId as pgtype.UUID for query
	var pgThreadId pgtype.UUID
	err = pgThreadId.Scan(threadId)
	if err != nil {
		utils.Log("CreateComment", "Unable to scan threadId", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateComment", "Unable to write response", err)
		}
		return
	}

	// Create the comment
	params := database.CreateCommentParams{
		Body:     body,
		Creator:  verifiedUsername,
		ThreadID: pgThreadId,
	}

	pgComment, err := queries.CreateComment(ctx, params)

	if err != nil {
		utils.Log("CreateComment", "Unable to create comment", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateComment", "Unable to write response", err)
		}
		return
	}

	comment := database.FormatPgComment(pgComment)

	// Return comment as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(comment)

	if jsonErr != nil {
		utils.Log("CreateComment", "Unable to encode comment as JSON", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateComment", "Unable to write response", err)
		}
		return
	}

	utils.Log("CreateComment", "Comment created on thread: "+threadId+"by: "+verifiedUsername, nil)

	return
}
