package persistence

import (
	"errors"

	"github.com/joshuabezaleel/small-microservices/pkg/user/user"
)

var allUsers = []*user.User{
	&user.User{ID: 1, Username: "joshua", Password: "joshua"},
	&user.User{ID: 2, Username: "zaky", Password: "zaky"},
	&user.User{ID: 3, Username: "gilang", Password: "gilang"},
}

type userRepository struct{}

// NewUserRepository returns initialized implementation
// of the repository for User domain model.
func NewUserRepository() user.Repository {
	return &userRepository{}
}

func (repo *userRepository) Get(ID int) (*user.User, error) {
	for _, eachUser := range allUsers {
		if eachUser.ID == ID {
			return eachUser, nil
		}
	}

	return nil, errors.New("User not found")
}

func (repo *userRepository) GetByUsername(username string) (*user.User, error) {
	for _, eachUser := range allUsers {
		if eachUser.Username == username {
			return eachUser, nil
		}
	}

	return nil, errors.New("User not found")
}

func (repo *userRepository) GetAll() ([]*user.User, error) {
	return allUsers, nil
}
