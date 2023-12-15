package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

// SECRETSTRING TODO: Environment variable
var SECRETSTRING = "cvwo-e8e2bda6-a005-41e2-8be7-2bae12888bba"

var SECRET = []byte(SECRETSTRING)

// CreateJWT Creates a new JWT token with the username as the payload. Valid for 24 hours.
func CreateJWT(username string) (string, error) {
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	signedToken, err := tokenWithClaims.SignedString(SECRET)

	if err != nil {
		log.Println("[ERROR] Error signing token: ", err)
		return "", err
	}
	return signedToken, nil
}

// VerifyJWT Verifies the JWT token and returns the username.
func VerifyJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Println("[ERROR] Unexpected signing method: ", token.Header["alg"])
			return nil, errors.New("unexpected signing method")
		}
		return SECRET, nil
	})

	if err != nil {
		log.Println("[ERROR] Error parsing token: ", err)
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	}

	return "", err
}
