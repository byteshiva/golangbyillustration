package main

import (
	"golang.org/x/tour/wc"
)

// WordCount returns count of words
func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
