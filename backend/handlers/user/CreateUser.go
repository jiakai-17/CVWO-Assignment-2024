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
// @Param data body models.AuthRequestJson true "Username and password"
// @Success 200 {object} models.AuthResponseJson
// @Failure 400 "Username already exists"
// @Failure 400 "Invalid data"
// @Failure 400 "Incorrect username/password"
// @Failure 405 "Method not allowed"
// @Failure 500 "Internal server error"
// @Router /user/create [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Only POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.Log("CreateUser", "Method not allowed", errors.New("method not allowed"))
		_, err := w.Write([]byte("Method not allowed"))
		if err != nil {
			utils.Log("CreateUser", "Unable to write response", err)
			return
		}
		return
	}

	// Get username and password from request
	var creds models.AuthRequestJson
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		utils.Log("CreateUser", "Unable to decode JSON", err)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Invalid data"))
		if err != nil {
			utils.Log("CreateUser", "Unable to write response", err)
			return
		}
		return
	}

	username := strings.TrimSpace(creds.Username)
	password := creds.Password

	// Validate username and password
	if len(username) < 1 ||
		len(username) > 30 ||
		len(password) < 6 ||
		regexp.MustCompile(`\s`).MatchString(username) {
		utils.Log("CreateUser", "Invalid username or password: "+username+"; "+password,
			errors.New("invalid username or password"))
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Incorrect username/password"))
		if err != nil {
			utils.Log("CreateUser", "Unable to write response", err)
			return
		}
		return
	}

	// Connect to database
	ctx := context.Background()
	conn := database.GetConnection()
	defer database.CloseConnection(conn)
	queries := tutorial.New(conn)

	// Check if username exists
	isExistingUser, err := queries.CheckUserExists(ctx, username)

	if err != nil {
		utils.Log("CreateUser", "Unable to check if user exists: "+username, err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateUser", "Unable to write response", err)
			return
		}
		return
	}

	if isExistingUser {
		utils.Log("CreateUser", "Username already exists: "+username, errors.New("username already exists"))
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Username already exists"))
		if err != nil {
			utils.Log("CreateUser", "Unable to write response", err)
			return
		}
		return
	}

	// Create user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.Log("CreateUser", "Unable to hash password", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateUser", "Unable to write response", err)
			return
		}
		return
	}

	err = queries.CreateUser(ctx, tutorial.CreateUserParams{
		Username: username,
		Password: string(hashedPassword)})

	if err != nil {
		utils.Log("CreateUser", "Unable to create user", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateUser", "Unable to write response", err)
			return
		}
		return
	}

	// Generate JWT token
	token, err := utils.CreateJWT(username)
	if err != nil {
		utils.Log("CreateUser", "Unable to generate JWT token", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateUser", "Unable to write response", err)
			return
		}
		return
	}

	// Return username and token as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(models.AuthResponseJson{
		Username: username,
		Token:    token})

	if jsonErr != nil {
		utils.Log("CreateUser", "Unable to encode JSON", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			utils.Log("CreateUser", "Unable to write response", err)
			return
		}
		return
	}

	utils.Log("CreateUser", "User created successfully: "+username, nil)

	return
}
