package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	postgresDNS = "postgres://root:secret@localhost:5455/blog_test?sslmode=disable"
	driver      = "postgres"
)

var (
	testDb *sql.DB
)

func TestMain(m *testing.M) {
	dbCon, err := sql.Open(driver, postgresDNS)

	if err != nil {
		log.Fatal(err)
	}

	err = dbCon.Ping()

	if err != nil {
		log.Fatal(err)
	}

	testDb = dbCon

	code := m.Run()

	err = testDb.Close()

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}
