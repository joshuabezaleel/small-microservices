package main

import (
	"github.com/joshuabezaleel/small-microservices/pkg/user/persistence"
	"github.com/joshuabezaleel/small-microservices/pkg/user/server"
	"github.com/joshuabezaleel/small-microservices/pkg/user/user"
)

func main() {
	userRepository := persistence.NewUserRepository()

	userService := user.NewUserService(userRepository)

	srv := server.NewServer(userService)
	srv.Run()
}
