package wc

import (
	"strings"
)

func WordCount(input string) map[string]uint8 {
	wordCounts := make(map[string]uint8)

	tokens := strings.Fields(input)
	for _, token := range tokens {
		wordCounts[token] += 1
	}

	return wordCounts
}
