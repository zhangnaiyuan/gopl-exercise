package main

import (
	"fmt"
	"strings"
	"time"
)

var args = []string{"The", "quick", "brown", "fox", "jump", "sover", "the", "lazy", "dog"}

func test(f func() string) {
	start := time.Now()
	s := f()
	cost := time.Since(start).Nanoseconds()
	fmt.Printf("result is %s, and cost %d nanoseconds\n", s, cost)
}

func concat() (s string) {
	var sep string
	for _, v := range args {
		s += sep + v
		sep = " "
	}
	return
}
func join() (s string) {
	s = strings.Join(args, " ")
	return
}

func main() {
	test(concat)
	test(join)
}
