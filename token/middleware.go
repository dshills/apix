package token

import (
	"log"
	"net/http"
	"strings"
)

// Handler returns a handler function that validates the authorization header
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if h == "" {
			log.Printf("[ERROR] %v %v %v %v", "No auth header included", r.RemoteAddr, r.Method, r.URL)
			http.Error(w, "Forbidden 403: Missing Authorization header", http.StatusForbidden)
			return
		}

		h = strings.Trim(h, " \t\n\r")
		parts := strings.Split(h, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			log.Printf("[ERROR] %v %v %v %v", "Authorization header format must be Bearer {token}", r.RemoteAddr, r.Method, r.URL)
			http.Error(w, "Forbidden 403: Authorization header format must be Bearer {token}", http.StatusForbidden)
		}

		_, err := Decode(parts[1])
		if err != nil {
			log.Printf("[ERROR] %v %v %v %v", err, r.RemoteAddr, r.Method, r.URL)
			http.Error(w, "Forbidden 403: "+err.Error(), http.StatusForbidden)
		}

		next.ServeHTTP(w, r)
	})
}
