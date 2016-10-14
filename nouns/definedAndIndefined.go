package nouns

import (
	"errors"
	"strings"
)

var invGender = "invalid gender"

// getDefinite return the definite of a singular word.
// If the gender string is invalid (m, f, n), then an error is returned.
// More information about the grammar: https://www.ntnu.edu/now/2/grammar#nouns
func getDefinite(noun string, gender string) (string, error) {
	switch gender {
	case "m":
		if endsWithE(noun) {
			return noun + "n", nil
		} else {
			return noun + "en", nil
		}
	case "f":
		if endsWithE(noun) {
			return strings.TrimSuffix(noun, "e") + "a", nil
		} else {
			return noun + "a", nil
		}
	case "n":
		if endsWithE(noun) {
			return noun + "t", nil
		} else {
			return noun + "et", nil
		}
	}
	return "", errors.New(invGender)
}

func endsWithE(noun string) bool {
	if strings.HasSuffix(noun, "e") {
		return true
	} else {
		return false
	}
}
