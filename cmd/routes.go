package main

import (
	"database/sql"

	"github.com/Sumith02/Practice/internals/handlers"
	"github.com/Sumith02/Practice/repository"
	"github.com/gorilla/mux"
)

func Routes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()
	authRepo := repository.NewAuthRepo(db)
	authHandler := handlers.NewAuthHandler(authRepo)
	router.HandleFunc("/register" , authHandler.RegisterHandler)
	return router
}