package server

import (
	"log"
	"net/http"

	"github.com/joshuabezaleel/small-microservices/pkg/auth/auth"

	"github.com/gorilla/mux"
)

// Server holds dependencies for the Auth HTTP server.
type Server struct {
	authService auth.Service

	Router *mux.Router
}

// NewServer returns a new Auth HTTP server
// with all of the necessary dependencies.
func NewServer(authService auth.Service) *Server {
	server := &Server{
		authService: authService,
	}

	authHandler := authHandler{authService}

	router := mux.NewRouter()

	authHandler.registerRouter(router)

	server.Router = router

	return server
}

// Run runs the HTTP server with the specified port and router.
func (srv *Server) Run() {
	var port = ":8083"

	log.Println("authService is running on port " + port)
	err := http.ListenAndServe(port, srv.Router)
	if err != nil {
		panic(err)
	}
}
