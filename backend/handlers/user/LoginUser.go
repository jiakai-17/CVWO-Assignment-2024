package user

import (
	"backend/tutorial"
	"backend/utils"
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// Handler for /api/v1/login
func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Only POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get username and password from request body
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Connect to database
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=postgres dbname=cvwo-1 password=cs2102")
	if err != nil {
		log.Println("[ERROR] Unable to connect to database: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close(ctx)

	queries := tutorial.New(conn)

	// Check if username exists
	userExists, err := queries.CheckUserExists(ctx, username)
	log.Println("exists: ", userExists)
	if err != nil {
		log.Println("[ERROR] Unable to check if user exists: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !userExists {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check password
	hashedPassword, err := queries.GetPasswordHash(ctx, username)
	if err != nil {
		log.Println("[ERROR] Unable to get password hash: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println("[ERROR] Incorrect password: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Generate JWT token
	token, err := utils.CreateJWT(username)
	if err != nil {
		log.Println("[ERROR] Unable to generate JWT token: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return username and token as JSON object
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(JwtJson{Username: username, Token: token})

	if jsonErr != nil {
		log.Println("[ERROR] Unable to encode JSON: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	return
}
