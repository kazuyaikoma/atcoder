package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func replaceAtIdx(in string, r rune, idx int) string {
	out := []rune(in)
	out[idx] = r
	return string(out)
}

func solve(n int, graph []string) {
	for i := len(graph) - 2; 0 <= i; i-- {
		under := graph[i+1]
		for j := 1; j < 2*n-2; j++ {
			if []rune(under)[j-1] == 'X' || []rune(under)[j] == 'X' || []rune(under)[j+1] == 'X' {
				if []rune(graph[i])[j] == '#' {
					graph[i] = replaceAtIdx(graph[i], 'X', j)
				}
			}
		}
	}

	for _, g := range graph {
		fmt.Println(g)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	var graph []string
	for i := 0; i < n; i++ {
		scanner.Scan()
		graph = append(graph, scanner.Text())
	}
	solve(n, graph)
}
