package main

import (
	"fmt"
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	f := strings.Fields(s)
	fmt.Println(s, f)
	m := make(map[string]int)
	for _, v := range f {
		m[v] = m[v] + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
