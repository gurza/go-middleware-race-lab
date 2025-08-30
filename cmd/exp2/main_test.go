package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestWithResponseWriter(t *testing.T) {
	ts := httptest.NewServer(withResponseWriter(handler()))
	defer ts.Close()

	const cnt = 100
	var wg sync.WaitGroup

	errCh := make(chan error, cnt)
	start := make(chan struct{})

	for i := range cnt {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			<-start

			resp, err := http.Get(ts.URL + "/")
			if err != nil {
				errCh <- fmt.Errorf("goroutine %d: %v", i, err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				errCh <- fmt.Errorf("goroutine %d: unexpected status %d", i, resp.StatusCode)
				return
			}

			found := false
			for _, c := range resp.Cookies() {
				if c.Name == "exp2" {
					found = true
					break
				}
			}
			if !found {
				errCh <- fmt.Errorf("goroutine %d: missing exp2 cookie", i)
			}
		}(i)
	}

	close(start)
	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			t.Error(err)
		}
	}
}
