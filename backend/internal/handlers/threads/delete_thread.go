package threads

import (
	"backend/internal/database"
	"backend/internal/utils"
	"context"
	"errors"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
)

// DeleteThread godoc
// @Summary Handles thread deletion requests
// @Description Deletes the thread with the given ID
// @Tags thread
// @Accept json
// @Produce json
// @Param id path string true "Thread ID"
// @Security ApiKeyAuth
// @Success 200
// @Failure 401 "Invalid JWT token"
// @Failure 403 "No permission to delete thread"
// @Failure 405 "Method not allowed"
// @Failure 500 "Internal server error"
// @Router /thread/{id} [delete]
func DeleteThread(w http.ResponseWriter, r *http.Request) {
	// Only DELETE
	if r.Method != http.MethodDelete {
		utils.Log("DeleteThread", "Method not allowed", errors.New("method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Method not allowed"))
		if err != nil {
			utils.Log("DeleteThread", "Unable to write response", err)
		}
		return
	}

	// Get details from request
	vars := mux.Vars(r)
	threadId := vars["id"]

	// Get and verify JWT token from request header
	token := r.Header.Get("Authorization")[7:]
	verifiedUsername, err := utils.VerifyJWT(token)

	if err != nil {
		utils.Log("DeleteThread", "Unable to verify JWT token", err)
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Invalid JWT token"))
		if err != nil {
			utils.Log("DeleteThread", "Unable to write response", err)
		}
		return
	}

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	defer database.CloseConnection(conn)
	queries := database.New(conn)

	// Create thread UUID for pg
	var pgThreadId pgtype.UUID

	err = pgThreadId.Scan(threadId)

	if err != nil {
		utils.Log("DeleteThread", "Unable to scan threadId", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("DeleteThread", "Unable to write response", err)
		}
		return
	}

	// Check if user is creator of thread
	isThreadCreator, err := queries.CheckThreadCreator(ctx, database.CheckThreadCreatorParams{
		Creator: verifiedUsername,
		ID:      pgThreadId})

	if err != nil || !isThreadCreator {
		utils.Log("DeleteThread", "Unable to check if user is creator of thread", err)
		w.WriteHeader(http.StatusForbidden)
		_, err := w.Write([]byte("No permission to delete thread"))
		if err != nil {
			utils.Log("DeleteThread", "Unable to write response", err)
		}
		return
	}

	// Delete the thread
	err = queries.DeleteThread(ctx, database.DeleteThreadParams{
		ID:      pgThreadId,
		Creator: verifiedUsername,
	})

	if err != nil {
		utils.Log("DeleteThread", "Unable to delete thread", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("DeleteThread", "Unable to write response", err)
		}
		return
	}

	utils.Log("DeleteThread", "Thread deleted: "+threadId+" by: "+verifiedUsername, nil)

	return
}
