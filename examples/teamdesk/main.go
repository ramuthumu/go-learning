// TeamDesk is a local teaching app. Add identity and authorization before exposing it.
package main

import (
	"bytes"
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"errors"
	"flag"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

//go:embed page.html
var pageHTML string
var page = template.Must(template.New("tickets").Parse(pageHTML))

type app struct{ store *Store }

func (a *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", a.home)
	mux.HandleFunc("POST /tickets", a.createForm)
	mux.HandleFunc("POST /tickets/{id}/complete", a.complete)
	mux.HandleFunc("GET /api/tickets", a.listAPI)
	mux.HandleFunc("POST /api/tickets", a.createAPI)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	// Modern browsers send Origin / Sec-Fetch-Site. This protects unsafe cross-origin requests.
	return http.NewCrossOriginProtection().Handler(mux)
}

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	tickets, err := a.store.List(r.Context())
	if err != nil {
		http.Error(w, "cannot load tickets", 500)
		return
	}
	var buf bytes.Buffer
	if err := page.Execute(&buf, tickets); err != nil {
		http.Error(w, "cannot render page", 500)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(buf.Bytes())
}

func (a *app) createForm(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", 400)
		return
	}
	title := strings.TrimSpace(r.PostForm.Get("title"))
	if title == "" || len(title) > 200 {
		http.Error(w, "title must be 1–200 bytes", 400)
		return
	}
	if _, err := a.store.Create(r.Context(), title); err != nil {
		http.Error(w, "cannot create ticket", 500)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *app) complete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		http.Error(w, "invalid id", 400)
		return
	}
	err = a.store.Complete(r.Context(), id)
	if errors.Is(err, sql.ErrNoRows) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, "cannot update ticket", 500)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *app) listAPI(w http.ResponseWriter, r *http.Request) {
	tickets, err := a.store.List(r.Context())
	if err != nil {
		http.Error(w, "cannot load tickets", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tickets); err != nil {
		slog.Error("encode tickets", "error", err)
	}
}

func (a *app) createAPI(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var input struct {
		Title string `json:"title"`
	}
	if err := decoder.Decode(&input); err != nil {
		http.Error(w, "invalid JSON", 400)
		return
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		http.Error(w, "send one JSON object", 400)
		return
	}
	input.Title = strings.TrimSpace(input.Title)
	if input.Title == "" || len(input.Title) > 200 {
		http.Error(w, "title must be 1–200 bytes", 400)
		return
	}
	ticket, err := a.store.Create(r.Context(), input.Title)
	if err != nil {
		http.Error(w, "cannot create ticket", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(ticket); err != nil {
		slog.Error("encode ticket", "error", err)
	}
}

func run() error {
	addr := flag.String("addr", "127.0.0.1:8080", "listen address")
	dbPath := flag.String("db", "teamdesk.db", "SQLite file")
	flag.Parse()
	store, err := openStore(*dbPath)
	if err != nil {
		return err
	}
	defer store.db.Close()
	srv := &http.Server{Addr: *addr, Handler: (&app{store}).routes(), ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 60 * time.Second}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	result := make(chan error, 1)
	go func() { result <- srv.ListenAndServe() }()
	slog.Info("TeamDesk started", "addr", *addr)
	select {
	case err := <-result:
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			_ = srv.Close()
			return err
		}
		return nil
	}
}

func main() {
	if err := run(); err != nil {
		slog.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
