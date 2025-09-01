package main

import (
	"fmt"
)

func main() {
	ss := []string{"cat", "cat", "dog", "cat", "tree"}

	unique := make(map[string]struct{})

	for _, s := range ss {
		unique[s] = struct{}{}
	}

	fmt.Println(unique)
}
