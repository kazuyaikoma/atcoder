package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const YES = "YES"
const NO = "NO"
const (
	eraser  = "eraser"
	erase   = "erase"
	dreamer = "dreamer"
	dream   = "dream"
)

func solve(S string) {
	replaced := strings.ReplaceAll(S, eraser, "")
	replaced = strings.ReplaceAll(replaced, erase, "")
	replaced = strings.ReplaceAll(replaced, dreamer, "")
	replaced = strings.ReplaceAll(replaced, dream, "")
	if replaced == "" {
		fmt.Println(YES)
	} else {
		fmt.Println(NO)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var S string
	scanner.Scan()
	S = scanner.Text()
	solve(S)
}
