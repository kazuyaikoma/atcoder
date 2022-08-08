package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(N, L int, x, T []int64) {
	t1, t2, t3 := float64(T[0]), float64(T[1]), float64(T[2])
	timeDp := make([]float64, L+1)
	for i, _ := range timeDp {
		timeDp[i] = math.Pow(10, 10)
	}
	timeDp[0] = 0

	hurdles := make([]bool, L+1)
	for i, _ := range hurdles {
		hurdles[i] = false
	}
	for _, b := range x {
		hurdles[b] = true
	}

	for i := 1; i <= L; i++ {
		timeDp[i] = math.Min(timeDp[i], timeDp[i-1]+t1)
		if i >= 2 {
			timeDp[i] = math.Min(timeDp[i], timeDp[i-2]+t1+t2)
		}
		if i >= 4 {
			timeDp[i] = math.Min(timeDp[i], timeDp[i-4]+t1+t2*3)
		}
		if hurdles[i] {
			timeDp[i] += t3
		}
	}
	ans := timeDp[L]

	if L-1 >= 0 {
		ans = math.Min(ans, timeDp[L-1]+(t1+t2)/2)
	}
	if L-2 >= 0 {
		ans = math.Min(ans, timeDp[L-2]+(t1+t2)/2+t2)
	}
	if L-3 >= 0 {
		ans = math.Min(ans, timeDp[L-3]+(t1+t2)/2+t2*2)
	}
	fmt.Println(int(ans))
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
	var L int
	scanner.Scan()
	L, _ = strconv.Atoi(scanner.Text())
	x := make([]int64, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		x[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	T := make([]int64, 3)
	for i := 0; i < 3; i++ {
		scanner.Scan()
		T[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, L, x, T)
}
