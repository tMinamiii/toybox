package main

import (
	"fmt"
	"iter"
)

func main() {
	s := []string{"hello", "world"}
	for i, x := range Backward(s) {
		fmt.Println(i, x)
	}
	for i := range Generator() {
		fmt.Println(i)
	}

}

func Backward[E any](s []E) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			fmt.Println("gen")
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

func Generator() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range 5 {
			if !yield(i) {
				return
			}
		}
	}
}
