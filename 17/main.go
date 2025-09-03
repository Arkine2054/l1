package main

import (
	"fmt"
)

func biFind(num int, arr []int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		half := (low + high) / 2

		if arr[half] == num {
			return half
		} else if arr[half] < num {
			low = half + 1
		} else {
			high = half - 1
		}
	}
	return -1
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	search := biFind(20, array)
	if search != -1 {
		fmt.Printf("Найдено %d на позиции %d\n", array[search], search)
	} else {
		fmt.Println("Элемент не найден")
	}
}
