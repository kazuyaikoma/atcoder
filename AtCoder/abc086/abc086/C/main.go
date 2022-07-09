package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(N int, t []int, x []int, y []int) {
	// 偶数秒(2x)後 → 2の倍数移動圏内であればOK
	// 奇数秒(2x+1)後 → 2の倍数移動圏内で移動した上で、±1どちらかに移動していればOK
	prevT := 0
	prevX, prevY := 0, 0
	for i := 0; i < len(t); i++ {
		deltaT := t[i] - prevT
		deltaPos := int(math.Abs(float64(x[i]-prevX))) + int(math.Abs(float64(y[i]-prevY)))
		if deltaT < deltaPos {
			fmt.Println(NO)
			return
		}
		if (deltaT)%2 == 0 {
			if deltaPos%2 != 0 {
				fmt.Println(NO)
				return
			}
		} else {
			if deltaPos%2 == 0 {
				fmt.Println(NO)
				return
			}
		}
		prevT, prevX, prevY = t[i], x[i], y[i]
	}
	fmt.Println(YES)
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
	t := make([]int, N)
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		t[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		x[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		y[i], _ = strconv.Atoi(scanner.Text())
	}
	solve(N, t, x, y)
}
