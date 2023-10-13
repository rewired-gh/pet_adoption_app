package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/rewired-gh/pet_adoption_app/api"
	"github.com/rewired-gh/pet_adoption_app/db/sqlc"
	"github.com/rewired-gh/pet_adoption_app/util"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	store := sqlc.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
