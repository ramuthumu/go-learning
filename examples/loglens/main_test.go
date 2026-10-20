package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	cases := []struct {
		name, input, needle string
		want                Summary
	}{
		{"empty", "", "ERROR", Summary{}},
		{"mixed", "INFO start\nERROR failed\n", "ERROR", Summary{2, 1}},
		{"no newline", "ERROR failed", "ERROR", Summary{1, 1}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := scan(strings.NewReader(tc.input), tc.needle)
			if err != nil || got != tc.want {
				t.Fatalf("got %+v, %v; want %+v", got, err, tc.want)
			}
		})
	}
}

func TestRunJSONAndUsage(t *testing.T) {
	var out, stderr bytes.Buffer
	code := run([]string{"-json"}, strings.NewReader("ERROR bad\n"), &out, &stderr)
	if code != 0 || out.String() != "{\"lines\":1,\"matches\":1}\n" {
		t.Fatalf("code=%d out=%q err=%q", code, out.String(), stderr.String())
	}
	if got := run([]string{"-unknown"}, strings.NewReader(""), &out, &stderr); got != 2 {
		t.Fatalf("usage code=%d", got)
	}
}

func FuzzScan(f *testing.F) {
	f.Add("INFO ok\nERROR failed\n", "ERROR")
	f.Fuzz(func(t *testing.T, input, needle string) {
		got, err := scan(strings.NewReader(input), needle)
		if err == nil && (got.Matches < 0 || got.Matches > got.Lines) {
			t.Fatalf("invalid summary: %+v", got)
		}
	})
}

func BenchmarkScan(b *testing.B) {
	input := strings.Repeat("INFO request complete\nERROR request failed\n", 100)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := scan(strings.NewReader(input), "ERROR"); err != nil {
			b.Fatal(err)
		}
	}
}
