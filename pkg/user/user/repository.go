package user

// Repository provides access to the User store.
type Repository interface {
	Get(ID int) (*User, error)
	GetAll() ([]*User, error)
}
