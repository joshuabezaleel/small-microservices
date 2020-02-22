package main

import (
	"net/http"

	"github.com/joshuabezaleel/small-microservices/pkg/auth/auth"
	"github.com/joshuabezaleel/small-microservices/pkg/auth/pkg"
	"github.com/joshuabezaleel/small-microservices/pkg/auth/server"
)

func main() {
	userService := pkg.UserServiceClient{
		Client:          http.DefaultClient,
		UserServiceAddr: ":8082",
	}

	authService := auth.NewAuthService(userService)

	srv := server.NewServer(authService)
	srv.Run()
}
