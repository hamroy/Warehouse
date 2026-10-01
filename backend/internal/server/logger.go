package server

import (
	"log"
	"net/http"
	"time"
)

// statusWriter wraps http.ResponseWriter to capture the status code.
type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// Logger is request-logging middleware.
// It logs METHOD, PATH, STATUS, and DURATION, e.g.:
//
//	POST /api/v1/items 201 12ms
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(sw, r)

		duration := time.Since(start)
		log.Printf("%s %s %d %s",
			r.Method, r.URL.Path, sw.status, duration.Round(time.Millisecond))
	})
}
