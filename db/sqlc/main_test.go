package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	DB_DRIVER = "postgres"
	DB_SOURCE = "postgresql://root:secret@127.0.0.1:5432/shop_management?sslmode=disable"
)

func TestMain(m *testing.M) {
	// config, err := util.LoadConfig("../..")
	// if err != nil {
	// 	log.Fatal("cannot cload config: ", err)
	// }
	var err error
	testDB, err = sql.Open(DB_DRIVER, DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
