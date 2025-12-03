package day02

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/macgeargear/advent-of-code-go/utils"
)

func SolvePart1(r io.Reader) any {
	fmt.Println("Running day 2 part 1...")
	lines := utils.ReadLines(r)
	var total int64 = 0

	for _, line := range lines {
		ranges := strings.SplitSeq(line, ",")
		for rangeStr := range ranges {
			parts := strings.Split(rangeStr, "-")
			l, _ := strconv.Atoi(parts[0])
			r, _ := strconv.Atoi(parts[1])
			for i := l; i <= r; i++ {
				numStr := strconv.Itoa(i)
				if len(numStr)%2 != 0 {
					continue
				}
				m := len(numStr) / 2
				ll, rr := numStr[:m], numStr[m:]
				if ll == rr {
					num, _ := strconv.Atoi(numStr)
					fmt.Println(num)
					total += int64(num)
				}
			}

		}
	}

	return int64(total)
}
