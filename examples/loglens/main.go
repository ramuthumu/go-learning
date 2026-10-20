// LogLens counts matching lines in a file or standard input.
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Summary struct {
	Lines   int `json:"lines"`
	Matches int `json:"matches"`
}

func scan(r io.Reader, needle string) (Summary, error) {
	var summary Summary
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 4096), 1024*1024)
	for scanner.Scan() {
		summary.Lines++
		if strings.Contains(scanner.Text(), needle) {
			summary.Matches++
		}
	}
	return summary, scanner.Err()
}

func run(args []string, in io.Reader, out, stderr io.Writer) int {
	fs := flag.NewFlagSet("loglens", flag.ContinueOnError)
	fs.SetOutput(stderr)
	needle := fs.String("contains", "ERROR", "substring to count")
	asJSON := fs.Bool("json", false, "write a JSON summary")
	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}
	if fs.NArg() > 1 {
		fmt.Fprintln(stderr, "usage: loglens [flags] [file]")
		return 2
	}
	if fs.NArg() == 1 {
		file, err := os.Open(fs.Arg(0))
		if err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		defer file.Close()
		in = file
	}
	summary, err := scan(in, *needle)
	if err != nil {
		fmt.Fprintln(stderr, "read log:", err)
		return 1
	}
	if *asJSON {
		err = json.NewEncoder(out).Encode(summary)
	} else {
		_, err = fmt.Fprintf(out, "%d matches in %d lines\n", summary.Matches, summary.Lines)
	}
	if err != nil {
		fmt.Fprintln(stderr, "write summary:", err)
		return 1
	}
	return 0
}

func main() { os.Exit(run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)) }
