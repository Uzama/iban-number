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

	// source https://www.iban.com/structure
	ibanLength := map[string]int{
		"AL": 28, // Albania
		"AD": 24, // Andorra
		"AT": 20, // Austria
		"AZ": 28, // Azerbaijan
		"BH": 22, // Bahrain
		"BY": 28, // Belarus
		"BE": 16, // Belgium
		"BA": 20, // Bosnia and Herzegovina
		"BR": 29, // Brazil
		"BG": 22, // Bulgaria
		"CR": 22, // Costa Rica
		"HR": 21, // Croatia
		"CY": 28, // Cyprus
		"CZ": 24, // Czech Republic
		"DK": 18, // Denmark
		"DO": 28, // Dominican Republic
		"EG": 29, // Egypt
		"SV": 28, // El Salvador
		"EE": 20, // Estonia
		"FO": 18, // Faroe Islands
		"FI": 18, // Finland
		"FR": 27, // France
		"GE": 22, // Georgia
		"DE": 22, // Germany
		"GI": 23, // Gibraltar
		"GR": 27, // Greece
		"GL": 18, // Greenland
		"GT": 28, // Guatemala
		"VA": 22, // Holy See (the)
		"HU": 28, // Hungary
		"IS": 26, // Iceland
		"IQ": 23, // Iraq
		"IE": 22, // Ireland
		"IL": 23, // Israel
		"IT": 27, // Italy
		"JO": 30, // Jordan
		"KZ": 20, // Kazakhstan
		"XK": 20, // Kosovo
		"KW": 30, // Kuwait
		"LV": 21, // Latvia
		"LB": 28, // Lebanon
		"LY": 25, // Libya
		"LI": 21, // Liechtenstein
		"LT": 20, // Lithuania
		"LU": 20, // Luxembourg
		"MT": 31, // Malta
		"MR": 27, // Mauritania
		"MU": 30, // Mauritius
		"MD": 24, // Moldova
		"MC": 27, // Monaco
		"ME": 22, // Montenegro
		"NL": 18, // Netherlands
		"MK": 19, // North Macedonia
		"NO": 15, // Norway
		"PK": 24, // Pakistan
		"PS": 29, // Palestine
		"PL": 28, // Poland
		"PT": 25, // Portugal
		"QA": 29, // Qatar
		"RO": 24, // Romania
		"LC": 32, // Saint Lucia
		"SM": 27, // San Marino
		"ST": 25, // Sao Tome and Principe
		"SA": 24, // Saudi Arabia
		"RS": 22, // Serbia
		"SC": 31, // Seychelles
		"SK": 24, // Slovak Republic
		"SI": 19, // Slovenia
		"ES": 24, // Spain
		"SD": 18, // Sudan
		"SE": 24, // Sweden
		"CH": 21, // Switzerland
		"TL": 23, // Timor-Leste
		"TN": 24, // Tunisia
		"TR": 26, // Turkey
		"UA": 29, // Ukraine
		"AE": 23, // United Arab Emirates
		"GB": 22, // United Kingdom
		"VG": 24, // Virgin Islands, British
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

	// remove spaces if found
	iban = v.removeSpaces(iban)

	// check iban length
	isMatched, err := v.checkIBANLength(iban)
	if err != nil {
		return false, err
	}

	if !isMatched {
		return false, nil
	}

	// move the four initial characters to the end of the string
	iban = v.reArrangeString(iban)

	// expand the string by replacing alphabatic with the two digit number
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

		// change to uppercase
		char = unicode.ToUpper(char)

		// check whether character is a alphabatic
		if 'A' <= char && char <= 'Z' {

			// get the two digit of the character
			digit := v.getTwoDigit(char)

			// replace the alphabatic with the two digit number
			expendedString += strconv.Itoa(digit)
			continue
		}

		expendedString += string(char)
	}

	return expendedString
}

func (v validator) getTwoDigit(char rune) int {

	// find the offset to start from 10
	offset := int('A') - 10

	return int(char) - offset
}

// get reminder for given numeric string value
func (v validator) getRemainder(str string) int {

	// use big int to process large numbers
	num, _ := new(big.Int).SetString(str, 10)

	reminder := num.Mod(num, big.NewInt(97))

	return int(reminder.Int64())
}
