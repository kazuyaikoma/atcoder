package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(c [][]int64) {
	if len(c) < 1 {
		fmt.Println(YES)
		return
	}
	// 先に最小の組み合わせを特定
	base := c[0]
	for _, arr := range c {
		if arr[0] < base[0] {
			base = arr
		}
	}

	for _, arr := range c {
		diff := arr[0] - base[0]
		for i := 0; i < 3; i++ {
			if base[i] != arr[i]-diff {
				fmt.Println(NO)
				return
			}
		}
	}
	fmt.Println(YES)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	c := make([][]int64, 3)
	for i := int64(0); i < 3; i++ {
		c[i] = make([]int64, 3)
	}
	for i := int64(0); i < 3; i++ {
		for j := int64(0); j < 3; j++ {
			scanner.Scan()
			c[i][j], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		}
	}
	solve(c)
}
