package debug

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"net/http"
)

// GetToken Handler for /api/v1/getToken
func GetToken(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")

	// Generates a token based on the username
	token, err := utils.CreateJWT(username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return username and token as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(models.AuthResponseJson{Username: username, Token: token})

	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
