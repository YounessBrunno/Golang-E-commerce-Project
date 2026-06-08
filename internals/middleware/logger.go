package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		log.Printf(
			"%s %s %s\n",
			r.Method,
			r.URL.Path,
			start.Format(time.RFC3339),
		)

		next.ServeHTTP(w, r)
	})
}