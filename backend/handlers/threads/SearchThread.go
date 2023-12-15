package threads

import (
	"backend/tutorial"
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

// SearchThread Handler for /api/v1/searchThread
func SearchThread(w http.ResponseWriter, r *http.Request) {
	// Only POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// How many threads per page
	size := 10
	// Available sorting orders
	availableSortOrders := []string{"created_time_asc", "created_time_desc", "num_comments_asc", "num_comments_desc"}

	// Get details from request body
	queryString := r.FormValue("query")
	page := r.FormValue("page")
	order := r.FormValue("order")

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

	conn, err := pgx.Connect(ctx, "user=postgres dbname=cvwo-1 password=cs2102")
	if err != nil {
		log.Println("[ERROR] Unable to connect to database: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close(ctx)

	queries := tutorial.New(conn)

	keywords := strings.Split(queryString, " ")

	if queryString == "" || len(keywords) == 0 {
		// Return all threads
		// Get threads
		log.Println("[INFO] Getting all threads", order)
		threads, err := queries.GetThreads(ctx, tutorial.GetThreadsParams{
			Limit:     int32(size),
			Offset:    int32(offset),
			Sortorder: order,
		})

		if err != nil {
			log.Println("[ERROR] Unable to get threads: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		jsonErr := json.NewEncoder(w).Encode(threads)

		if jsonErr != nil {
			log.Println("[ERROR] Unable to encode threads as JSON: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		return

	} else {

		var parsedKeywords []string
		var tags []string

		for _, keyword := range keywords {
			if strings.HasPrefix(keyword, "tag:") {
				tags = append(tags, keyword[4:])
			} else {
				parsedKeywords = append(parsedKeywords, keyword)
			}
		}

		if len(parsedKeywords) > 0 && len(tags) > 0 {
			// Complex Search with tags and text
			log.Println("[INFO] Search by both keyword and tag", order)

			threads, err := queries.GetThreadsByMultipleTagsAndKeyword(ctx,
				tutorial.GetThreadsByMultipleTagsAndKeywordParams{
					Limit:     int32(size),
					Offset:    int32(offset),
					Sortorder: order,
					Keywords:  strings.Join(parsedKeywords, " & "),
					Tagarray:  tags,
				})

			if err != nil {
				log.Println("[ERROR] Unable to get threads: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			jsonErr := json.NewEncoder(w).Encode(threads)

			if jsonErr != nil {
				log.Println("[ERROR] Unable to encode threads as JSON: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			return
		} else if len(parsedKeywords) > 0 {
			// Search with text only
			log.Println("[INFO] Search by keyword only", order)

			threads, err := queries.GetThreadsByMultipleKeyword(ctx, tutorial.GetThreadsByMultipleKeywordParams{
				Limit:     int32(size),
				Offset:    int32(offset),
				Sortorder: order,
				Keywords:  strings.Join(parsedKeywords, " & "),
			})

			if err != nil {
				log.Println("[ERROR] Unable to get threads: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			jsonErr := json.NewEncoder(w).Encode(threads)

			if jsonErr != nil {
				log.Println("[ERROR] Unable to encode threads as JSON: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			return
		} else {
			// Search with tags only

			log.Println("[INFO] Search by tag only", order)

			threads, err := queries.GetThreadsByMultipleTags(ctx, tutorial.GetThreadsByMultipleTagsParams{
				Limit:     int32(size),
				Offset:    int32(offset),
				Sortorder: order,
				Tagarray:  tags,
			})

			if err != nil {
				log.Println("[ERROR] Unable to get threads: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			jsonErr := json.NewEncoder(w).Encode(threads)

			if jsonErr != nil {
				log.Println("[ERROR] Unable to encode threads as JSON: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			return
		}
	}
}
