package day04

import (
	"io"

	"github.com/macgeargear/advent-of-code-go/utils"
)

type pos struct{ r, c int }

func countRoll(R, C int, grid [][]rune) ([][]rune, int) {
	total := 0
	positions := []pos{}

	for r, line := range grid {
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
					if grid[nr][nc] == '@' {
						roll++
					}
				}
			}
			if roll < 4 {
				positions = append(positions, pos{r, c})
				total++
			}
		}
	}
	for _, p := range positions {
		grid[p.r][p.c] = '.'
	}

	return grid, total
}

func SolvePart2(r io.Reader) any {
	lines := utils.ReadLines(r)
	var total = 0

	R := len(lines)
	C := len(lines[0])

	grid := make([][]rune, R)
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	for {
		roll := 0
		grid, roll = countRoll(R, C, grid)
		if roll == 0 {
			break
		}
		total += roll
	}

	return total
}
