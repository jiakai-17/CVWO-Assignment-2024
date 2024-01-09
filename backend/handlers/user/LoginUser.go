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
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} AuthResponseJson
// @Failure 400 "Username does not exist"
// @Failure 401 "Incorrect password"
// @Failure 500
// @Router /user/login [post]
func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Only POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	var creds models.AuthRequestJson
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		utils.Log("LoginUser", "Unable to decode JSON: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid data"))
		return
	}

	// Get username and password from request json
	username := creds.Username
	password := creds.Password

	username = strings.TrimSpace(username)

	utils.Log("LoginUser", "Username: "+username+", Password: "+password, nil)

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	queries := tutorial.New(conn)

	// Check if username exists
	isExistingUser, err := queries.CheckUserExists(ctx, username)

	if err != nil {
		utils.Log("LoginUser", "Unable to check if user exists: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	if !isExistingUser {
		utils.Log("LoginUser", "Username does not exist: "+username, errors.New("username does not exist"))
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid username/password"))
		return
	}

	// Check password
	user, err := queries.GetPasswordHash(ctx, username)
	if err != nil {
		utils.Log("LoginUser", "Unable to get password hash: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	hashedPassword := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		utils.Log("LoginUser", "Incorrect password: ", err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Incorrect username/password"))
		return
	}

	// Generate JWT token
	token, err := utils.CreateJWT(user.Username)
	if err != nil {
		utils.Log("LoginUser", "Unable to create JWT token: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Return username and token as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(models.AuthResponseJson{
		Username: user.Username,
		Token:    token})

	if jsonErr != nil {
		utils.Log("LoginUser", "Unable to encode JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	utils.Log("LoginUser", "User logged in: "+user.Username, nil)
	return
}
