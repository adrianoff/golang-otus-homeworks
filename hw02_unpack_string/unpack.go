package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func repeatChar(builder *strings.Builder, char rune, count int) {
	for i := 0; i < count; i++ {
		builder.WriteRune(char)
	}
}

func Unpack(input string) (string, error) {
	var builder strings.Builder
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
			repeatChar(&builder, prevRune, count)
		} else if !unicode.IsDigit(nextRune) {
			builder.WriteRune(curRune)
		}
	}

	return builder.String(), nil
}
