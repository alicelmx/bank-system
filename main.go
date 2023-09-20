package main

import (
	"database/sql"
	"github/alicelmx/simplebank/api"
	db "github/alicelmx/simplebank/db/sqlc"
	"github/alicelmx/simplebank/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadCofig(".")
	if err != nil {
		log.Fatal("cannot connect config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
