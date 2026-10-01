package item

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// ErrNotFound signals the item does not exist (or is soft-deleted).
var ErrNotFound = errors.New("item not found")

// Repository handles all database access for items.
// It contains no business logic — only persistence.
type Repository struct {
	db *sql.DB
}

// NewRepository returns a Repository bound to the given connection.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// Create inserts a new item row.
func (r *Repository) Create(ctx context.Context, it *Item) error {
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO items (sku, name, category, unit)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, created_at, updated_at`,
		it.SKU, it.Name, it.Category, it.Unit,
	).Scan(&it.ID, &it.CreatedAt, &it.UpdatedAt)
	if err != nil {
		return fmt.Errorf("create item: %w", err)
	}
	return nil
}

// FindByID returns an active item by its UUID.
// Soft-deleted items are treated as not found.
func (r *Repository) FindByID(ctx context.Context, id string) (*Item, error) {
	var it Item
	var deletedAt *time.Time
	err := r.db.QueryRowContext(ctx,
		`SELECT id, sku, name, category, unit, created_at, updated_at, deleted_at
		 FROM items
		 WHERE id = $1 AND deleted_at IS NULL`,
		id,
	).Scan(&it.ID, &it.SKU, &it.Name, &it.Category, &it.Unit,
		&it.CreatedAt, &it.UpdatedAt, &deletedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("find item: %w", err)
	}
	it.DeletedAt = deletedAt
	return &it, nil
}

// FindAll returns a page of active items, optionally filtered by
// search (SKU or name ILIKE) and category.
func (r *Repository) FindAll(ctx context.Context, q ListQuery) ([]Item, int, error) {
	where := `deleted_at IS NULL`
	args := []any{}
	argIdx := 1

	if q.Search != "" {
		where += fmt.Sprintf(` AND (sku ILIKE $%d OR name ILIKE $%d)`, argIdx, argIdx+1)
		like := "%" + q.Search + "%"
		args = append(args, like, like)
		argIdx += 2
	}
	if q.Category != "" {
		where += fmt.Sprintf(` AND category = $%d`, argIdx)
		args = append(args, q.Category)
		argIdx++
	}

	var total int
	countQuery := `SELECT count(*) FROM items WHERE ` + where
	if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count items: %w", err)
	}

	offset := (q.Page - 1) * q.Limit
	args = append(args, q.Limit, offset)

	rows, err := r.db.QueryContext(ctx,
		`SELECT id, sku, name, category, unit, created_at, updated_at, deleted_at
		 FROM items
		 WHERE `+where+`
		 ORDER BY created_at DESC
		 LIMIT $`+fmt.Sprint(argIdx)+` OFFSET $`+fmt.Sprint(argIdx+1),
		args...,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("list items: %w", err)
	}
	defer rows.Close()

	items := []Item{}
	for rows.Next() {
		var it Item
		var deletedAt *time.Time
		if err := rows.Scan(&it.ID, &it.SKU, &it.Name, &it.Category, &it.Unit,
			&it.CreatedAt, &it.UpdatedAt, &deletedAt); err != nil {
			return nil, 0, fmt.Errorf("scan item: %w", err)
		}
		it.DeletedAt = deletedAt
		items = append(items, it)
	}
	return items, total, rows.Err()
}

// Update modifies an active item's fields.
func (r *Repository) Update(ctx context.Context, id string, req UpdateRequest) (*Item, error) {
	var it Item
	var deletedAt *time.Time
	err := r.db.QueryRowContext(ctx,
		`UPDATE items
		 SET sku = $1, name = $2, category = $3, unit = $4
		 WHERE id = $5 AND deleted_at IS NULL
		 RETURNING id, sku, name, category, unit, created_at, updated_at, deleted_at`,
		req.SKU, req.Name, req.Category, req.Unit, id,
	).Scan(&it.ID, &it.SKU, &it.Name, &it.Category, &it.Unit,
		&it.CreatedAt, &it.UpdatedAt, &deletedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("update item: %w", err)
	}
	it.DeletedAt = deletedAt
	return &it, nil
}

// SoftDelete marks an item as deleted by setting deleted_at.
func (r *Repository) SoftDelete(ctx context.Context, id string) error {
	res, err := r.db.ExecContext(ctx,
		`UPDATE items SET deleted_at = now()
		 WHERE id = $1 AND deleted_at IS NULL`,
		id,
	)
	if err != nil {
		return fmt.Errorf("soft delete item: %w", err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("soft delete affected rows: %w", err)
	}
	if affected == 0 {
		return ErrNotFound
	}
	return nil
}

// ExistsBySKU reports whether an active item already uses the SKU.
func (r *Repository) ExistsBySKU(ctx context.Context, sku string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx,
		`SELECT EXISTS (SELECT 1 FROM items WHERE sku = $1 AND deleted_at IS NULL)`,
		sku,
	).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check sku exists: %w", err)
	}
	return exists, nil
}
