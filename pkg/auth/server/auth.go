package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshuabezaleel/small-microservices/pkg/auth/auth"
)

type authHandler struct {
	authService auth.Service
}

func (handler *authHandler) registerRouter(router *mux.Router) {
	router.HandleFunc("/login", handler.login).Methods("POST")
}

func (handler *authHandler) login(w http.ResponseWriter, r *http.Request) {
	loginReq := auth.LoginRequest{}

	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	ctx := r.Context()

	isAuthorized, err := handler.authService.Login(ctx, &loginReq)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !isAuthorized {
		respondWithJSON(w, http.StatusOK, "Failed logging in")
		return
	}

	respondWithJSON(w, http.StatusOK, "Logged in!")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"Error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
