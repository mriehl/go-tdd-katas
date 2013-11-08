package wc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnNothingForEmptyString(t *testing.T) {
	assert.Equal(t, WordCount(""),
		map[string]uint8{})
}

func TestShouldReturnNothingForWhitespace(t *testing.T) {
	assert.Equal(t, WordCount("        "),
		map[string]uint8{})
}

func TestShouldCountOneWord(t *testing.T) {
	assert.Equal(t, WordCount("foo"),
		map[string]uint8{"foo": 1})
}

func TestShouldCountOneWordMultipleTimes(t *testing.T) {
	assert.Equal(t, WordCount("foo foo"),
		map[string]uint8{"foo": 2})
}

func TestShouldCountTwoWords(t *testing.T) {
	assert.Equal(t, WordCount("foo bar"),
		map[string]uint8{"foo": 1, "bar": 1})
}

func TestShouldCountMultipleWordsMultipleTimes(t *testing.T) {
	assert.Equal(t, WordCount("foo bar yo foo yo test yarrr"),
		map[string]uint8{
			"foo": 2, "bar": 1, "yo": 2,
			"test": 1, "yarrr": 1})
}
