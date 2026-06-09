package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		slog.Info("Incoming Request",
		   slog.String("method", r.Method),
		   slog.String("path", r.URL.Path),
		   slog.Time("timestamp", start),
		)

		next.ServeHTTP(w, r)
	})
}