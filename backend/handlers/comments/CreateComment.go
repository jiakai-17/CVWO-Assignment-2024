package comments

import (
	"backend/database"
	"backend/tutorial"
	"backend/utils"
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"net/http"
)

// CreateComment godoc
// @Summary Handles comment creation requests
// @Description Creates a new comment for the given thread
// @Tags comment
// @Accept json
// @Produce json
// @Param username formData string true "Username"
// @Param thread formData string true "Thread UUID"
// @Param body formData string true "Comment body"
// @Success 200 "JSON of Created comment"
// @Failure 401 "Invalid JWT token"
// @Failure 500
// @Router /comment/create [post]
func CreateComment(w http.ResponseWriter, r *http.Request) {
	// Only POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	type CommentCreate struct {
		ThreadId string `json:"thread_id"`
		Body     string `json:"body"`
	}

	var commentCreate CommentCreate

	err := json.NewDecoder(r.Body).Decode(&commentCreate)

	if err != nil {
		log.Println("[ERROR] Unable to decode JSON: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Malformed JSON"))
		return
	}

	thread := commentCreate.ThreadId
	body := commentCreate.Body

	// Get JWT token from request header
	token := r.Header.Get("Authorization")

	// Remove "Bearer " from token
	token = token[7:]

	log.Println("[DEBUG] Token: ", token)

	// Verify token
	verifiedUsername, err := utils.VerifyJWT(token)

	if err != nil {
		log.Println("[ERROR] Unable to verify JWT token: ", err, verifiedUsername)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	queries := tutorial.New(conn)

	// Create thread UUID for pg
	var threadID pgtype.UUID

	threadID.Scan(thread)

	// Create the comment
	params := tutorial.CreateCommentParams{
		Body:     body,
		Creator:  verifiedUsername,
		ThreadID: threadID,
	}

	comment, err := queries.CreateComment(ctx, params)

	if err != nil {
		log.Println("[ERROR] Unable to create comment: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return comment as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(comment)

	if jsonErr != nil {
		log.Println("[ERROR] Unable to encode comment as JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
