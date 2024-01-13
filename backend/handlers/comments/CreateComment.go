package comments

import (
	"backend/database"
	"backend/tutorial"
	"backend/utils"
	"context"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
)

type CreateCommentRequestJson struct {
	ThreadId string `json:"thread_id"`
	Body     string `json:"body"`
}

// CreateComment godoc
// @Summary Handles comment creation requests
// @Description Creates a new comment for the given thread
// @Tags comment
// @Accept json
// @Produce json
// @Param data body CreateCommentRequestJson true "Comment data"
// @Security ApiKeyAuth
// @Success 200 {object} tutorial.Comment
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
			return
		}
		return
	}

	// Get thread ID and comment body from request
	var createCommentRequest CreateCommentRequestJson

	err := json.NewDecoder(r.Body).Decode(&createCommentRequest)

	if err != nil {
		utils.Log("CreateComment", "Unable to decode JSON", err)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Invalid data"))
		if err != nil {
			utils.Log("CreateComment", "Unable to write response", err)
			return
		}
		return
	}

	threadId := createCommentRequest.ThreadId
	body := createCommentRequest.Body

	// Get and verify JWT token from request header
	token := r.Header.Get("Authorization")[7:]

	verifiedUsername, err := utils.VerifyJWT(token)

	if err != nil {
		utils.Log("CreateComment", "Unable to verify JWT token", err)
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Invalid JWT token"))
		if err != nil {
			utils.Log("CreateComment", "Unable to write response", err)
			return
		}
		return
	}

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	defer database.CloseConnection(conn)
	queries := tutorial.New(conn)

	// Format threadId as pgtype.UUID for query
	var pgThreadId pgtype.UUID
	err = pgThreadId.Scan(threadId)
	if err != nil {
		utils.Log("CreateComment", "Unable to scan threadId", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateComment", "Unable to write response", err)
			return
		}
		return
	}

	// Create the comment
	params := tutorial.CreateCommentParams{
		Body:     body,
		Creator:  verifiedUsername,
		ThreadID: pgThreadId,
	}

	comment, err := queries.CreateComment(ctx, params)

	if err != nil {
		utils.Log("CreateComment", "Unable to create comment", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateComment", "Unable to write response", err)
			return
		}
	}

	// Return comment as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(comment)

	if jsonErr != nil {
		utils.Log("CreateComment", "Unable to encode comment as JSON", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateComment", "Unable to write response", err)
			return
		}
	}

	utils.Log("CreateComment", "Comment created on thread: "+threadId+"by: "+verifiedUsername, nil)

	return
}
