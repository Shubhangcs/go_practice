package repository

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Sumith02/Practice/internals/models"
	"github.com/Sumith02/Practice/pkg/database"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepo struct{
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo{
	return &AuthRepo{
		db,
	}
}

func (ar *AuthRepo) Register(r *http.Request) error {
	var newUser models.AuthDetails
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		return err
	}
	newUser.UserId = uuid.NewString()
	hash , err := bcrypt.GenerateFromPassword([]byte(newUser.Password) , bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(hash)
	query := database.NewQuery(ar.db)
	if err := query.RegisterQuery(newUser); err != nil {
		return err
	}
	return nil
}

func (ar *AuthRepo) Login() {

}