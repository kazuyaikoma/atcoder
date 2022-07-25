package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var cnt int

func solve(K int64) {
	// index 0: num
	// index 1: cnt
	q := []int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
	}
	for {
		num := q[0]
		cnt++
		if cnt == int(K) {
			fmt.Println(num)
			return
		}
		q = q[1:]

		str := strconv.Itoa(num)

		last := str[len(str)-1]
		if last == '0' {
			n0, _ := strconv.Atoi(str + "0")
			n1, _ := strconv.Atoi(str + "1")
			q = append(q, n0)
			q = append(q, n1)
		} else if last < '9' {
			s0 := str + strconv.Itoa(int(last-'0'-1))
			s1 := str + strconv.Itoa(int(last-'0'))
			s2 := str + strconv.Itoa(int(last-'0'+1))
			n0, _ := strconv.Atoi(s0)
			n1, _ := strconv.Atoi(s1)
			n2, _ := strconv.Atoi(s2)
			q = append(q, n0)
			q = append(q, n1)
			q = append(q, n2)
		} else if last == '9' {
			n0, _ := strconv.Atoi(str + "8")
			n1, _ := strconv.Atoi(str + "9")
			q = append(q, n0)
			q = append(q, n1)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(K)
}
