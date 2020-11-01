package main

import (
	"fmt"
	"time"
)

func sender(out chan int, exit chan bool) {
	for i := 1; i <= 10; i++ {
		out <- i
	}

	// time.Sleep(2000 * time.Millisecond)
	out <- 11
	exit <- true
}

func main() {
	out := make(chan int, 20)
	exit := make(chan bool)

	go sender(out, exit)

	time.Sleep(500 * time.Millisecond)

L:
	for {
		select {
		case i := <-out:
			fmt.Printf("S1, Value: %d\n", i)
		default:
			select {
			case i := <-out:
				fmt.Printf("S2 in default, Value: %d\n", i)
			case <-exit:
				select {
				case i := <-out:
					fmt.Printf("S3 in exit, Value: %d\n", i)
				default:
					fmt.Println("Exiting")
					break L
				}
			}
		}
	}
	fmt.Println("Did we get all 10? Yes.")
	fmt.Println("Did we get 11? DEFINITELY YES")
}
