package fields

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

type CPFTag struct{}
type CPF = Field[CPFTag]

var ErrInvalidCPF = errors.New("invalid cpf")

// Constructor
func NewCPF(value string) (cpf CPF, err error) {
	return NewField[CPFTag](value, cpf.Tag.Validate, ErrInvalidCPF)
}

func (c CPFTag) Validate(value string) bool {
	// Remove non-digit characters
	value = strings.ReplaceAll(value, ".", "")
	value = strings.ReplaceAll(value, "-", "")

	// Must be 11 digits
	if len(value) != 11 {
		return false
	}

	// Invalid known value (all digits the same)
	if slices.Contains([]string{
		"00000000000", "11111111111", "22222222222",
		"33333333333", "44444444444", "55555555555",
		"66666666666", "77777777777", "88888888888",
		"99999999999",
	}, value) {
		return false
	}

	// Calculate first digit
	sum := 0
	for i := range 9 {
		num, _ := strconv.Atoi(string(value[i]))
		sum += num * (10 - i)
	}
	firstDigit := 11 - (sum % 11)
	if firstDigit >= 10 {
		firstDigit = 0
	}
	if firstDigit != int(value[9]-'0') {
		return false
	}

	// Calculate second digit
	sum = 0
	for i := range 10 {
		num, _ := strconv.Atoi(string(value[i]))
		sum += num * (11 - i)
	}
	secondDigit := 11 - (sum % 11)
	if secondDigit >= 10 {
		secondDigit = 0
	}
	if secondDigit != int(value[10]-'0') {
		return false
	}

	return true
}
