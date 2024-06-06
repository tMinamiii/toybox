package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	const max = 1024 * 1024
	var buf = make([]byte, max)
	sc.Buffer(buf, 5000000)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var buf = make([]byte, 1024*1024)
	sc.Buffer(buf, 500000)

	NKLine := NextLine(sc)
	NKStrings := strings.Split(NKLine, " ")
	N := Atoi(NKStrings[0])
	K := Atoi(NKStrings[1])

	hLine := NextLine(sc)
	hstrings := strings.Split(hLine, " ")
	h := make([]int, 0, N)
	for _, v := range hstrings {
		h = append(h, Atoi(v))
	}

	dp := make([]int, N)
	for i := 1; i < N; i++ {
		dp[i] = 1 << 30
	}

	for i := 0; i < N; i++ {
		for j := 1; j <= K; j++ {
			if i+j < N {
				dp[i+j] = min(dp[i+j], dp[i]+abs(h[i+j], h[i]))
			}
		}
	}

	fmt.Println(dp[N-1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func NextLine(sc *bufio.Scanner) string {
	sc.Scan()
	line := sc.Text()
	return line
}

func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln("failed to atoi")
	}

	return v
}
