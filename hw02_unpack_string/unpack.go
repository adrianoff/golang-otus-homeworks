package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func repeatChar(char rune, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += string(char)
	}
	return repeated
}

func Unpack(input string) (string, error) {
	result := ""
	runes := []rune(input)

	for i := 0; i < len(runes); i++ {
		curRune := runes[i]
		var nextRune rune
		var prevRune rune

		if i+1 < len(runes) {
			nextRune = runes[i+1]
		}
		if i > 0 {
			prevRune = runes[i-1]
		}

		if unicode.IsDigit(curRune) {
			if prevRune == 0 || unicode.IsDigit(prevRune) {
				return "", ErrInvalidString
			}

			count, _ := strconv.Atoi(string(curRune))
			result += repeatChar(prevRune, count)
		} else if !unicode.IsDigit(nextRune) {
			result += string(curRune)
		}
	}

	return result, nil
}
