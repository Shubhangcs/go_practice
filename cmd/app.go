package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Sumith02/Practice/pkg/database"
)


func Start(db *sql.DB) {
	server := &http.Server{
		Addr: ":8000",
		Handler: Routes(db),
	}
	query := database.NewQuery(db)
	err := query.InitilizeDatabase()
	if err != nil {
		log.Fatalf("unable to initilize database: %s" , err)
	}
	log.Println("Server has started...")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("unable to start the server: %s" , err)
	}
}