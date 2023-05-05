package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Dev-El-badry/wallet-system/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	conf, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("can't load config file", err)
	}

	testDB, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("can not connect to db", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
