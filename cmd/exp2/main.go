package main

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
)

type contextKey string

const httpResponseWriterKey contextKey = "httpResponseWriter"

func withResponseWriter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), httpResponseWriterKey, w)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wx := r.Context().Value(httpResponseWriterKey).(http.ResponseWriter)
		http.SetCookie(wx, &http.Cookie{
			Name:  "exp2",
			Value: fmt.Sprintf("%d", rand.IntN(1_000_000)),
			Path:  "/",
		})

		w.Write([]byte("kek"))
	})
}

func handler2() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cs, ok := cookieSetterFrom(r.Context())
		if !ok {
			w.Write([]byte("fail"))
			return
		}
		cs.Set(&http.Cookie{
			Name:  "exp2",
			Value: fmt.Sprintf("%d", rand.IntN(1_000_000)),
			Path:  "/",
		})

		w.Write([]byte("kek"))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /", withCookieSetter(handler2()))

	fmt.Print("http://localhost:3001\n")
	err := http.ListenAndServe(":3001", mux)
	if err != nil {
		log.Fatal(err)
	}
}
