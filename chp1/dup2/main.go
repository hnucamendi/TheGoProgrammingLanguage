// dup2 checks for files passed in from the local path
// if no files passed in use stdin for calculating duplicate lines
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		calculateDuplicates(os.Stdin, counts)
		return
	}

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("[ERROR]%v", err)
			os.Exit(1)
		}

		calculateDuplicates(f, counts)
	}
}

func calculateDuplicates(f *os.File, c map[string]int) error {
	s := bufio.NewScanner(f)
	for s.Scan() {
		c[s.Text()]++
	}

	if err := s.Err(); err != nil {
		return err
	}

	return nil
}
