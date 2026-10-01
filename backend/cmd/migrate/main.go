// Command migrate applies or rolls back database migrations.
//
// Usage:
//
//	go run ./cmd/migrate up
//	go run ./cmd/migrate down
package main

import (
	"context"
	"log"
	"os"

	"warehouse-management/backend/internal/config"
	"warehouse-management/backend/internal/db"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: migrate <up|down>")
	}

	cfg := config.Load()
	ctx := context.Background()

	conn, err := db.Connect(cfg.DSN())
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer conn.Close()

	switch os.Args[1] {
	case "up":
		if err := db.Migrate(ctx, conn); err != nil {
			log.Fatalf("migrate: %v", err)
		}
	case "down":
		if err := db.Rollback(ctx, conn); err != nil {
			log.Fatalf("rollback: %v", err)
		}
	default:
		log.Fatalf("unknown command %q (use up|down)", os.Args[1])
	}
}
