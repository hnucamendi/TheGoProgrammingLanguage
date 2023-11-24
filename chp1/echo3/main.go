package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	nE := testFunc(nonEff, 1000)
	e := testFunc(eff, 1000)
	fmt.Printf("func: nonEff\nTime:%v NS\n", nE)
	fmt.Printf("func: eff\nTime:%v NS\n", e)

	nEF := nonEff()
	eF := eff()

	fmt.Println(nEF)
	fmt.Println(eF)
}

func nonEff() string {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += string(rune(i)) + sep + os.Args[i]
		sep = " "
	}
	return s
}

func eff() string {
	return strings.Join(os.Args[0:], " ")
}

func testFunc(f func() string, t int) int64 {
	c := 0
	times := make([]int64, t)

	for c != t {
		start := time.Now().UnixNano()
		f()
		times[c] = time.Now().UnixNano() - start
		c++
	}
	return calcAverage(times)
}

func calcAverage(s []int64) int64 {
	var total int64 = 0
	for _, n := range s {
		total += n
	}

	return total / int64(len(s))
}
