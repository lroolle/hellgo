package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		// time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id, n)
	}

}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func gen() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500/10)) * time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
}

func main() {
	var c1, c2 = gen(), gen()

	n := 0
	w := createWorker(0)
	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			// fmt.Println("Received from c1: ", n)
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
			// fmt.Println("Received from c2: ", n)
			// default:
			// fmt.Println("No value received")
		case activeWorker <- activeValue:
			// send to w
			values = values[1:]
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Reach 800ms time out!")
		case <-tick:
			fmt.Printf("Remaining Values are: %v\n", values)
		case <-tm:
			fmt.Println("Good bye")
			return
		}
	}
}
