package main

import (
	"fmt"
	"time"
)

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func chanDemo() {
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
}

func bufferedChan() {
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3
}

func chanClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c) // 关闭
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("=== Channel as first-class citizen ===")
	chanDemo()

	fmt.Println("=== Buffered channel ===")
	bufferedChan()

	fmt.Println("=== Channel close and range ===")
	chanClose()
}
