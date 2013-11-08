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

func asRoman(decimal uint16) (roman string) {

	var mapping = []RomanAtom{
		RomanAtom{DecimalValue: 1000, Rune: "M"},
		RomanAtom{DecimalValue: 900, Rune: "CM"},
		RomanAtom{DecimalValue: 500, Rune: "D"},
		RomanAtom{DecimalValue: 400, Rune: "CD"},
		RomanAtom{DecimalValue: 100, Rune: "C"},
		RomanAtom{DecimalValue: 90, Rune: "XC"},
		RomanAtom{DecimalValue: 50, Rune: "L"},
		RomanAtom{DecimalValue: 40, Rune: "XL"},
		RomanAtom{DecimalValue: 10, Rune: "X"},
		RomanAtom{DecimalValue: 9, Rune: "IX"},
		RomanAtom{DecimalValue: 5, Rune: "V"},
		RomanAtom{DecimalValue: 4, Rune: "IV"},
		RomanAtom{DecimalValue: 1, Rune: "I"},
	}

	for _, romanAtom := range mapping {
		amount_of_runes, remaining_decimal := DivMod(decimal, romanAtom.DecimalValue)
		decimal = remaining_decimal
		// TODO overload with Rune * amount_of_runes would be cleaner
		for count := uint16(1); count <= amount_of_runes; count++ {
			roman += romanAtom.Rune
		}

	}

	return
}
