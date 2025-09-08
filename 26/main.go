package main

import (
	"fmt"
	"strings"
)

func onlyUnique(register string) bool {
	register = strings.ToLower(register)
	symbols := make(map[rune]bool)

	for _, s := range register {
		if symbols[s] {
			return false
		}
		symbols[s] = true
	}
	return true
}

func main() {
	check := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
		"GoLang",
		"Hello",
	}

	for _, c := range check {
		fmt.Printf("%q -> %v\n", c, onlyUnique(c))
	}
}
