package sqrt

import (
	"math"
)

func Sqrt(x float64) (approx float64) {
	if x == 0 {
		return 0
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
