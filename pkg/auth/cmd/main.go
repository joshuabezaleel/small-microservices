package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/joshuabezaleel/small-microservices/pkg/auth/auth"
	"github.com/joshuabezaleel/small-microservices/pkg/auth/pkg"
	"github.com/joshuabezaleel/small-microservices/pkg/auth/server"
)

func main() {
	err := godotenv.Load("build/.env")
	if err != nil {
		panic(err)
	}

	userService := pkg.UserServiceClient{
		Client:          http.DefaultClient,
		UserServiceHost: os.Getenv("USER_SERVICE_HOST"),
		UserServicePort: ":" + os.Getenv("USER_SERVICE_PORT"),
	}

	authService := auth.NewAuthService(userService)

	srv := server.NewServer(authService)
	srv.Run(os.Getenv("AUTH_SERVICE_PORT"))
}
