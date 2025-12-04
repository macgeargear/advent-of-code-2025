package day03

import (
	"io"

	"github.com/macgeargear/advent-of-code-go/utils"
)

func SolvePart1(r io.Reader) any {
	lines := utils.ReadLines(r)
	var total int32 = 0

	for _, line := range lines {
		var num1 int32 = -1
		var num2 int32 = -1
		num1Idx := 0

		for i, ch := range line {
			num := int32(ch - '0')
			if num1 == -1 {
				num1 = num
			}
			if num1 != -1 && i < len(line)-1 {
				if num > num1 {
					num1 = num
					num1Idx = i
				}
			}
		}
		for i := num1Idx + 1; i < len(line); i++ {
			num := int32(line[i] - '0')
			num2 = max(num2, num)
		}
		total += num1*10 + num2
	}

	return int64(total)
}
