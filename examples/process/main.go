package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	output, err := exec.CommandContext(ctx, "go", "version").CombinedOutput()
	if err != nil {
		return fmt.Errorf("go version failed: %w (%s)", err, output)
	}
	fmt.Print(string(output))
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
