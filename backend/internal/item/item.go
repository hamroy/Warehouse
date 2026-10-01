// Package item contains the item domain model and related types.
package item

import (
	"time"

	"github.com/google/uuid"
)

// Item is a catalog entry in the warehouse.
type Item struct {
	ID        uuid.UUID `json:"id"`
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Unit      string    `json:"unit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// IsActive reports whether the item is live (not soft-deleted).
func (i Item) IsActive() bool {
	return i.DeletedAt == nil
}

// CreateRequest is the payload for creating an item.
type CreateRequest struct {
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Unit     string `json:"unit"`
}

// UpdateRequest is the payload for updating an item.
type UpdateRequest struct {
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Unit     string `json:"unit"`
}

// ListQuery holds list/search/filter parameters.
type ListQuery struct {
	Search   string
	Category string
	Page     int
	Limit    int
}
