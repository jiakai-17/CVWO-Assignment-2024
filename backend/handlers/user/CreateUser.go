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
	"regexp"
	"strings"
)

// CreateUser godoc
// @Summary Handles registration requests
// @Description Registers a new user with the given username and password
// @Tags user
// @Accept json
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} AuthResponseJson
// @Failure 400 "Username already exists"
// @Failure 500
// @Router /user/create [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Only POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	var creds models.AuthRequestJson
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		utils.Log("CreateUser", "Unable to decode JSON: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid data"))
		return
	}

	// Get username and password from request body
	username := creds.Username
	password := creds.Password

	username = strings.TrimSpace(username)

	// Validate username and password
	if len(username) < 1 || len(username) > 30 || len(password) < 6 || regexp.MustCompile(`\s`).MatchString(username) {
		utils.Log("CreateUser", "Invalid username or password: "+username+"; "+password,
			errors.New("invalid username or password"))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid username or password"))
		return
	}

	utils.Log("CreateUser", "Username: "+username+", Password: "+password, nil)

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	queries := tutorial.New(conn)

	// Check if username exists
	isExistingUser, err := queries.CheckUserExists(ctx, username)

	if err != nil {
		utils.Log("CreateUser", "Unable to check if user exists: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	if isExistingUser {
		utils.Log("CreateUser", "Username already exists: "+username, errors.New("username already exists"))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username already exists"))
		return
	}

	// Create user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.Log("CreateUser", "Unable to hash password: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	err = queries.CreateUser(ctx, tutorial.CreateUserParams{
		Username: username,
		Password: string(hashedPassword)})

	if err != nil {
		utils.Log("CreateUser", "Unable to create user: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Generate JWT token
	token, err := utils.CreateJWT(username)
	if err != nil {
		utils.Log("CreateUser", "Unable to generate JWT token: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Return username and token as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(models.AuthResponseJson{
		Username: username,
		Token:    token})

	if jsonErr != nil {
		utils.Log("CreateUser", "Unable to encode JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	utils.Log("CreateUser", "User created successfully: "+username, nil)

	return
}
