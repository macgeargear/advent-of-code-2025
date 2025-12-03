package day02

import (
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/macgeargear/advent-of-code-go/utils"
)

func SolvePart2(r io.Reader) any {
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
				if checkNum([]rune(numStr), len(numStr)-1) {
					total += int64(i)
				}
			}
		}
	}

	return total
}

func checkNum(numStr []rune, n int) bool {
	for i := 1; i <= n; i++ {
		prev := -1
		found := true
		if i > 1 && len(numStr)%i != 0 {
			continue
		}
		for chunk := range slices.Chunk(numStr, i) {
			curNum, _ := strconv.Atoi(string(chunk))
			if prev == -1 {
				prev = curNum
				continue
			} else {
				found = false
				break
			}

		}
		if found {
			return true
		}
	}
	return false
}
