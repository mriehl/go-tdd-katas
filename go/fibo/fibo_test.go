package fibo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnFibonacciSequence(t *testing.T) {
	nextFibo := Fibonacci()
	assert.Equal(t, nextFibo(), 0)
	assert.Equal(t, nextFibo(), 1)
	assert.Equal(t, nextFibo(), 1)
	assert.Equal(t, nextFibo(), 2)
	assert.Equal(t, nextFibo(), 3)
	assert.Equal(t, nextFibo(), 5)
	assert.Equal(t, nextFibo(), 8)
	assert.Equal(t, nextFibo(), 13)
	assert.Equal(t, nextFibo(), 21)
	assert.Equal(t, nextFibo(), 34)
	assert.Equal(t, nextFibo(), 55)
	assert.Equal(t, nextFibo(), 89)
	assert.Equal(t, nextFibo(), 144)
	assert.Equal(t, nextFibo(), 233)
}

func TestShouldNotRelyOnGlobalState(t *testing.T) {
	fibo1 := Fibonacci()
	fibo2 := Fibonacci()
	assert.Equal(t, fibo1(), 0)
	assert.Equal(t, fibo2(), 0)
	assert.Equal(t, fibo1(), 1)
	assert.Equal(t, fibo2(), 1)
}
