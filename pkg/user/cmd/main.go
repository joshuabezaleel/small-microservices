package main

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/joshuabezaleel/small-microservices/pkg/user/persistence"
	"github.com/joshuabezaleel/small-microservices/pkg/user/server"
	"github.com/joshuabezaleel/small-microservices/pkg/user/user"
)

func main() {
	err := godotenv.Load("build/.env")
	if err != nil {
		panic(err)
	}

	userRepository := persistence.NewUserRepository()

	userService := user.NewUserService(userRepository)

	srv := server.NewServer(userService)
	srv.Run(os.Getenv("USER_SERVICE_PORT"))
}
