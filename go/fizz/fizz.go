package fizz

import (
	"strconv"
)

func FizzBuzz (number int) string {
	switch {
	case number % 3 == 0 && number % 5 == 0:
		return "fizzbuzz"
	case number % 3 == 0:
		return "fizz"
	case number % 5 == 0:
		return "buzz"
	default:
		return  strconv.Itoa(number)
	}
}
