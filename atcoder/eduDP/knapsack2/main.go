package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const MaxV = 100000
const INF = 1 << 30

func main() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, 500000)

	NWLine := NextLine(sc)
	NWStrs := strings.Split(NWLine, " ")
	N := Atoi(NWStrs[0])
	W := Atoi(NWStrs[1])

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
		dp[i] = make([]int, MaxV+1)
		for j := 0; j <= MaxV; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0

	// dp[ i ][ sum_w ] := i-1 番目までの品物から重さが sum_w 以下となるように選んだときの、価値の総和の最大値
	// dp[ i ][ sum_v ] := i-1 番目までの品物から価値が sum_v となるように選んだときの、重さの総和の最小値
	for i := 0; i < N; i++ {
		for sumV := 0; sumV <= MaxV; sumV++ {
			if sumV-value[i] >= 0 {
				dp[i+1][sumV] = min(dp[i+1][sumV], dp[i][sumV-value[i]]+weight[i])
			}
			dp[i+1][sumV] = min(dp[i+1][sumV], dp[i][sumV])
		}
	}

	result := 0
	for sumV := 0; sumV <= MaxV; sumV++ {
		if dp[N][sumV] <= W {
			result = sumV
		}
	}

	fmt.Println(result)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
