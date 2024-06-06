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
	Nstr := NextLine(sc)
	N := Atoi(Nstr)

	a := make([][]int, N)
	for i := 0; i < N; i++ {
		astr := NextLine(sc)
		astrs := strings.Split(astr, " ")
		for _, v := range astrs {
			aij := Atoi(v)
			a[i] = append(a[i], aij)
		}
	}

	// メモするのは 3パターン
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		for j := 0; j < 3; j++ {
			dp[i] = append(dp[i], 0)
		}
	}

	// 全ルート6パターンしらべてメモしていく
	for i := 0; i < N; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j == k {
					continue
				}
				dp[i+1][k] = max(dp[i+1][k], dp[i][j]+a[i][k])
			}
		}
	}

	result := 0
	for i := 0; i < 3; i++ {
		result = max(result, dp[N][i])
	}

	fmt.Println(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func NextLine(sc *bufio.Scanner) string {
	sc.Scan()
	line := sc.Text()
	return line
}

func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to atoi")
	}
	return v
}
