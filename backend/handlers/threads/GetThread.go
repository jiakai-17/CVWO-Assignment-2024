package threads

import (
	"backend/tutorial"
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"net/http"
)

// GetThread Handler for /api/v1/getThread
func GetThread(w http.ResponseWriter, r *http.Request) {
	// Only POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get details from request body
	id := r.FormValue("id")

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

	// Create the thread
	thread, err := queries.GetThreadDetails(ctx, threadUUID)

	if err != nil {
		log.Println("[ERROR] Unable to get thread: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return comment as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(thread)

	if jsonErr != nil {
		log.Println("[ERROR] Unable to encode thread as JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
