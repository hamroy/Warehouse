package stock

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// ErrNotFound signals the stock record does not exist.
var ErrNotFound = errors.New("stock not found")

// Repository handles database access for stocks and logs.
type Repository struct {
	db *sql.DB
}

// NewRepository returns a Repository bound to the given connection.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// FindByItemAndLocation returns the stock row for a pair, or
// ErrNotFound when none exists yet.
func (r *Repository) FindByItemAndLocation(ctx context.Context, itemID, locationID string) (*Stock, error) {
	var s Stock
	err := r.db.QueryRowContext(ctx,
		`SELECT id, item_id, location_id, qty, created_at, updated_at
		 FROM stocks
		 WHERE item_id = $1 AND location_id = $2`,
		itemID, locationID,
	).Scan(&s.ID, &s.ItemID, &s.LocationID, &s.Qty, &s.CreatedAt, &s.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("find stock: %w", err)
	}
	return &s, nil
}

// Create inserts a new stock row.
func (r *Repository) Create(ctx context.Context, itemID, locationID string, qty int) (*Stock, error) {
	var s Stock
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO stocks (item_id, location_id, qty)
		 VALUES ($1, $2, $3)
		 RETURNING id, item_id, location_id, qty, created_at, updated_at`,
		itemID, locationID, qty,
	).Scan(&s.ID, &s.ItemID, &s.LocationID, &s.Qty, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("create stock: %w", err)
	}
	return &s, nil
}

// IncreaseQuantity adds qty to an existing stock row.
func (r *Repository) IncreaseQuantity(ctx context.Context, id string, delta int) (*Stock, error) {
	var s Stock
	err := r.db.QueryRowContext(ctx,
		`UPDATE stocks SET qty = qty + $1
		 WHERE id = $2
		 RETURNING id, item_id, location_id, qty, created_at, updated_at`,
		delta, id,
	).Scan(&s.ID, &s.ItemID, &s.LocationID, &s.Qty, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("increase stock: %w", err)
	}
	return &s, nil
}

// FindByItem returns all stock rows for an item with location details.
func (r *Repository) FindByItem(ctx context.Context, itemID string) ([]StockDetail, int, error) {
	var total int
	err := r.db.QueryRowContext(ctx,
		`SELECT COALESCE(SUM(qty), 0) FROM stocks WHERE item_id = $1`,
		itemID,
	).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("sum stock: %w", err)
	}

	rows, err := r.db.QueryContext(ctx,
		`SELECT s.item_id, i.sku, i.name, s.location_id, l.code, s.qty
		 FROM stocks s
		 JOIN items i ON i.id = s.item_id
		 JOIN locations l ON l.id = s.location_id
		 WHERE s.item_id = $1
		 ORDER BY l.code`,
		itemID,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("list stock: %w", err)
	}
	defer rows.Close()

	details := []StockDetail{}
	for rows.Next() {
		var d StockDetail
		if err := rows.Scan(&d.ItemID, &d.ItemSKU, &d.ItemName, &d.LocationID, &d.LocationCode, &d.Qty); err != nil {
			return nil, 0, fmt.Errorf("scan stock: %w", err)
		}
		details = append(details, d)
	}
	return details, total, rows.Err()
}

// InsertLog writes a stock transaction log inside the caller's transaction.
func (r *Repository) InsertLog(ctx context.Context, tx *sql.Tx, typ LogType, itemID, locationID string, qty int) error {
	_, err := tx.ExecContext(ctx,
		`INSERT INTO stock_logs (type, item_id, location_id, qty)
		 VALUES ($1, $2, $3, $4)`,
		typ, itemID, locationID, qty,
	)
	if err != nil {
		return fmt.Errorf("insert stock log: %w", err)
	}
	return nil
}

