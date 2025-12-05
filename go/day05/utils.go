package day05

import (
	"strconv"
	"strings"
)

type db struct{ l, r int64 }

func getIngredientRanges(lines []string) []db {
	dbs := []db{}
	for _, line := range lines {
		parts := strings.Split(line, "-")
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		dbs = append(dbs, db{int64(l), int64(r)})
	}
	return dbs
}

func splitInput(lines []string) ([]string, []int64) {
	ranges := []string{}
	ings := []int64{}
	idx := 0
	for i, line := range lines {
		if line == "" {
			idx = i
			break
		}
		ranges = append(ranges, lines[i])
	}
	for i := idx + 1; i < len(lines); i++ {
		num, _ := strconv.Atoi(lines[i])
		ings = append(ings, int64(num))
	}

	return ranges, ings
}
