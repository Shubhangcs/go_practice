package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewConnection() *sql.DB {
	con , err := sql.Open("postgres" , "postgresql://biometric_eb0b_user:nIWEea8rwP3JgSbO8Eg84L0bcO1LvZlq@dpg-ctf7kr1u0jms739ju1gg-a.oregon-postgres.render.com/biometric_eb0b")
	if err != nil {
		log.Fatalf("unable to connect to database: %s" , err)
	}
	log.Println("connected to database")
	return con
}