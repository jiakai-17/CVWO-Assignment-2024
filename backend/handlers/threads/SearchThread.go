package threads

import (
	"backend/database"
	"backend/tutorial"
	"backend/utils"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

type Pagination struct {
	TotalPages  int `json:"total_pages"`
	CurrentPage int `json:"current_page"`
}

type SearchThreadResponse struct {
	TotalThreads int                                `json:"total_threads"`
	Threads      []tutorial.GetThreadsByCriteriaRow `json:"threads"`
}

// SearchThread godoc
// @Summary Handles thread search requests
// @Description Retrieves threads matching the given query
// @Tags thread
// @Accept json
// @Produce json
// @Param q query string true "Search query"
// @Param order query string false "Sorting order, default 'created_time_desc'" Enums(created_time_asc, created_time_desc, num_comments_asc, num_comments_desc)
// @Param p query string false "Page number, default '1'"
// @Success 200 {object} SearchThreadResponse
// @Failure 405 "Method not allowed"
// @Failure 500 "Internal server error"
// @Router /thread/search [get]
func SearchThread(w http.ResponseWriter, r *http.Request) {
	// Only GET
	if r.Method != http.MethodGet {
		utils.Log("SearchThread", "Method not allowed", errors.New("method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Method not allowed"))
		if err != nil {
			utils.Log("SearchThread", "Unable to write response", err)
			return
		}
		return
	}

	pageSize := 10
	availableSortOrders := []string{"created_time_asc", "created_time_desc", "num_comments_asc", "num_comments_desc"}

	// Get details from request
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
		offset = (pageNumber - 1) * pageSize
	}

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	defer database.CloseConnection(conn)
	queries := tutorial.New(conn)

	keywords := strings.Split(queryString, " ")

	var parsedKeywords []string
	var parsedTagArray []string

	for _, keyword := range keywords {
		if strings.HasPrefix(keyword, "tag:") {
			parsedTagArray = append(parsedTagArray, keyword[4:])
		} else {
			parsedKeywords = append(parsedKeywords, keyword)
		}
	}

	formattedKeywords := strings.Join(parsedKeywords, " & ")

	// Get threads
	threads, err := queries.GetThreadsByCriteria(ctx, tutorial.GetThreadsByCriteriaParams{
		Limit:     int32(pageSize),
		Offset:    int32(offset),
		Sortorder: order,
		Keywords:  formattedKeywords,
		Tagarray:  parsedTagArray,
	})

	if err != nil {
		utils.Log("SearchThread", "Unable to get threads", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("SearchThread", "Unable to write response", err)
			return
		}
		return
	}

	totalThreads, err := queries.GetThreadsByCriteriaCount(ctx, tutorial.GetThreadsByCriteriaCountParams{
		Keywords: formattedKeywords,
		Tagarray: parsedTagArray,
	})

	if err != nil {
		utils.Log("SearchThread", "Unable to get threads count", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("SearchThread", "Unable to write response", err)
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(SearchThreadResponse{int(totalThreads), threads})

	if jsonErr != nil {
		utils.Log("SearchThread", "Unable to encode threads as JSON", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("SearchThread", "Unable to write response", err)
			return
		}
		return
	}

	utils.Log("SearchThread", "Threads retrieved for query: "+queryString, nil)

	return
}
