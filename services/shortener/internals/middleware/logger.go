package middleware

import (
	"bytes"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		log.Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Str("remote_addr", r.RemoteAddr).
			Msg("Incoming request")

		ww := &statusWriter{ResponseWriter: w, body: &bytes.Buffer{}}

		defer func() {
			if err := recover(); err != nil {
				log.Error().
					Str("method", r.Method).
					Str("url", r.URL.String()).
					Interface("error", err).
					Msg("Recovered from panic")
				http.Error(ww, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(ww, r)

		log.Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", ww.status).
			Dur("duration", time.Since(startTime)).
			Str("response", ww.body.String()).
			Msg("Completed request")
	})
}

type statusWriter struct {
	http.ResponseWriter
	status int
	body   *bytes.Buffer
}

func (w *statusWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
