package server

import (
	"fmt"
	"net/http"
	"time"

	"warehouse-management/backend/internal/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// New builds the HTTP router with all registered routes.
func New() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/healthz"))

	r.Get("/health", handler.Health)

	// API versioning mount point; item & stock routes attach here in later tasks.
	r.Route("/api/v1", func(api chi.Router) {})

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
