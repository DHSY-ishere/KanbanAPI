package main

import (
	"log"

	"github.com/DHSY-ishere/KanbanAPI/internal/api"
	"github.com/DHSY-ishere/KanbanAPI/internal/config"
)

func main() {
	cfg := config.Load()

	server := api.NewServer(cfg)
	log.Printf("listening on %s", cfg.Addr)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
