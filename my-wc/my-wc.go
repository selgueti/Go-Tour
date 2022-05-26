package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	list := strings.Split(s, " ")

	for _, elem := range list {
		if _, ok := m[elem]; !ok {
			m[elem] = 1
		} else {
			m[elem]++
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
