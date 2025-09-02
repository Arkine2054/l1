package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 14, 52, 123, 132}
	b := []int{2, 3, 4, 5, 12, 13, 14, 15, 123}

	cross := []int{}

	for _, vala := range a {
		for _, valb := range b {
			if vala == valb {
				cross = append(cross, vala)
			}

		}
	}

	fmt.Printf("Пересечение a и b: %v", cross)
}
