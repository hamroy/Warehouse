package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"warehouse-management/backend/internal/item"
	"warehouse-management/backend/internal/stock"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// fakeItemRepo is an in-memory item.Repo for integration tests. It exercises
// the real HTTP handlers + service validation without touching a database.
type fakeItemRepo struct {
	bySKU map[string]*item.Item
	byID  map[string]*item.Item
	order []*item.Item
}

func newFakeItemRepo() *fakeItemRepo {
	return &fakeItemRepo{
		bySKU: map[string]*item.Item{},
		byID:  map[string]*item.Item{},
	}
}

func (f *fakeItemRepo) Create(ctx context.Context, it *item.Item) error {
	it.ID = uuid.New()
	clone := *it
	f.bySKU[it.SKU] = &clone
	f.byID[it.ID.String()] = &clone
	f.order = append(f.order, &clone)
	return nil
}

func (f *fakeItemRepo) FindByID(ctx context.Context, id string) (*item.Item, error) {
	it, ok := f.byID[id]
	if !ok || !it.IsActive() {
		return nil, item.ErrNotFound
	}
	return it, nil
}

func (f *fakeItemRepo) FindAll(ctx context.Context, q item.ListQuery) ([]item.Item, int, error) {
	out := []item.Item{}
	for _, it := range f.order {
		if it.IsActive() {
			out = append(out, *it)
		}
	}
	return out, len(out), nil
}

func (f *fakeItemRepo) Update(ctx context.Context, id string, req item.UpdateRequest) (*item.Item, error) {
	it, ok := f.byID[id]
	if !ok || !it.IsActive() {
		return nil, item.ErrNotFound
	}
	it.SKU = req.SKU
	it.Name = req.Name
	it.Category = req.Category
	it.Unit = req.Unit
	return it, nil
}

func (f *fakeItemRepo) SoftDelete(ctx context.Context, id string) error {
	it, ok := f.byID[id]
	if !ok || !it.IsActive() {
		return item.ErrNotFound
	}
	now := time.Now()
	it.DeletedAt = &now
	return nil
}

func (f *fakeItemRepo) ExistsBySKU(ctx context.Context, sku string) (bool, error) {
	it, ok := f.bySKU[sku]
	return ok && it.IsActive(), nil
}

// stockItemFinder satisfies the item lookup surface the stock service needs.
type stockItemFinder struct {
	notFound bool
}

func (f *stockItemFinder) FindByID(ctx context.Context, id string) (*item.Item, error) {
	if f.notFound {
		return nil, item.ErrNotFound
	}
	return &item.Item{SKU: "SKU-001"}, nil
}

// stockLocationFinder satisfies the location lookup surface the stock service needs.
type stockLocationFinder struct {
	notFound bool
}

func (f *stockLocationFinder) ExistsByID(ctx context.Context, id string) (bool, error) {
	return !f.notFound, nil
}

// newTestRouter mounts the real handlers over chi with fake repositories.
func newTestRouter(ir item.Repo, sf itemFinder, lf locationFinder) *chi.Mux {
	r := chi.NewRouter()

	ih := NewItemHandler(item.NewService(ir))
	sh := NewStockHandler(stock.NewService(nil, sf, lf))

	r.Post("/api/v1/items", ih.Create)
	r.Get("/api/v1/items", ih.List)
	r.Get("/api/v1/items/{id}", ih.GetByID)
	r.Put("/api/v1/items/{id}", ih.Update)
	r.Delete("/api/v1/items/{id}", ih.Delete)

	r.Post("/api/v1/stock/receive", sh.Receive)
	r.Get("/api/v1/stock/{item_id}", sh.GetByItem)

	return r
}

// itemFinder / locationFinder mirror the subset the stock service accepts.
type itemFinder interface {
	FindByID(ctx context.Context, id string) (*item.Item, error)
}

type locationFinder interface {
	ExistsByID(ctx context.Context, id string) (bool, error)
}

func doRequest(t *testing.T, r *chi.Mux, method, path, body string) *httptest.ResponseRecorder {
	t.Helper()
	var req *http.Request
	if body == "" {
		req = httptest.NewRequest(method, path, nil)
	} else {
		req = httptest.NewRequest(method, path, strings.NewReader(body))
	}
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	return rec
}

func decodeResponse(t *testing.T, rec *httptest.ResponseRecorder) Response {
	t.Helper()
	var resp Response
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	return resp
}

func TestItemAPI_Create_Success(t *testing.T) {
	r := newTestRouter(newFakeItemRepo(), &stockItemFinder{}, &stockLocationFinder{})

	rec := doRequest(t, r, http.MethodPost, "/api/v1/items",
		`{"sku":"SKU-001","name":"Laptop","category":"Electronics","unit":"PCS"}`)
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", rec.Code, rec.Body.String())
	}
	if resp := decodeResponse(t, rec); !resp.Success {
		t.Fatalf("expected success envelope, got %+v", resp)
	}
}

func TestItemAPI_Create_Invalid(t *testing.T) {
	r := newTestRouter(newFakeItemRepo(), &stockItemFinder{}, &stockLocationFinder{})

	rec := doRequest(t, r, http.MethodPost, "/api/v1/items",
		`{"sku":"","name":"Laptop","category":"Electronics","unit":"PCS"}`)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestItemAPI_Create_Duplicate(t *testing.T) {
	r := newTestRouter(newFakeItemRepo(), &stockItemFinder{}, &stockLocationFinder{})

	body := `{"sku":"SKU-001","name":"Laptop","category":"Electronics","unit":"PCS"}`
	first := doRequest(t, r, http.MethodPost, "/api/v1/items", body)
	if first.Code != http.StatusCreated {
		t.Fatalf("expected first create 201, got %d", first.Code)
	}
	second := doRequest(t, r, http.MethodPost, "/api/v1/items", body)
	if second.Code != http.StatusConflict {
		t.Fatalf("expected 409, got %d: %s", second.Code, second.Body.String())
	}
}

func TestItemAPI_GetByID_Success(t *testing.T) {
	r := newTestRouter(newFakeItemRepo(), &stockItemFinder{}, &stockLocationFinder{})

	created := doRequest(t, r, http.MethodPost, "/api/v1/items",
		`{"sku":"SKU-001","name":"Laptop","category":"Electronics","unit":"PCS"}`)
	if created.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", created.Code)
	}
	resp := decodeResponse(t, created)
	it := resp.Data.(map[string]any)
	id := it["id"].(string)

	rec := doRequest(t, r, http.MethodGet, "/api/v1/items/"+id, "")
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestItemAPI_GetByID_NotFound(t *testing.T) {
	r := newTestRouter(newFakeItemRepo(), &stockItemFinder{}, &stockLocationFinder{})

	rec := doRequest(t, r, http.MethodGet, "/api/v1/items/"+uuid.NewString(), "")
	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestItemAPI_List_Success(t *testing.T) {
	r := newTestRouter(newFakeItemRepo(), &stockItemFinder{}, &stockLocationFinder{})

	rec := doRequest(t, r, http.MethodGet, "/api/v1/items?page=1&limit=10", "")
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	resp := decodeResponse(t, rec)
	if resp.Meta == nil {
		t.Fatalf("expected pagination meta, got %+v", resp)
	}
}

func TestItemAPI_Update_Success(t *testing.T) {
	r := newTestRouter(newFakeItemRepo(), &stockItemFinder{}, &stockLocationFinder{})

	created := doRequest(t, r, http.MethodPost, "/api/v1/items",
		`{"sku":"SKU-001","name":"Laptop","category":"Electronics","unit":"PCS"}`)
	id := decodeResponse(t, created).Data.(map[string]any)["id"].(string)

	rec := doRequest(t, r, http.MethodPut, "/api/v1/items/"+id,
		`{"sku":"SKU-001","name":"Laptop Dell","category":"Electronics","unit":"PCS"}`)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestItemAPI_Update_NotFound(t *testing.T) {
	r := newTestRouter(newFakeItemRepo(), &stockItemFinder{}, &stockLocationFinder{})

	rec := doRequest(t, r, http.MethodPut, "/api/v1/items/"+uuid.NewString(),
		`{"sku":"SKU-001","name":"Laptop","category":"Electronics","unit":"PCS"}`)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestItemAPI_Delete_Success(t *testing.T) {
	r := newTestRouter(newFakeItemRepo(), &stockItemFinder{}, &stockLocationFinder{})

	created := doRequest(t, r, http.MethodPost, "/api/v1/items",
		`{"sku":"SKU-001","name":"Laptop","category":"Electronics","unit":"PCS"}`)
	id := decodeResponse(t, created).Data.(map[string]any)["id"].(string)

	rec := doRequest(t, r, http.MethodDelete, "/api/v1/items/"+id, "")
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	// soft-deleted item must no longer resolve
	again := doRequest(t, r, http.MethodGet, "/api/v1/items/"+id, "")
	if again.Code != http.StatusNotFound {
		t.Fatalf("expected 404 after delete, got %d", again.Code)
	}
}

func TestStockAPI_Receive_Validation(t *testing.T) {
	r := newTestRouter(newFakeItemRepo(), &stockItemFinder{}, &stockLocationFinder{})

	rec := doRequest(t, r, http.MethodPost, "/api/v1/stock/receive",
		`{"item_id":"`+uuid.NewString()+`","location_id":"`+uuid.NewString()+`","qty":0}`)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestStockAPI_Receive_ItemNotFound(t *testing.T) {
	r := newTestRouter(newFakeItemRepo(), &stockItemFinder{notFound: true}, &stockLocationFinder{})

	rec := doRequest(t, r, http.MethodPost, "/api/v1/stock/receive",
		`{"item_id":"`+uuid.NewString()+`","location_id":"`+uuid.NewString()+`","qty":10}`)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestStockAPI_Receive_LocationNotFound(t *testing.T) {
	r := newTestRouter(newFakeItemRepo(), &stockItemFinder{}, &stockLocationFinder{notFound: true})

	rec := doRequest(t, r, http.MethodPost, "/api/v1/stock/receive",
		`{"item_id":"`+uuid.NewString()+`","location_id":"`+uuid.NewString()+`","qty":10}`)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestStockAPI_GetByItem_NotFound(t *testing.T) {
	r := newTestRouter(newFakeItemRepo(), &stockItemFinder{notFound: true}, &stockLocationFinder{})

	rec := doRequest(t, r, http.MethodGet, "/api/v1/stock/"+uuid.NewString(), "")
	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d: %s", rec.Code, rec.Body.String())
	}
}
