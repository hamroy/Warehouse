// Package stock contains the stock domain model, repository,
// and receive-stock service with transaction support.
package stock

import (
	"time"

	"github.com/google/uuid"
)

// Stock is the quantity of an item at a location.
type Stock struct {
	ID         uuid.UUID `json:"id"`
	ItemID     uuid.UUID `json:"item_id"`
	LocationID uuid.UUID `json:"location_id"`
	Qty        int       `json:"qty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// StockDetail joins stock with item and location display fields.
type StockDetail struct {
	ItemID       uuid.UUID `json:"item_id"`
	ItemSKU      string    `json:"item_sku"`
	ItemName     string    `json:"item_name"`
	LocationID   uuid.UUID `json:"location_id"`
	LocationCode string    `json:"location_code"`
	Qty          int       `json:"qty"`
}

// ReceiveRequest is the payload for receiving stock.
type ReceiveRequest struct {
	ItemID     string `json:"item_id"`
	LocationID string `json:"location_id"`
	Qty        int    `json:"qty"`
}

// LogType enumerates stock transaction types.
type LogType string

const (
	LogTypeReceive LogType = "RECEIVE_STOCK"
)

// TransactionLog records every stock movement.
type TransactionLog struct {
	ID         uuid.UUID `json:"id"`
	Type       LogType   `json:"type"`
	ItemID     uuid.UUID `json:"item_id"`
	LocationID uuid.UUID `json:"location_id"`
	Qty        int       `json:"qty"`
	CreatedAt  time.Time `json:"created_at"`
}
