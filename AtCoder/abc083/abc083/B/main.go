package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum(nums []int) int {
	var ret int
	for _, num := range nums {
		ret += num
	}
	return ret
}

func isNumFromAtoB(num, A, B int) bool {
	str := strconv.Itoa(num)
	var sum int
	for _, s := range str {
		delta := s - '0'
		if B < sum+int(delta) {
			return false
		}
		sum += int(delta)
	}

	return A <= sum && sum <= B
}

func solve(N, A, B int) {
	var nums []int
	for i := 1; i <= N; i++ {
		if isNumFromAtoB(i, A, B) {
			nums = append(nums, i)
		}
	}
	fmt.Println(sum(nums))
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
	var A int
	scanner.Scan()
	A, _ = strconv.Atoi(scanner.Text())
	var B int
	scanner.Scan()
	B, _ = strconv.Atoi(scanner.Text())
	solve(N, A, B)
}
