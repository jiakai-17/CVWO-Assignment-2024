package comments

import (
	"backend/tutorial"
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"net/http"
	"slices"
	"strconv"
)

// GetComment Handler for /api/v1/getComment
func GetComment(w http.ResponseWriter, r *http.Request) {
	// Only POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// How many comments per page
	size := 10

	// Get details from request body
	id := r.FormValue("id")
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

	conn, err := pgx.Connect(ctx, "user=postgres dbname=cvwo-1 password=cs2102")
	if err != nil {
		log.Println("[ERROR] Unable to connect to database: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close(ctx)

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
