package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sulemaanhamza/gobank/api"
	db "github.com/sulemaanhamza/gobank/db/sqlc"
	"github.com/sulemaanhamza/gobank/util"
	"log"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configurations", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to the database", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("Cannot create the server", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start the server", err)
	}
}
