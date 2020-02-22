package persistence

import (
	"errors"

	"github.com/joshuabezaleel/small-microservices/pkg/user/user"
)

func populateUsers() []*user.User {
	users := make([]*user.User, 3)
	users[0] = &user.User{ID: 1, Username: "joshua", Password: "joshua"}
	users[1] = &user.User{ID: 2, Username: "zaky", Password: "zaky"}
	users[2] = &user.User{ID: 3, Username: "gilang", Password: "gilang"}

	return users
}

// var users = make([]user.User, 3)
// users = append(users, user.User{ID: 0, Username: "joshua", Password: "joshua"})

type userRepository struct{}

// NewUserRepository returns initialized implementation
// of the repository for User domain model.
func NewUserRepository() user.Repository {
	return &userRepository{}
}

func (repo *userRepository) Get(ID int) (*user.User, error) {
	allUsers := populateUsers()

	for _, eachUser := range allUsers {
		if eachUser.ID == ID {
			return eachUser, nil
		}
	}

	return nil, errors.New("User not found")
}

func (repo *userRepository) GetByUsername(username string) (*user.User, error) {
	allUsers := populateUsers()

	for _, eachUser := range allUsers {
		if eachUser.Username == username {
			return eachUser, nil
		}
	}

	return nil, errors.New("User not found")
}

func (repo *userRepository) GetAll() ([]*user.User, error) {
	allUsers := populateUsers()

	return allUsers, nil
}
