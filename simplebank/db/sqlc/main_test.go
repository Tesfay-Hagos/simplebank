package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	DBDriver = "postgres"
	DBSource = "postgresql://tesfay2f:tsionawi@2121@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testdb *sql.DB

func TestMain(m *testing.M) {
	var err error
	testdb, err = sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testdb)

	os.Exit(m.Run())
}
