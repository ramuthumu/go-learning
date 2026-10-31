package main

import (
	"encoding/json"
	"fmt"
)

// ServerConfig models our server settings with JSON struct tags.
// Note: Fields MUST be capitalized (exported) so encoding/json can see them!
// The `json:"..."` tag maps Go PascalCase fields to JSON snake_case or camelCase keys.
type ServerConfig struct {
	Host     string   `json:"host"`
	Port     int      `json:"port"`
	Services []string `json:"services"`
	IsActive bool     `json:"is_active"`
	Secret   string   `json:"-"` // "-" tells Go to NEVER include this field in JSON output!
}

func main() {
	// 1. Marshalling: Converting Go Struct -> JSON String (like Python's json.dumps)
	cfg := ServerConfig{
		Host:     "192.168.1.236",
		Port:     32400,
		Services: []string{"plex", "pihole", "homeassistant"},
		IsActive: true,
		Secret:   "super-secret-token",
	}

	// json.MarshalIndent formats with clean indentation
	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	fmt.Println("--- 1. Marshaled JSON Output ---")
	fmt.Println(string(jsonData))

	// 2. Unmarshaling: Converting JSON String -> Go Struct (like Python's json.loads)
	incomingRawJSON := `{
		"host": "192.168.1.197",
		"port": 8080,
		"services": ["bravia-tv", "kodi"],
		"is_active": false
	}`

	var parsedCfg ServerConfig
	// Notice we pass &parsedCfg (pointer) so Unmarshal can fill its fields!
	err = json.Unmarshal([]byte(incomingRawJSON), &parsedCfg)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return
	}

	fmt.Println("\n--- 2. Unmarshaled Struct ---")
	fmt.Printf("Parsed Host: %s\n", parsedCfg.Host)
	fmt.Printf("Parsed Port: %d\n", parsedCfg.Port)
	fmt.Printf("Services (%d): %v\n", len(parsedCfg.Services), parsedCfg.Services)

	// 3. Parsing Dynamic / Unknown JSON using map[string]any
	// Python dicts hold arbitrary types; in Go, map[string]any does the same!
	fmt.Println("\n--- 3. Dynamic JSON without a fixed struct ---")
	dynamicJSON := `{"cluster": "home-lab", "nodes": 3, "healthy": true}`

	var dynamicMap map[string]any
	_ = json.Unmarshal([]byte(dynamicJSON), &dynamicMap)
	for k, v := range dynamicMap {
		fmt.Printf("Key: %-10s -> Value: %v (%T)\n", k, v, v)
	}
}
