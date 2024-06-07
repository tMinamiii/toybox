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

func Split(s string) []string { return strings.Split(s, " ") }

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

	ALine := Split(NextLine(sc))
	A := make([]int, N)
	for i, v := range ALine {
		A[i+1] = Atoi(v)
	}

	BLine := Split(NextLine(sc))
	B := make([]int, N)
	for i, v := range BLine {
		B[i+2] = Atoi(v)
	}

	dp := make([]int, N)
	for i := 1; i < N; i++ {
		dp[i] = 1 << 30
	}

	for i := 0; i < N; i++ {
		if i+1 < N {
			dp[i+1] = Min(dp[i+1], dp[i]+A[i+1])
		}
		if i+2 < N {
			dp[i+2] = Min(dp[i+2], dp[i]+B[i+2])
		}
	}

	fmt.Println(dp[N-1])
}
