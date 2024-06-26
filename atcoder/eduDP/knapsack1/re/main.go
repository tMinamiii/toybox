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

	NWLine := SplitNextLine(sc)
	N := Atoi(NWLine[0])
	W := Atoi(NWLine[1])

	w := make([]int, 0, N)
	v := make([]int, 0, N)
	for i := 0; i < N; i++ {
		wvLine := SplitNextLine(sc)
		w = append(w, Atoi(wvLine[0]))
		v = append(v, Atoi(wvLine[1]))
	}

	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, 100010)
	}

	for i := 0; i < N; i++ {
		for sumW := 0; sumW <= W; sumW++ {
			// i番目の品物をはこぶ場合
			if sumW-w[i] >= 0 {
				dp[i+1][sumW] = Max(dp[i+1][sumW], dp[i][sumW-w[i]]+v[i])
			}
			dp[i+1][sumW] = Max(dp[i+1][sumW], dp[i][sumW])
		}
	}

	fmt.Println(dp[N][W])
}
