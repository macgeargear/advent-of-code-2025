package day04

import (
	"fmt"
	"io"

	"github.com/macgeargear/advent-of-code-go/utils"
)

func SolvePart1(r io.Reader) any {
	lines := utils.ReadLines(r)
	var total int32 = 0

	R := len(lines)
	C := len(lines[0])

	for r, line := range lines {
		for c, ch := range line {
			if ch == '.' {
				continue
			}
			roll := 0
			for _, dr := range []int{-1, 0, 1} {
				for _, dc := range []int{-1, 0, 1} {
					if dr == 0 && dc == 0 {
						continue
					}
					nr := r + dr
					nc := c + dc
					if !valid(nr, nc, R, C) {
						continue
					}
					if lines[nr][nc] == '@' {
						roll++
					}
				}
			}
			if roll < 4 {
				fmt.Println(r, c, roll)
				total++
			}
		}
	}

	return int64(total)
}
