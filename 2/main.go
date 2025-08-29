package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	for i := 0; i < len(array); i++ {
		go fmt.Fprintf(os.Stdout, "square = %v\n", array[i]*array[i])
	}
	time.Sleep(1 * time.Second)
}
