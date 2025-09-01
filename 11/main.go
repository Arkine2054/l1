package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	cross := make(map[int][]int)

	for _, c := range a {
		key := c % 2
		cross[key] = append(cross[key], c)
	}
	for _, c := range b {
		key := c % 3
		cross[key] = append(cross[key], c)
	}

	fmt.Printf("Пересечение a и b: %v", cross[0])
}
