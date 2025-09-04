package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Counter struct {
	count int64
}

func (c *Counter) Incrementation() {
	atomic.AddInt64(&c.count, 1)
}

func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.count)
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
