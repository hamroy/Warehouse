// Command seed inserts development data into the database.
//
// It is idempotent: re-running it does not create duplicates.
//
// Usage:
//
//	go run ./cmd/seed
package main

import (
	"context"
	"log"

	"warehouse-management/backend/internal/config"
	"warehouse-management/backend/internal/db"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()

	conn, err := db.Connect(cfg.DSN())
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer conn.Close()

	if err := db.Seed(ctx, conn); err != nil {
		log.Fatalf("seed: %v", err)
	}
	log.Println("seed completed")
}
