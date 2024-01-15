package models

// AuthResponse Provides the layout for the JSON object returned by CreateUser and LoginUser
type AuthResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
