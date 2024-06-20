package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const (
	postgresDNS = "postgres://root:secret@localhost:5455/blog?sslmode=disable"
	driver      = "postgres"
)

func main() {
	_, err := connectToDB(driver)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Hello World!")
}

func connectToDB(driver string) (*sql.DB, error) {
	db, err := sql.Open(driver, postgresDNS)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
