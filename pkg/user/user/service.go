package user

import "errors"

// Errors definition.
var (
	ErrGetAllUsers = errors.New("Error retrieving all users")
	ErrGetUser     = errors.New("Error retrieving user")
)

// Service provides basic operations on User domain model.
type Service interface {
	Get(ID int) (*User, error)
	GetByUsername(username string) (*User, error)
	GetAll() ([]*User, error)
}

type service struct {
	userRepository Repository
}

// NewUserService creates an instance of the service for
// User domain model with all the necessary dependencies.
func NewUserService(userRepository Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) Get(ID int) (*User, error) {
	user, err := s.userRepository.Get(ID)
	if err != nil {
		return nil, ErrGetUser
	}

	return user, nil
}

func (s *service) GetByUsername(username string) (*User, error) {
	user, err := s.userRepository.GetByUsername(username)
	if err != nil {
		return nil, ErrGetUser
	}

	return user, nil
}

func (s *service) GetAll() ([]*User, error) {
	users, err := s.userRepository.GetAll()
	if err != nil {
		return nil, ErrGetAllUsers
	}

	return users, nil
}
