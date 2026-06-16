package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


func Connect() *sql.DB {

	conn := "host=localhost port=5432 user=yiwu password=yiwu123 dbname=yiwu_mall sslmode=disable"


	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}


	err = db.Ping()

	if err != nil {
		panic(err)
	}


	fmt.Println("PostgreSQL connected")

	return db
}