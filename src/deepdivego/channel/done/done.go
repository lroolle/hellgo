package main

import (
	"fmt"
)

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		// 通知打印完了
		// done <- true
		go func() { done <- true }()

	}
}

func chanDemo() {
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	for _, worker := range workers {
		<-worker.done
	}
}

func bufferedChan() {
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3
}

func main() {
	fmt.Println("=== Channel as first-class citizen ===")
	chanDemo()
}
