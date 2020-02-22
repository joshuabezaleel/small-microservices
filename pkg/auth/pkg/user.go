package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// User domain model.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserServiceClient creates an instance of HTTP client
// for accessing user service.
type UserServiceClient struct {
	Client          *http.Client
	UserServiceAddr string
}

// Get retrieves user using user service with specified username.
func (us *UserServiceClient) Get(ctx context.Context, username string) (*User, error) {
	var user *User

	url := fmt.Sprintf("http://localhost%s/users/%s", us.UserServiceAddr, username)
	req, _ := http.NewRequest("GET", url, nil)
	req = req.WithContext(ctx)

	// urlWithNoPort := fmt.Sprintf("http://users/%s", username)
	// req, _ := http.NewRequest("GET", urlWithNoPort, nil)
	// req = req.WithContext(ctx)
	log.Println(req)
	log.Println(url)

	resp, err := us.Client.Do(req)
	if err != nil {
		log.Printf("error = %v\n", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("could not get user data: %s", string(body))
	}

	err = json.Unmarshal(body, user)

	return user, nil
}
