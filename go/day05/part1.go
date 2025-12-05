package day05

import (
	"io"
	"sort"

	"github.com/macgeargear/advent-of-code-go/utils"
)

func SolvePart1(r io.Reader) any {
	lines := utils.ReadLines(r)
	var total int = 0

	ranges, ings := splitInput(lines)
	ingRanges := getIngredientRanges(ranges)

	sort.Slice(ingRanges, func(i, j int) bool {
		return ingRanges[i].l < ingRanges[j].l
	})

	for _, ing := range ings {
		if isInAnyRange(ing, ingRanges) {
			total++
		}
	}

	return total
}

func isInAnyRange(val int64, ranges []db) bool {
	for _, r := range ranges {
		if val < r.l {
			return false
		}
		if val >= r.l && val <= r.r {
			return true
		}
	}
	return false
}
