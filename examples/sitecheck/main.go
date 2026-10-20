// Sitecheck probes operator-supplied URLs with bounded concurrency.
package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"sync"
	"time"
)

type Result struct {
	URL    string
	Status int
	Err    error
}

func check(ctx context.Context, client *http.Client, urls []string, workers int) []Result {
	if workers < 1 {
		workers = 1
	}
	jobs := make(chan string)
	results := make(chan Result)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range jobs {
				result := Result{URL: url}
				req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
				if err == nil {
					var resp *http.Response
					resp, err = client.Do(req)
					if err == nil {
						result.Status = resp.StatusCode
						resp.Body.Close()
					}
				}
				result.Err = err
				results <- result
			}
		}()
	}
	go func() {
		defer close(jobs)
		for _, url := range urls {
			select {
			case jobs <- url:
			case <-ctx.Done():
				return
			}
		}
	}()
	go func() { wg.Wait(); close(results) }()
	out := make([]Result, 0, len(urls))
	for result := range results {
		out = append(out, result)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].URL < out[j].URL })
	return out
}

func main() {
	workers := flag.Int("workers", 3, "maximum simultaneous requests")
	timeout := flag.Duration("timeout", 3*time.Second, "timeout for each request")
	flag.Parse()
	if *workers < 1 || *workers > 64 || *timeout <= 0 || flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "usage: sitecheck -workers 3 -timeout 3s URL...")
		os.Exit(2)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	client := &http.Client{Timeout: *timeout}
	failed := false
	for _, result := range check(ctx, client, flag.Args(), *workers) {
		if result.Err != nil {
			fmt.Printf("FAIL %s: %v\n", result.URL, result.Err)
			failed = true
			continue
		}
		fmt.Printf("%d %s\n", result.Status, result.URL)
		if result.Status >= 400 {
			failed = true
		}
	}
	if failed || ctx.Err() != nil {
		os.Exit(1)
	}
}
