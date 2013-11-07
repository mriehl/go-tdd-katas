package sqrt

import (
	"math"
	"testing"
)

func TestShouldSolveFor0(t *testing.T) {
	if actual := Sqrt(0); actual != 0 {
		t.Errorf("Sqrt(0) is not 0 but %v instead", actual)
	}
}

func TestShouldSolveFor1(t *testing.T) {
	assertSqrtDeltaOk(1, 0.0000001, t)
}

func TestShouldSolveForPerfectSquares(t *testing.T) {
	if Sqrt(float64(4)) != float64(2) {
		t.Errorf("Sqrt(4) should be 2 but is %v", Sqrt(4))
	}
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
	actual := Sqrt(value)
	expected := math.Sqrt(value)
	delta := math.Abs(expected - actual)
	if deltaTooBig := delta > allowedDelta; deltaTooBig {
		t.Errorf("Delta for %v is too big (%v), at most %v is allowed\nExpected: %v - Actual: %v", value, delta, allowedDelta, expected, actual)
	}
}
