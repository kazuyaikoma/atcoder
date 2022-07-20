package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N, X, Y int, xs, ys []int) {
	const margin = 201
	// 壁となる2次元配列 (403 x 403)
	// 問題制約として"無限に広がる二次元グリッド"とあるので、201 x 201の二次元空間をはみ出して別のマスに移動が可能。そのため403 x 403。
	// true: 通路, false: 壁
	var walls [][]bool
	for i := 0; i < margin*2+1; i++ {
		var blk []bool
		for j := 0; j < margin*2+1; j++ {
			blk = append(blk, true)
		}
		walls = append(walls, blk)
	}
	// 問題から与えられた壁の反映
	walls[0+margin][0+margin] = false
	for i := 0; i < N; i++ {
		walls[xs[i]+margin][ys[i]+margin] = false
	}

	// index 0: x
	// index 1: y
	// index 2: cnt
	q := [][]int{{0, 0, 0}}
	deltas := [][]int{
		{1, 1},
		{0, 1},
		{-1, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}
	for len(q) > 0 {
		x, y, cnt := q[0][0], q[0][1], q[0][2]
		q = q[1:]
		if x == X && y == Y {
			fmt.Println(cnt)
			return
		}
		for _, delta := range deltas {
			nx, ny, nCnt := x+delta[0], y+delta[1], cnt+1
			if -201 <= nx && nx <= 201 && -201 <= ny && ny <= 201 && walls[nx+margin][ny+margin] {
				q = append(q, []int{nx, ny, nCnt})
				walls[nx+margin][ny+margin] = false
			}
		}
	}
	fmt.Println(-1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	var X int
	scanner.Scan()
	X, _ = strconv.Atoi(scanner.Text())
	var Y int
	scanner.Scan()
	Y, _ = strconv.Atoi(scanner.Text())
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		x[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		y[i], _ = strconv.Atoi(scanner.Text())
	}
	solve(N, X, Y, x, y)
}
