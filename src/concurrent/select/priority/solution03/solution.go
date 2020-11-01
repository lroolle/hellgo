package main

import "fmt"

func sender(out chan int) {
	defer close(out)
	for i := 1; i <= 10; i++ {
		out <- i
	}
}

func main() {
	out := make(chan int, 10)
	go sender(out)

	for i := range out {
		fmt.Printf("Received: %v\n", i)
	}
}
