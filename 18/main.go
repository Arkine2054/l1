package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Incrementation() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	counter := &Counter{}

	for i := 0; i < 1000; i++ {
		go func() {
			counter.Incrementation()
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("Итоговое значение счетчика:", counter.Value())
}
