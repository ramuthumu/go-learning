package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	output, err := exec.CommandContext(ctx, "go", "version").CombinedOutput()
	if err != nil {
		fmt.Printf("go version failed: %v (%s)\n", err, output)
		return
	}
	fmt.Print(string(output))
}
