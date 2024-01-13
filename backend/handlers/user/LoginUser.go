package user

import (
	"backend/database"
	"backend/models"
	"backend/tutorial"
	"backend/utils"
	"context"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

// LoginUser godoc
// @Summary Handles login requests
// @Description Logs in a user with the given username and password
// @Tags user
// @Accept json
// @Produce json
// @Param data body models.AuthRequestJson true "Username and password"
// @Success 200 {object} models.AuthResponseJson
// @Failure 400 "Invalid data"
// @Failure 401 "Incorrect username/password"
// @Failure 405 "Method not allowed"
// @Failure 500 "Internal server error"
// @Router /user/login [post]
func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Only POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Method not allowed"))
		if err != nil {
			utils.Log("LoginUser", "Unable to write response", err)
			return
		}
		return
	}

	// Get username and password from request
	var creds models.AuthRequestJson
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		utils.Log("LoginUser", "Unable to decode JSON", err)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Invalid data"))
		if err != nil {
			utils.Log("LoginUser", "Unable to write response", err)
			return
		}
		return
	}

	username := strings.TrimSpace(creds.Username)
	password := creds.Password

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	defer database.CloseConnection(conn)
	queries := tutorial.New(conn)

	// Check if username exists
	isExistingUser, err := queries.CheckUserExists(ctx, username)

	if err != nil {
		utils.Log("LoginUser", "Unable to check if user exists", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("LoginUser", "Unable to write response", err)
			return
		}
		return
	}

	if !isExistingUser {
		utils.Log("LoginUser", "Username does not exist: "+username, errors.New("username does not exist"))
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Incorrect username/password"))
		if err != nil {
			utils.Log("LoginUser", "Unable to write response", err)
			return
		}
		return
	}

	// Check password
	user, err := queries.GetPasswordHash(ctx, username)
	if err != nil {
		utils.Log("LoginUser", "Unable to get password hash", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("LoginUser", "Unable to write response", err)
			return
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		utils.Log("LoginUser", "Incorrect password for user: "+username, err)
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Incorrect username/password"))
		if err != nil {
			utils.Log("LoginUser", "Unable to write response", err)
			return
		}
		return
	}

	// Generate JWT token
	token, err := utils.CreateJWT(user.Username)
	if err != nil {
		utils.Log("LoginUser", "Unable to create JWT token", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("LoginUser", "Unable to write response", err)
			return
		}
		return
	}

	// Return username and token as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(models.AuthResponseJson{
		Username: user.Username,
		Token:    token})

	if jsonErr != nil {
		utils.Log("LoginUser", "Unable to encode JSON", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("LoginUser", "Unable to write response", err)
			return
		}
		return
	}

	utils.Log("LoginUser", "User logged in: "+user.Username, nil)
	return
}
