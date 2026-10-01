package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"warehouse-management/backend/internal/item"

	"github.com/go-chi/chi/v5"
)

// ItemHandler handles item HTTP endpoints.
type ItemHandler struct {
	svc *item.Service
}

// NewItemHandler returns an ItemHandler bound to the given service.
func NewItemHandler(svc *item.Service) *ItemHandler {
	return &ItemHandler{svc: svc}
}

// Create handles POST /api/v1/items.
func (h *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req item.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, http.StatusBadRequest, "Validation failed", "invalid JSON body")
		return
	}

	created, err := h.svc.Create(r.Context(), req)
	if err != nil {
		HandleError(w, err)
		return
	}
	WriteJSON(w, http.StatusCreated, "Item created successfully", created)
}

// List handles GET /api/v1/items with search, category filter, pagination.
func (h *ItemHandler) List(w http.ResponseWriter, r *http.Request) {
	q := item.ListQuery{
		Search:   r.URL.Query().Get("search"),
		Category: r.URL.Query().Get("category"),
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}
	q.Page, q.Limit = page, limit

	items, total, err := h.svc.List(r.Context(), q)
	if err != nil {
		HandleError(w, err)
		return
	}
	WritePaginatedJSON(w, http.StatusOK, "Items retrieved successfully", items, Meta{
		Page: page, Limit: limit, Total: total,
	})
}

// GetByID handles GET /api/v1/items/:id.
func (h *ItemHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	it, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		HandleError(w, err)
		return
	}
	WriteJSON(w, http.StatusOK, "Item retrieved successfully", it)
}

// Update handles PUT /api/v1/items/:id.
func (h *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req item.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, http.StatusBadRequest, "Validation failed", "invalid JSON body")
		return
	}

	updated, err := h.svc.Update(r.Context(), id, req)
	if err != nil {
		HandleError(w, err)
		return
	}
	WriteJSON(w, http.StatusOK, "Item updated successfully", updated)
}

// Delete handles DELETE /api/v1/items/:id (soft delete).
func (h *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.svc.Delete(r.Context(), id); err != nil {
		HandleError(w, err)
		return
	}
	WriteJSON(w, http.StatusOK, "Item deleted successfully", nil)
}
