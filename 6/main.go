package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//1 return
	go func() {
		fmt.Println("start")
		time.Sleep(1 * time.Second)
		fmt.Println("end")
		return
	}()
	time.Sleep(2 * time.Second)

	//2* канал для сигнала
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("goroutine stopped")
				return
			default:
				fmt.Println("working...")
			}
		}
	}()

	time.Sleep(2 * time.Second)
	close(done)

	//3* context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopped:", ctx.Err())
				return
			default:
				fmt.Println("ready to work...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	//4 time.after
	go func() {
		timer := time.After(3 * time.Second)
		for {
			select {
			case <-timer:
				fmt.Println("time expired")
				return
			default:
				fmt.Println("tick")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	//5* закрытие канала
	ch := make(chan int)

	go func() {
		for val := range ch {
			fmt.Println("got:", val)
		}
		fmt.Println("channel closed, goroutine exit")
	}()

	ch <- 1
	ch <- 2
	close(ch)

	//6 WaitGroup
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("work")
		time.Sleep(time.Second)
	}()

	wg.Wait()

	//7 Flag
	var stop int32

	go func() {
		for {
			if atomic.LoadInt32(&stop) == 1 {
				fmt.Println("stopped")
				return
			}
			fmt.Println("need more gold...")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
	atomic.StoreInt32(&stop, 1)

}
