package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func min(arr []int) int {
	m := arr[0]
	for _, v := range arr {
		if v < m {
			m = v
		}
	}
	return m
}

func solve(N, L int, x, T []int64) {
	t1, t2, t3 := int(T[0]), int(T[1]), int(T[2])
	//timeDp := make([]float64, L+1)
	//for i, _ := range timeDp {
	//	timeDp[i] = math.Pow(10, 10)
	//}
	//timeDp[0] = 0
	//
	hurdles := make([]bool, L+1)
	for i, _ := range hurdles {
		hurdles[i] = false
	}
	for _, b := range x {
		hurdles[b] = true
	}
	//
	//for i := 1; i <= L; i++ {
	//	timeDp[i] = math.Min(timeDp[i], timeDp[i-1]+t1)
	//	if i >= 2 {
	//		timeDp[i] = math.Min(timeDp[i], timeDp[i-2]+t1+t2)
	//	}
	//	if i >= 4 {
	//		timeDp[i] = math.Min(timeDp[i], timeDp[i-4]+t1+t2*3)
	//	}
	//	if hurdles[i] {
	//		timeDp[i] += t3
	//	}
	//}
	//ans := timeDp[L]
	//
	//if L-1 >= 0 {
	//	ans = math.Min(ans, timeDp[L-1]+(t1+t2)/2)
	//}
	//if L-2 >= 0 {
	//	ans = math.Min(ans, timeDp[L-2]+(t1+t2)/2+t2)
	//}
	//if L-3 >= 0 {
	//	ans = math.Min(ans, timeDp[L-3]+(t1+t2)/2+t2*2)
	//}

	mod := 1000000007
	time := make([]int, L+1)
	inf := mod * mod
	for i := 1; i <= L; i++ {
		t := []int{inf, inf, inf}
		t[0] = time[i-1] + t1
		if i-2 >= 0 {
			t[1] = time[i-2] + t1 + t2
		}
		if i-4 >= 0 {
			t[2] = time[i-4] + t1 + 3*t2
		}
		time[i] = min(t)
		if hurdles[i] {
			time[i] += t3
		}
	}
	t := []int{time[L], time[L-1] + t1/2 + t2/2, inf, inf}
	if L-2 >= 0 {
		t[2] = time[L-2] + t1/2 + t2*3/2
	}
	if L-3 >= 0 {
		t[3] = time[L-3] + t1/2 + t2*5/2
	}
	fmt.Println(min(t))
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
