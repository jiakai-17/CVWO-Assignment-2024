package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var SECRET []byte = nil

// InitJwtSecret Initializes the secret used to sign JWT tokens. Must be called before any JWT operations.
func InitJwtSecret() {
	secretString := os.Getenv("JWT_SECRETSTRING")
	if secretString == "" {
		Log("main", "No JWT secret string provided, using 'secretstring'", nil)
		secretString = "secretstring"
	}

	SECRET = []byte(secretString)
}

// CreateJWT Creates a new JWT token with the username as the payload. Valid for 24 hours.
func CreateJWT(username string) (string, error) {
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"iat":      time.Now().Unix(),
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	signedToken, err := tokenWithClaims.SignedString(SECRET)

	if err != nil {
		Log("jwt", "Unable to sign token", err)
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
			Log("jwt", "Unexpected signing method",
				errors.New("got: "+token.Header["alg"].(string)+", expected: HS256"))
			return nil, errors.New("unexpected signing method")
		}
		return SECRET, nil
	})

	if err != nil {
		Log("jwt", "Error parsing token", err)
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	}

	return "", err
}
