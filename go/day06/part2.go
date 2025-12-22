package day06

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/macgeargear/advent-of-code-go/utils"
)

func SolvePart2(r io.Reader) any {
	lines := utils.ReadLines(r)

	maxLineLength := math.MinInt

	for _, line := range lines {
		maxLineLength = max(maxLineLength, len(line))
	}

	opsIdx := []int{}
	ops := []string{}
	for i, ch := range lines[len(lines)-1] {
		if ch == '*' || ch == '+' {
			opsIdx = append(opsIdx, i)
			ops = append(ops, string(ch))
		}
	}
	opsIdx = append(opsIdx, maxLineLength+1)

	// fix pad
	for i, line := range lines {
		lines[i] = lines[i] + strings.Repeat(" ", (maxLineLength-len(line)))
	}

	numsList := [][]string{}
	for i := range len(lines) - 1 {
		numsStr := []string{}
		prev := 0
		opIdx := 1
		for j := range maxLineLength + 2 {
			if j == opsIdx[opIdx] {
				numsStr = append(numsStr, lines[i][prev:j-1])
				prev = j
				opIdx += 1
			}
		}
		numsList = append(numsList, numsStr)
	}

	numsListT := [][]string{}
	for j := range len(numsList[0]) {
		numsT := []string{}
		for i := range len(numsList) {
			numsT = append(numsT, numsList[i][j])
		}
		numsListT = append(numsListT, numsT)
	}

	ans := 0

	for idx, nums := range numsListT {
		for _, num := range nums {
			fmt.Printf("[%s]\n", num)
		}

		op := ops[idx]
		var total int
		if op == "*" {
			total = 1
		} else {
			total = 0
		}
		for j := range len(nums[0]) {
			cur := 0
			exp := 0.0
			for i := len(nums) - 1; i >= 0; i-- {
				num, _ := strconv.Atoi(string(nums[i][j]))
				cur += int(math.Pow(10, exp)) * num
				if num != 0 {
					exp += 1
				}
			}
			if op == "*" {
				total *= cur
			} else {
				total += cur
			}
			fmt.Println(cur)
		}
		ans += total
		fmt.Println("total: ", total)
		fmt.Println("---------")
	}

	return ans
}
