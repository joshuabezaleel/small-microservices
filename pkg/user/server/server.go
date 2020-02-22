package server

import (
	"log"
	"net/http"

	"github.com/joshuabezaleel/small-microservices/pkg/user/user"

	"github.com/gorilla/mux"
)

// Server holds dependencies for the User HTTP server.
type Server struct {
	userService user.Service

	Router *mux.Router
}

// NewServer returns a new User HTTP server
// with all of the necessary dependencies.
func NewServer(userService user.Service) *Server {
	server := &Server{
		userService: userService,
	}

	userHandler := userHandler{userService}

	router := mux.NewRouter()

	userHandler.registerRouter(router)

	server.Router = router

	return server
}

// Run runs the HTTP server with the specified port and router.
func (srv *Server) Run() {
	var port = ":8082"

	log.Println("userService is running on port " + port)
	err := http.ListenAndServe(port, srv.Router)
	if err != nil {
		panic(err)
	}
}
