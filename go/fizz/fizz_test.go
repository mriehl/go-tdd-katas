package fizz

import (
	"testing"
)

func TestNoFizzBuzz(t *testing.T) {
	assertEqualStr(FizzBuzz(1), "", t)
	assertEqualStr(FizzBuzz(2), "", t)
	assertEqualStr(FizzBuzz(4), "", t)
	assertEqualStr(FizzBuzz(7), "", t)
	assertEqualStr(FizzBuzz(8), "", t)
	assertEqualStr(FizzBuzz(11), "", t)
}

func TestFizzNoBuzz(t *testing.T) {
	assertEqualStr(FizzBuzz(3), "fizz", t)
	assertEqualStr(FizzBuzz(6), "fizz", t)
	assertEqualStr(FizzBuzz(9), "fizz", t)
	assertEqualStr(FizzBuzz(12), "fizz", t)
	assertEqualStr(FizzBuzz(18), "fizz", t)
}

func TestBuzzNoFizz(t *testing.T) {
	assertEqualStr(FizzBuzz(5), "buzz", t)
	assertEqualStr(FizzBuzz(20), "buzz", t)
	assertEqualStr(FizzBuzz(25), "buzz", t)
}

func TestFizzBuzz(t *testing.T) {
	assertEqualStr(FizzBuzz(15), "fizzbuzz", t)
}

func assertEqualBool(actual, expected bool, t *testing.T) {
	if actual != expected {
		t.Errorf("Got %v but expected %v", actual, expected)
	}
}

func assertEqualStr(actual, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("Got %v but expected %v", actual, expected)
	}
}
