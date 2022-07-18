package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(R, C, sy, sx, gy, gx int, cs []string) {
	// init maze
	var maze [][]bool

	for i := 0; i <= R; i++ {
		var arr []bool
		for j := 0; j <= C; j++ {
			arr = append(arr, false)
		}
		maze = append(maze, arr)
	}

	for i, c := range cs {
		for j, s := range c {
			if s == '.' {
				maze[i+1][j+1] = true
			}
		}
	}

	// index 0: y
	// index 1: x
	// index 2: cnt
	q := [][]int{{sy, sx, 0}}

	deltas := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		y, x, cnt := n[0], n[1], n[2]

		if y == gy && x == gx {
			fmt.Println(cnt)
			return
		}

		for _, d := range deltas {
			ny, nx, nCnt := y+d[0], x+d[1], cnt+1
			if maze[ny][nx] {
				maze[ny][nx] = false
				q = append(q, []int{ny, nx, nCnt})
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var R int
	scanner.Scan()
	R, _ = strconv.Atoi(scanner.Text())
	var C int
	scanner.Scan()
	C, _ = strconv.Atoi(scanner.Text())
	var sy int
	scanner.Scan()
	sy, _ = strconv.Atoi(scanner.Text())
	var sx int
	scanner.Scan()
	sx, _ = strconv.Atoi(scanner.Text())
	var gy int
	scanner.Scan()
	gy, _ = strconv.Atoi(scanner.Text())
	var gx int
	scanner.Scan()
	gx, _ = strconv.Atoi(scanner.Text())

	var cs []string
	for i := 0; i < R; i++ {
		var c string
		scanner.Scan()
		c = scanner.Text()
		cs = append(cs, c)
	}

	solve(R, C, sy, sx, gy, gx, cs)
}
