package day03

import (
	"strings"
	"testing"
)

var input1 = `987654321111111
811111111111119
234234234234278
818181911112111
`

var answer1 int64 = 357

func TestSolvePart1(t *testing.T) {
	tests := []struct {
		input  string
		answer int64
	}{
		{input1, answer1},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := SolvePart1(r)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
