package auth

import (
	"errors"

	"github.com/joshuabezaleel/small-microservices/pkg/auth/pkg"
)

// Errors definition.
var (
	ErrLogin = errors.New("Error logging in user")
	ErrAuth  = errors.New("Wrong combination of username/password")
)

// Service provides basic operations for Auth service.
type Service interface {
	Login(loginReq *LoginRequest) (bool, error)
	// GetStoredPasswordByUsername(username string) (string, error)
	// ComparePassword(incomingPassword, storedPassword string) (bool, error)
}

type service struct {
	userService pkg.UserServiceClient
}

// NewAuthService creates an instance of the service for
// Auth domain model with necessary dependencies.
func NewAuthService(userService pkg.UserServiceClient) Service {
	return &service{
		userService: userService,
	}
}

func (s *service) Login(loginReq *LoginRequest) (bool, error) {
	var user *pkg.User

	user, err := s.userService.Get(loginReq.Username)
	if err != nil {
		return false, ErrLogin
	}

	if user.Password != loginReq.Password {
		return false, ErrAuth
	}

	return true, nil
}

// func (s *service) ComparePassword(incomingPassword, storedPassword string) (bool, error) {

// }
