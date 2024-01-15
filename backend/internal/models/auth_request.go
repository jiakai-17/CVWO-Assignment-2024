package models

// AuthRequest Provides the layout for the JSON object sent by frontend to login or signup
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
