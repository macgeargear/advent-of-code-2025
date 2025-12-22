package day06

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/macgeargear/advent-of-code-go/utils"
)

func SolvePart1(r io.Reader) any {
	lines := utils.ReadLines(r)

	numsLists := [][]int{}
	ops := []string{}
	for op := range strings.SplitSeq(lines[len(lines)-1], " ") {
		if op == "" {
			continue
		}
		ops = append(ops, op)
	}
	for _, line := range lines[:len(lines)-1] {
		nums := []int{}
		for numStr := range strings.SplitSeq(line, " ") {
			if numStr == "" {
				continue
			}
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}
		numsLists = append(numsLists, nums)
	}

	fmt.Println(numsLists)
	fmt.Println(ops)

	total := 0
	for j := 0; j < len(numsLists[0]); j++ {
		cur := 0
		if ops[j] == "*" {
			cur = 1
		}

		for i := 0; i < len(numsLists); i++ {
			if ops[j] == "*" {
				cur *= numsLists[i][j]
			} else {
				cur += numsLists[i][j]
			}
		}

		fmt.Println(cur)
		total += cur
	}

	return total
}
