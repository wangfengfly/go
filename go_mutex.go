// go_mutex.go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//SafeCounter is safe to user concurrently
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	defer wg.Done()

	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go c.Inc("somekey")
	}

	wg.Wait()

	fmt.Println(c.Value("somekey"))
}
