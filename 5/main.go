package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for res := range ch {
			fmt.Fprintf(os.Stdout, "Result: %v\n", res)
		}
	}()

	<-time.After(3 * time.Second)

	fmt.Println("Done")
	close(ch)
}
