package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

// Unpack is unpack string...
func Unpack(str string) (string, error) {
	// if string start with digit
	r, _ := utf8.DecodeRuneInString(str)
	if unicode.IsDigit(r) {
		return "", ErrInvalidString
	}

	var strBuilder strings.Builder
	strR := []rune(str)

	for i := 0; i < len(strR); i++ {
		currentRune := strR[i]
		// if end rune
		if i == len(strR)-1 {
			strBuilder.WriteRune(currentRune)
			return strBuilder.String(), nil
		}

		rightRune := strR[i+1]

		if unicode.IsDigit(currentRune) && unicode.IsDigit(rightRune) {
			// if two rune is digit
			return "", ErrInvalidString
		}
		if unicode.IsDigit(currentRune) {
			count, err := strconv.Atoi(string(currentRune))
			if err != nil {
				return "", ErrInvalidString
			}
			if currentRune == '0' {
				strTempR := []rune(strBuilder.String())
				strBuilder.Reset()
				strTemp := string(strTempR[:len(strTempR)-1])
				strBuilder.WriteString(strTemp)
				continue
			}
			leftRune := strR[i-1]
			strBuilder.WriteString(strings.Repeat(string(leftRune), count-1))
		} else {
			strBuilder.WriteRune(currentRune)
		}
	}

	return strBuilder.String(), nil
}
