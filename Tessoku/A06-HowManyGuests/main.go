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

	NQ := SplitToInt(NextLine(sc))
	N := NQ[0]
	Q := NQ[1]

	A := SplitToInt(NextLine(sc))

	L := make([]int, 0, Q)
	R := make([]int, 0, Q)
	for i := 0; i < Q; i++ {
		LR := SplitToInt(NextLine(sc))
		L = append(L, LR[0])
		R = append(R, LR[1])
	}
	// 累積計算メモ
	S := make([]int, N+1)
	for i := 0; i < N; i++ {
		S[i+1] = S[i] + A[i]
	}

	for i := 0; i < Q; i++ {
		sum := S[R[i]] - S[L[i]-1]
		fmt.Println(sum)
	}
}
