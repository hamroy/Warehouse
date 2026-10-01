package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"warehouse-management/backend/internal/config"
	"warehouse-management/backend/internal/db"
	"warehouse-management/backend/internal/server"
)

func main() {
	cfg := config.Load()

	conn, err := db.Connect(cfg.DSN())
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	if err := db.Migrate(ctx, conn); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	router := server.New(conn)

	addr := ":" + cfg.Port

	// graceful shutdown
	idle := make(chan os.Signal, 1)
	signal.Notify(idle, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.Run(addr, router); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-idle
	log.Println("shutting down")
}
