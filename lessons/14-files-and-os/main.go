package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 1. Command-line args and Environment variables
	fmt.Println("--- 1. OS & Environment ---")
	fmt.Printf("Program binary: %s\n", os.Args[0])
	user := os.Getenv("USER")
	if user == "" {
		user = "Developer"
	}
	fmt.Printf("Current system user: %s\n", user)

	// 2. Writing a file using os.WriteFile (Quick one-shot write)
	tempDir := os.TempDir()
	filePath := filepath.Join(tempDir, "go_learning_demo.txt")

	content := "Line 1: Go makes I/O fast and explicit\nLine 2: defer ensures resources are closed\nLine 3: Python developers love Go's speed\n"
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	fmt.Printf("\n--- 2. File written to: %s ---\n", filePath)

	// 3. Opening and reading line-by-line using bufio.Scanner (Idiomatic Go)
	// Equivalent to Python's:
	// with open(filePath, "r") as f:
	//     for line in f: ...
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	// Crucial Go pattern: ALWAYS defer file.Close() immediately after checking err != nil!
	defer file.Close()

	fmt.Println("\n--- 3. Reading line-by-line with bufio.Scanner ---")
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		text := scanner.Text()
		fmt.Printf("Line %d [%d chars]: %s\n", lineCount, len(text), strings.ToUpper(text))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
	}

	// Clean up temporary file
	_ = os.Remove(filePath)
	fmt.Println("\nCleaned up temporary test file.")
}
