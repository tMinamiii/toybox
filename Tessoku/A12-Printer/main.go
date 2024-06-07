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

	NKLine := SplitToInt(NextLine(sc))
	N := NKLine[0]
	K := NKLine[1]

	A := SplitToInt(NextLine(sc))

	left := 1
	right := 1_000_000_000 // 1000_000_000 秒を探索範囲の最大にする
	for left < right {
		mid := (left + right) / 2
		sum := 0
		for i := 0; i < N; i++ {
			// 5_000_000_00秒 / A秒で何枚プリントされたか調べる
			sum += mid / A[i]
		}
		if sum >= K {
			right = mid
		} else {
			left = mid + 1
		}
	}

	fmt.Println(left)
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
