package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestCheckBoundsAndFailures(t *testing.T) {
	var active, peak atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := active.Add(1)
		defer active.Add(-1)
		for old := peak.Load(); n > old; old = peak.Load() {
			if peak.CompareAndSwap(old, n) {
				break
			}
		}
		time.Sleep(time.Millisecond)
		if r.URL.Path == "/bad" {
			w.WriteHeader(503)
			return
		}
		w.WriteHeader(204)
	}))
	defer srv.Close()
	urls := []string{srv.URL, srv.URL, srv.URL, srv.URL + "/bad"}
	got := check(context.Background(), srv.Client(), urls, 2)
	if len(got) != 4 || peak.Load() > 2 {
		t.Fatalf("results=%d peak=%d", len(got), peak.Load())
	}
	bad := 0
	for _, r := range got {
		if r.Err != nil {
			t.Fatal(r.Err)
		}
		if r.Status == 503 {
			bad++
		}
	}
	if bad != 1 {
		t.Fatalf("503 count=%d", bad)
	}
}

func TestCancelledRequest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	got := check(ctx, &http.Client{Timeout: time.Second}, []string{"http://127.0.0.1:1"}, 1)
	for _, r := range got {
		if r.Err == nil {
			t.Fatal("cancelled request succeeded")
		}
	}
}
