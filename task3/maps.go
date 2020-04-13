package task3

import "strings"

func WordCounts(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		m[word]++
	}
	return m
}
