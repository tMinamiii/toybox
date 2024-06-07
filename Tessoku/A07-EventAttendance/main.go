package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	const max = 1024 * 1024
	var buf = make([]byte, max)
	sc.Buffer(buf, 5000000)

	D := Atoi(NextLine(sc))
	N := Atoi(NextLine(sc))

	L := make([]int, N+10)
	R := make([]int, N+10)
	for i := 1; i <= N; i++ {
		LR := SplitToInt(NextLine(sc))
		L[i] = LR[0]
		R[i] = LR[1]
	}

	B := make([]int, N+10)
	for i := 1; i <= N; i++ {
		B[L[i]] += 1
		B[R[i]+1] -= 1
	}

	fmt.Println(B)

	// 塁積和

	ans := make([]int, D+10)
	for d := 1; d <= D; d++ {
		ans[d] = ans[d-1] + B[d]
	}

	for d := 1; d <= D; d++ {
		fmt.Println(ans[d])
	}
}

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
