package roman

type RomanAtom struct {
	DecimalValue uint16
	Rune         string
}

func DivMod(number uint16, divisor uint16) (quotient, remainder uint16) {
	quotient = number / divisor
	remainder = number % divisor
	return
}

func AsRoman(decimal uint16) (roman string) {

	var mapping = []RomanAtom{
		{DecimalValue: 1000, Rune: "M"},
		{DecimalValue: 900, Rune: "CM"},
		{DecimalValue: 500, Rune: "D"},
		{DecimalValue: 400, Rune: "CD"},
		{DecimalValue: 100, Rune: "C"},
		{DecimalValue: 90, Rune: "XC"},
		{DecimalValue: 50, Rune: "L"},
		{DecimalValue: 40, Rune: "XL"},
		{DecimalValue: 10, Rune: "X"},
		{DecimalValue: 9, Rune: "IX"},
		{DecimalValue: 5, Rune: "V"},
		{DecimalValue: 4, Rune: "IV"},
		{DecimalValue: 1, Rune: "I"},
	}

	for _, romanAtom := range mapping {
		amount_of_runes, remaining_decimal := DivMod(decimal, romanAtom.DecimalValue)
		decimal = remaining_decimal
		// TODO overload with Rune * amount_of_runes would be cleaner
		// is it possible?
		for count := uint16(1); count <= amount_of_runes; count++ {
			roman += romanAtom.Rune
		}

	}

	return
}
