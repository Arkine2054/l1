package main

import "fmt"

func setBit(num int64, pos uint) int64 {
	num |= 1 << pos
	return num
}

func clearBit(num int64, pos uint) int64 {
	num &^= 1 << pos
	return num
}

func main() {

	for i := 1; i < 25; i++ {

		fmt.Printf("Исходное число: %d (%b)\n", i, i)
		newNum := int64(i)

		newNum = setBit(newNum, 1)
		fmt.Printf("После установки i-го бита в 1: %d (%b)\n", newNum, newNum)

		newNum = clearBit(newNum, 2)
		fmt.Printf("После установки i-го бита в 0: %d (%b)\n\n", newNum, newNum)

	}
}
