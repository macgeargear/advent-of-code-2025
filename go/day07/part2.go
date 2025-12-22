package day07

import (
	"io"

	"github.com/macgeargear/advent-of-code-go/utils"
)

func traverse(lines []string, N, M int, cur pos, memo, visited map[pos]int) int {
	if cur.r < 0 || cur.c < 0 || cur.c >= M {
		return 0
	}
	if cnt, ok := memo[cur]; ok {
		return cnt
	}
	if visited[cur] == 1 {
		return 0
	}
	if cur.r == N-1 {
		return 1
	}

	visited[cur] = 1
	defer func() { visited[cur] = 0 }()

	cnt := 0

	if lines[cur.r][cur.c] == '^' {
		cnt += traverse(lines, N, M, pos{cur.r, cur.c - 1}, memo, visited)
		cnt += traverse(lines, N, M, pos{cur.r, cur.c + 1}, memo, visited)
	} else {
		cnt += traverse(lines, N, M, pos{cur.r + 1, cur.c}, memo, visited)
	}

	memo[cur] = cnt
	return cnt
}

func SolvePart2(r io.Reader) any {
	lines := utils.ReadLines(r)

	byteLines := make([][]byte, len(lines))
	for i, line := range lines {
		byteLines[i] = []byte(line)
	}

	start := 0
	for i, ch := range lines[0] {
		if ch == 'S' {
			start = i
			break
		}
	}

	N := len(lines)
	M := len(lines[0])

	memo := map[pos]int{}
	visited := map[pos]int{}

	return traverse(lines, N, M, pos{0, start}, memo, visited)
}
