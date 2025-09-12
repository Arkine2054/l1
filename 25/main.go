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
func SleepAfter(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Start")
	Sleep(2 * time.Second)
	fmt.Println("End")

	fmt.Println("StartAfter")
	SleepAfter(2 * time.Second)
	fmt.Println("EndAfter")

}
