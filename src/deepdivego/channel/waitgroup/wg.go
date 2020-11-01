package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	// wg *sync.WaitGroup
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		// wg: wg,
		done: func() { wg.Done() },
	}
	go doWork(id, w)
	return w
}

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		// 通知打印完了
		// done <- true
		w.done()

	}
}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
		wg.Add(1)
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
		wg.Add(1)
	}
	wg.Wait()
	// for _, worker := range workers {
	// 	<-worker.done
	// }
}

func main() {
	fmt.Println("=== Channel as first-class citizen ===")
	chanDemo()
}
