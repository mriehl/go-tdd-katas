package fizz

import (
	"testing"
)

func TestNoFizzBuzz(t *testing.T) {
	assertFizzbuzzIs(1, "", t)
	assertFizzbuzzIs(2, "", t)
	assertFizzbuzzIs(4, "", t)
	assertFizzbuzzIs(7, "", t)
	assertFizzbuzzIs(8, "", t)
	assertFizzbuzzIs(11, "", t)
}

func TestFizzNoBuzz(t *testing.T) {
	assertFizzbuzzIs(3, "fizz", t)
	assertFizzbuzzIs(6, "fizz", t)
	assertFizzbuzzIs(9, "fizz", t)
	assertFizzbuzzIs(12, "fizz", t)
	assertFizzbuzzIs(18, "fizz", t)
}

func TestBuzzNoFizz(t *testing.T) {
	assertFizzbuzzIs(5, "buzz", t)
	assertFizzbuzzIs(20, "buzz", t)
	assertFizzbuzzIs(25, "buzz", t)
}

func TestFizzBuzz(t *testing.T) {
	assertFizzbuzzIs(15, "fizzbuzz", t)
}

func assertFizzbuzzIs(number uint32, expected string, t *testing.T) {
	if actual := FizzBuzz(number); actual != expected {
		t.Errorf("Got %q but expected %q", actual, expected)
	}
}
