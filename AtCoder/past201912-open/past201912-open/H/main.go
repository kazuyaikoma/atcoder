package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func minAndOddMin(ns []int) (int, int) {
	setMin, allMin := int(math.Pow(10, 9))+1, int(math.Pow(10, 9))+1
	for i, n := range ns {
		if i%2 == 0 && n < setMin {
			setMin = n
		} else if n < allMin {
			allMin = n
		}
	}

	// setMin: セット販売対象カードのうち最小枚数
	// allMin: 全カードのうち最小枚数
	return setMin, allMin
}

func solve(N, Q int, cs []int, ss [][]int) {
	// setMin: セット販売対象カードのうち最小枚数
	// allMin: 全カードのうち最小枚数
	setMin, allMin := minAndOddMin(cs)
	// totalSetSell: セット販売対象カードの合計販売枚数
	// totalAllSell: オール販売対象カードの合計販売枚数
	totalSetSell, totalAllSell := 0, 0

	// 合計販売枚数
	var sell int

	for _, s := range ss {
		switch s[0] {
		case 1:
			x := s[1] - 1
			c := cs[x]

			c -= totalAllSell
			if x%2 == 0 {
				c -= totalSetSell
			}

			if s[2] <= c {
				c -= s[2]
				cs[x] -= s[2]
				sell += s[2]

				// 各最小枚数の更新
				if x%2 == 0 && c < setMin {
					setMin = c
				}
				if c < allMin {
					allMin = c
				}
			}
		case 2:
			if s[1] <= setMin {
				sell += s[1] * ((len(cs) + 1) / 2)

				// 各最小枚数の更新
				setMin -= s[1]
				if setMin < allMin {
					allMin = setMin
				}
				// 販売合計枚数の更新
				totalSetSell += s[1]
			}
		case 3:
			if s[1] <= setMin && s[1] <= allMin {
				sell += s[1] * len(cs)
				// 各最小枚数の更新
				setMin -= s[1]
				allMin -= s[1]
				// 販売合計枚数の更新
				totalAllSell += s[1]
			}
		}
	}
	fmt.Println(sell)
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

	var cs []int
	for i := 0; i < N; i++ {
		var c int
		scanner.Scan()
		c, _ = strconv.Atoi(scanner.Text())
		cs = append(cs, c)
	}

	var Q int
	scanner.Scan()
	Q, _ = strconv.Atoi(scanner.Text())

	var ss [][]int
	for i := 0; i < Q; i++ {
		scanner.Scan()
		s1, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		s2, _ := strconv.Atoi(scanner.Text())
		s := []int{s1, s2}
		if s1 == 1 {
			scanner.Scan()
			s3, _ := strconv.Atoi(scanner.Text())
			s = append(s, s3)
		}
		ss = append(ss, s)
	}

	solve(N, Q, cs, ss)
}
