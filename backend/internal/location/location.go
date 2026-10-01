// Package location contains the location domain model and repository.
package location

import (
	"time"

	"github.com/google/uuid"
)

// Location is a warehouse storage location.
type Location struct {
	ID        uuid.UUID  `json:"id"`
	Code      string     `json:"code"`
	Zone      string     `json:"zone"`
	Type      string     `json:"type"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
