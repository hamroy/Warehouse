package stock

import (
	"context"

	"errors"
	"fmt"

	"warehouse-management/backend/internal/item"
	"warehouse-management/backend/internal/location"
)

// ValidationError signals invalid receive input (HTTP 400).
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

// itemRepo is the subset of the item repository the stock service needs.
// *item.Repository and unit-test fakes both satisfy it.
type itemRepo interface {
	FindByID(ctx context.Context, id string) (*item.Item, error)
}

// locationRepo is the subset of the location repository the stock service needs.
type locationRepo interface {
	ExistsByID(ctx context.Context, id string) (bool, error)
}

// Service holds stock business logic.
type Service struct {
	repo      *Repository
	items     itemRepo
	locations locationRepo
}

// NewService returns a Service bound to the given repositories.
func NewService(repo *Repository, items itemRepo, locations locationRepo) *Service {
	return &Service{repo: repo, items: items, locations: locations}
}

// Receive validates item, location and quantity, then applies the
// receipt inside a single database transaction:
//
//	BEGIN
//	  find existing stock
//	  create or increase
//	  write log
//	COMMIT
//
// On any error the transaction is rolled back.
func (s *Service) Receive(ctx context.Context, req ReceiveRequest) (*Stock, error) {
	// validate quantity
	if req.Qty <= 0 {
		return nil, &ValidationError{Field: "qty", Msg: "must be greater than 0"}
	}

	// validate item exists and is active
	if _, err := s.items.FindByID(ctx, req.ItemID); err != nil {
		if item.IsNotFound(err) {
			return nil, item.ErrNotFound
		}
		return nil, fmt.Errorf("validate item: %w", err)
	}

	// validate location exists
	exists, err := s.locations.ExistsByID(ctx, req.LocationID)
	if err != nil {
		return nil, fmt.Errorf("validate location: %w", err)
	}
	if !exists {
		return nil, location.ErrNotFound
	}

	tx, err := s.repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// find existing stock within the transaction
	existing, err := s.repo.FindByItemAndLocation(ctx, req.ItemID, req.LocationID)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return nil, fmt.Errorf("find existing stock: %w", err)
	}

	var result *Stock
	if errors.Is(err, ErrNotFound) {
		// stock does not exist yet -> create
		result, err = s.repo.Create(ctx, req.ItemID, req.LocationID, req.Qty)
		if err != nil {
			return nil, fmt.Errorf("create stock: %w", err)
		}
	} else {
		// existing stock -> increase
		result, err = s.repo.IncreaseQuantity(ctx, existing.ID.String(), req.Qty)
		if err != nil {
			return nil, fmt.Errorf("increase stock: %w", err)
		}
	}

	// write transaction log
	if err = s.repo.InsertLog(ctx, tx, LogTypeReceive, result.ItemID.String(), result.LocationID.String(), req.Qty); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit transaction: %w", err)
	}

	return result, nil
}

// GetByItem returns all stock rows for an item plus the total quantity.
func (s *Service) GetByItem(ctx context.Context, itemID string) ([]StockDetail, int, error) {
	if _, err := s.items.FindByID(ctx, itemID); err != nil {
		if item.IsNotFound(err) {
			return nil, 0, item.ErrNotFound
		}
		return nil, 0, fmt.Errorf("validate item: %w", err)
	}
	return s.repo.FindByItem(ctx, itemID)
}

// IsValidation reports whether err is a ValidationError.
func IsValidation(err error) bool {
	var target *ValidationError
	return errors.As(err, &target)
}
