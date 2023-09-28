package middlewares

import (
	"fmt"
	"net/http"
)

// Middleware function for logging requests
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Logging request details
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// Middleware function for adding headers
func AddHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add custom headers
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// Middleware function for authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Perform authentication logic here
		// Example: Check if the request contains a valid token

		// If authentication fails
		// w.WriteHeader(http.StatusUnauthorized)
		// fmt.Fprintf(w, "Unauthorized")

		// If authentication succeeds, call the next handler
		next.ServeHTTP(w, r)
	})
}
