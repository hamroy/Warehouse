package stock_test

import (
	"context"
	"testing"

	"warehouse-management/backend/internal/item"
	"warehouse-management/backend/internal/location"
	"warehouse-management/backend/internal/stock"
)

type fakeItemRepo struct {
	notFound bool
}

func (f *fakeItemRepo) Create(ctx context.Context, it *item.Item) error { return nil }
func (f *fakeItemRepo) FindByID(ctx context.Context, id string) (*item.Item, error) {
	if f.notFound {
		return nil, item.ErrNotFound
	}
	return &item.Item{SKU: "SKU-1"}, nil
}
func (f *fakeItemRepo) FindAll(ctx context.Context, q item.ListQuery) ([]item.Item, int, error) {
	return nil, 0, nil
}
func (f *fakeItemRepo) Update(ctx context.Context, id string, req item.UpdateRequest) (*item.Item, error) {
	return nil, nil
}
func (f *fakeItemRepo) SoftDelete(ctx context.Context, id string) error { return nil }
func (f *fakeItemRepo) ExistsBySKU(ctx context.Context, sku string) (bool, error) {
	return false, nil
}

type fakeLocationRepo struct {
	notFound bool
}

func (f *fakeLocationRepo) FindByID(ctx context.Context, id string) (*location.Location, error) {
	return nil, nil
}
func (f *fakeLocationRepo) ExistsByID(ctx context.Context, id string) (bool, error) {
	return !f.notFound, nil
}

// newTestService builds a Service with fake item/location repos.
// The stock repository is nil: Receive returns before touching it
// when validation fails, which is the path covered here. Transactional
// happy-path behavior is covered by integration tests (task 1103).
func newTestService(itemNotFound, locNotFound bool) *stock.Service {
	return stock.NewService(
		nil,
		&fakeItemRepo{notFound: itemNotFound},
		&fakeLocationRepo{notFound: locNotFound},
	)
}

func TestReceive_Validation_QtyZero(t *testing.T) {
	svc := newTestService(false, false)

	_, err := svc.Receive(context.Background(), stock.ReceiveRequest{
		ItemID: "i", LocationID: "l", Qty: 0,
	})
	if !stock.IsValidation(err) {
		t.Fatalf("expected ValidationError for qty=0, got %v", err)
	}
}

func TestReceive_Validation_QtyNegative(t *testing.T) {
	svc := newTestService(false, false)

	_, err := svc.Receive(context.Background(), stock.ReceiveRequest{
		ItemID: "i", LocationID: "l", Qty: -5,
	})
	if !stock.IsValidation(err) {
		t.Fatalf("expected ValidationError for qty=-5, got %v", err)
	}
}

func TestReceive_ItemNotFound(t *testing.T) {
	svc := newTestService(true, false)

	_, err := svc.Receive(context.Background(), stock.ReceiveRequest{
		ItemID: "i", LocationID: "l", Qty: 10,
	})
	if !item.IsNotFound(err) {
		t.Fatalf("expected item not found, got %v", err)
	}
}

func TestReceive_LocationNotFound(t *testing.T) {
	svc := newTestService(false, true)

	_, err := svc.Receive(context.Background(), stock.ReceiveRequest{
		ItemID: "i", LocationID: "l", Qty: 10,
	})
	if !location.IsNotFound(err) {
		t.Fatalf("expected location not found, got %v", err)
	}
}
