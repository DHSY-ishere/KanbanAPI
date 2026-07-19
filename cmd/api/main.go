package main

import (
	"log"

	"github.com/DHSY-ishere/kanbanAPI/internal/api"
	"github.com/DHSY-ishere/kanbanAPI/internal/config"
	"github.com/DHSY-ishere/kanbanAPI/internal/db"
)

func main() {
	cfg := config.Load()

	database, err := db.New(cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	server := api.NewServer(cfg, database)
	log.Printf("listening on %s", cfg.Addr)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
