package day03

import (
	"strings"
	"testing"
)

var input2 = `987654321111111
811111111111119
234234234234278
818181911112111
`

var answer2 int64 = 3121910778619

func TestSolvePart2(t *testing.T) {
	tests := []struct {
		input  string
		answer int64
	}{
		{input2, answer2},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := SolvePart2(r)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
