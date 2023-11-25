// dup2 checks for files passed in from the local path
// if no files passed in use stdin for calculating duplicate lines
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("Input a file name")
		return
	}

	for _, arg := range files {
		file, err := os.ReadFile(arg)
		if err != nil {
			fmt.Printf("[ERROR]%v", err)
			continue
		}

		for _, f := range strings.Split(string(file), "\n") {
			counts[f]++
		}
	}

	for v, n := range counts {
		if n > 1 {
			fmt.Printf("%v\t%v\n", v, n)
		}
	}
}
