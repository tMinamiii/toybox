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

func SplitNextLine(sc *bufio.Scanner) []string {
	sc.Scan()
	s := sc.Text()
	return strings.Split(s, " ")
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

	NSLine := SplitNextLine(sc)
	N := Atoi(NSLine[0])
	S := Atoi(NSLine[1])

	A := make([]int, 0, N)
	ALine := SplitNextLine(sc)
	for _, v := range ALine {
		A = append(A, Atoi(v))
	}

	dp := make([][]bool, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]bool, 10010)
	}
	dp[0][0] = true

	for i := 1; i < N+1; i++ {
		for j := 0; j < S+1; j++ {
			if dp[i-1][j] {
				dp[i][j] = dp[i-1][j]
				dp[i][j+A[i-1]] = true
			}
		}
	}

	if dp[N][S] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
