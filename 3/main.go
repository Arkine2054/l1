package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d done\n", id)
			return
		case job := <-jobs:
			fmt.Printf("worker# %d print %d \n", id, job)
		}
	}

}

func main() {
	var countStr string
	fmt.Println("Enter count of goroutines")
	fmt.Scan(&countStr)
	fmt.Printf("Goroutines: %v\n", countStr)

	count, err := strconv.Atoi(countStr)
	if err != nil || count <= 0 {
		fmt.Printf("Invalid count of goroutines\n")
	}

	jobs := make(chan int)

	cleanupFns := make([]func(), 0)

	for i := 0; i < count; i++ {
		ctx, cancel := context.WithCancel(context.Background())
		cleanupFns = append(cleanupFns, cancel)
		go worker(ctx, i, jobs)
	}

	n := 0

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh

		for _, fn := range cleanupFns {
			fn()
		}

		time.Sleep(time.Second)
		os.Exit(0)
	}()

	for {
		jobs <- n
		n++
	}

}
