package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func DbConnection() *sql.DB {
	URI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"0.0.0.0", 5051, "postgres", "qwerty", "loja")
	db, err := sql.Open("postgres", URI)

	if err != nil {
		panic(err.Error())
	}
	return db
}
