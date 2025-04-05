package main

import (
	"log"
	"realtime-chat/internal/common"
	"realtime-chat/internal/db"
	"realtime-chat/internal/server"
)

func main() {
	common.LoadEnv()
	db.ConnectPostgres()
	db.RunMigrations()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
