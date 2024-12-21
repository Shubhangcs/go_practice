package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Sumith02/Practice/internals/models"
)

type AuthHandler struct {
	authRepo models.AuthInterface
}

func NewAuthHandler(authRepo models.AuthInterface) *AuthHandler {
	return &AuthHandler{
		authRepo,
	}
}

func (ah *AuthHandler) RegisterHandler(w http.ResponseWriter , r *http.Request) {
	err := ah.authRepo.Register(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message":err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message":"Success"})
}