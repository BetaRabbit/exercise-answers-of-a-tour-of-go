package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	m := make(map[string]int)

	for i := range arr {
		m[arr[i]]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
