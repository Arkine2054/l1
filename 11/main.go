package main

import (
	"fmt"
)

func cross(a, b []int) []int {
	m := make(map[int]bool)
	var result []int

	for _, val := range a {
		m[val] = true
	}

	for _, val := range b {
		if m[val] {
			result = append(result, val)
			m[val] = false

		}
	}

	return result
}

func main() {
	A := []int{1, 2, 3, 14, 52, 123, 132}
	B := []int{2, 3, 4, 5, 12, 13, 14, 15, 123}

	fmt.Println("Пересечение A и Б:", cross(A, B))
}
