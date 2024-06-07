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

	NXLine := SplitToInt(NextLine(sc))
	// N := NXLine[0]
	X := NXLine[1]

	A := SplitToInt(NextLine(sc))

	left := 0
	right := len(A)
	for left < right {
		mid := (left + right) / 2
		if A[mid] >= X {
			right = mid
		} else {
			left = mid + 1
		}
	}

	fmt.Println(left + 1)
}
