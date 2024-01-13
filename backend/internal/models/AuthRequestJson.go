package models

// AuthRequestJson Provides the layout for the JSON object sent by frontend to login or signup
type AuthRequestJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
