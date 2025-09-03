package main

import "fmt"

func convertString(word string) string {
	runes := []rune(word)
	size := len(runes)
	for i := 0; i < size/2; i++ {
		runes[i], runes[size-1-i] = runes[size-1-i], runes[i]
	}

	return string(runes)
}

func main() {
	fmt.Println(convertString("Никита"))
	fmt.Println(convertString("Welcome home"))
	fmt.Println(convertString("Еmoji 😊"))
}
