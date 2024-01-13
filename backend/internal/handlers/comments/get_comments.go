package comments

import (
	"backend/internal/database"
	"backend/internal/utils"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
	"slices"
	"strconv"
)

type GetCommentResponse struct {
	Comments []database.Comment `json:"comments"`
	Count    int32              `json:"count"`
}

// GetComments godoc
// @Summary Handles comment retrieval requests
// @Description Retrieves comments for the given thread
// @Tags comment
// @Param thread_id path string true "Thread UUID"
// @Param order query string false "Sorting order, default 'created_time_asc'" Enums(created_time_asc, created_time_desc)
// @Param p query string false "Page number, default '1'"
// @Success 200 {object} GetCommentResponse
// @Failure 405 "Method not allowed"
// @Failure 500 "Internal server error"
// @Router /thread/{thread_id}/comments [get]
func GetComments(w http.ResponseWriter, r *http.Request) {
	// Only GET
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Method not allowed"))
		if err != nil {
			utils.Log("GetComments", "Unable to write response", err)
			return
		}
		return
	}

	// Number of comments per page
	pageSize := 10

	// Get details from request body
	threadId := mux.Vars(r)["thread_id"]
	order := r.FormValue("order")
	page := r.FormValue("p")

	// Available sorting orders
	availableSortOrders := []string{"created_time_asc", "created_time_desc"}

	// Check sort order
	if order == "" || !slices.Contains(availableSortOrders, order) {
		// Default sorting order - latest threads first
		order = "created_time_asc"
	}

	// Check page number
	pageNumber, err := strconv.Atoi(page)
	offset := 0
	if err == nil && pageNumber > 1 {
		offset = (pageNumber - 1) * pageSize
	}

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	defer database.CloseConnection(conn)
	queries := database.New(conn)

	var pgThreadId pgtype.UUID
	err = pgThreadId.Scan(threadId)
	if err != nil {
		utils.Log("GetComments", "Unable to scan threadId", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("GetComments", "Unable to write response", err)
			return
		}
		return
	}

	// Get the comments
	comments, err := queries.GetComments(ctx, database.GetCommentsParams{
		ThreadID:  pgThreadId,
		Sortorder: order,
		Offset:    int32(offset),
		Limit:     int32(pageSize),
	})

	if err != nil {
		utils.Log("GetComments", "Unable to get comments", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("GetComments", "Unable to write response", err)
			return
		}
		return
	}

	commentsCount, err := queries.GetCommentCount(ctx, pgThreadId)

	if err != nil {
		utils.Log("GetComments", "Unable to get comment count", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("GetComments", "Unable to write response", err)
			return
		}
		return
	}

	var response GetCommentResponse
	response.Comments = comments
	response.Count = int32(commentsCount)

	// Return comments as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(response)

	if jsonErr != nil {
		utils.Log("GetComments", "Unable to encode comments as JSON", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("GetComments", "Unable to write response", err)
			return
		}
		return
	}

	utils.Log("GetComments", "Comments retrieved for thread: "+threadId, nil)

	return
}
