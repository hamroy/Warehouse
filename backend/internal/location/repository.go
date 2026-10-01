package location

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// ErrNotFound signals the location does not exist.
var ErrNotFound = errors.New("location not found")

// Repository handles database access for locations.
type Repository struct {
	db *sql.DB
}

// NewRepository returns a Repository bound to the given connection.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// FindByID returns a location by its UUID.
func (r *Repository) FindByID(ctx context.Context, id string) (*Location, error) {
	var loc Location
	err := r.db.QueryRowContext(ctx,
		`SELECT id, code, zone, type, created_at, updated_at
		 FROM locations
		 WHERE id = $1`,
		id,
	).Scan(&loc.ID, &loc.Code, &loc.Zone, &loc.Type, &loc.CreatedAt, &loc.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("find location: %w", err)
	}
	return &loc, nil
}

// ExistsByID reports whether a location with the given UUID exists.
func (r *Repository) ExistsByID(ctx context.Context, id string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx,
		`SELECT EXISTS (SELECT 1 FROM locations WHERE id = $1)`,
		id,
	).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check location exists: %w", err)
	}
	return exists, nil
}
