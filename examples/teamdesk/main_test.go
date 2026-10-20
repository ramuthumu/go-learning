package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

func TestAPIAndTemplates(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tickets.db")
	s, err := openStore(path)
	if err != nil {
		t.Fatal(err)
	}
	defer s.db.Close()
	h := (&app{s}).routes()
	for _, tc := range []struct {
		body   string
		status int
	}{
		{`{"title":"export"}`, 201}, {`{"title":""}`, 400},
		{`{"title":"x","admin":true}`, 400}, {`{"title":"x"}{"title":"y"}`, 400},
	} {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest("POST", "/api/tickets", strings.NewReader(tc.body)))
		if rec.Code != tc.status {
			t.Errorf("body=%s status=%d want=%d", tc.body, rec.Code, tc.status)
		}
	}
	if _, err = s.Create(context.Background(), "<script>alert(1)</script>"); err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
	if rec.Code != 200 || strings.Contains(rec.Body.String(), "<script>alert") {
		t.Fatalf("template escaped incorrectly: %s", rec.Body.String())
	}
	rec = httptest.NewRecorder()
	req := httptest.NewRequest("POST", "http://example.com/tickets", strings.NewReader("title=x"))
	req.Header.Set("Origin", "https://attacker.example")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusForbidden {
		t.Fatalf("cross-origin POST status=%d", rec.Code)
	}
	if err := s.db.Close(); err != nil {
		t.Fatal(err)
	}
	reopened, err := openStore(path)
	if err != nil {
		t.Fatal(err)
	}
	defer reopened.db.Close()
	tickets, err := reopened.List(context.Background())
	if err != nil || len(tickets) != 2 {
		t.Fatalf("persisted=%v err=%v", tickets, err)
	}
}
