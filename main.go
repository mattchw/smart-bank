package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mattchw/smart-bank/api"
	db "github.com/mattchw/smart-bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://postgres:test1234@localhost:5432/smart_bank?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Fail to connect to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Fail to start server:", err)
	}
}
