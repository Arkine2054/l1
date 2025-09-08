package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	start := time.Now()
	for {
		if time.Since(start) >= d {
			return
		}
	}
}

func main() {
	fmt.Println("Start")
	Sleep(2 * time.Second)
	fmt.Println("End")
}
