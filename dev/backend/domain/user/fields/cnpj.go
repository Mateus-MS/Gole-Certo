package fields

import (
	"errors"
	"regexp"
	"slices"
	"strconv"
)

type CNPJTag struct{}
type CNPJ = Field[CNPJTag]

var ErrInvalidCNPJ = errors.New("invalid cnpj")

// Constructor
func NewCNPJ(value string) (cnpj CNPJ, err error) {
	value = cnpj.Tag.sanitize(value)
	return NewField[CNPJTag](value, cnpj.Tag.Validate, ErrInvalidCNPJ)
}

func (c CNPJTag) Validate(value string) bool {
	// Must be 14 digits
	if len(value) != 14 {
		return false
	}

	// Reject known invalid sequences (all digits the same)
	invalids := []string{
		"00000000000000", "11111111111111", "22222222222222",
		"33333333333333", "44444444444444", "55555555555555",
		"66666666666666", "77777777777777", "88888888888888",
		"99999999999999",
	}
	if slices.Contains(invalids, value) {
		return false
	}

	// Calculate first check digit
	if !checkCNPJDigit(value, 12) {
		return false
	}

	// Calculate second check digit
	if !checkCNPJDigit(value, 13) {
		return false
	}

	return true
}

func (c CNPJTag) sanitize(value string) (clean string) {
	// Remove non-digit characters
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(value, "")
}

func checkCNPJDigit(cnpj string, pos int) bool {
	weights := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	if pos == 13 {
		weights = append([]int{6}, weights...)
	}

	sum := 0
	for i := 0; i < pos; i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		sum += num * weights[i]
	}

	r := sum % 11
	var digit int
	if r < 2 {
		digit = 0
	} else {
		digit = 11 - r
	}

	expectedDigit, _ := strconv.Atoi(string(cnpj[pos]))
	return digit == expectedDigit
}
