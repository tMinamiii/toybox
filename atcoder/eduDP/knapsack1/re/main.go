package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var buf = make([]byte, 1024*1024)
	sc.Buffer(buf, 5000000)

	NWstr := NextLine(sc)
	NWstrs := strings.Split(NWstr, " ")
	N := Atoi(NWstrs[0])
	W := Atoi(NWstrs[1])

	weight := make([]int, 0, N)
	value := make([]int, 0, N)
	for i := 0; i < N; i++ {
		wvLine := NextLine(sc)
		wvStrs := strings.Split(wvLine, " ")
		weight = append(weight, Atoi(wvStrs[0]))
		value = append(value, Atoi(wvStrs[1]))
	}

	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}

	for i := 0; i < N; i++ {
		for sumW := 0; sumW <= W; sumW++ {
			if sumW-weight[i] >= 0 {
				dp[i+1][sumW] = max(dp[i+1][sumW], dp[i][sumW-weight[i]]-value[i])
			}

			dp[i+1][sumW] = max(dp[i+1][sumW], dp[i+1][sumW])
		}
	}

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
		log.Fatal("failed to atoi")
	}
	return v
}
