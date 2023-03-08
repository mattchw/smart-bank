package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mattchw/smart-bank/api"
	db "github.com/mattchw/smart-bank/db/sqlc"
	"github.com/mattchw/smart-bank/util"
)

func main() {
	config, err := util.LoadConfig(".")

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Fail to connect to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Fail to start server:", err)
	}
}
