package day06

import (
	"strings"
	"testing"
)

var input2 = `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +
`

var answer2 int = 3263827

func TestSolvePart2(t *testing.T) {
	tests := []struct {
		input  string
		answer int
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
