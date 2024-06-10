package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := sc.Text()

	lineValues := strings.Split(line, " ")
	asWords := lineValues[:len(lineValues)-1]
	mulwords := make(map[int]string, len(asWords))
	for _, v := range asWords {
		vals := strings.Split(v, ":")
		a, err := strconv.Atoi(vals[0])
		if err != nil {
			log.Fatalf("failed to multiple Atoi %v\n", err)
		}
		s := vals[1]
		mulwords[a] = s
	}

	mWord := lineValues[len(lineValues)-1]
	m, err := strconv.Atoi(mWord)
	if err != nil {
		log.Fatalf("failed to m Atoi %v\n", err)
	}

	matched := false
	for i := 1; i <= 20; i++ {
		if w, ok := mulwords[i]; ok {
			if m%i == 0 {
				fmt.Print(w)
				matched = true
			}
		}
	}

	if matched {
		fmt.Println()
	} else {
		fmt.Println(m)
	}

}
