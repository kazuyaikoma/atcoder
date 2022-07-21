package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

var flag bool

func dfs(H, W int, city [][]bool, cur, goal []int) {
	if cur[0] == goal[0] && cur[1] == goal[1] {
		flag = true
	}
	deltas := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	for _, delta := range deltas {
		nx, ny := cur[0]+delta[0], cur[1]+delta[1]
		if 0 <= nx && nx < H && 0 <= ny && ny < W && city[nx][ny] {
			city[nx][ny] = false
			dfs(H, W, city, []int{nx, ny}, goal)
		}
	}
}

func solve(H, W int, ss []string) {
	var city [][]bool
	for i := 0; i < H; i++ {
		var row []bool
		for j := 0; j < W; j++ {
			row = append(row, false)
		}
		city = append(city, row)
	}

	var start, goal []int
	for i, s := range ss {
		str := s
		for j, c := range str {
			switch c {
			case 's':
				start = []int{i, j}
				city[i][j] = true
			case 'g':
				goal = []int{i, j}
				city[i][j] = true
			case '.':
				city[i][j] = true
			}
		}
	}

	dfs(H, W, city, start, goal)
	if flag {
		fmt.Println(YES)
	} else {
		fmt.Println(NO)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var H int
	scanner.Scan()
	H, _ = strconv.Atoi(scanner.Text())
	var W int
	scanner.Scan()
	W, _ = strconv.Atoi(scanner.Text())
	var ss []string
	for i := 0; i < H; i++ {
		scanner.Scan()
		s := scanner.Text()
		ss = append(ss, s)
	}
	solve(H, W, ss)
}
