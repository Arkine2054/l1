package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func (s *SafeMap) Store(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SafeMap) Load(key string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.m[key]
	return val, ok
}

func main() {

	wg := sync.WaitGroup{}
	safeMap := &SafeMap{
		m: make(map[string]int),
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			safeMap.Store(key, id*10)
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			if v, ok := safeMap.Load(key); ok {
				fmt.Printf("Горутина %d прочитала: %v\n", id, v)
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("Все значения в карте:")
	for k, v := range safeMap.m {
		fmt.Println(k, "=", v)
	}
}
