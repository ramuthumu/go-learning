package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, word := range strings.Fields("Go ships Go tools") {
		counts[strings.ToLower(word)]++
	}
	keys := make([]string, 0, len(counts))
	for word := range counts {
		keys = append(keys, word)
	}
	sort.Strings(keys)
	for _, word := range keys {
		fmt.Println(word, counts[word])
	}
	original := []int{1, 2, 3}
	view := original[:2]
	view[0] = 99
	fmt.Println("shared array:", original)
	copyOf := append([]int(nil), original...)
	copyOf[0] = 1
	fmt.Println("independent copy:", copyOf)
}
