package user

import (
	"backend/tutorial"
	"backend/utils"
	"context"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

// CreateUser Handles registration requests
func CreateUser(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	// Only POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get username and password from request body
	username := r.FormValue("username")
	password := r.FormValue("password")

	username = strings.TrimSpace(username)

	utils.Log("CreateUser", "Username: "+username+", Password: "+password, nil)

	// Connect to database
	ctx := context.Background()
	queries := tutorial.New(conn)

	// Check if username exists
	isExistingUser, err := queries.CheckUserExists(ctx, username)

	if err != nil {
		utils.Log("CreateUser", "Unable to check if user exists: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if isExistingUser {
		utils.Log("CreateUser", "Username already exists: "+username, errors.New("username already exists"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.Log("CreateUser", "Unable to hash password: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = queries.CreateUser(ctx, tutorial.CreateUserParams{
		Username: username,
		Password: string(hashedPassword)})

	if err != nil {
		utils.Log("CreateUser", "Unable to create user: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Generate JWT token
	token, err := utils.CreateJWT(username)
	if err != nil {
		utils.Log("CreateUser", "Unable to generate JWT token: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return username and token as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(AuthResponseJson{
		Username: username,
		Token:    token})

	if jsonErr != nil {
		utils.Log("CreateUser", "Unable to encode JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Log("CreateUser", "User created successfully: "+username, nil)

	return
}
