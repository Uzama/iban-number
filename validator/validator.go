package validator

import (
	"errors"
	"math/big"
	"strconv"
	"strings"
	"unicode"
)

type validator struct {
	// country code -> IBAN length
	ibanLength map[string]int
}

func NewValidator() Validator {
	ibanLength := map[string]int{
		"GB": 22, // United Kingdom
	}

	return &validator{
		ibanLength: ibanLength,
	}
}

/*
1. Check that the total IBAN length is correct as per the country. If not, the IBAN is invalid
2. Move the four initial characters to the end of the string
3. Replace each letter in the string with two digits, thereby expanding the string, where A = 10, B = 11, ..., Z = 35
4. Interpret the string as a decimal integer and compute the remainder of that number on division by 97

If the remainder is 1, the check digit test is passed and the IBAN might be valid.
*/
func (v validator) Validate(iban string) (bool, error) {

	iban = v.removeSpaces(iban)

	isMatched, err := v.checkIBANLength(iban)
	if err != nil {
		return false, err
	}

	if !isMatched {
		return false, nil
	}

	iban = v.reArrangeString(iban)

	iban = v.expandString(iban)

	reminder := v.getRemainder(iban)

	isValid := reminder == 1

	return isValid, nil
}

func (v validator) removeSpaces(iban string) string {

	splitted := strings.Split(iban, " ")

	joined := strings.Join(splitted, "")

	return joined
}

func (v validator) checkIBANLength(iban string) (bool, error) {

	countryCode := iban[0:2]
	ibanLength, ok := v.ibanLength[countryCode]
	if !ok {
		return false, errors.New("information not exists for given country code")
	}

	isMatched := ibanLength == len(iban)

	return isMatched, nil
}

func (v validator) reArrangeString(iban string) string {

	temp := iban[4:]

	temp += iban[0:4]

	return temp
}

func (v validator) expandString(iban string) string {

	expendedString := ""

	for _, char := range iban {

		char = unicode.ToUpper(char)

		if 'A' <= char && char <= 'Z' {

			digit := v.getTwoDigit(char)

			expendedString += strconv.Itoa(digit)
			continue
		}

		expendedString += string(char)
	}

	return expendedString
}

func (v validator) getTwoDigit(char rune) int {

	offset := int('A') - 10

	return int(char) - offset
}

func (v validator) getRemainder(str string) int {

	num, _ := new(big.Int).SetString(str, 10)

	reminder := num.Mod(num, big.NewInt(97))

	return int(reminder.Int64())
}
