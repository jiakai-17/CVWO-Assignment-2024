package user

// JwtJson Provides the layout for the JSON object returned by CreateUser and LoginUser
type JwtJson struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
