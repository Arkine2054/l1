package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	go func() {
		for _, val := range array {
			ch1 <- val

		}
		close(ch1)
	}()

	go func() {
		for x := range ch1 {
			ch2 <- x * 2
		}
		close(ch2)
	}()

	for val := range ch2 {
		fmt.Fprintf(os.Stdout, "Channel 2  %d\n", val)
		time.Sleep(500 * time.Millisecond)
	}
}
