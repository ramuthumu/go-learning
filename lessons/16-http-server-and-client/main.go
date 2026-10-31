package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type StatusResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	UptimeSec int    `json:"uptime_sec"`
}

func main() {
	mux := http.NewServeMux()

	// 1. Defining an HTTP Handler (FastAPI / Flask equivalent in Go's standard library)
	mux.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		// Set JSON header
		w.Header().Set("Content-Type", "application/json")

		res := StatusResponse{
			Status:    "healthy",
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			UptimeSec: 3600,
		}

		// Stream JSON directly into the HTTP response writer!
		_ = json.NewEncoder(w).Encode(res)
	})

	server := &http.Server{
		Addr:    "127.0.0.1:8989",
		Handler: mux,
	}

	// 2. Launch the HTTP server in a background Goroutine
	go func() {
		fmt.Println("HTTP Server listening on http://127.0.0.1:8989/api/status ...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	// Give the server 50 milliseconds to bind to port
	time.Sleep(50 * time.Millisecond)

	// 3. Built-in HTTP Client (equivalent to Python's requests.get())
	fmt.Println("\n--- Making HTTP Client Request ---")
	resp, err := http.Get("http://127.0.0.1:8989/api/status")
	if err != nil {
		fmt.Printf("HTTP GET failed: %v\n", err)
		return
	}
	// Golden Rule: ALWAYS defer closing the response body!
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Printf("HTTP Status Code: %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
	fmt.Printf("Response Body:\n%s", string(bodyBytes))

	// 4. Decode into Go struct
	var status StatusResponse
	_ = json.Unmarshal(bodyBytes, &status)
	fmt.Printf("Decoded Status: %s at %s\n", status.Status, status.Timestamp)

	// Shutdown server cleanly
	_ = server.Close()
	fmt.Println("\nServer shut down cleanly.")
}
