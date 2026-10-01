package db

import (
	"context"
	"database/sql"
	"fmt"
)

// seedLocation is a fixed master-data row for development.
type seedLocation struct {
	code string
	zone string
	typ  string
}

// seedItem is a fixed catalog row for development.
type seedItem struct {
	sku      string
	name     string
	category string
	unit     string
}

// seedLocations and seedItems match the PRD seed specification.
var (
	seedLocations = []seedLocation{
		{code: "A-01-01", zone: "A", typ: "rack"},
		{code: "A-01-02", zone: "A", typ: "rack"},
		{code: "B-01-01", zone: "B", typ: "rack"},
	}

	seedItems = []seedItem{
		{sku: "SKU-001", name: "Laptop Dell", category: "Electronics", unit: "PCS"},
		{sku: "SKU-002", name: "Mouse", category: "Electronics", unit: "PCS"},
		{sku: "SKU-003", name: "Office Chair", category: "Furniture", unit: "PCS"},
	}
)

// Seed inserts development data. The strategy is ON CONFLICT DO NOTHING so
// running Seed repeatedly never creates duplicate rows (idempotent).
func Seed(ctx context.Context, db *sql.DB) error {
	for _, loc := range seedLocations {
		if _, err := db.ExecContext(ctx,
			`INSERT INTO locations (code, zone, type)
			 VALUES ($1, $2, $3)
			 ON CONFLICT (code) DO NOTHING`,
			loc.code, loc.zone, loc.typ); err != nil {
			return fmt.Errorf("seed location %s: %w", loc.code, err)
		}
	}

	for _, item := range seedItems {
		if _, err := db.ExecContext(ctx,
			`INSERT INTO items (sku, name, category, unit)
			 VALUES ($1, $2, $3, $4)
			 ON CONFLICT (sku) DO NOTHING`,
			item.sku, item.name, item.category, item.unit); err != nil {
			return fmt.Errorf("seed item %s: %w", item.sku, err)
		}
	}
	return nil
}
