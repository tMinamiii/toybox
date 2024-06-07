package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NextLine(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func SplitToInt(s string) []int {
	strs := strings.Split(s, " ")
	ary := make([]int, 0, len(strs))
	for _, v := range strings.Split(s, " ") {
		ary = append(ary, Atoi(v))
	}
	return ary
}

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func Abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func Min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	const max = 1024 * 1024
	var buf = make([]byte, max)
	sc.Buffer(buf, 5000000)

	N := Atoi(NextLine(sc))
	h := SplitToInt(NextLine(sc))

	dp := make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = 1 << 30
	}
	dp[0] = 0

	for i := 0; i < N; i++ {
		if i+1 < N {
			dp[i+1] = min(dp[i+1], dp[i]+Abs(h[i+1], h[i]))
		}
		if i+2 < N {
			dp[i+2] = min(dp[i+2], dp[i]+Abs(h[i+2], h[i]))
		}
	}

	fmt.Println(dp[N-1])
}
