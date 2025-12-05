package day05

import (
	"strings"
	"testing"
)

var input2 = `3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

var answer2 int = 14

func TestSolvePart2(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{input1, answer2},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := SolvePart2(r)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
