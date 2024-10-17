package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

// LoggingMiddleware logs each incoming request and its response
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Log request details
		log.Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Str("remote_addr", r.RemoteAddr).
			Msg("Incoming request")

		// Wrap the ResponseWriter to capture the status code
		ww := &statusWriter{ResponseWriter: w}
		next.ServeHTTP(ww, r)

		// Log response details
		log.Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", ww.status).
			Dur("duration", time.Since(startTime)).
			Msg("Completed request")
	})
}

// statusWriter is a wrapper for http.ResponseWriter that captures the status code
type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
