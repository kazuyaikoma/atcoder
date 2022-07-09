package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(N int64) {
	digit := math.Ceil(float64(N) / 9)
	num := N % 9
	if num == 0 {
		num = 9
	}
	var r []rune
	for n := 0; n < int(digit); n++ {
		r = append(r, []rune(strconv.Itoa(int(num)))...)
	}
	fmt.Println(string(r))
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
	solve(N)
}
