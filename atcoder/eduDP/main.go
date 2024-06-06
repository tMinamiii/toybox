package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
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
	n := StrToInt(NextLine(sc))
	h := SplitIntlist(NextLine(sc))
	dp := make([]int, n+1)
	for i := 2; i < n+1; i++ {
		if i == 2 {
			dp[i] = abs(h[i-1], h[i-2])
		} else {
			dp[i] = min(dp[i-1]+abs(h[i-2], h[i-1]), dp[i-2]+abs(h[i-3], h[i-1]))
		}
	}

	fmt.Println(dp[n])
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func bisect(a []int, target int) int {
	left, right := 0, len(a)
	for left < right {
		mid := (left + right) / 2
		if a[mid] == target {
			return mid
		}
		if a[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func judge(a []int, target int) bool {
	left, right := 0, len(a)
	for left < right {
		mid := (left + right) / 2
		if a[mid] == target {
			return true
		}
		if a[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return false
}

func lmin(a []int) int {
	res := 1 << 62
	for _, i := range a {
		if i < res {
			res = i
		}
	}
	return res
}

func lmax(a []int) int {
	res := 1 >> 30
	for _, i := range a {
		if i > res {
			res = i
		}
	}
	return res
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// NextLine buinfo.Scanのポインタを渡し、標準入力の次の行を読み込み
// ex. sc := buinfo.NewScanner(os.stdin)
//
//	GetNextLine(sc)
func NextLine(sc *bufio.Scanner) string {
	sc.Scan()
	s := sc.Text()
	return strings.TrimSpace(s)
}

// SplitStrList 文字列を空白区切りの文字列のリストに変換して返却
func SplitStrList(s string) []string {
	return strings.Split(s, " ")
}

// SplitIntlist 文字列を空白区切りの整数値に変換して返却
func SplitIntlist(s string) []int {
	strList := strings.Split(s, " ")
	return StrListToIntList(strList)
}

// StrListToIntList string型のスライスを渡してint型の配列に変換
func StrListToIntList(strList []string) (intList []int) {
	for _, str := range strList {
		str = strings.TrimRight(str, "\n")
		i := StrToInt(str)
		intList = append(intList, i)
	}
	return
}

// StrToInt string型をint型に変換
func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// Sort int型スライスの並び替え
func Sort(slice []int, order string) []int {
	sort.SliceStable(slice, func(i, j int) bool {
		if order == "desc" {
			return slice[i] > slice[j]
		} else {
			return slice[i] < slice[j]
		}
	})
	return slice
}

// FindMaxAndMin 最大値最小値を返す
func FindMaxAndMin(slice []int) (max, min int) {
	max = slice[0]
	min = slice[0]
	for _, elm := range slice {
		if elm > max {
			max = elm
		}
		if elm < min {
			min = elm
		}
	}
	return max, min
}

// Sum 合計値を返す
func Sum(slice []int) (sum int) {
	for _, i := range slice {
		sum += i
	}
	return sum
}

// Permutation Pの値を計算
func Permutation(n int, k int) *big.Int {
	v := big.NewInt(1)
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			k := big.NewInt(int64(n - i))
			v = v.Mul(v, k)
		}
	} else if k > n {
		v = big.NewInt(0)
	}
	return v
}

// Factorial Fの値を計算
func Factorial(n int) *big.Int {
	return Permutation(n, n-1)
}

// Combination Cの計算
func Combination(n int, k int) *big.Int {
	child := Permutation(n, k)
	mother := Factorial(k)
	return child.Div(child, mother)
}

// Homogeneous Hの計算
func Homogeneous(n int, k int) *big.Int {
	return Combination(n+k-1, k)
}
