package nouns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckDefinitePlural(t *testing.T) {
	var nouns = []struct {
		noun     string
		gender   string
		plural   string
		expected bool
	}{
		{"brus", "m", "brusen", true},
		{"avis", "f", "avisa", true},
		{"tog", "n", "toget", true},
		{"pose", "m", "posen", true},
		{"jente", "f", "jenta", true},
		{"frimerke", "n", "frimerket", true},
		{"brus", "m", "brusa", false},
		{"avis", "f", "aviset", false},
		{"tog", "n", "togen", false},
		{"pose", "m", "posa", false},
		{"jente", "f", "jentet", false},
		{"frimerke", "n", "frimerken", false},
	}

	for _, td := range nouns {
		eq := CheckDefinitePlural(td.noun, td.gender, td.plural)
		assert.Equal(t, eq, td.expected)
	}
}

// Examples: https://www.ntnu.edu/now/2/grammar#nouns
func TestDefinedVariables(t *testing.T) {
	// Test correct input
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
			assert.Equal(t, td.expected, plrl)
		}
	}

	// Test invalid gender
	plrl, err := getDefinite("brus", "M")
	assert.Equal(t, "", plrl)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), invGender)
}
