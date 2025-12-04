package day01

import (
	"fmt"
	"io"
	"strconv"

	"github.com/macgeargear/advent-of-code-go/utils"
)

func SolvePart1(r io.Reader) any {
	lines := utils.ReadLines(r)
	fmt.Println(lines)

	pos := 50
	cnt := 0
	for _, line := range lines {
		dir := line[:1]
		dist, _ := strconv.Atoi(line[1:])
		switch dir {
		case "L":
			pos = ((pos-dist)%100 + 100) % 100
		case "R":
			pos = (pos + dist) % 100
		}
		if pos == 0 {
			cnt++
		}
	}

	return cnt
}
