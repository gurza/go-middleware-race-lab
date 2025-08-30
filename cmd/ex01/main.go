package main

import (
	"fmt"
	"net/http"
	"strings"
)

func NewMiddleware(handler http.Handler, rateLimitEnabled bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/admin") {
			rateLimitEnabled = false
		}
		if rateLimitEnabled {
			fmt.Printf("path=%s rate_limit_enabled=yes\n", r.URL.Path)
		} else {
			fmt.Printf("path=%s rate_limit_enabled=no\n", r.URL.Path)
		}
		handler.ServeHTTP(w, r)
	})
}

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("kek"))
	})
	m := NewMiddleware(h, true)
	http.Handle("GET /", m)

	http.ListenAndServe(":3001", nil)
}
