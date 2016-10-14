package nouns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Examples: https://www.ntnu.edu/now/2/grammar#nouns
func TestDefinedVariables(t *testing.T) {
	var nouns = []struct {
		noun     string
		gender   string
		expected string
	}{
		{"brus", "m", "brusen"},
		{"avis", "f", "avisa"},
		{"tog", "n", "toget"},
		{"pose", "m", "posen"},
		{"jente", "f", "jenta"},
		{"frimerke", "n", "frimerket"},
	}

	for _, td := range nouns {
		plrl, err := getDefinite(td.noun, td.gender)
		if assert.Nil(t, err) {
			assert.Equal(t, plrl, td.expected)
		}
	}
}
