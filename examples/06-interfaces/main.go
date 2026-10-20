package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func report(w io.Writer, name string) error {
	_, err := fmt.Fprintf(w, "report: %s\n", name)
	return err
}

func unique[T comparable](items []T) []T {
	seen := make(map[T]bool)
	out := make([]T, 0, len(items))
	for _, item := range items {
		if !seen[item] {
			out = append(out, item)
			seen[item] = true
		}
	}
	return out
}

func main() {
	var buffer bytes.Buffer
	if err := report(&buffer, "daily"); err != nil {
		panic(err)
	}
	fmt.Print(buffer.String())
	if err := report(os.Stdout, "weekly"); err != nil {
		panic(err)
	}
	fmt.Println(unique([]string{"go", "sql", "go"}))
}
