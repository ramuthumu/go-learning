package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Addr    string `json:"addr"`
	Workers int    `json:"workers"`
}

func main() {
	cfg := Config{Addr: "127.0.0.1:8080", Workers: 2}
	decoder := json.NewDecoder(strings.NewReader(`{"workers":4}`))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if addr := os.Getenv("APP_ADDR"); addr != "" {
		cfg.Addr = addr
	}
	if cfg.Workers < 1 {
		fmt.Fprintln(os.Stderr, "workers must be positive")
		os.Exit(1)
	}
	fmt.Printf("addr=%s workers=%d\n", cfg.Addr, cfg.Workers)
}
