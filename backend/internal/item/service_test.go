package item_test

import (
	"context"
	"errors"
	"testing"

	"warehouse-management/backend/internal/item"

	"github.com/google/uuid"
)

// stubRepo implements the repository surface used by Service.
// It is a fake, not the real repository, so service tests
// never touch a database.
type stubRepo struct {
	exists   bool
	notFound bool
	created  *item.Item
}

func (s *stubRepo) Create(ctx context.Context, it *item.Item) error {
	it.ID = uuid.New()
	s.created = it
	return nil
}

func (s *stubRepo) FindByID(ctx context.Context, id string) (*item.Item, error) {
	if s.notFound {
		return nil, item.ErrNotFound
	}
	return &item.Item{ID: uuid.New(), SKU: "SKU-1"}, nil
}

func (s *stubRepo) FindAll(ctx context.Context, q item.ListQuery) ([]item.Item, int, error) {
	return []item.Item{}, 0, nil
}

func (s *stubRepo) Update(ctx context.Context, id string, req item.UpdateRequest) (*item.Item, error) {
	return &item.Item{ID: uuid.New()}, nil
}

func (s *stubRepo) SoftDelete(ctx context.Context, id string) error {
	return nil
}

func (s *stubRepo) ExistsBySKU(ctx context.Context, sku string) (bool, error) {
	return s.exists, nil
}

func newService(exists bool) *item.Service {
	return item.NewService(&stubRepo{exists: exists})
}

func TestCreateItem_Valid(t *testing.T) {
	svc := newService(false)

	created, err := svc.Create(context.Background(), item.CreateRequest{
		SKU: "SKU-1", Name: "Laptop", Category: "Electronics", Unit: "PCS",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if created == nil || created.SKU != "SKU-1" {
		t.Fatalf("expected created item, got %+v", created)
	}
}

func TestCreateItem_DuplicateSKU(t *testing.T) {
	svc := newService(true)

	_, err := svc.Create(context.Background(), item.CreateRequest{
		SKU: "SKU-1", Name: "Laptop", Category: "Electronics", Unit: "PCS",
	})

	if !item.IsConflict(err) {
		t.Fatalf("expected ConflictError, got %v", err)
	}
}

func TestCreateItem_InvalidInput(t *testing.T) {
	svc := newService(false)

	cases := []item.CreateRequest{
		{SKU: "", Name: "x", Category: "c", Unit: "u"},
		{SKU: "s", Name: "", Category: "c", Unit: "u"},
		{SKU: "s", Name: "n", Category: "", Unit: "u"},
		{SKU: "s", Name: "n", Category: "c", Unit: ""},
	}

	for i, req := range cases {
		_, err := svc.Create(context.Background(), req)
		if !item.IsValidation(err) {
			t.Fatalf("case %d: expected ValidationError, got %v", i, err)
		}
	}
}

func TestGetByID_NotFound(t *testing.T) {
	svc2 := item.NewService(&stubRepo{notFound: true})

	_, err := svc2.GetByID(context.Background(), "x")
	if !item.IsNotFound(err) {
		t.Fatalf("expected NotFound, got %v", err)
	}
	if !errors.Is(err, item.ErrNotFound) {
		t.Fatal("expected errors.Is(err, ErrNotFound)")
	}
}
