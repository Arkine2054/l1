package main

import (
	"fmt"
)

func checkType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("int: %d\n", v)
	case string:
		fmt.Printf("string: %q\n", v)
	case bool:
		fmt.Printf("bool: %t\n", v)
	case chan int:
		fmt.Printf("chan: %T\n", v)

	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	ch := make(chan int)
	ch2 := make(chan string)
	truth := true

	checkType(123)
	checkType("string")
	checkType(truth)
	checkType(ch)
	checkType(ch2)

}
