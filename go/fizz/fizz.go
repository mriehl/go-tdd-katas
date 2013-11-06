package fizz

const (
	FizzDiv  = 3
	FizzName = "fizz"
	BuzzDiv  = 5
	BuzzName = "buzz"
)

func FizzBuzz(number uint32) (accu string) {
	if number%FizzDiv == 0 {
		accu += FizzName
	}
	if number%BuzzDiv == 0 {
		accu += BuzzName
	}

	return
}
