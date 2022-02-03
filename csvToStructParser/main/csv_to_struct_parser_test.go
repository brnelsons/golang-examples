package main_test

import (
	"Examples/csvToStructParser/main"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseCsvUsingHeaders(t *testing.T) {
	err := main.ParseCsvUsingHeaders(
		strings.NewReader("Name,Phone Number,Age\n\"Stev\nen\",\"192,890,1234\",45"),
		func(p main.Person) {
			assert.Equal(t,
				main.Person{
					Name:        "Stev\nen",
					PhoneNumber: "192,890,1234",
					Age:         45,
				},
				p)
		},
	)
	assert.Nil(t, err)
}
