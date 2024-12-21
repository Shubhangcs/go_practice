package database

import (
	"database/sql"

	"github.com/Sumith02/Practice/internals/models"
)

type Query struct{
	db *sql.DB
}

func NewQuery(db *sql.DB) *Query {
	return &Query{
		db,
	}
}

func (q *Query) RegisterQuery(user models.AuthDetails) error {
	_ , err := q.db.Exec("INSERT INTO users(user_id , email , password) VALUES($1,$2,$3)" , user.UserId , user.Email , user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (q *Query) InitilizeDatabase() error {
	queries := []string{
		`CREATE TABLE users(
			user_id VARCHAR(100),
			email VARCHAR(200),
			password VARCHAR(100)
		)`,
	}
	for _ , query := range queries {
		_ , err := q.db.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}