package threads

import (
	"backend/database"
	"backend/tutorial"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

type Pagination struct {
	TotalPages  int `json:"total_pages"`
	CurrentPage int `json:"current_page"`
}

// SearchThread Handler for /api/v1/searchThread
func SearchThread(w http.ResponseWriter, r *http.Request) {
	// Only GET
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// TODO Add pagination object to response
	// How many threads per page
	size := 10
	// Available sorting orders
	availableSortOrders := []string{"created_time_asc", "created_time_desc", "num_comments_asc", "num_comments_desc"}

	// Get details from request body
	params := r.URL.Query()
	queryString := params.Get("q")
	page := params.Get("p")
	order := params.Get("order")

	// Check sort order
	if order == "" || !slices.Contains(availableSortOrders, order) {
		// Default sorting order - latest threads first
		order = "created_time_desc"
	}

	// Check page number
	pageNumber, err := strconv.Atoi(page)
	offset := 0
	if err == nil && pageNumber > 1 {
		offset = (pageNumber - 1) * size
	}

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	defer conn.Close(ctx)
	queries := tutorial.New(conn)

	keywords := strings.Split(queryString, " ")

	var parsedKeywords []string
	var tagArray []string

	for _, keyword := range keywords {
		if strings.HasPrefix(keyword, "tag:") {
			tagArray = append(tagArray, keyword[4:])
		} else {
			parsedKeywords = append(parsedKeywords, keyword)
		}
	}

	formattedSearchQuery := strings.Join(parsedKeywords, " & ")

	// Get threads
	log.Println("[INFO] Getting threads", order, formattedSearchQuery, tagArray)
	threads, err := queries.GetThreadsByCriteria(ctx, tutorial.GetThreadsByCriteriaParams{
		Limit:     int32(size),
		Offset:    int32(offset),
		Sortorder: order,
		Keywords:  formattedSearchQuery,
		Tagarray:  tagArray,
	})

	totalThreads, err := queries.GetThreadsByCriteriaCount(ctx, tutorial.GetThreadsByCriteriaCountParams{
		Keywords: formattedSearchQuery,
		Tagarray: tagArray,
	})

	type response struct {
		TotalThreads int                                `json:"total_threads"`
		Threads      []tutorial.GetThreadsByCriteriaRow `json:"threads"`
	}

	if err != nil {
		log.Println("[ERROR] Unable to get threads: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(response{int(totalThreads), threads})

	if jsonErr != nil {
		log.Println("[ERROR] Unable to encode threads as JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
