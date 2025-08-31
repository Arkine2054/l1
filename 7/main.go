package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

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

func NewConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{
		m: make(map[K]V),
	}
}

func (c *ConcurrentMap[K, V]) Store(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}

func (c *ConcurrentMap[K, V]) Load(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *ConcurrentMap[K, V]) Range(f func(key K, value V) bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	for k, v := range c.m {
		if !f(k, v) {
			break
		}
	}
}

func main() {
	//ConcurentMap
	cm := NewConcurrentMap[string, int]()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			cm.Store(key, id*100)
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			if v, ok := cm.Load(key); ok {
				fmt.Printf("Reader %d прочитал: %s = %v\n", id, key, v)
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("Все значения в ConcurrentMap:")
	cm.Range(func(k string, v int) bool {
		fmt.Println(k, "=", v)
		return true
	})

	//Mutex
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
