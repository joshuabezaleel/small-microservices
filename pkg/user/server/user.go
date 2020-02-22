package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/joshuabezaleel/small-microservices/pkg/user/user"

	"github.com/gorilla/mux"
)

type userHandler struct {
	userService user.Service
}

func (handler *userHandler) registerRouter(router *mux.Router) {
	router.HandleFunc("/users", handler.getAllUsers).Methods("GET")
	router.HandleFunc("/users/{ID}", handler.getUser).Methods("GET")
}

func (handler *userHandler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := handler.userService.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

func (handler *userHandler) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDString, ok := vars["ID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid URL path")
		return
	}

	userID, _ := strconv.Atoi(userIDString)

	user, err := handler.userService.Get(userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)
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
