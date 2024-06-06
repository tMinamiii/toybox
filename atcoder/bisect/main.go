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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, 500000)

	NKLine := NextLine(sc)
	NKStr := strings.Split(NKLine, " ")
	N := Atoi(NKStr[0])
	K := Atoi(NKStr[1])

	ALine := NextLine(sc)
	AStr := strings.Split(ALine, " ")
	A := make([]int, 0, N)
	for _, v := range AStr {
		A = append(A, Atoi(v))
	}

	i := bisect(A, K)
	fmt.Println(i)
}

func bisect(a []int, target int) int {
	left, right := 0, len(a)
	for left < right {
		mid := (left + right) / 2
		if a[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if right >= len(a) {
		return -1
	}

	return left
}

func NextLine(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
