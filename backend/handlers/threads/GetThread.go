package threads

import (
	"backend/database"
	"backend/tutorial"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"net/http"
)

// GetThread Handler for /api/v1/getThread
func GetThread(w http.ResponseWriter, r *http.Request) {
	// Only GET
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get details from request url
	vars := mux.Vars(r)
	id := vars["id"]

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	defer conn.Close(ctx)
	queries := tutorial.New(conn)

	var threadUUID pgtype.UUID
	threadUUID.Scan(id)

	// Create the thread
	thread, err := queries.GetThreadDetails(ctx, threadUUID)

	if err != nil {
		if err.Error() == "no rows in result set" {
			log.Println("[INFO] Thread not found: ", err)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Thread not found"))
			return
		} else {
			log.Println("[ERROR] Unable to get thread: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
			return
		}
	}

	// Return comment as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(thread)

	if jsonErr != nil {
		log.Println("[ERROR] Unable to encode thread as JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
}
