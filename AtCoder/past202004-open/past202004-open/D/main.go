package main

import (
	"bufio"
	"fmt"
	"os"
)

// 同じ長さの文字列の比較
func check(s, t string) bool {
	for i := 0; i < len(s); i++ {
		if t[i] == '.' {
			continue
		}
		if s[i] != t[i] {
			return false
		}
	}

	return true

}

func isMatch(s, t string) bool {
	for i := 0; i <= len(s)-len(t); i++ {
		if check(s[i:i+len(t)], t) {
			return true
		}
	}
	return false
}

func solve(S string) {
	str := "abcdefghijklmnopqrstuvwxyz."
	var cnt int
	for _, s1 := range str {
		t := string(s1)
		if isMatch(S, t) {
			cnt++
		}
		for _, s2 := range str {
			t = string(s1) + string(s2)
			if isMatch(S, t) {
				cnt++
			}
			for _, s3 := range str {
				t = string(s1) + string(s2) + string(s3)
				if isMatch(S, t) {
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)
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
