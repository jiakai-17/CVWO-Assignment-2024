package debug

import (
	"backend/handlers/user"
	"backend/utils"
	"encoding/json"
	"net/http"
)

// VerifyToken Handler for /api/v1/verifyToken
func VerifyToken(w http.ResponseWriter, r *http.Request) {
	// Get token from request body
	token := r.FormValue("token")

	// Verify token
	username, err := utils.VerifyJWT(token)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Returns username and token as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(user.JwtJson{Username: username, Token: token})

	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
