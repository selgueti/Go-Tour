package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	list := strings.Split(s, " ")

	for _, word := range list {
		if _, ok := m[word]; !ok {
			m[word] = 1
		} else {
			m[word]++
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
