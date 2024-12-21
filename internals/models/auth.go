package models

import "net/http"

type AuthDetails struct{
	UserId string `json:"user_id"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type AuthInterface interface{
	Register(*http.Request) error
	Login()
}