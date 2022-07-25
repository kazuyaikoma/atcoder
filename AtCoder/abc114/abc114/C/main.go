package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var cnt int
var N int

func dfs(num int, use3, use5, use7 bool) {
	if num > N {
		return
	}
	if use3 && use5 && use7 {
		cnt++
	}

	d := float64(len(strconv.Itoa(num)))
	digit := int(math.Pow(10, d))

	dfs(int(3*digit+num), true, use5, use7)
	dfs(int(5*digit+num), use3, true, use7)
	dfs(int(7*digit+num), use3, use5, true)
}

func solve(N int) {
	dfs(3, true, false, false)
	dfs(5, false, true, false)
	dfs(7, false, false, true)
	fmt.Println(cnt)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	solve(N)
}
