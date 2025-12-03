package day01

import (
	"strings"
	"testing"
)

var testInput1 = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func TestSolvePart1(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput1, 3},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := SolvePart1(r)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
