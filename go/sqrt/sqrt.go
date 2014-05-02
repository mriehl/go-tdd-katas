package sqrt

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt struct {
	input float64
}

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt() negative number %v", err.input)
}

func Sqrt(x float64) (approx float64, err error) {
	if x < 0 {
		return 0, ErrNegativeSqrt{input: x}
	}
	if x == 0 {
		return 0, nil
	}
	approx = float64(x / 2)

	delta := float64(1000)

	for delta > float64(0.01) {
		oldApprox := approx
		approx = approx - (approx*approx-x)/(2*approx)
		newApprox := approx
		delta = math.Abs(newApprox - oldApprox)
	}
	return
}
