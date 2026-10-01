package item

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

// ValidationError signals invalid input (HTTP 400).
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

// ConflictError signals a duplicate SKU (HTTP 409).
type ConflictError struct {
	SKU string
}

func (e *ConflictError) Error() string {
	return fmt.Sprintf("sku %q already exists", e.SKU)
}

// Repo is the persistence surface the Service needs. It is satisfied by
// *Repository and by fakes in unit tests, so the service can be tested
// without a database.
type Repo interface {
	Create(ctx context.Context, it *Item) error
	FindByID(ctx context.Context, id string) (*Item, error)
	FindAll(ctx context.Context, q ListQuery) ([]Item, int, error)
	Update(ctx context.Context, id string, req UpdateRequest) (*Item, error)
	SoftDelete(ctx context.Context, id string) error
	ExistsBySKU(ctx context.Context, sku string) (bool, error)
}

// Service holds item business logic.
type Service struct {
	repo Repo
}

// NewService returns a Service bound to the given repository.
func NewService(repo Repo) *Service {
	return &Service{repo: repo}
}

// validateFields checks the required item fields shared by
// create and update requests.
func validateFields(sku, name, category, unit string) error {
	if strings.TrimSpace(sku) == "" {
		return &ValidationError{Field: "sku", Msg: "required"}
	}
	if strings.TrimSpace(name) == "" {
		return &ValidationError{Field: "name", Msg: "required"}
	}
	if strings.TrimSpace(category) == "" {
		return &ValidationError{Field: "category", Msg: "required"}
	}
	if strings.TrimSpace(unit) == "" {
		return &ValidationError{Field: "unit", Msg: "required"}
	}
	return nil
}

// validateCreate checks required fields.
func validateCreate(req CreateRequest) error {
	return validateFields(req.SKU, req.Name, req.Category, req.Unit)
}

// Create validates input, rejects duplicate SKU, then persists the item.
func (s *Service) Create(ctx context.Context, req CreateRequest) (*Item, error) {
	if err := validateCreate(req); err != nil {
		return nil, err
	}

	exists, err := s.repo.ExistsBySKU(ctx, req.SKU)
	if err != nil {
		return nil, fmt.Errorf("check duplicate sku: %w", err)
	}
	if exists {
		return nil, &ConflictError{SKU: req.SKU}
	}

	it := &Item{
		SKU:      req.SKU,
		Name:     req.Name,
		Category: req.Category,
		Unit:     req.Unit,
	}
	if err := s.repo.Create(ctx, it); err != nil {
		return nil, fmt.Errorf("create item: %w", err)
	}
	return it, nil
}

// GetByID returns an active item or ErrNotFound.
func (s *Service) GetByID(ctx context.Context, id string) (*Item, error) {
	return s.repo.FindByID(ctx, id)
}

// List returns a page of items plus the total count.
func (s *Service) List(ctx context.Context, q ListQuery) ([]Item, int, error) {
	return s.repo.FindAll(ctx, q)
}

// Update validates then persists changes to an active item.
func (s *Service) Update(ctx context.Context, id string, req UpdateRequest) (*Item, error) {
	if err := validateFields(req.SKU, req.Name, req.Category, req.Unit); err != nil {
		return nil, err
	}
	return s.repo.Update(ctx, id, req)
}

// Delete soft-deletes an item.
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.SoftDelete(ctx, id)
}

// IsNotFound / IsConflict / IsValidation classify domain errors
// so the HTTP layer maps them to the right status code.
func IsNotFound(err error) bool { return errors.Is(err, ErrNotFound) }

func IsConflict(err error) bool {
	var target *ConflictError
	return errors.As(err, &target)
}

func IsValidation(err error) bool {
	var target *ValidationError
	return errors.As(err, &target)
}
