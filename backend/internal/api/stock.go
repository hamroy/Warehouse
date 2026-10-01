package api

import (
	"encoding/json"
	"net/http"

	"warehouse-management/backend/internal/stock"

	"github.com/go-chi/chi/v5"
)

// StockHandler handles stock HTTP endpoints.
type StockHandler struct {
	svc *stock.Service
}

// NewStockHandler returns a StockHandler bound to the given service.
func NewStockHandler(svc *stock.Service) *StockHandler {
	return &StockHandler{svc: svc}
}

// Receive handles POST /api/v1/stock/receive.
func (h *StockHandler) Receive(w http.ResponseWriter, r *http.Request) {
	var req stock.ReceiveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, http.StatusBadRequest, "Validation failed", "invalid JSON body")
		return
	}

	result, err := h.svc.Receive(r.Context(), req)
	if err != nil {
		if stock.IsValidation(err) {
			WriteError(w, http.StatusBadRequest, "Validation failed", err.Error())
			return
		}
		HandleError(w, err)
		return
	}
	WriteJSON(w, http.StatusOK, "Stock received successfully", result)
}

// GetByItem handles GET /api/v1/stock/:item_id.
func (h *StockHandler) GetByItem(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "item_id")

	details, total, err := h.svc.GetByItem(r.Context(), itemID)
	if err != nil {
		HandleError(w, err)
		return
	}

	payload := struct {
		Total int            `json:"total"`
		Items []stock.StockDetail `json:"items"`
	}{Total: total, Items: details}
	WriteJSON(w, http.StatusOK, "Stock retrieved successfully", payload)
}
