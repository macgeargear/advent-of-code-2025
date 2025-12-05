package day05

import (
	"io"
	"sort"

	"github.com/macgeargear/advent-of-code-go/utils"
)

func SolvePart2(r io.Reader) any {
	lines := utils.ReadLines(r)

	ranges, _ := splitInput(lines)
	ingRanges := getIngredientRanges(ranges)

	sort.Slice(ingRanges, func(i, j int) bool {
		return ingRanges[i].l < ingRanges[j].l
	})

	merged := mergeRanges(ingRanges)

	var total int64 = 0
	for _, r := range merged {
		total += r.r - r.l + 1
	}

	return total
}

func mergeRanges(ingRanges []db) []db {
	merged := []db{ingRanges[0]}

	for i := 1; i < len(ingRanges); i++ {
		cur := ingRanges[i]
		prev := &merged[len(merged)-1]

		if prev.r+1 >= cur.l {
			if cur.r > prev.r {
				prev.r = cur.r
			}
		} else {
			merged = append(merged, cur)
		}
	}

	return merged
}
