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

	NWLine := Split(NextLine(sc))
	N := Atoi(NWLine[0])
	W := Atoi(NWLine[1])

	weight := make([]int, 0, N)
	value := make([]int, 0, N)
	for i := 0; i < N; i++ {
		ALine := Split(NextLine(sc))
		weight = append(weight, Atoi(ALine[0]))
		value = append(value, Atoi(ALine[1]))
	}

	// 動的計画法のメモの初期化
	// 最大化問題なので、初期値は0にする
	// 最小化はINF, 最大化は0
	// N+10, W+10としてメモリ領域を多めに確保しておく
	dp := make([][]int, N+10)
	for i := 0; i < N+10; i++ {
		dp[i] = make([]int, W+10)
	}

	for i := 0; i < N; i++ {
		for sumW := 0; sumW <= W; sumW++ {
			if sumW-weight[i] >= 0 {
				dp[i+1][sumW] = Max(dp[i+1][sumW], dp[i][sumW-weight[i]]+value[i])
			}
			dp[i+1][sumW] = Max(dp[i+1][sumW], dp[i][sumW])
		}
	}

	fmt.Println(dp[N][W])
}
