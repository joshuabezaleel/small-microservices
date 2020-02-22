package auth

// LoginRequest reflects the username and password
// input by user when logging in.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
