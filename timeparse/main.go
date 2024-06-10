package main

import (
	"fmt"
	"time"
)

func main() {
	val := "25:10:10"
	layout := "03:04:05"
	t, err := time.Parse(layout, val)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(t)
}
