package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(A [][]int64, N int64, b []int64) {
	bingo := [][]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	}
	// 操作
	for n := 0; n < int(N); n++ {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if A[i][j] == b[n] {
					bingo[i][j] = true
				}
			}
		}
	}
	// 判定
	for i := 0; i < 3; i++ {
		if bingo[i][0] && bingo[i][1] && bingo[i][2] {
			fmt.Println(YES)
			return
		}
		if bingo[0][i] && bingo[1][i] && bingo[2][i] {
			fmt.Println(YES)
			return
		}
	}
	if bingo[0][0] && bingo[1][1] && bingo[2][2] {
		fmt.Println(YES)
		return
	}
	if bingo[0][2] && bingo[1][1] && bingo[2][0] {
		fmt.Println(YES)
		return
	}
	fmt.Println(NO)

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	A := make([][]int64, 3)
	for i := int64(0); i < 3; i++ {
		A[i] = make([]int64, 3)
	}
	for i := int64(0); i < 3; i++ {
		for j := int64(0); j < 3; j++ {
			scanner.Scan()
			A[i][j], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		}
	}
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	b := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		b[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(A, N, b)
}
