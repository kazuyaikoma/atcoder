package main

import (
	"bufio"
	"os"
	"strconv"
)

const YES = "Yes"
const NO = "No"

func solve(c [][]int64) {

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
