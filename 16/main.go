package main

import (
	"fmt"
)

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]
	left, right := 0, len(arr)-1

	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := range arr {
		if arr[i] < pivot {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])
}

func main() {
	array := []int{1, 3, 2, 10, 7, 15, 5, 4, 12, 15, 14, 24, 152, 111}
	quickSort(array)
	fmt.Println(array)
}
