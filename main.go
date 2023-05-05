package main

import (
	"database/sql"
	"log"

	"github.com/Dev-El-badry/wallet-system/api"
	db "github.com/Dev-El-badry/wallet-system/db/sqlc"
	"github.com/Dev-El-badry/wallet-system/util"

	_ "github.com/lib/pq"
)

func main() {
	conf, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("can't load config file", err)
	}

	conn, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(conf, store)
	if err != nil {
		log.Fatal("cann't create a server", err)
	}

	err = server.Start(conf.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
