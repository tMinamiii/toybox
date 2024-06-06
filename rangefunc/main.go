package main

import "fmt"

func main() {
	s := []string{"hello", "world"}
	for i, x := range Backward(s) {
		fmt.Println(i, x)
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
