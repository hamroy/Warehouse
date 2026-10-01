package main

import (
	"log"

	"warehouse-management/backend/internal/config"
	"warehouse-management/backend/internal/server"
)

func main() {
	cfg := config.Load()

	router := server.New()

	addr := ":" + cfg.Port
	if err := server.Run(addr, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
