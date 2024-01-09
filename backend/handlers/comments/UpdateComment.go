package comments

import (
	"backend/database"
	"backend/tutorial"
	"backend/utils"
	"context"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"net/http"
)

// UpdateComment godoc
// @Summary Handles comment update requests
// @Description Updates a comment
// @Tags comment
// @Param id path string true "Comment UUID"
// @Param body formData string true "Comment body"
// @Success 200
// @Failure 401 "Invalid JWT token"
// @Failure 403 "User is not the creator of the comment"
// @Failure 500
// @Router /comment/{id} [put]
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	// Only PUT
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get details from request body
	commentId := mux.Vars(r)["id"]
	utils.Log("updateComment", "[DEBUG] Comment ID: "+commentId, nil)
	body := r.FormValue("body")

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

	// Create comment UUID for pg
	var commentUUID pgtype.UUID

	commentUUID.Scan(commentId)

	// Check if the user is the creator of the comment
	isCreator, err := queries.CheckCommentCreator(ctx, tutorial.CheckCommentCreatorParams{Creator: verifiedUsername,
		ID: commentUUID})

	if err != nil || !isCreator {
		log.Println("[ERROR] Unable to verify creator: ", err, isCreator)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Update the comment
	err = queries.UpdateComment(ctx, tutorial.UpdateCommentParams{
		Body:    body,
		Creator: verifiedUsername,
		ID:      commentUUID,
	})

	if err != nil {
		log.Println("[ERROR] Unable to update comment: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
