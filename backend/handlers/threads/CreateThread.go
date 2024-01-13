package threads

import (
	"backend/database"
	"backend/tutorial"
	"backend/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"net/http"
)

type CreateThreadRequestJson struct {
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

// CreateThread godoc
// @Summary Handles thread creation requests
// @Description Creates a new thread
// @Tags thread
// @Accept json
// @Produce json
// @Param data body CreateThreadRequestJson true "Thread data"
// @Security ApiKeyAuth
// @Success 200 {object} tutorial.GetThreadDetailsRow
// @Failure 400 "Invalid data"
// @Failure 401 "Invalid JWT token"
// @Failure 405 "Method not allowed"
// @Failure 413 "Input too large"
// @Failure 500 "Internal server error"
// @Router /thread/create [post]
func CreateThread(w http.ResponseWriter, r *http.Request) {
	// Only POST
	if r.Method != http.MethodPost {
		utils.Log("CreateThread", "Method not allowed", errors.New("method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Method not allowed"))
		if err != nil {
			utils.Log("CreateThread", "Unable to write response", err)
			return
		}
		return
	}

	// Get details from request
	var threadCreate CreateThreadRequestJson
	err := json.NewDecoder(r.Body).Decode(&threadCreate)
	if err != nil {
		utils.Log("CreateThread", "Unable to decode JSON", err)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Malformed JSON"))
		if err != nil {
			utils.Log("CreateThread", "Unable to write response", err)
			return
		}
		return
	}

	title := threadCreate.Title
	body := threadCreate.Body

	if len(threadCreate.Tags) > 3 {
		utils.Log("CreateThread", "Too many tags", errors.New("too many tags"))
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		_, err := w.Write([]byte("Input too large"))
		if err != nil {
			utils.Log("CreateThread", "Unable to write response", err)
			return
		}
		return
	}

	if title == "" || body == "" {
		utils.Log("CreateThread", "Title or body is empty",
			errors.New("title or body is empty"))
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Invalid data"))
		if err != nil {
			utils.Log("CreateThread", "Unable to write response", err)
			return
		}
		return
	}

	// Get and verify JWT token from request header
	token := r.Header.Get("Authorization")[7:]
	verifiedUsername, err := utils.VerifyJWT(token)

	if err != nil {
		utils.Log("CreateThread", "Unable to verify JWT token", err)
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Invalid JWT token"))
		if err != nil {
			utils.Log("CreateThread", "Unable to write response", err)
			return
		}
		return
	}

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	defer database.CloseConnection(conn)
	queries := tutorial.New(conn)

	// Begin a new transaction
	tx, err := conn.Begin(ctx)
	if err != nil {
		utils.Log("CreateThread", "Unable to begin transaction", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateThread", "Unable to write response", err)
			return
		}
		return
	}

	var hasCommitted = false

	defer func(tx pgx.Tx, ctx context.Context) {
		if hasCommitted {
			return
		}
		err := tx.Rollback(ctx)
		if err != nil {
			utils.Log("CreateThread", "Unable to rollback transaction", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("Internal server error"))
			if err != nil {
				utils.Log("CreateThread", "Unable to write response", err)
				return
			}
		}
	}(tx, ctx)

	qtx := queries.WithTx(tx)

	// Create the thread
	thread, err := qtx.CreateThread(ctx, tutorial.CreateThreadParams{
		Creator: verifiedUsername,
		Title:   title,
		Body:    body})

	if err != nil {
		utils.Log("CreateThread", "Unable to create thread", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateThread", "Unable to write response", err)
			return
		}
		return
	}

	var pgThreadId = thread.ID

	// Create tags for the thread
	err = qtx.AddNewTags(ctx, threadCreate.Tags)
	if err != nil {
		utils.Log("CreateThread", "Unable to create tags", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateThread", "Unable to write response", err)
			return
		}
		return
	}

	err = qtx.AddThreadTags(ctx, tutorial.AddThreadTagsParams{
		ThreadID: pgThreadId,
		Tagarray: threadCreate.Tags})

	if err != nil {
		utils.Log("CreateThread", "Unable to add tags to thread", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateThread", "Unable to write response", err)
			return
		}
		return
	}

	err = tx.Commit(ctx)

	if err != nil {
		utils.Log("CreateThread", "Unable to commit transaction", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateThread", "Unable to write response", err)
			return
		}
		return
	}

	hasCommitted = true

	createdThread, err := queries.GetThreadDetails(ctx, pgThreadId)

	if err != nil {
		utils.Log("CreateThread", "Unable to get thread details", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateThread", "Unable to write response", err)
			return
		}
		return
	}

	// Return thread as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(createdThread)

	if jsonErr != nil {
		utils.Log("CreateThread", "Unable to encode thread as JSON", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateThread", "Unable to write response", err)
			return
		}
		return
	}

	// Adapted from https://stackoverflow.com/a/71134336
	var threadId = fmt.Sprintf("%x-%x-%x-%x-%x",
		pgThreadId.Bytes[0:4],
		pgThreadId.Bytes[4:6],
		pgThreadId.Bytes[6:8],
		pgThreadId.Bytes[8:10],
		pgThreadId.Bytes[10:16])

	utils.Log("CreateThread", "Thread: "+threadId+" created by: "+verifiedUsername, nil)

	return
}
