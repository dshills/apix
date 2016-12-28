package logrequest

import (
	"log"
	"net/http"
)

// Handler will log the HTTP requests
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO] %v %v %v", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
