package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// Device represents a network client in our home lab (e.g. BRAVIA TV, Plex, Pi-hole)
type Device struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	IP       string    `json:"ip"`
	Status   string    `json:"status"` // "online" or "offline"
	LastSeen time.Time `json:"last_seen"`
}

// DeviceStore is an interface for managing devices (Lesson 9: Interfaces)
type DeviceStore interface {
	GetAll() []Device
	Add(d Device) error
	GetByID(id string) (Device, bool)
}

// InMemoryStore is a thread-safe implementation of DeviceStore
type InMemoryStore struct {
	mu      sync.RWMutex      // Protects concurrent map reads and writes
	devices map[string]Device // Lesson 6: Maps
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		devices: make(map[string]Device),
	}
}

// Add safely adds a new device (Pointer receiver: Lesson 8 & 12)
func (s *InMemoryStore) Add(d Device) error {
	s.mu.Lock()
	defer s.mu.Unlock() // Lesson 13: Defer cleanup

	if d.ID == "" || d.IP == "" {
		return errors.New("device id and ip are required")
	}

	d.LastSeen = time.Now().UTC()
	s.devices[d.ID] = d
	return nil
}

// GetAll safely returns all devices as a slice (Lesson 5: Slices)
func (s *InMemoryStore) GetAll() []Device {
	s.mu.RLock() // Read lock allows concurrent readers
	defer s.mu.RUnlock()

	result := make([]Device, 0, len(s.devices))
	for _, dev := range s.devices {
		result = append(result, dev)
	}
	return result
}

func (s *InMemoryStore) GetByID(id string) (Device, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	dev, ok := s.devices[id]
	return dev, ok
}

// Server holds HTTP routing and dependencies
type Server struct {
	store DeviceStore
	mux   *http.ServeMux
}

func NewServer(store DeviceStore) *Server {
	s := &Server{
		store: store,
		mux:   http.NewServeMux(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("/health", s.handleHealth)
	s.mux.HandleFunc("/api/devices", s.handleDevices)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
		"time":   time.Now().UTC().Format(time.RFC3339),
	})
}

func (s *Server) handleDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		devices := s.store.GetAll()
		_ = json.NewEncoder(w).Encode(devices)

	case http.MethodPost:
		var req Device
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error": "invalid json body"}`, http.StatusBadRequest)
			return
		}

		if err := s.store.Add(req); err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": "device registered successfully",
			"id":      req.ID,
		})

	default:
		http.Error(w, `{"error": "method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

func main() {
	store := NewInMemoryStore()

	// Seed with an initial device
	_ = store.Add(Device{
		ID:     "plex-minilink",
		Name:   "Plex Media Server",
		IP:     "192.168.1.236",
		Status: "online",
	})

	srv := NewServer(store)
	httpServer := &http.Server{
		Addr:    "127.0.0.1:9095",
		Handler: srv.mux,
	}

	// 1. Run HTTP Server in a Goroutine (Lesson 10)
	go func() {
		fmt.Println("🚀 Capstone Service running at http://127.0.0.1:9095 ...")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP server error: %v\n", err)
		}
	}()

	time.Sleep(50 * time.Millisecond) // Wait briefly for server startup

	// 2. Client Test: Health Check
	fmt.Println("\n--- 1. Testing GET /health ---")
	resp, err := http.Get("http://127.0.0.1:9095/health")
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Health response: %s", string(body))
	}

	// 3. Client Test: POST /api/devices (Register BRAVIA TV)
	fmt.Println("--- 2. Testing POST /api/devices ---")
	newDev := Device{
		ID:     "bravia-vh21",
		Name:   "Sony BRAVIA 4K TV",
		IP:     "192.168.1.197",
		Status: "online",
	}
	payload, _ := json.Marshal(newDev)
	postResp, err := http.Post("http://127.0.0.1:9095/api/devices", "application/json", bytes.NewBuffer(payload))
	if err == nil {
		defer postResp.Body.Close()
		body, _ := io.ReadAll(postResp.Body)
		fmt.Printf("POST Status: %d\nResponse: %s", postResp.StatusCode, string(body))
	}

	// 4. Client Test: GET /api/devices (List all registered devices)
	fmt.Println("--- 3. Testing GET /api/devices ---")
	getResp, err := http.Get("http://127.0.0.1:9095/api/devices")
	if err == nil {
		defer getResp.Body.Close()
		body, _ := io.ReadAll(getResp.Body)
		var devices []Device
		_ = json.Unmarshal(body, &devices)

		fmt.Printf("Registered Devices Count: %d\n", len(devices))
		for _, d := range devices {
			fmt.Printf("  • [%s] %s (%s) - Status: %s\n", d.ID, d.Name, d.IP, d.Status)
		}
	}

	// 5. Graceful Shutdown with Context Timeout (Lesson 18: Context)
	fmt.Println("\n--- 4. Graceful Shutdown ---")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("Shutdown error: %v\n", err)
	} else {
		fmt.Println("✅ Capstone Service stopped gracefully with zero dropped requests.")
	}
}
