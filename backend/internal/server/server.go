package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"warehouse-management/backend/internal/api"
	"warehouse-management/backend/internal/handler"
	"warehouse-management/backend/internal/item"
	"warehouse-management/backend/internal/location"
	"warehouse-management/backend/internal/stock"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// New builds the HTTP router with all registered routes.
// The db handle is optional; when nil only the health endpoint is mounted.
func New(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/healthz"))

	r.Get("/health", handler.Health)

	if db == nil {
		return r
	}

	itemRepo := item.NewRepository(db)
	itemSvc := item.NewService(itemRepo)
	itemHandler := api.NewItemHandler(itemSvc)

	locRepo := location.NewRepository(db)
	stockRepo := stock.NewRepository(db)
	stockSvc := stock.NewService(stockRepo, itemRepo, locRepo)
	stockHandler := api.NewStockHandler(stockSvc)

	r.Route("/api/v1", func(api chi.Router) {
		api.Post("/items", itemHandler.Create)
		api.Get("/items", itemHandler.List)
		api.Get("/items/{id}", itemHandler.GetByID)
		api.Put("/items/{id}", itemHandler.Update)
		api.Delete("/items/{id}", itemHandler.Delete)

		api.Post("/stock/receive", stockHandler.Receive)
		api.Get("/stock/{item_id}", stockHandler.GetByItem)
	})

	return r
}

// Run starts the HTTP server on the given address and blocks until shutdown.
func Run(addr string, handler http.Handler) error {
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("server listening on %s\n", addr)
	return srv.ListenAndServe()
}
