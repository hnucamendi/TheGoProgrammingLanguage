package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := 0
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() && count < 10 {
		counts[input.Text()]++
		count++
	}
	if err := input.Err(); err != nil {
		fmt.Printf("ERROR:%v\n", err)
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
		}
	}
}
