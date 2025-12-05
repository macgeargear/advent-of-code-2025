package day05

import (
	"strings"
	"testing"
)

var input1 = `3-5
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

var answer1 = 3

func TestSolvePart1(t *testing.T) {
	tests := []struct {
		input  string
		answer int
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
