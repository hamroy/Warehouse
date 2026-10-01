package handler

import (
	"encoding/json"
	"net/http"
)

// HealthResponse is the payload returned by GET /health.
type HealthResponse struct {
	Status string `json:"status"`
}

// Health handles GET /health and reports service liveness.
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(HealthResponse{Status: "ok"})
}
