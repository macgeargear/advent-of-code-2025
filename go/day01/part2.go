package day01

import (
	"fmt"
	"io"
	"strconv"

	"github.com/macgeargear/advent-of-code-go/utils"
)

func SolvePart2(r io.Reader) any {
	fmt.Println("Running day 1 part 2...")
	lines := utils.ReadLines(r)
	fmt.Println(lines)

	pos := 50
	cnt := 0

	for _, line := range lines {
		dir := line[:1]
		dist, _ := strconv.Atoi(line[1:])

		cnt += dist / 100
		rem := dist % 100
		for range rem {
			if dir == "L" {
				pos--
				if pos < 0 {
					pos = 99
				}
				if pos == 0 {
					cnt++
				}
			} else {
				pos++
				if pos > 99 {
					pos = 0
				}
				if pos == 0 {
					cnt++
				}
			}
		}
	}

	return cnt
}
