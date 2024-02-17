package repo

import (
	"database/sql"
	"log"
)

func NewDB(dbstring string) *sql.DB {
	db, err := sql.Open("postgres", dbstring)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
