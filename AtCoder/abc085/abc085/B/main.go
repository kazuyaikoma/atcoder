package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N int64, d []int64) {
	sort.Slice(d, func(i, j int) bool { return d[i] < d[j] })
	cur := d[0]
	cnt := 1
	for _, n := range d {
		if n == cur {
			continue
		}
		cur = n
		cnt++
	}
	fmt.Println(cnt)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	d := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		d[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, d)
}
