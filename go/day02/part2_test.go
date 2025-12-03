package day02

import (
	"strings"
	"testing"
)

var input2 = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

// var input2 = `11-22,95-115,998-1012`

var answer2 int64 = 4174379265

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
