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
func LoginUser(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	// Only POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get username and password from request body
	username := r.FormValue("username")
	password := r.FormValue("password")

	username = strings.TrimSpace(username)

	utils.Log("LoginUser", "Username: "+username+", Password: "+password, nil)

	// Connect to database
	ctx := context.Background()
	queries := tutorial.New(conn)

	// Check if username exists
	isExistingUser, err := queries.CheckUserExists(ctx, username)

	if err != nil {
		utils.Log("LoginUser", "Unable to check if user exists: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !isExistingUser {
		utils.Log("LoginUser", "Username does not exist: "+username, errors.New("username does not exist"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check password
	user, err := queries.GetPasswordHash(ctx, username)
	if err != nil {
		utils.Log("LoginUser", "Unable to get password hash: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hashedPassword := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		utils.Log("LoginUser", "Incorrect password: ", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := utils.CreateJWT(user.Username)
	if err != nil {
		utils.Log("LoginUser", "Unable to create JWT token: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return username and token as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(AuthResponseJson{
		Username: user.Username,
		Token:    token})

	if jsonErr != nil {
		utils.Log("LoginUser", "Unable to encode JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Log("LoginUser", "User logged in: "+user.Username, nil)
	return
}
