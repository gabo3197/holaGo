package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Ejecuta el siguiente handler
		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("[%s] %s%s - %v", r.Method, r.Host, r.URL.Path, duration)
	})
}
