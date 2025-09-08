package main

import (
	"fmt"
)

func remove(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Printf("Ошибка: индекс %d вне диапазона [0..%d]\n", i, len(slice)-1)
		return slice
	}

	copy(slice[i:], slice[i+1:])
	fmt.Println("Массив после удаления i-го элемента: ", slice[:len(slice)-1])
	return slice[:len(slice)-1]
}

func main() {
	nums := []int{30, 20, 22, 30, 40, 50}
	fmt.Println("Базовый массив:", nums)

	nums = remove(nums, 11)
}
