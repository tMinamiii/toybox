package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	var str *string
	fmt.Println(lo.Min([]int{9, 8, 2, 1, 4, 5}))
	fmt.Println(lo.Max([]int{9, 8, 2, 1, 4, 5}))

	fmt.Println(lo.FromPtr(str)) // ゼロ値。つまり空文字。
	fmt.Println(lo.FromPtrOr(str, "別の値"))
	// Output:
	//
	// 別の値
}
