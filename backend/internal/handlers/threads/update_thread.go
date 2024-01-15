package threads

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/utils"
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
	"strings"
)

// UpdateThread godoc
// @Summary Handles thread update requests
// @Description Updates a thread
// @Tags thread
// @Param id path string true "Thread UUID"
// @Param data body models.UpdateThreadRequest true "Thread data"
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200
// @Failure 400 "Invalid data"
// @Failure 401 "Invalid JWT token"
// @Failure 403 "No permission to update thread"
// @Failure 405 "Method not allowed"
// @Failure 500 "Internal server error"
// @Router /thread/{id} [put]
func UpdateThread(w http.ResponseWriter, r *http.Request) {
	// Only PUT
	if r.Method != http.MethodPut {
		utils.Log("UpdateThread", "Method not allowed", errors.New("method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Method not allowed"))
		if err != nil {
			utils.Log("UpdateThread", "Unable to write response", err)
			return
		}
		return
	}

	// Get details from request
	vars := mux.Vars(r)
	threadId := vars["id"]
	var updatedThread models.UpdateThreadRequest
	err := json.NewDecoder(r.Body).Decode(&updatedThread)

	if err != nil {
		utils.Log("UpdateThread", "Unable to decode JSON", err)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Invalid data"))
		if err != nil {
			utils.Log("UpdateThread", "Unable to write response", err)
			return
		}
		return
	}

	title := strings.TrimSpace(updatedThread.Title)
	body := strings.TrimSpace(updatedThread.Body)
	var tags []string
	for _, tag := range updatedThread.Tags {
		trimmedTag := strings.TrimSpace(tag)
		if len(trimmedTag) > 0 {
			tags = append(tags, trimmedTag)
		}
	}

	// Check if fields are valid
	if len(tags) > 3 || len(title) == 0 || len(body) == 0 || len(title) > 100 || len(body) > 3000 {
		utils.Log("UpdateThread", "Invalid inputs", errors.New("invalid input"))
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Invalid data"))
		if err != nil {
			utils.Log("UpdateThread", "Unable to write response", err)
			return
		}
		return
	}

	// Get and verify JWT token from request header
	token := r.Header.Get("Authorization")[7:]
	verifiedUsername, err := utils.VerifyJWT(token)

	if err != nil {
		utils.Log("UpdateThread", "Unable to verify JWT token", err)
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Invalid JWT token"))
		if err != nil {
			utils.Log("UpdateThread", "Unable to write response", err)
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
		utils.Log("UpdateThread", "Unable to begin transaction", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("UpdateThread", "Unable to write response", err)
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
			utils.Log("UpdateThread", "Unable to rollback transaction", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("Internal server error"))
			if err != nil {
				utils.Log("UpdateThread", "Unable to write response", err)
				return
			}
			return
		}
	}(tx, ctx)

	qtx := queries.WithTx(tx)

	// Create thread UUID for pg
	var pgThreadId pgtype.UUID

	err = pgThreadId.Scan(threadId)
	if err != nil {
		utils.Log("UpdateThread", "Unable to scan threadId", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("UpdateThread", "Unable to write response", err)
			return
		}
		return
	}

	// Check if user is creator of thread
	isThreadCreator, err := qtx.CheckThreadCreator(ctx, database.CheckThreadCreatorParams{
		Creator: verifiedUsername,
		ID:      pgThreadId})
	if err != nil || !isThreadCreator {
		utils.Log("UpdateThread", "Unable to check if user is creator of thread", err)
		w.WriteHeader(http.StatusForbidden)
		_, err := w.Write([]byte("No permission to update thread"))
		if err != nil {
			utils.Log("UpdateThread", "Unable to write response", err)
			return
		}
		return
	}

	// Update the thread
	err = qtx.UpdateThread(ctx, database.UpdateThreadParams{
		ID:      pgThreadId,
		Title:   title,
		Body:    body,
		Creator: verifiedUsername})
	if err != nil {
		utils.Log("UpdateThread", "Unable to update thread", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("UpdateThread", "Unable to write response", err)
			return
		}
	}

	// Update the tags
	err = qtx.DeleteThreadTags(ctx, pgThreadId)
	if err != nil {
		utils.Log("UpdateThread", "Unable to delete thread tags", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("UpdateThread", "Unable to write response", err)
			return
		}
	}

	err = qtx.DeleteUnusedTags(ctx)
	if err != nil {
		utils.Log("UpdateThread", "Unable to delete unused tags", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("UpdateThread", "Unable to write response", err)
			return
		}
	}

	err = qtx.AddNewTags(ctx, tags)
	if err != nil {
		utils.Log("UpdateThread", "Unable to add new tags", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("UpdateThread", "Unable to write response", err)
			return
		}
	}

	err = qtx.AddThreadTags(ctx, database.AddThreadTagsParams{
		ThreadID: pgThreadId,
		Tagarray: tags})
	if err != nil {
		utils.Log("UpdateThread", "Unable to add thread tags", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("UpdateThread", "Unable to write response", err)
			return
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		utils.Log("UpdateThread", "Unable to commit transaction", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("UpdateThread", "Unable to write response", err)
			return
		}
		return
	}

	hasCommitted = true

	utils.Log("UpdateThread", "Thread "+threadId+" updated", nil)

	return
}
