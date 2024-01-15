package threads

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/utils"
	"context"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5"
	"net/http"
	"strings"
)

// CreateThread godoc
// @Summary Handles thread creation requests
// @Description Creates a new thread
// @Tags thread
// @Accept json
// @Produce json
// @Param data body models.CreateThreadRequest true "Thread data"
// @Security ApiKeyAuth
// @Success 200 {object} models.Thread
// @Failure 400 "Invalid data"
// @Failure 401 "Invalid JWT token"
// @Failure 405 "Method not allowed"
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
	var threadCreate models.CreateThreadRequest
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
	var tags []string
	for _, tag := range threadCreate.Tags {
		trimmedTag := strings.TrimSpace(tag)
		if len(trimmedTag) > 0 {
			tags = append(tags, trimmedTag)
		}
	}

	// Check if fields are valid
	if len(tags) > 3 || len(title) == 0 || len(body) == 0 || len(title) > 100 || len(body) > 3000 {
		utils.Log("CreateThread", "Invalid inputs", errors.New("invalid input"))
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
	queries := database.New(conn)

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
	thread, err := qtx.CreateThread(ctx, database.CreateThreadParams{
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

	err = qtx.AddThreadTags(ctx, database.AddThreadTagsParams{
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

	pgCreatedThread, err := queries.GetThreadDetails(ctx, pgThreadId)

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

	createdThread := database.FormatPgThread(pgCreatedThread)

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

	utils.Log("CreateThread", "Thread: "+createdThread.ID+" created by: "+verifiedUsername, nil)

	return
}
