package day07

import (
	"fmt"
	"io"

	"github.com/macgeargear/advent-of-code-go/utils"
)

type pos struct {
	r, c int
}

func SolvePart1(r io.Reader) any {
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

	visited := map[string]bool{}
	visited[fmt.Sprintf("%d%d", 0, start)] = true

	queue := []pos{}

	cnt := 0

	queue = append(queue, pos{r: 0, c: start})

	for len(queue) > 0 {
		cur := queue[0]
		if cur.r < 0 || cur.r >= N || cur.c < 0 || cur.c >= M {
			break
		}
		if byteLines[cur.r][cur.c] != '^' {
			byteLines[cur.r][cur.c] = '|'
		}
		queue = queue[1:]
		for {
			cur.r += 1
			if cur.r >= N {
				break
			}
			if byteLines[cur.r][cur.c] != '^' {
				byteLines[cur.r][cur.c] = '|'
			}

			key := fmt.Sprintf("%d,%d", cur.r, cur.c)
			if visited[key] {
				break
			}
			visited[key] = true

			if lines[cur.r][cur.c] == '^' {
				byteLines[cur.r][cur.c] = '^'
				cnt += 1
				queue = append(queue, pos{r: cur.r, c: cur.c - 1})
				queue = append(queue, pos{r: cur.r, c: cur.c + 1})
				visited[fmt.Sprintf("%d,%d", cur.r, cur.c-1)] = true
				visited[fmt.Sprintf("%d,%d", cur.r, cur.c+1)] = true
				break
			}
		}
	}

	byteLines[0][start] = 'S'
	for _, line := range byteLines {
		fmt.Println(string(line))
	}

	return cnt
}
