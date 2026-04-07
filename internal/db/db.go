package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDB(dbURL string) {
	var err error
	DB, err = sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	log.Println("DB connected successfully")
}
