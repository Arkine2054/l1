package main

import (
	"fmt"
)

func reverseBytes(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func reverseText(s string) string {
	b := []byte(s)

	reverseBytes(b, 0, len(b)-1)

	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverseBytes(b, start, i-1)
			start = i + 1
		}
	}

	return string(b)
}

func main() {
	s := "snow dog sun"
	fmt.Println(reverseText(s))

	s2 := "Nikita love to play basketball"
	fmt.Println(reverseText(s2))
}
