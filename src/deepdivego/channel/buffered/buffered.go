package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c)
	var wg sync.WaitGroup
	go func(wg *sync.WaitGroup) {
		fmt.Println(<-c)
		wg.Done()
	}(&wg)
	wg.Add(1)
	wg.Wait()
}
