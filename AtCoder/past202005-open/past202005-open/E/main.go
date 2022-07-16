package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N, M, Q int, colors []int, uvs, ss [][]int) {
	// u, vの隣接リスト
	adjList := map[int][]int{}
	for _, uv := range uvs {
		// 0-indexedなので-1
		u, v := uv[0]-1, uv[1]-1
		if _, ok := adjList[u]; !ok {
			adjList[u] = []int{v}
		} else {
			adjList[u] = append(adjList[u], v)
		}

		if _, ok := adjList[v]; !ok {
			adjList[v] = []int{u}
		} else {
			adjList[v] = append(adjList[v], u)
		}
	}

	// クエリ処理
	for _, s := range ss {
		// 0-indexedなので-1tes
		x := s[1] - 1
		fmt.Println(colors[x])

		if s[0] == 1 {
			adjIndices := adjList[x]
			for _, adj := range adjIndices {
				colors[adj] = colors[x]
			}
		} else if s[0] == 2 {
			y := s[2]
			colors[x] = y
		}
	}
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
	var M int
	scanner.Scan()
	M, _ = strconv.Atoi(scanner.Text())
	var Q int
	scanner.Scan()
	Q, _ = strconv.Atoi(scanner.Text())

	var uvs [][]int
	for i := 0; i < M; i++ {
		scanner.Scan()
		u, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		uvs = append(uvs, []int{u, v})
	}

	var colors []int
	for i := 0; i < N; i++ {
		scanner.Scan()
		c, _ := strconv.Atoi(scanner.Text())
		colors = append(colors, c)
	}

	var ss [][]int
	for i := 0; i < Q; i++ {
		scanner.Scan()
		s1, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		s2, _ := strconv.Atoi(scanner.Text())
		s := []int{s1, s2}
		if s1 == 2 {
			scanner.Scan()
			s3, _ := strconv.Atoi(scanner.Text())
			s = append(s, s3)
		}
		ss = append(ss, s)
	}

	solve(N, M, Q, colors, uvs, ss)
}
