package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Unsafe counter
type Counter struct {
	v map[string]int
}

func (c *Counter) Inc(key string) {
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	return c.v[key]
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	// c := SafeCounter{v: make(map[string]int)}
	c := new(SafeCounter)
	c.v = make(map[string]int)
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))

	unsafeCounter := new(Counter)
	unsafeCounter.v = make(map[string]int)
	for i := 0; i < 2; i++ {
		// go unsafeCounter.Inc("unsafecount")
		go func(i int) {
			unsafeCounter.Inc(strconv.Itoa(i))
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println(unsafeCounter.Value("0"))
}
