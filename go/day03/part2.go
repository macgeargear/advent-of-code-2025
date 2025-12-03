package day03

import (
	"fmt"
	"io"
	"strconv"

	"github.com/macgeargear/advent-of-code-go/utils"
)

func SolvePart2(r io.Reader) any {
	fmt.Println("Running day 3 part 1...")
	lines := utils.ReadLines(r)
	var total int64 = 0

	for _, line := range lines {
		num := getMaxNumFromLine(line, 12)
		fmt.Println(num)
		total += num
	}

	return total
}

func getMaxNumFromLine(line string, k int) int64 {
	st := make([]byte, 0, k)
	n := len(line)
	for i := range n {
		cur := line[i]
		if len(st) > 0 && st[len(st)-1] < cur && len(st)-1+(n-i) >= k {
			st = st[:len(st)-1] // pop
		}
		if len(st) < k {
			st = append(st, cur)
		}
	}
	num, _ := strconv.Atoi(string(st))
	return int64(num)
}
