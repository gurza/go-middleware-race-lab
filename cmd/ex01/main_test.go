package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestMiddlewareDataRace(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	m := NewMiddleware(h, true)

	var wg sync.WaitGroup
	start := make(chan struct{})

	const cnt = 100
	for range cnt {
		wg.Add(2)

		go func() {
			<-start
			req := httptest.NewRequest(http.MethodGet, "/admin/panel", nil)
			rr := httptest.NewRecorder()
			m.ServeHTTP(rr, req)
			wg.Done()
		}()

		go func() {
			<-start
			req := httptest.NewRequest(http.MethodGet, "/user/home", nil)
			rr := httptest.NewRecorder()
			m.ServeHTTP(rr, req)
			wg.Done()
		}()
	}

	close(start)
	wg.Wait()
}
