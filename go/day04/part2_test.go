package day04

import (
	"strings"
	"testing"
)

var input2 = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`

var answer2 int = 43

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
