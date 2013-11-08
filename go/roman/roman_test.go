package roman

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnRomanRepresentationOf1(t *testing.T) {
	assert.Equal(t, asRoman(1), "I")
}

func TestShouldReturnRomanRepresentationOf2(t *testing.T) {
	assert.Equal(t, asRoman(2), "II")
}

func TestShouldReturnRomanRepresentationOf3(t *testing.T) {
	assert.Equal(t, asRoman(3), "III")
}

func TestShouldReturnRomanRepresentationOf4(t *testing.T) {
	assert.Equal(t, asRoman(4), "IV")
}

func TestShouldReturnRomanRepresentationOf5(t *testing.T) {
	assert.Equal(t, asRoman(5), "V")
}

func TestShouldReturnRomanRepresentationOf6(t *testing.T) {
	assert.Equal(t, asRoman(6), "VI")
}

func TestShouldReturnRomanRepresentationOf9(t *testing.T) {
	assert.Equal(t, asRoman(9), "IX")
}

func TestShouldReturnRomanRepresentationOf10(t *testing.T) {
	assert.Equal(t, asRoman(10), "X")
}

func TestShouldReturnRomanRepresentationOf11(t *testing.T) {
	assert.Equal(t, asRoman(11), "XI")
}

func TestShouldReturnRomanRepresentationOf14(t *testing.T) {
	assert.Equal(t, asRoman(14), "XIV")
}

func TestShouldReturnRomanRepresentationOf20(t *testing.T) {
	assert.Equal(t, asRoman(20), "XX")
}

func TestShouldReturnRomanRepresentationOf40(t *testing.T) {
	assert.Equal(t, asRoman(40), "XL")
}

func TestShouldReturnRomanRepresentationOf50(t *testing.T) {
	assert.Equal(t, asRoman(50), "L")
}

func TestShouldReturnRomanRepresentationOf60(t *testing.T) {
	assert.Equal(t, asRoman(60), "LX")
}

func TestShouldReturnRomanRepresentationOf90(t *testing.T) {
	assert.Equal(t, asRoman(90), "XC")
}

func TestShouldReturnRomanRepresentationOf100(t *testing.T) {
	assert.Equal(t, asRoman(100), "C")
}

func TestShouldReturnRomanRepresentationOf400(t *testing.T) {
	assert.Equal(t, asRoman(400), "CD")
}

func TestShouldReturnRomanRepresentationOf500(t *testing.T) {
	assert.Equal(t, asRoman(500), "D")
}

func TestShouldReturnRomanRepresentationOf900(t *testing.T) {
	assert.Equal(t, asRoman(900), "CM")
}

func TestShouldReturnRomanRepresentationOf1000(t *testing.T) {
	assert.Equal(t, asRoman(1000), "M")
}
