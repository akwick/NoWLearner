package nouns

import (
	"errors"
	"strings"
)

type Noun struct {
	*noun
}

type noun struct {
	defSig string
	gender string
}

func NewNoun(defSig, gender string) Noun {
	return Noun{&noun{defSig, gender}}
}

var invGender = "invalid gender"

// CheckDefinitePlural checks whether the plural of noun with gender is build correctly.
func (n Noun) CheckDefinite(definite string) (bool, error) {
	s, err := getDefinite(n.defSig, n.gender)
	if err != nil {
		return false, err
	}
	return definite == s, nil
}

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
