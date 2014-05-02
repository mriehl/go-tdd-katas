package sqrt

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestShouldSolveFor0(t *testing.T) {
	actual, _ := Sqrt(0)
	assert.Equal(t, 0, actual)
}

func TestShouldRaiseErrorWhenInputIsNegative(t *testing.T) {
	_, err := Sqrt(-5)
	assert.Equal(t, "Cannot Sqrt() negative number -5", err.Error())
}

func TestShouldSolveFor1(t *testing.T) {
	assertSqrtDeltaOk(1, 0.0000001, t)
}

func TestShouldSolveForPerfectSquares(t *testing.T) {
	actual, _ := Sqrt(4)
	assert.Equal(t, float64(2), actual)
}

func TestShouldSolveFor9Roughly(t *testing.T) {
	assertSqrtDeltaOk(9, 0.01, t)
}

func TestShouldSolveFor9Precisely(t *testing.T) {
	assertSqrtDeltaOk(9, 0.0001, t)
}

func TestShouldSolveFor18Roughly(t *testing.T) {
	assertSqrtDeltaOk(18, 0.01, t)
}

func TestShouldSolveFor18Precisely(t *testing.T) {
	assertSqrtDeltaOk(18, 0.000001, t)
}

func assertSqrtDeltaOk(value float64, allowedDelta float64, t *testing.T) {
	actual, err := Sqrt(value)
	if err != nil {
		t.Errorf("Got an error calling Sqrt(%v) : %v", value, err)
	}
	expected := math.Sqrt(value)
	delta := math.Abs(expected - actual)
	if deltaTooBig := delta > allowedDelta; deltaTooBig {
		t.Errorf("Delta for %v is too big (%v), at most %v is allowed\nExpected: %v - Actual: %v", value, delta, allowedDelta, expected, actual)
	}
}
