package user

// AuthResponseJson Provides the layout for the JSON object returned by CreateUser and LoginUser
type AuthResponseJson struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
