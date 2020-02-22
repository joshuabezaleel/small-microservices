package user

// User domain model.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewUser creates a new instance of
// User domain model.
func NewUser(id int, username string, password string) *User {
	return &User{
		ID:       id,
		Username: username,
		Password: password,
	}
}
