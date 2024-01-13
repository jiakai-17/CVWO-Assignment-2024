package threads

import (
	"backend/database"
	"backend/tutorial"
	"backend/utils"
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
)

// GetThread godoc
// @Summary Handles thread retrieval requests
// @Description Retrieves the thread with the given ID
// @Tags thread
// @Accept json
// @Produce json
// @Param id path string true "Thread ID"
// @Success 200 {object} tutorial.Thread
// @Failure 404 "Thread not found"
// @Failure 405 "Method not allowed"
// @Failure 500 "Internal server error"
// @Router /thread/{id} [get]
func GetThread(w http.ResponseWriter, r *http.Request) {
	// Only GET
	if r.Method != http.MethodGet {
		utils.Log("GetThread", "Method not allowed", errors.New("method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Method not allowed"))
		if err != nil {
			utils.Log("GetThread", "Unable to write response", err)
			return
		}
		return
	}

	// Get details from request url
	vars := mux.Vars(r)
	id := vars["id"]

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	defer database.CloseConnection(conn)
	queries := tutorial.New(conn)

	var pgThreadId pgtype.UUID
	err := pgThreadId.Scan(id)

	if err != nil {
		utils.Log("GetThread", "Unable to scan threadId", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("GetThread", "Unable to write response", err)
			return
		}
	}

	// Create the thread
	thread, err := queries.GetThreadDetails(ctx, pgThreadId)

	if err != nil {
		if err.Error() == "no rows in result set" {
			utils.Log("GetThread", "Thread"+id+" not found", err)
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte("Thread not found"))
			if err != nil {
				utils.Log("GetThread", "Unable to write response", err)
				return
			}
			return
		} else {
			utils.Log("GetThread", "Unable to get thread "+id, err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("Internal server error"))
			if err != nil {
				utils.Log("GetThread", "Unable to write response", err)
				return
			}
			return
		}
	}

	// Return thread as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(thread)

	if jsonErr != nil {
		utils.Log("GetThread", "Unable to encode thread as JSON", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("GetThread", "Unable to write response", err)
			return
		}
		return
	}

	utils.Log("GetThread", "Thread "+id+" retrieved", nil)
}
