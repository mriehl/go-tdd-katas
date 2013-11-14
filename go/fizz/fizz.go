package fizz

import (
	"strconv"
)

const (
	FizzDiv  = 3
	FizzName = "fizz"
	BuzzDiv  = 5
	BuzzName = "buzz"
)

func FizzBuzz(number int) (accu string) {
	if number%FizzDiv == 0 {
		accu += FizzName
	}
	if number%BuzzDiv == 0 {
		accu += BuzzName
	}

	if isNeitherFizzNorBuzz := (len(accu) == 0); isNeitherFizzNorBuzz {
		accu += strconv.Itoa(number)
	}

	return
}
