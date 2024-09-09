// Package api provides HTTP handlers and middleware for the financial transaction system.
package api

import (
	"net/http"
	"time"
	"sync"
	"golang.org/x/time/rate"

	"github.com/adnanlabib1509/go-transaction-engine/pkg/logger"
)

// LoggingMiddleware logs information about each HTTP request.
func LoggingMiddleware(next http.Handler, l logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log the request details
		l.Info("Request processed",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start),
			"remote_addr", r.RemoteAddr,
		)
	})
}

// AuthenticationMiddleware checks for a valid API key in the request header.
func AuthenticationMiddleware(next http.Handler, l logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		
		// In a real application, you would validate the API key against a database or service
		if apiKey != "your-secret-api-key" {
			l.Warn("Authentication failed", "remote_addr", r.RemoteAddr)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// RateLimitingMiddleware implements a simple rate limiting mechanism.
func RateLimitingMiddleware(next http.Handler, l logger.Logger) http.Handler {
	// Create a map to store rate limiters for each IP
	var (
		limiterMu sync.Mutex
		limiters  = make(map[string]*rate.Limiter)
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the IP address from the request
		ip := r.RemoteAddr

		// Create or get rate limiter for this IP
		limiterMu.Lock()
		limiter, exists := limiters[ip]
		if !exists {
			limiter = rate.NewLimiter(rate.Every(time.Second), 10) // 10 requests per second
			limiters[ip] = limiter
		}
		limiterMu.Unlock()

		if !limiter.Allow() {
			l.Warn("Rate limit exceeded", "remote_addr", ip)
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// CORSMiddleware adds Cross-Origin Resource Sharing headers to responses.
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
