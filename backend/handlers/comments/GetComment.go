package comments

import (
	"backend/tutorial"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"net/http"
	"slices"
	"strconv"
)

// GetComment godoc
// @Summary Handles comment retrieval requests
// @Description Retrieves comments for the given thread
// @Tags comment
// @Param thread_id path string true "Thread UUID"
// @Param order query string false "Sorting order" Enums(created_time_asc, created_time_desc)
// @Param page query string false "Page number"
// @Success 200 "JSON array of comments"
// @Failure 500
// @Router /thread/{thread_id}/comments [get]
func GetComment(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	// Only GET
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// How many comments per page
	size := 10

	// Get details from request body
	id := mux.Vars(r)["thread_id"]
	order := r.FormValue("order")
	page := r.FormValue("page")

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
		offset = (pageNumber - 1) * size
	}

	// Connect to database
	ctx := context.Background()
	queries := tutorial.New(conn)

	var threadUUID pgtype.UUID
	threadUUID.Scan(id)

	// Get the comments
	comments, err := queries.GetComments(ctx, tutorial.GetCommentsParams{
		ThreadID:  threadUUID,
		Sortorder: order,
		Offset:    int32(offset),
		Limit:     int32(size),
	})

	if err != nil {
		log.Println("[ERROR] Unable to get comments of thread: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return comments as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(comments)

	if jsonErr != nil {
		log.Println("[ERROR] Unable to encode comments as JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
