package wc

import (
	"strings"
)

func WordCount(input string) map[string]uint8 {
	wordCounts := make(map[string]uint8)

	words := strings.Fields(input)
	for _, word := range words {
		wordCounts[word] += 1
	}

	return wordCounts
}
