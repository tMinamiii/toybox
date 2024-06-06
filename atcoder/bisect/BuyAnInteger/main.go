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

	ABXLine := NextLine(sc)
	ABXStr := strings.Split(ABXLine, " ")
	A := Atoi(ABXStr[0])
	B := Atoi(ABXStr[1])
	X := Atoi(ABXStr[1])

	ok := 0
	ng := 1000_000_000 + 1

	for abs(ok, ng) > 1 {
		mid := (ok + ng) / 2
		strMid := strconv.Itoa(mid)
		d := len(strMid)
		price := A*mid + B*d
		if price <= X {
			ok = mid
		} else {
			ng = mid
		}
	}

	fmt.Println(ok)
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func NextLine(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
