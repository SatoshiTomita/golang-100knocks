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

// sync.Mutexを使うことで、複数のゴルーチンが同時にカウンタを操作しても競合状態を回避できる

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	time.Sleep(1 * time.Second)
	c.count++
	fmt.Println("Increment", c.count)
}

func (c *Counter) Decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	time.Sleep(1 * time.Second)
	c.count--
	fmt.Println("Decrement", c.count)
}

func (c *Counter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	var counter Counter

	// 並行してカウンタを操作
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
			counter.Decrement()
		}()
	}
	wg.Wait()

	fmt.Println("Final count:", counter.GetCount())
}
